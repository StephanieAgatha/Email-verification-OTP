package middleware

import "github.com/gin-gonic/gin"

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("token")
		if err != nil || token == "" {
			c.AbortWithStatusJSON(400, gin.H{"Error": "Unathorized"})
			return
		}
		c.Next()
	}
}
