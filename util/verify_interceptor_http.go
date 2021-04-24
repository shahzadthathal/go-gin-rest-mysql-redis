package util

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// TokenAuthMiddleware ...
func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.GetHeader("Authorization")) == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Authorization is required Header"})
			c.Abort()
			return
		}
		accessDetails, err := ExtractFromRedis(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Verify Token failure. Reason: " + err.Error()})
			c.Abort()
			return
		}

		s := strconv.FormatInt(accessDetails.UserID, 10)
		c.Request.Header.Set("userId", s)
		//c.Request.Header.Set("hellowWrold", s)
		//How to access this userId var from request in services
		//c *gin.Context
		//c.Request.Header["Userid"] //output will be userid id 1
		c.Next()
	}
}
