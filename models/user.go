// models.user.go

package models

import (
	"strings"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// For this demo, we are storing the user list in memory
// We also have some users predefined
// In a real application, this list will most likely be fetched
// from a database. Moreover in, production settings, we should
// store passwords securely by storing and hashing them instead of
// using them as we are doing in this demo
var UserList = []User{
	User{Username: "user1", Password: "pass1"},
	User{Username: "user2", Password: "pass2"},
	User{Username: "user3", Password: "pass3"},
}

// Check if the username and password combination is valid
func IsUserValid(username, password string) error {
	db := psqlDB()
	defer db.Close()

	var dbPassword string
	// Get the existing entry present in the database for the given username
	result := db.QueryRow("select password from userDetail where username=$1", username)
	err := result.Scan(&dbPassword)

	if err != nil {
		// If an entry with the username does not exist, send an "Unauthorized"(401) status
		fmt.Println("Error:")
		fmt.Println(err)
		return err
	}

	// Compare the stored hashed password, with the hashed version of the password that was received
	if err = bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(password)); err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
		return err
	}

	return nil
}

// Regsiter a new user with the given username and password
func RegisterNewUser(username, password string) (error) {
	if strings.TrimSpace(password) == "" {
		return errors.New("The password cannot be empty")
	} 
	// else if isUsernameAvailable(username){
	// 	return errors.New("The username is already available")
	// }

	db := psqlDB()
	defer db.Close()
	
	_, err := db.Query("SELECT * FROM userDetail")
	if err != nil {
		fmt.Println("userDetail table is not present")
		fmt.Println("Creating userDetail table in postgres db........")

		_, err = db.Exec("CREATE TABLE userDetail ( username varchar(50) PRIMARY KEY, password varchar(200) NOT NULL );")
		if err != nil {
			fmt.Println("Error while creating userDetail table")
			return err;
		}
	}
	
	// Salt and hash the password using the bcrypt algorithm
	// The second argument is the cost of hashing, which we arbitrarily set as 8 (this value can be more or less, depending on the computing power you wish to utilize)
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 8)
	
	// Next, insert the username, along with the hashed password into the database
	sqlStatement := `INSERT INTO userDetail (username, password)
	VALUES ($1, $2)`
	if _, err := db.Exec(sqlStatement, username, hashedPassword); err != nil {
		// If there is any issue with inserting into the database, return a 500 error
		return err
	}
	
	// UserList = append(UserList, u)
	return nil
}

// func isUsernameAvailable(username string) bool {
// 	db := psqlDB()
// 	defer db.Close()

// 	sqlStatement := `SELECT EXISTS (SELECT * FROM userDetail WHERE username=$1)`
// 	if result, err := db.Exec(sqlStatement, username); err == nil {
// 		fmt.Println(result)
// 		return true
// 	}
// 	// for _, u := range UserList {
// 	// 	if u.Username == username {
// 	// 		return true
// 	// 	}
// 	// }
// 	return false
// }

