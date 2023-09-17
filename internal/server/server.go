package server

import (
	"goapi/internal/api/routes"
	"goapi/internal/config"
	"strconv" // Import the strconv package

	"github.com/gin-gonic/gin"
)

// StartServer initializes and starts the HTTP server.
func StartServer() error {
	// Load application configuration
	config, err := config.LoadConfig()
	if err != nil {
		return err
	}

	// Create a Gin router
	router := gin.Default()

	// Setup API routes
	routes.SetupHelloRoutes(router) // Add more routes as needed

	// Start the server
	port := strconv.Itoa(config.Port) // Convert port to string
	if err := router.Run(":" + port); err != nil {
		return err
	}

	return nil
}
