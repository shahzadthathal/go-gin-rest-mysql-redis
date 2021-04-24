package service

import (
	"JWT_REST_Gin_MySQL/model"
	"JWT_REST_Gin_MySQL/repository"
	"JWT_REST_Gin_MySQL/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// RoutesPost ...
func RoutesPost(rg *gin.RouterGroup) {
	post := rg.Group("/post")

	post.GET("/:id", util.TokenAuthMiddleware(), getPostByID)
	post.GET("/", getPosts)
	post.POST("/", util.TokenAuthMiddleware(), createPost)
	post.PUT("/", util.TokenAuthMiddleware(), updatePost)
	post.DELETE("/:id", util.TokenAuthMiddleware(), deletePostByID)
}

// getPostByID godoc
// @Summary show Post by id
// @Description get string by ID
// @Tags Post
// @Accept  json
// @Produce  json
// @Param id path int true "Post ID"
// @Success 200 {object} model.MPost
// @Failure 400 {string} string
// @Failure 404 {object} model.MPost
// @Failure 500 {string} string
// @Security bearerAuth
// @Router /post/{id} [get]
func getPostByID(c *gin.Context) {
	var post model.MPost
	paramID := c.Param("id")
	varID, err := strconv.ParseInt(paramID, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	post, err = repository.GetPostByID(varID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if (model.MPost{}) == post {
		c.JSON(http.StatusNotFound, post)
	} else {
		c.JSON(http.StatusOK, post)
	}
}

// getPosts godoc
// @Summary show list post
// @Description get posts
// @Tags Post
// @Accept  json
// @Produce  json
// @Success 200 {array} model.MPost
// @Failure 400 {string} string
// @Failure 404 {object} model.MPost
// @Failure 500 {string} string
// @Router /post/ [get]
func getPosts(c *gin.Context) {

	var posts []model.MPost
	posts, err := repository.GetPostAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, posts)

}

// createPost godoc
// @Summary create post
// @Description add by json post
// @Tags Post
// @Accept  json
// @Produce  json
// @Param post body model.MPost true "Post ID"
// @Success 200 {object} model.MPost
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Security bearerAuth
// @Router /post/ [post]
func createPost(c *gin.Context) {

	// c.JSON(200, gin.H{
	// 	"token data": c.Request.Header["Userid"][0],
	// })
	// fmt.Println("userid: ", c.Request.Header["Userid"][0])
	// return
	userId := c.Request.Header["Userid"][0]
	if userId == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Somethnig went wrong"})
		return
	}
	userIdParsedInt, _ := strconv.ParseInt(userId, 10, 64)

	// userIdLogin, err := strconv.ParseInt(userId, 10, 64)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	// 	return
	// }

	var post model.MPost

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "invalid json", "error": err.Error()})
		return
	}

	//Pointer to the Post struct
	//ps := &post
	//fmt.Println(ps)

	//fmt.Println("x")
	//fmt.Println(x)
	//ps.UserId = x
	post.UserId = userIdParsedInt
	post, err := repository.CreatePost(post)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, post)
}

// updatePost godoc
// @Summary update post
// @Description update by json post
// @Tags Post
// @Accept  json
// @Produce  json
// @Param post body model.MPost true "Post ID"
// @Success 200 {object} model.MPost
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Security bearerAuth
// @Router /post/ [put]
func updatePost(c *gin.Context) {

	var post model.MPost

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "invalid json"})
		return
	}

	pst, err := repository.UpdatePost(post)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, pst)
}

// deletePostByID godoc
// @Summary delete a post by id
// @Description delete post by ID
// @Tags Post
// @Accept  json
// @Produce  json
// @Param id path int true "Post ID" Format(int64)
// @Success 200 {object} model.MPost
// @Failure 400 {string} string
// @Failure 404 {object} model.MPost
// @Failure 500 {string} string
// @Security bearerAuth
// @Router /post/{id} [delete]
func deletePostByID(c *gin.Context) {

	var post model.MPost

	paramID := c.Param("id")
	varID, err := strconv.ParseInt(paramID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err = repository.DeletePostByID(varID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, post)
}
