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
	// BODY DE LA PETICION HTTP.
	var body struct {
		Email    string
		Password string
	}

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== VALIDACIONES ==================================
	// OBTENCION Y VALIDACION DEL BODY.
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"STATUS": "Failed to read body...",
			"ERROR":  c.Errors,
		})
		c.Abort()
		return
	}

	// VERFICAMOS QUE EL EMAIL EXISTA.
	var userExist domain.User
	userExistResult := postgres.DB.First(&userExist, "email	=	?", body.Email)
	if userExist.ID == 0 || userExistResult.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"STATUS": "Invalid email or password.",
			"ERROR":  userExistResult.Error.Error(),
		})
		c.Abort()
		return
	}

	// VERIFICAR QUE LA CONSTRASEÑA SEA LA MISMA.
	errPasswordVerification := bcrypt.CompareHashAndPassword([]byte(userExist.Passsword), []byte(body.Password))
	if errPasswordVerification != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"STATUS": "Invalid email or password.",
			"ERROR":  errPasswordVerification.Error(),
		})
		c.Abort()
		return
	}

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== TOKEN =========================================
	// DEFINICION DEL TOKEN.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userExist.ID,
		"exp": time.Now().Add(time.Minute).Unix(),
	})

	// GENERACION DE UN JWT TOKEN.
	tokenString, errToken := token.SignedString([]byte(os.Getenv("SECRET")))
	if errToken != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"STATUS": "Failed to create token.",
			"ERROR":  errToken.Error(),
		})
		c.Abort()
		return
	}

	// REGRESAR EL TOKEN.
	c.JSON(http.StatusOK, gin.H{"TOKEN": tokenString})
}
