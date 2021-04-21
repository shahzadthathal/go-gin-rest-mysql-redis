package service

import (
	"JWT_REST_Gin_MySQL/model"
	"JWT_REST_Gin_MySQL/repository"
	"JWT_REST_Gin_MySQL/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// RoutesUserDetail ...
func RoutesUserDetail(rg *gin.RouterGroup) {
	user := rg.Group("/userDetail")

	user.GET("/:id", util.TokenAuthMiddleware(), getUserDetailByID)
	user.GET("/", getAllUserDetails)
}

func getUserDetailByID(c *gin.Context) {

	var userDtl model.MUserDetail
	paramID := c.Param("id")
	varID, err := strconv.ParseInt(paramID, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	userDtl, err = repository.GetUserDetailByID(varID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if (model.MUserDetail{}) == userDtl {
		c.JSON(http.StatusNotFound, userDtl)
	} else {
		c.JSON(http.StatusOK, userDtl)
	}
}

func getAllUserDetails(c *gin.Context) {

	var userDtls []model.MUserDetail
	userDtls, err := repository.GetAllUserDetail()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, userDtls)
}

func createUserDetail(c *gin.Context) {

	var userDtl model.MUserDetail

	if err := c.ShouldBindJSON(&userDtl); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "invalid json"})
		return
	}

	userDtl, err := repository.CreateUserDetail(userDtl)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, userDtl)
}

func updateUserDetail(c *gin.Context) {

	var userDtl model.MUserDetail

	if err := c.ShouldBindJSON(&userDtl); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "invalid json"})
		return
	}

	usrDtl, err := repository.UpdateUserDetail(userDtl)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, usrDtl)
}

func deleteUserDetailByID(c *gin.Context) {

	var userDtl model.MUserDetail

	paramID := c.Param("id")
	varID, err := strconv.ParseInt(paramID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err = repository.DeleteUserDetailByID(varID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, userDtl)
}
