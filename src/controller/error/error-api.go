package error

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseError struct {
	Reason string `json:"reason"`
}

func BadRequest(c *gin.Context) {
	var user interface{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, ResponseError{
			Reason: "Bad request",
		})
	}
}
