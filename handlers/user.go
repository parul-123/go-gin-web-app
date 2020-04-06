// handlers.user.go

package handlers

import (
	"math/rand"
	"net/http"
	"strconv"

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
	if models.IsUserValid(username, password){
		// If the username and password is valid set the token in a cookie
		token := generateSessionToken()
		c.SetCookie("token", token, 3600, "", "", false, true)
		c.Set("is_logged_in", true)

		utils.Render(c, gin.H{
			"title": "Successful Login",
		}, "login-successful.html")
	} else {
		// If the username and password combination is invalid,
		// show the error message on the login page
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"ErrorTitle": "Login Failed",
			"ErrorMessage": "Invalid credentials provided",
		})
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

	// Clear the cookie
	c.SetCookie("token", "", -1, "", "", false, true)

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

	if _, err := models.RegisterNewUser(username, password); err == nil {
		// If the user is created, set the token in a cookie and log the user in
		token := generateSessionToken()
		c.SetCookie("token", token, 3600, "", "", false, true)
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
