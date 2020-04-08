// handlers.user.go

package handlers

import (
	"math/rand"
	"net/http"
	"strconv"
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	"go-gin-web-app/utils"
	"go-gin-web-app/models"
)

func ShowLoginPage(c *gin.Context){
	// Call the render function with the name of the
	// template to render
	utils.Render(c, gin.H{
		"title": "Login",
	}, "login.html")
}

func PerformLogin(c *gin.Context){
	// Obtain the posted username and password values
	username := c.PostForm("username")
	password := c.PostForm("password")

	// var sameSiteCookie http.SameSite;

	// Check if the username and password combination is valid
	if err := models.IsUserValid(username, password); err == nil {
		// If the username and password is valid set the token in a cookie
		sessionToken := generateSessionToken()
		// Set the token in the cache, along with the user whom it represents
		// The token has an expiry time of 3600 seconds
		cache := models.InitRedisCache()
		fmt.Println("Adding sessionToken in redis cache")
		_, error := cache.Do("SETEX", sessionToken, "3600", username)
		if error != nil {
			// If there is an error in setting the cache, return an internal server error
			c.HTML(http.StatusInternalServerError, "login.html", gin.H{
				"ErrorTitle": "Login Failed",
				"ErrorMessage": "Internal Server Error",})
		}
		_, error = cache.Do("SETEX", username, "3600", sessionToken)
		if error != nil {
			// If there is an error in setting the cache, return an internal server error
			c.HTML(http.StatusInternalServerError, "login.html", gin.H{
				"ErrorTitle": "Login Failed",
				"ErrorMessage": "Internal Server Error",})
		}
		c.SetCookie("session_token", sessionToken, 3600, "", "", false, true)
		c.Set("is_logged_in", true)

		utils.Render(c, gin.H{
			"title": "Successful Login",
		}, "login-successful.html")
	} else {
		if err == sql.ErrNoRows {
			c.HTML(http.StatusUnauthorized, "login.html", gin.H{
				"ErrorTitle": "Login Failed",
				"ErrorMessage": "Invalid credentials provided",
			})
		} else {
			// If the username and password combination is invalid,
			// show the error message on the login page
			c.HTML(http.StatusInternalServerError, "login.html", gin.H{
				"ErrorTitle": "Login Failed",
				"ErrorMessage": "Invalid credentials provided",
			})
		}
	}
}

func generateSessionToken() string{
	// We are using a random 16 character string as the session token
	// This is not a secure way of generating session tokens
	// Do not use this in production
	return strconv.FormatInt(rand.Int63(), 16)
}

func Logout(c *gin.Context){

	// var sameSiteCookie http.SameSite;

	// Get the session value from redis cache
	cache := models.InitRedisCache()
	sessionToken, err := c.Cookie("session_token")

	// Delete the older session token from redis cache
	fmt.Println("Deleting older sessionToken in redis cache")
	username, err := cache.Do("GET", sessionToken)
	_, err = cache.Do("DEL", sessionToken)
	if err != nil {
		c.Redirect(http.StatusInternalServerError, "/")
	}
	_, err = cache.Do("DEL", username)
	if err != nil {
		c.Redirect(http.StatusInternalServerError, "/")
	}

	// Clear the cookie
	c.SetCookie("session_token", "", -1, "", "", false, true)
	// Redirect to the home page
	c.Redirect(http.StatusTemporaryRedirect, "/")
}

func ShowRegistrationPage(c *gin.Context){
	// Call the render function with the name of the
	// template to render
	utils.Render(c, gin.H{
		"title": "Register",
	}, "register.html")
}

func Register(c *gin.Context){
	// Obtain the Posted username and password values
	username := c.PostForm("username")
	password := c.PostForm("password")

	// var sameSiteCookie http.SameSite;

	if err := models.RegisterNewUser(username, password); err == nil {
		// If the user is created, set the token in a cookie and log the user in
		sessionToken := generateSessionToken()

		// Set the token in the cache, along with the user whom it represents
		// The token has an expiry time of 3600 seconds
		cache := models.InitRedisCache()
		fmt.Println("Adding sessionToken in redis cache")
		_, error := cache.Do("SETEX", sessionToken, "3600", username)
		if error != nil {
			// If there is an error in setting the cache, return an internal server error
			c.HTML(http.StatusInternalServerError, "register.html", gin.H{
				"ErrorTitle": "Registration Failed",
				"ErrorMessage": "Internal Server Error",})
		}
		_, error = cache.Do("SETEX", username, "3600", sessionToken)
		if error != nil {
			// If there is an error in setting the cache, return an internal server error
			c.HTML(http.StatusInternalServerError, "register.html", gin.H{
				"ErrorTitle": "Registration Failed",
				"ErrorMessage": "Internal Server Error",})
		}
		c.SetCookie("session_token", sessionToken, 3600, "", "", false, true)
		c.Set("is_logged_in", true)

		utils.Render(c, gin.H{
			"title": "Successful registration and login",
		}, "login-successful.html")
	} else {
		// If the username/password combination is invalid,
		// show the error message on the login page
		c.HTML(http.StatusBadRequest, "register.html", gin.H{
			"ErrorTitle":   "Registration Failed",
			"ErrorMessage": err.Error()})
	}
}
