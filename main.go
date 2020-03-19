// main.go

package main

import (
	"go-gin-web-app/routes"
)



func main() {

	// Set the router as the default one provided by Gin
	router := routes.GetRouter()

	//Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("templates/*")

	// Initialize the routes
	routes.InitializeRoutes(router)

	// Start serving the application
	router.Run()
}
