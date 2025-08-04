package handler

import (
	"net/http"

	"github.com/Aftab-web-dev/learningproject/internal/controller"
	"github.com/Aftab-web-dev/learningproject/internal/models"
	"github.com/gin-gonic/gin"
)

// POST /users
func CreateUserHandler(c *gin.Context) {
	var user models.User
	
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	insertedId, err := controller.CreateUserController(user)
	if err != nil {
		// ✅ Print the real error to your terminal
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create user",
			"details": err.Error(), // ✅ send actual error to Postman
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user_id": insertedId.Hex(),
	})
}
