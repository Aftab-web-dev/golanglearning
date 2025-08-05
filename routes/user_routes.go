package routes

import (
	"github.com/Aftab-web-dev/learningproject/internal/handler"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {

	// Grouping user routes under /auth
	userGroup := r.Group("/auth")
	userGroup.POST("/users", handler.CreateUserHandler )
	userGroup.GET("/users/:id", handler.GetUserbyidHandler)
	userGroup.GET("/allusers", handler.GetallUsersHandler)

	// User routes

}