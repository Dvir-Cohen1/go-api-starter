package main

import (
	"net/http"

	"goapi/internal/api/routes" // Import your custom route configuration

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a Gin router
	r := gin.Default()

	// Serve static files from the "static" folder
	r.Static("/static", "./static")

	// Serve script files from the "scripts" folder
	r.Static("/scripts", "./scripts")

	// Define a route that responds with "Hello, World!"
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})

	// Include your custom route configuration
	routes.SetupHelloRoutes(r) // Use your custom route configuration

	// Run the server on port 8080
	r.Run(":8080")
}
