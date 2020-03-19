// routes.go

package routes

import (
	"go-gin-web-app/handlers"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine{

	var router *gin.Engine

	// Set the router as the default one provided by Gin
	router = gin.Default()

	return router
}

func InitializeRoutes(router *gin.Engine) {
	
	// Handle the index route
	router.GET("/", handlers.ShowIndexPage)

	// Handle GET requests at /article/view/some_article_id
	router.GET("/article/view/:article_id", handlers.GetArticle)
}
