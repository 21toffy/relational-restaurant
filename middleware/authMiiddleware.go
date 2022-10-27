package middleware

import (
	"net/http"

	"github.com/21toffy/relational-restaurant/helpers"
	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		var IPWhitelist = map[string]bool{
			"127.0.0.1": true,
		}

		ip := c.ClientIP()
		if !IPWhitelist[ip] {
			c.IndentedJSON(http.StatusForbidden, gin.H{
				"error": "You are not authorised to use this endpoint",
			})
			return
		}
		err := helpers.TokenValid(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}
