package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
  
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "im chaning into new one let see does it work or not",
		})
	})

	r.GET("/goodbye", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Goodbye, World!",
		})
	})
}