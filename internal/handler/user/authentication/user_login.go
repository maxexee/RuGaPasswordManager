package handler

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	postgres "github.com/maxexee/rugaPasswordManager/infrastructure/db"
	"github.com/maxexee/rugaPasswordManager/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	// RECIBE LA INFORMACION DE LA PETICION, ESPERANDO UN EMAIL Y UN PASSWORD. -- DUPLICADA --
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	// VERFICAMOS QUE EL EMAIL EXISTA...
	var user domain.User
	postgres.DB.First(&user, "email	=	?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password.",
		})
		return
	}

	// VERIFICAR QUE LA CONSTRASEÑA SEA LA MISMA.
	err := bcrypt.CompareHashAndPassword([]byte(user.Passsword), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password.",
		})
		return
	}

	// GENERACION DE UN JWT TOKEN.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Minute).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token.",
		})
		return
	}

	// REGRESAR EL TOKEN.
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}
