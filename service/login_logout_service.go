package service

import (
	"JWT_REST_Gin_MySQL/model"
	"JWT_REST_Gin_MySQL/repository"
	"JWT_REST_Gin_MySQL/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CredentialsLogin Create a struct to read the username and password from the request body
type CredentialsLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// RoutesLoginLogout ...
func RoutesLoginLogout(rg *gin.RouterGroup) {
	cred := rg.Group("/")

	cred.POST("login", getUserLogin)
	cred.GET("logout", getUserLogout)
}

// getUserLogin godoc
// @Summary Auth user
// @Description login user
// @Accept  json
// @Produce  json
// @Param user body CredentialsLogin true "Input username & password"
// @Success 200 {object} util.TokenDetails
// @Failure 400 {string} string
// @Failure 401 {string} string
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router /login [post]
func getUserLogin(c *gin.Context) {

	var creds CredentialsLogin
	var user model.MUser
	var err error

	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "invalid json"})
		return
	}

	user, err = repository.GetUserLogin(creds.Username, creds.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	jwt, err := util.CreateToken(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err = util.SaveToRedis(user.ID, jwt)
	if err != nil {
		ad := &util.AccessDetails{
			AccessUUID: jwt.AccessUUID,
			UserID:     user.ID,
		}
		util.DeleteToken(ad)
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, jwt)
}

// getUserLogout godoc
// @Summary Logout
// @Description logout
// @Tags Logout
// @Accept  json
// @Produce  json
// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Security bearerAuth
// @Router /logout [get]
func getUserLogout(c *gin.Context) {

	accessDetails, err := util.ExtractFromRedis(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Verify Token failure. Reason: " + err.Error()})
		return
	}

	err = util.DeleteToken(accessDetails)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Delete Token failure. Reason: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out!"})
}
