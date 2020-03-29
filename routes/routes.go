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

	// Handle POST requests at /article/create
	router.POST("/article/create", handlers.CreateArticle)
	
	// Handle the GET request at /article/create 
	// Show the Article Creation Page
	router.GET("/article/create", handlers.ArticleCreationPage)

	// Handle the GET request at /article/update/some_article_id
	// Show the Article Update Page
	router.GET("/article/update/:article_id", handlers.ArticleUpdatePage)

	// Handle the POST request at /article/update
	router.POST("/article/update/:article_id", handlers.UpdateArticle)

	// Handle the GET request at /article/delete/some_article_id
	// Delete the Article
	router.GET("/article/delete/:article_id", handlers.DeleteArticle)

}
