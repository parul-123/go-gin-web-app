// models.article.go

package models

import (
	"errors"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "postgres"
	dbname = "articles_go_db"
)

type Article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// For this demo, we are storing the article list in memory
// In a real application, this list will most likely be fetched
// from a database or from static files
var ArticleList = []Article{
	Article{ID: 1, Title: "Article 1", Content: "Article 1 body"},
	Article{ID: 2, Title: "Article 2", Content: "Article 2 body"},
}

var ArticleListDB []Article

// Return a list of articles
func GetAllArticles() []Article {
	db := psqlDB();
	ArticleListDB = []Article{}
	defer db.Close()

	_, err := db.Query("SELECT * FROM article")
	if err != nil {
		fmt.Println("article table is not present")
		fmt.Println("Creating article table in postgres db........")

		 _, err = db.Exec("CREATE TABLE article ( id serial PRIMARY KEY, title varchar(50) NOT NULL, content varchar(200) NOT NULL );")
		 if err != nil {
			 fmt.Println("Error while creating article table")
			 fmt.Println(err.Error())
			 return ArticleListDB
		 }

		 fmt.Println("Adding sample data in article table .......")

		 _, err = db.Exec("INSERT INTO article (title, content) VALUES" +
		  "('Article 1', 'Article 1 Content'), ('Article 2', 'Article 2 Content'), ('Article 3', 'Article 3 Content')")
		 if err != nil {
			 fmt.Println("Error while adding sample data in article table")
			 panic(err.Error())
		 }
	}

	rows, err := db.Query("SELECT * FROM article")
	defer rows.Close()

	for rows.Next() {
		var id int
		var title string
		var content string
		err = rows.Scan(&id, &title, &content)
		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Println("Zero rows found");
			} else {
				panic(err)
			}
		}
		articleItem := Article{
			ID: id,
			Title: title,
			Content: content,
		}

		ArticleListDB = append(ArticleListDB, articleItem)
	}

	return ArticleListDB
}

// Get Article Content by Id
func GetArticleByID(id int) (*Article, error) {
	for _, a := range ArticleListDB {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("Article Not Found")
}

// Setup a connection to the database
func psqlDB() (*sql.DB) {
	fmt.Println(port, host, user, password, dbname);
	psqlInfo := fmt.Sprintf("port=%d host=%s user=%s " +
		"password=%s dbname=%s sslmode=disable",
		port, host, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err.Error())
	}

	// err = db.Ping()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Successfully Connected!")
	return db
}