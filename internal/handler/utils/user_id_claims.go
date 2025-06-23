package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SaveUserIdFromClaims(c *gin.Context) (bool, int) {
	// OBTENCION DEL ID DEL USUARIO.
	// SE RECUPERA EL "user_id" DEL CONTEXTO, ES DECIR, DE LOS CLAIMS DEL TOKEN DEL LOGIN.
	userIdClaims, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"STATUS": "Unauthorized",
			"ERROR":  "There is no user_id on the token...",
		})
		return false, 0
	}

	// SE VERIFICA QUE SEA ENTERO/INT.
	// SI SI ES, ENTONCES SE ALMACENA EN "userId".
	userId, ok := userIdClaims.(int)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"STATUS": "Unauthorized",
			"ERROR":  "user_id in context is not an integer...",
		})
		return false, 0
	}

	return true, userId
}
