package validationsauthentication

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BodyValidation(c *gin.Context, body any) bool {
	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"STATUS": "Failed to read body...",
			"ERROR":  c.Errors.Errors(),
		})
		return false
	}
	return true
}
