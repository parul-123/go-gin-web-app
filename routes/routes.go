// routes.go

package routes

import (
	"go-gin-web-app/handlers"
	"go-gin-web-app/middlewares"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine{

	var router *gin.Engine

	// Set the router as the default one provided by Gin
	router = gin.Default()

	return router
}

func InitializeRoutes(router *gin.Engine) {
	
	// Use the setUserStatus middleware for every route to set a flag
	// indicating whether the request was from an authenticated user or not
	router.Use(middlewares.SetUserStatus())

	// Handle the index route
	router.GET("/", handlers.ShowIndexPage)

	// Group user related routes together
	userRoutes := router.Group("/user")
	{
		// Handle the GET requests at /user/login 
		// Show the login page
		// Ensure that the user is not logged in by using the middleware
		userRoutes.GET("/login", middlewares.EnsureNotLoggedIn(), handlers.ShowLoginPage)

		// Handle the POST requests at /user/login
		// Ensure that the user is not logged in by using the middleware
		userRoutes.POST("/login", middlewares.EnsureNotLoggedIn(), handlers.PerformLogin)

		// Handle the GET requests at /user/logout
		// Ensure that the user is logged in by using the middleware
		userRoutes.GET("/logout", middlewares.EnsuredLoggedIn(), handlers.Logout)

		// Handle the GET requests at /user/register
		// Show the registration page
		// Ensure that the user is not logged in by using the middleware
		userRoutes.GET("/register", middlewares.EnsureNotLoggedIn(), handlers.ShowRegistrationPage)

		// Handle the POST requests at /user/register
		// Ensure that the user isnot logged in by using the middleware
		userRoutes.POST("/register", middlewares.EnsureNotLoggedIn(), handlers.Register)
	}

	// Group article related routes together
	articleRoutes := router.Group("/article")
	{
		// Handle GET requests at /article/view/some_article_id
		articleRoutes.GET("/view/:article_id", handlers.GetArticle)

		// Handle POST requests at /article/create
		articleRoutes.POST("/create", middlewares.EnsuredLoggedIn(), handlers.CreateArticle)
	
		// Handle the GET request at /article/create 
		// Show the Article Creation Page
		// Ensure that the user is logged in by using the middleware
		articleRoutes.GET("/create", middlewares.EnsuredLoggedIn(), handlers.ArticleCreationPage)

		// Handle the GET request at /article/update/some_article_id
		// Show the Article Update Page
		// Ensure that the user is logged in by using the middleware
		articleRoutes.GET("/update/:article_id", middlewares.EnsuredLoggedIn(), handlers.ArticleUpdatePage)

		// Handle the POST request at /article/update
		// Ensure that the user is logged in by using the middleware
		articleRoutes.POST("/update/:article_id", middlewares.EnsuredLoggedIn(), handlers.UpdateArticle)

		// Handle the GET request at /article/delete/some_article_id
		// Delete the Article
		// Ensure that the user is logged in by using the middleware
		articleRoutes.GET("/delete/:article_id", middlewares.EnsuredLoggedIn(), handlers.DeleteArticle)
	}

}
