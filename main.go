// main.go

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {

	// Set the router as the default one provided by Gin
	router = gin.Default()

	//Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("templates/*")

	//Define the route for the index page and display the index.html template
	//To start with, we will use an inline route handler. Later on, we will create
	//standalone functions that will be used as route handlers.
	// router.GET("/", func(c *gin.Context) {

	// 	// Call the HTML method of the Context to render a template
	// 	c.HTML(
	// 		// Set the HTTP status to 200 (OK)
	// 		http.StatusOK,
	// 		// Use the index.html template
	// 		"index.html",
	// 		// Pass the data that the page uses (in this case, 'titile')
	// 		gin.H{
	// 			"title": "Home Page",
	// 		},
	// 	)
	// })

	// Initialize the routes
	initializeRoutes()

	// Start serving the application
	router.Run()
}

// Render one of the HTML, JSON or CSV based on the 'Accept' header of the request
// If the header doesn't specify this, HTML is rendered, provided that
// the template name is present
func render(c *gin.Context, data gin.H, templateName string) {

	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with JSON
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// Respond with XML
		c.XML(http.StatusOK, data["payload"])
	default:
		c.HTML(http.StatusOK, templateName, data)
	}
}
