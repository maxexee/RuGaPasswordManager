package validationsauthentication

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorValidations(c *gin.Context, err error, message string) bool {
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"STATUS": message,
			"ERROR":  err.Error(),
		})
		return false
	}
	return true
}
