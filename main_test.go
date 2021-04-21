package main

import (
	"JWT_REST_Gin_MySQL/configuration"
	"JWT_REST_Gin_MySQL/model"
	"JWT_REST_Gin_MySQL/router"
	"JWT_REST_Gin_MySQL/util"
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

var db sql.DB
var client = &http.Client{}
var userID int64

func TestMain(m *testing.T) {
	viper.SetConfigFile("./resource/properties-test.yaml")

	viper.ReadInConfig()

	util.Pool = util.SetupRedisJWT()

	var err error

	// Setup database
	configuration.DB, err = configuration.SetupDB()
	if err != nil {
		log.Fatal(err)
	}
}

var user = &model.MUser{
	ID:       1999999999,
	UserName: "haha",
	Password: "pass123",
}

func NewMock() (*sql.DB, sqlmock.Sqlmock, error) {

	db, mock, err := sqlmock.New()

	return db, mock, err
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {

	rr := httptest.NewRecorder()

	r := router.NewRoutes()

	var transport http.RoundTripper = &http.Transport{
		DisableKeepAlives: true,
	}
	client.Transport = transport

	r.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

// ====================== test API ======================
func TestAPILoginUser(t *testing.T) {
	payload := []byte(`{"username":"admin", "password":"admin1234"}`)

	req, _ := http.NewRequest("POST", "/api/login", bytes.NewBuffer(payload))
	resp := executeRequest(req)

	checkResponseCode(t, http.StatusOK, resp.Code)

	var m map[string]interface{}
	json.Unmarshal(resp.Body.Bytes(), &m)

	assert.NotEqual(t, m["accessToken"], "")
}

func TestAPILogoutUser(t *testing.T) {
	// doing login first to get token
	payloadLogin := []byte(`{"username":"admin", "password":"admin1234"}`)

	reqLogin, _ := http.NewRequest("POST", "/api/login", bytes.NewBuffer(payloadLogin))
	respLogin := executeRequest(reqLogin)

	checkResponseCode(t, http.StatusOK, respLogin.Code)

	var mLogin map[string]interface{}
	json.Unmarshal(respLogin.Body.Bytes(), &mLogin)

	var token = mLogin["accessToken"]
	tokenString := fmt.Sprintf("Bearer %v", token)
	// ----------------------------------------------

	payload := []byte(`{"username":"admin"}`)

	req, _ := http.NewRequest("POST", "/api/logout", bytes.NewBuffer(payload))
	req.Header.Set("Authorization", tokenString)
	resp := executeRequest(req)

	checkResponseCode(t, http.StatusOK, resp.Code)

	var m map[string]interface{}
	json.Unmarshal(resp.Body.Bytes(), &m)

	assert.Equal(t, m["message"], "Successfully logged out!")

}

func TestAPICreateUser(t *testing.T) {
	// doing login first to get token
	payloadLogin := []byte(`{"username":"admin", "password":"admin1234"}`)

	reqLogin, _ := http.NewRequest("POST", "/api/login", bytes.NewBuffer(payloadLogin))
	respLogin := executeRequest(reqLogin)

	checkResponseCode(t, http.StatusOK, respLogin.Code)

	var mLogin map[string]interface{}
	json.Unmarshal(respLogin.Body.Bytes(), &mLogin)

	var token = mLogin["accessToken"]
	tokenString := fmt.Sprintf("Bearer %v", token)
	// ----------------------------------------------

	payload := []byte(`{"userName":"wiro", "password":"pass345", "accountExpired":false, "accountLocked":false, "credentialsExpired":false, "enabled":true}`)

	req, _ := http.NewRequest("POST", "/api/user/", bytes.NewBuffer(payload))
	req.Header.Set("Authorization", tokenString)
	resp := executeRequest(req)

	checkResponseCode(t, http.StatusCreated, resp.Code)

	var m map[string]interface{}
	json.Unmarshal(resp.Body.Bytes(), &m)

	if m["userName"] != "wiro" {
		t.Errorf("Expected user name to be 'wiro'. Got '%v'", m["userName"])
	}

	// var userIDstr = mLogin["ID"]
	userIDstr := fmt.Sprintf("%v", m["id"])
	userID, _ = strconv.ParseInt(userIDstr, 10, 64)
}
func TestAPIGetByIDUser(t *testing.T) {

	// doing login first to get token
	payloadLogin := []byte(`{"username":"admin", "password":"admin1234"}`)

	reqLogin, _ := http.NewRequest("POST", "/api/login", bytes.NewBuffer(payloadLogin))
	respLogin := executeRequest(reqLogin)

	checkResponseCode(t, http.StatusOK, respLogin.Code)

	var mLogin map[string]interface{}
	json.Unmarshal(respLogin.Body.Bytes(), &mLogin)

	var token = mLogin["accessToken"]
	tokenString := fmt.Sprintf("Bearer %v", token)
	// ----------------------------------------------

	req, err := http.NewRequest("GET", "/api/user/"+strconv.FormatInt(userID, 10), nil)
	if err != nil {
		t.Errorf("Expected sdf code %v. Go", err)
	}
	req.Header.Set("Authorization", tokenString)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

}

func TestAPIGetByIDUserNotFound(t *testing.T) {

	// doing login first to get token
	payloadLogin := []byte(`{"username":"admin", "password":"admin1234"}`)

	reqLogin, _ := http.NewRequest("POST", "/api/login", bytes.NewBuffer(payloadLogin))
	respLogin := executeRequest(reqLogin)

	checkResponseCode(t, http.StatusOK, respLogin.Code)

	var mLogin map[string]interface{}
	json.Unmarshal(respLogin.Body.Bytes(), &mLogin)

	var token = mLogin["accessToken"]
	tokenString := fmt.Sprintf("Bearer %v", token)
	// ----------------------------------------------

	req, _ := http.NewRequest("GET", "/api/user/9999999999", nil)
	req.Header.Set("Authorization", tokenString)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusNotFound, response.Code)
}

func TestAPIGetAllUser(t *testing.T) {

	// doing login first to get token
	payloadLogin := []byte(`{"username":"admin", "password":"admin1234"}`)

	reqLogin, _ := http.NewRequest("POST", "/api/login", bytes.NewBuffer(payloadLogin))
	respLogin := executeRequest(reqLogin)

	checkResponseCode(t, http.StatusOK, respLogin.Code)

	var mLogin map[string]interface{}
	json.Unmarshal(respLogin.Body.Bytes(), &mLogin)

	var token = mLogin["accessToken"]
	tokenString := fmt.Sprintf("Bearer %v", token)
	// ----------------------------------------------

	req, _ := http.NewRequest("GET", "/api/user/", nil)
	req.Header.Set("Authorization", tokenString)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestAPIDeleteByID(t *testing.T) {

	// doing login first to get token
	payloadLogin := []byte(`{"username":"admin", "password":"admin1234"}`)

	reqLogin, _ := http.NewRequest("POST", "/api/login", bytes.NewBuffer(payloadLogin))
	respLogin := executeRequest(reqLogin)

	checkResponseCode(t, http.StatusOK, respLogin.Code)

	var mLogin map[string]interface{}
	json.Unmarshal(respLogin.Body.Bytes(), &mLogin)

	var token = mLogin["accessToken"]
	tokenString := fmt.Sprintf("Bearer %v", token)
	// ----------------------------------------------

	req, _ := http.NewRequest("DELETE", "/api/user/"+strconv.FormatInt(userID, 10), nil)
	req.Header.Set("Authorization", tokenString)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusNoContent, response.Code)
}

// ====================== Test repo ======================
func TestFindUserById(t *testing.T) {
	_, mock, err := NewMock()
	if err != nil {
		fmt.Printf("error mock: " + err.Error())
	}

	// simulate any sql driver behavior in tests, without needing a real database connection
	query := "select id, user_name, password from m_user where id = \\?"

	rows := sqlmock.NewRows([]string{"id", "user_name", "password"}).
		AddRow(user.ID, user.UserName, user.Password)

	mock.ExpectQuery(query).WithArgs(user.ID).WillReturnRows(rows)
	// ------------ end of mock ---------------

	assert.NotNil(t, user)
}

func TestFindUserByIdError(t *testing.T) {
	db, mock, err := NewMock()
	if err != nil {
		fmt.Printf("error mock: " + err.Error())
	}

	db, err = configuration.SetupDB()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// simulate any sql driver behavior in tests, without needing a real database connection
	query := "select id, user_name, password from m_user where id = \\?"

	rows := sqlmock.NewRows([]string{"id", "user_name", "password"}).
		AddRow(user.ID, user.UserName, user.Password)

	mock.ExpectQuery(query).WithArgs(user.ID).WillReturnRows(rows)
	// ------------ end of mock ---------------

	// Context like a timeout or deadline or a channel to indicate stop working and return
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res := new(model.MUser)
	err = db.QueryRowContext(ctx, "select id, user_name, password from m_user where id = ?", user.ID).Scan(&res.ID, &res.UserName, &res.Password)

	assert.Empty(t, res)
	assert.Error(t, err)
}

func TestFindAllUser(t *testing.T) {
	users := make([]*model.MUser, 0)

	db, mock, err := NewMock()
	if err != nil {
		fmt.Printf("error mock: " + err.Error())
	}

	db, err = configuration.SetupDB()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// simulate any sql driver behavior in tests, without needing a real database connection
	query := "select id, user_name, password from m_user where id = ?"
	rows := sqlmock.NewRows([]string{"id", "user_name", "password"}).
		AddRow(user.ID, user.UserName, user.Password)

	mock.ExpectQuery(query).WithArgs(user.ID).WillReturnRows(rows)
	// ------------ end of mock ---------------

	// Context like a timeout or deadline or a channel to indicate stop working and return
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := db.QueryContext(ctx, "select id, user_name, password from m_user")
	defer res.Close()

	for res.Next() {
		user := new(model.MUser)
		err = res.Scan(
			&user.ID,
			&user.UserName,
			&user.Password,
		)

		users = append(users, user)
	}

	assert.NotEmpty(t, users)
	assert.NoError(t, err)
	// assert.Len(t, users, 1)
}

func TestCreateUser(t *testing.T) {
	_, mock, err := NewMock()
	if err != nil {
		fmt.Printf("error mock: " + err.Error())
	}

	query := "insert into m_user \\(user_name, password\\) values \\(\\?, \\?\\)"

	hash, _ := util.HashPassword(user.Password, bcrypt.DefaultCost)
	user.Password = hash

	rows := mock.ExpectPrepare(query)
	rows.ExpectExec().WithArgs(user.UserName, user.Password).WillReturnResult(sqlmock.NewResult(0, 0))
}

func TestUpdateUser(t *testing.T) {
	_, mock, err := NewMock()
	if err != nil {
		fmt.Printf("error mock: " + err.Error())
	}

	query := "update m_user set user_name =\\?, password =\\? where id =\\?"

	hash, _ := util.HashPassword(user.Password, bcrypt.DefaultCost)
	user.Password = hash

	rows := mock.ExpectPrepare(query)
	rows.ExpectExec().WithArgs(user.UserName, user.Password, user.ID).WillReturnResult(sqlmock.NewResult(0, 1))
}

func TestDeleteUser(t *testing.T) {
	_, mock, err := NewMock()
	if err != nil {
		fmt.Printf("error mock: " + err.Error())
	}

	query := "delete from m_user where id =\\?"

	rows := mock.ExpectPrepare(query)
	rows.ExpectExec().WithArgs(user.ID).WillReturnResult(sqlmock.NewResult(0, 1))
}
