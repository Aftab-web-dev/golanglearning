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
	userGroup.DELETE("/users/:id", handler.DeleteUserbyidHandler)
	userGroup.PUT("/users/:id", handler.UpdateUserHandler)

	//login routes
	userGroup.POST("/login", handler.LoginHandler)

	// User routes

}