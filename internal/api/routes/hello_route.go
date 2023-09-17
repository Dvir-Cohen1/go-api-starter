package routes

import (
	"goapi/internal/api/handlers"
	"goapi/internal/api/middleware"

	"github.com/gin-gonic/gin"
)

// SetupHelloRoutes configures routes related to the "Hello, World!" example.
func SetupHelloRoutes(router *gin.Engine) {
	// Group routes under a common middleware, e.g., authentication
	helloGroup := router.Group("/hello")
	helloGroup.Use(middleware.AuthMiddleware()) // Apply middleware to this group if needed

	// Define the "/hello" route and specify the corresponding handler
	helloGroup.GET("/", handlers.HelloHandler)

	// Define the "/hello/template" route and specify the corresponding handler
	helloGroup.GET("/template", handlers.HelloTemplateHandler)
}
