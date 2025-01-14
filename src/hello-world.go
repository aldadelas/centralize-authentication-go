package main

import "github.com/gin-gonic/gin"

func main() {
	route := gin.Default()
	route.GET("/hello-world", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world 🖐️",
		})
	})
	err := route.Run()

	if err != nil {
		route.Run(":2025")
	}
}
