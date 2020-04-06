// models.user.go

package models

import (
	"strings"
	"errors"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"-"`
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
func IsUserValid(username, password string) bool {
	for _, user := range UserList {
		if user.Username == username && user.Password == password {
			return true
		}
	}
	return false
}

// Regsiter a new user with the given username and password
func RegisterNewUser(username, password string) (*User, error) {
	if strings.TrimSpace(password) == "" {
		return nil, errors.New("The password cannot be empty")
	} else if isUsernameAvailable(username){
		return nil, errors.New("The username is already available")
	}

	u := User{Username: username, Password: password}

	UserList = append(UserList, u)
	return &u, nil
}

func isUsernameAvailable(username string) bool {
	for _, u := range UserList {
		if u.Username == username {
			return true
		}
	}
	return false
}

