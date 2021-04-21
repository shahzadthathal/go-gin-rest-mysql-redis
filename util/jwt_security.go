package util

import (
	"JWT_REST_Gin_MySQL/model"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gomodule/redigo/redis"

	"github.com/spf13/viper"

	"github.com/segmentio/ksuid"

	"github.com/dgrijalva/jwt-go"
)

var accessSecret string
var refreshSecret string
var redisDSN string
var maxActive, maxIdle int
var timeToken = int(15)       // get from system param (minutes)
var timeRefreshToken = int(3) // get from system param (hours)

// Pool ...
var Pool *redis.Pool

// TokenDetails ...
type TokenDetails struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	AccessUUID   string `json:"-"`
	RefreshUUID  string `json:"-"`
	AtExpires    int64  `json:"atExpires"`
	RtExpires    int64  `json:"rtExpires"`
}

// AccessDetails ...
type AccessDetails struct {
	AccessUUID string
	UserID     int64
}

// SetupRedisJWT ...
func SetupRedisJWT() *redis.Pool {
	accessSecret = viper.GetString("JWT.ACCESS_SECRET")
	refreshSecret = viper.GetString("JWT.REFRESH_SECRET")
	redisDSN = viper.GetString("REDIS.DSN")
	maxActive = viper.GetInt("REDIS.MAX_ACTIVE")
	maxIdle = viper.GetInt("REDIS.MAX_IDLE")

	Pool = &redis.Pool{
		MaxActive: maxActive,
		MaxIdle:   maxIdle,
		// IdleTimeout: 240*time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", redisDSN)
		},
	}

	return Pool

}

// CreateToken ...
func CreateToken(u model.MUser) (*TokenDetails, error) {

	td := &TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * time.Duration(timeToken)).Unix()
	td.AccessUUID = ksuid.New().String()

	td.RtExpires = time.Now().Add(time.Hour * time.Duration(timeRefreshToken)).Unix()
	td.RefreshUUID = td.AccessUUID + "++" + strconv.FormatInt(u.ID, 10)

	var err error

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":         td.AtExpires,
		"access_uuid": td.AccessUUID,
		"user_id":     u.ID,
		"name":        u.UserName,
		"authorized":  true,
	})
	td.AccessToken, err = at.SignedString([]byte(accessSecret))
	if err != nil {
		return nil, err
	}

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":          td.RtExpires,
		"refresh_uuid": td.RefreshUUID,
		"user_id":      u.ID,
		"name":         u.UserName,
	})
	td.RefreshToken, err = rt.SignedString([]byte(refreshSecret))
	if err != nil {
		return nil, err
	}

	return td, nil
}

// SaveToRedis ...
func SaveToRedis(userID int64, td *TokenDetails) error {

	// Use the connection pool's Get() method to fetch a single Redis
	// connection from the pool.
	conn := Pool.Get()

	// Importantly, use defer and the connection's Close() method to
	// ensure that the connection is always returned to the pool.
	defer conn.Close()

	_, errAccess := conn.Do("SET", td.AccessUUID, userID)
	if errAccess != nil {
		return errAccess
	}

	_, errReferesh := conn.Do("SET", td.RefreshUUID, userID)
	if errReferesh != nil {
		return errReferesh
	}

	return nil
}

// ExtractFromRedis ...
func ExtractFromRedis(r *http.Request) (*AccessDetails, error) {
	tokenStr := ExtractToken(r)
	// verify token
	token, err := VerifyToken(r, tokenStr)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUUID, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}

		userID := claims["user_id"].(float64)

		// get pool connection redigo
		conn := Pool.Get()
		defer conn.Close()

		redisIDUser, err := redis.Int64(conn.Do("GET", accessUUID))
		if err != nil {
			return nil, err
		}

		if int64(userID) != redisIDUser {
			return nil, errors.New("Authentification failure")
		}

		return &AccessDetails{
			AccessUUID: accessUUID,
			UserID:     redisIDUser,
		}, nil
	}
	return nil, err
}

// ExtractToken ...
func ExtractToken(r *http.Request) string {

	bearToken := r.Header.Get("Authorization")
	if len(bearToken) == 0 {
		return ""
	}

	// extract token
	tokenArr := strings.Split(bearToken, " ")
	if len(tokenArr) == 2 {
		return tokenArr[1]
	}

	return ""
}

// VerifyToken ...
func VerifyToken(r *http.Request, tokenStr string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			s := fmt.Sprintf("unexpected signing method: %v", token.Header["alg"])
			return nil, errors.New(s)
		}

		return []byte(accessSecret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// DeleteAuthByUUID ...
func DeleteAuthByUUID(UUID string) error {

	conn := Pool.Get()
	defer conn.Close()

	_, err := conn.Do("DEL", UUID)
	if err != nil {
		return err
	}
	return nil
}

// DeleteToken ...
func DeleteToken(authD *AccessDetails) error {
	// get refresh uuid
	refreshUUID := fmt.Sprintf("%s++%d", authD.AccessUUID, authD.UserID)

	var err error
	// delete access token
	err = DeleteAuthByUUID(authD.AccessUUID)
	if err != nil {
		return err
	}

	// delete refresh token
	err = DeleteAuthByUUID(refreshUUID)
	if err != nil {
		return err
	}

	return nil
}
