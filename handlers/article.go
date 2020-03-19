// handlers.article.go

package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go-gin-web-app/models"
	"go-gin-web-app/utils"
)

func ShowIndexPage(c *gin.Context) {
	articles := models.GetAllArticles()

	// Call the render function with the name of the
	// template to render
	utils.Render(c, gin.H{
		"title":   "Home Page",
		"payload": articles}, "index.html")
}

func GetArticle(c *gin.Context) {
	// Check if the article ID is valid
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		// Check if the article exists
		if article, err := models.GetArticleByID(articleID); err == nil {

			// Call the render function with the name of the
			// template to render
			utils.Render(c, gin.H{
				"title":   article.Title,
				"payload": article}, "article.html")

		} else {
			// If the article is not found, abort with an error
			c.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		// If an invalid article ID is specified in the URL, abort with an error
		c.AbortWithStatus(http.StatusNotFound)
	}
}
