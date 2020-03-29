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

func ArticleCreationPage(c *gin.Context){
	// Call the render function with the name of the
	// template to render
	utils.Render(c, gin.H{
		"title": "Create New Article"}, "form-article-create.html")
}

func CreateArticle(c *gin.Context){
	// Obtain the Posted title and content values
	title := c.PostForm("title")
	content := c.PostForm("content")

	if article, err := models.CreateNewArticle(title, content); err == nil {
		// If the article is successfully created, show success message
		utils.Render(c, gin.H{
			"title": "Article " + title + " Created Successfully",
			"payload": article}, "article-created.html")
	} else {
		// If an article is not created successfully, abort with an error
		c.AbortWithStatus(http.StatusBadRequest)
	}
	
}

func ArticleUpdatePage(c *gin.Context){
	// Check if the article ID is valid
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		// Check if the article exists
		if article, err := models.GetArticleByID(articleID); err == nil {
			// Call the render function with the name of the
			// template to render
			utils.Render(c, gin.H{
				"title": "Update Existing Article",
				"payload": article, }, "form-article-update.html")
		} else {
				// If the article is not found, abort with an error
				c.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		// If an invalid article ID is specified in the URL, abort with an error
		c.AbortWithStatus(http.StatusNotFound)
	}	
}

func UpdateArticle(c *gin.Context){
	// Obtain the Posted title and content values
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		title := c.PostForm("title")
		content := c.PostForm("content")

		if article, err := models.UpdateExistingArticle(articleID, title, content); err == nil {
			// Call the render function with the name of the
			// template to render
			utils.Render(c, gin.H{
				"title": "Article " + title + " Updated Successfully",
				"payload": article, }, "article-updated.html")
		} else {
			// If an article is not updated successfully, abort with an error
			c.AbortWithStatus(http.StatusBadRequest)
		}
	} else {
		// If an invalid article ID is specified in the URL, abort with an error
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func DeleteArticle(c *gin.Context){
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		// Check if the article exists
		if article, err := models.GetArticleByID(articleID); err == nil {

			if err := models.DeleteExistingArticle(articleID); err == nil {
				// Call the render function with the name of the
				// template to render
				utils.Render(c, gin.H{
					"title": "Article " + article.Title + " Deleted Successfully",
					"payload": article, }, "article-deleted.html")
			} else {
				c.AbortWithError(http.StatusNotFound, err)
			}
		} else {
			// If the article is not found, abort with an error
			c.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		// If an invalid article ID is specified in the URL, abort with an error
		c.AbortWithStatus(http.StatusNotFound)
	}
}
