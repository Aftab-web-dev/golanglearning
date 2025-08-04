package controller

import (
	"github.com/Aftab-web-dev/learningproject/internal/models"
	"github.com/gin-gonic/gin"
)

type SignupController struct{
	// You can add any dependencies here, like a database client
}

func (c *SignupController) Signup(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	// Here you would typically save the user to the database
	// For now, we just return the user data
	ctx.JSON(201, gin.H{"message": "User created successfully", "user": user})
}