package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SecureRoutesMiddleware(e *gin.Engine) {
	fmt.Println("=== Etapa de Validacion ===")

	protected := e.Group(("/"))
	protected.Use(RequireAuth)

	protected.GET("/validate", func(c *gin.Context) {
		userID, _ := c.Get("user_id")
		c.JSON(http.StatusOK, gin.H{
			"message": "Valid token...",
			"user_id": userID,
		})
	})
}
