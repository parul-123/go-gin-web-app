// models.article_test.go

package test

import (
	"testing"
	"go-gin-web-app/models"
)

// Test the function that fetches all articles
func TestGetAllArticles(t *testing.T) {
	alist := models.GetAllArticles()

	//Check that the length of the list of articles returned is the
	//same as the length of the global variable holding the list.
	if len(alist) != len(models.ArticleListDB) {
		t.Fail()
	}

	//Check that each member is identical
	for i, v := range alist {
		if v.Content != models.ArticleListDB[i].Content ||
			v.ID != models.ArticleListDB[i].ID ||
			v.Title != models.ArticleListDB[i].Title {

			t.Fail()
			break
		}
	}
}

// Test the function that fetch an Article by its ID
func TestGetArticleByID(t *testing.T) {
	a, err := models.GetArticleByID(1)

	if err != nil || a.ID != 1 || a.Title != models.ArticleListDB[0].Title || a.Content != models.ArticleListDB[0].Content {
		t.Fail()
	}
}
