package validationsauthentication

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DataBaseValidations(c *gin.Context, databaseValidation *gorm.DB, id uint, message string) bool {
	if id == 0 || databaseValidation.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"STATUS": message,
			"ERROR":  databaseValidation.Error.Error(),
		})
		return false
	}
	return true
}
