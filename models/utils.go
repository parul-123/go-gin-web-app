package models

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/gomodule/redigo/redis"
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "postgres"
	dbname = "articles_go_db"
)

// Setup a connection to the database
func psqlDB() (*sql.DB) {
	// fmt.Println(port, host, user, password, dbname);
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

func InitRedisCache() (redis.Conn){
	// Initialize the redis connection to a redis instance in local
	conn, err := redis.DialURL("redis://localhost")
	if err != nil {
		panic(err.Error())
	}
	return conn
}
