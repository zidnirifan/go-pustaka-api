package utils

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RequestError struct {
	StatusCode int
	Err        error
}

func (r *RequestError) Error() string {
	return r.Err.Error()
}

func HandleError(err error, c *gin.Context) {
	log.Println(err)
	statusCode := http.StatusInternalServerError
	re, ok := err.(*RequestError)
	if ok {
		statusCode = re.StatusCode
		c.JSON(statusCode, gin.H{
			"message": re.Error(),
		})
		return
	}
	c.JSON(statusCode, gin.H{
		"message": "Internal server error",
	})
}
