package main

import (
	"centralize-authentication-go/src/controller/authentication"
	"centralize-authentication-go/src/middleware"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	/**
	route:
		- /auth -> this route is for login
		- /auth/refresh -> this route is for refresh token
		- /auth/create -> this roucte is for sign up
	*/
	router := gin.New()
	router.Use(middleware.ResponseBuilder())
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v\n", httpMethod, absolutePath)
	}
	router.GET("/hello-world", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world 🖐️",
		})
	})

	auth := router.Group("/auth")
	{
		auth.POST("", authentication.LoginRoute)
	}

	err := router.Run()

	if err != nil {
		router.Run(":2025")
	}
}
