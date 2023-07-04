// Package controllers contains every REST API route logic
package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/twelvee/boxie/internal/boxie"
	"github.com/twelvee/boxie/pkg/boxie/structs"
	"net/http"
	"os"
)

// Register method will create a new user and generate invite link
func Register(c *gin.Context) {
	var input structs.CreateUserRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shelf := boxie.GetShelf(os.Getenv("BOXIE_SHELF_DRIVER"), os.Getenv("SHELF_DSN"))
	code, err := shelf.CreateUser(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"invite_code": code})
}

// Login method will return user token by email and password combination
func Login(c *gin.Context) {
	var input structs.LoginRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shelf := boxie.GetShelf(os.Getenv("BOXIE_SHELF_DRIVER"), os.Getenv("SHELF_DSN"))
	user, err := shelf.CreateToken(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

// DeleteUser will delete user by ID
func DeleteUser(c *gin.Context) {
	var input structs.DeleteUserRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shelf := boxie.GetShelf(os.Getenv("BOXIE_SHELF_DRIVER"), os.Getenv("SHELF_DSN"))
	err := shelf.DeleteUser(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

// GetUser will return a serialized user struct as json
func GetUser(c *gin.Context) {
	token := c.GetHeader("x-auth-token")
	shelf := boxie.GetShelf(os.Getenv("BOXIE_SHELF_DRIVER"), os.Getenv("SHELF_DSN"))
	user, err := shelf.GetUser(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

// GetUsers will return a all users
func GetUsers(c *gin.Context) {
	shelf := boxie.GetShelf(os.Getenv("BOXIE_SHELF_DRIVER"), os.Getenv("SHELF_DSN"))
	users, err := shelf.GetUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

// AcceptInvite will return a serialized user struct as json by invite code
func AcceptInvite(c *gin.Context) {
	var input structs.AcceptInviteRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	shelf := boxie.GetShelf(os.Getenv("BOXIE_SHELF_DRIVER"), os.Getenv("SHELF_DSN"))
	user, err := shelf.AcceptInvite(input.Code)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

// SetUserPassword will return a serialized user struct as json
func SetUserPassword(c *gin.Context) {
	var input structs.SetPasswordRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if input.Password != input.PasswordConfirmation {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("Password missmatch.")})
		return
	}

	shelf := boxie.GetShelf(os.Getenv("BOXIE_SHELF_DRIVER"), os.Getenv("SHELF_DSN"))
	user, err := shelf.SetUserPassword(input.Code, input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
