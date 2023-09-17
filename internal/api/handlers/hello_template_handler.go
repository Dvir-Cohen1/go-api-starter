package handlers

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HelloTemplateHandler handles requests to render the hello template.
func HelloTemplateHandler(c *gin.Context) {
	// Load the template
	tmpl, err := template.ParseFiles("templates/hello.html")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Define data to be passed to the template
	data := struct {
		Name string
	}{
		Name: "John", // Replace with the name you want to display
	}

	// Execute the template with the data
	err = tmpl.Execute(c.Writer, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
}
