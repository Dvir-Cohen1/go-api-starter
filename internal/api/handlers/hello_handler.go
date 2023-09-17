package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HelloHandler handles requests to the "/" route.
func HelloHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello, World!")
}
