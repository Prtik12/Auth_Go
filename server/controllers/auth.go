package controllers

import (
	"net/http"

	"github.com/Prtik12/auth_Golang/models"
	"github.com/Prtik12/auth_Golang/utils"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var req models.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	req.Password = hashedPassword

	if err := models.CreateUser(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

func SignIn(c *gin.Context) {
	var req models.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := models.GetUserByEmail(req.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := utils.GenerateJWT(*user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.SetCookie("token", token, 3600, "/", "localhost", false, true)

	session := sessions.Default(c)
	session.Set("user", user.Email)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}

func SignOut(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Options(sessions.Options{
		MaxAge: -1,
	})
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to clear session"})
		return
	}
	c.SetCookie("token", "", -1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}
