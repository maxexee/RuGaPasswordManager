package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	fmt.Println("=== I am middleware Requiering Authorization ===")

	authHeader := c.GetHeader("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer") {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Miss or invalid token...",
		})
		c.Abort()
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil || !token.Valid {
		fmt.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
		c.Abort()
		return
	}

	// VALIDAR EXPIRACION (SI SE USAN CLAIMS ESTANDAR).
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if exp, ok := claims["exp"].(float64); ok {
			if time.Unix(int64(exp), 0).Before(time.Now()) {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Token expirado"})
				c.Abort()
				return
			}
		}

		// CONVERTIR EL "sub" A INT (POSIBLEMENTE ERA floeat64).
		if sub, ok := claims["sub"].(float64); ok {
			c.Set("user_id", int(sub))
		} else {
			c.Set("user_id", nil)
		}
	}

	c.Next()
}
