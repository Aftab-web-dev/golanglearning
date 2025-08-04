package routes

import (
	"github.com/Aftab-web-dev/learningproject/internal/handler"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	// User routes
    r.POST("/users", handler.CreateUserHandler )

}