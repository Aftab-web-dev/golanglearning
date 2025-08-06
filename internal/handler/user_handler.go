package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Aftab-web-dev/learningproject/internal/controller"
	"github.com/Aftab-web-dev/learningproject/internal/middleware"
	"github.com/Aftab-web-dev/learningproject/internal/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func CreateUserHandler(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	insertedId, err := controller.CreateUserController(c.Request.Context(), user)
	if err != nil {
		fmt.Printf("Controller error: %v\n", err)
		if strings.Contains(err.Error(), "exists") || strings.Contains(err.Error(), "taken") {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Other DB errors → 500
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user_id": insertedId.Hex(),
	})
}

func GetallUsersHandler(c *gin.Context) {
	users, err := controller.GetallUsersController(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, users)
}

func GetUserbyidHandler(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	user, err := controller.GetUserbyidController(c.Request.Context(), id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)

}

func DeleteUserbyidHandler(c *gin.Context) {
	id := c.Param("id")
    
	err := controller.DeleteUserbyIdController(c.Request.Context(), id); 
	 if err != nil {
        if err.Error() == "user not found" || err.Error() == "invalid user ID format" {
            c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        }
        return
    }

	
    c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})

}

func LoginHandler(c *gin.Context) {
	var loginreq models.LoginUser

	// Parse JSON from request
	if err := c.ShouldBindJSON(&loginreq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	//call controller to verify login
	err := controller.LoginuserController(c, loginreq)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
    
	token , err := middleware.GenerateToken(loginreq.Username)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"failed to generate token"})
		return
	}
    
	//login  sucessfully
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token": token,
	})

}

func UpdateUserHandler(c *gin.Context) {
    id := c.Param("id")

    var userupdate models.UserUpdate
    if err := c.ShouldBindJSON(&userupdate); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
        return
    }

    err := controller.UpdatedetailsController(c.Request.Context(), id, userupdate)
    if err != nil {
        if err.Error() == "user not found" {
            c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        } else {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        }
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}
