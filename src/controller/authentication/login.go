package authentication

import (
	"github.com/gin-gonic/gin"
)

func LoginRoute(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "OK",
	})
}
