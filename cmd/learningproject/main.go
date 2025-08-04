package main

import (
	"github.com/Aftab-web-dev/learningproject/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.RegisterRoutes(r)
    
	// Custom 404 handler
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"error":   "Route not found",
			"message": "The route you are trying to access does not exist",
		})
	})
	
	r.Run(":8080")
}