package middleware

import (
	"github.com/gin-gonic/gin"
)

// AuthMiddleware is an example middleware.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Add your authentication logic here if needed
		// For this example, we're not implementing actual authentication
		c.Next()
	}
}
