package configuration

import (
	"github.com/gin-gonic/gin"
)

// CORS configuration
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Set Headers
		// c.Writer.Header().Set("Access-Control-Allow-Headers:", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
