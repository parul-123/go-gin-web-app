//middlewares.auth.go

package middlewares

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"go-gin-web-app/models"
	"github.com/gomodule/redigo/redis"
)

// This middleware ensures that a request will be aborted with an error
// if the user is not logged in
func EnsuredLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context){
		// If there is an error or if the token is empty
		// the user is not logged in
		loggedInInterface, _ := c.Get("is_logged_in")
		loggedIn := loggedInInterface.(bool)
		if !loggedIn {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

// This middleware ensures that a request will be aborted with an error
// if the user is already logged in
func EnsureNotLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context){
		// If there is not error or if the token is not empty
		// the user is already logged in
		loggedInInterface, _ := c.Get("is_logged_in")
		loggedIn := loggedInInterface.(bool)
		if loggedIn {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}


// This middleware sets whether the user is logged in or not.
func SetUserStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		if sessionToken, err := c.Cookie("session_token"); err == nil && isTokenValid(sessionToken){
			c.Set("is_logged_in", true)
		} else {
			c.Set("is_logged_in", false)
		}
	}
}

func isTokenValid(sessionToken string) (bool){
	cache := models.InitRedisCache()
	username, err := redis.String(cache.Do("GET", sessionToken))
	if err != nil {
		return false
	}
	redisSessionToken, err := redis.String(cache.Do("GET", username))
	if err != nil {
		return false
	}
	if redisSessionToken == sessionToken {
		return true
	}
	return false
}
