package handler

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	postgres "github.com/maxexee/rugaPasswordManager/infrastructure/db"
	"github.com/maxexee/rugaPasswordManager/internal/domain"
	validations "github.com/maxexee/rugaPasswordManager/internal/handler/validations"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	// BODY DE LA PETICION HTTP.
	var body struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=8"`
	}

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== VALIDACIONES ==================================
	// OBTENCION Y VALIDACION DEL BODY.
	if !validations.BodyValidation(c, &body) {
		c.Abort()
		return
	}

	// VERFICAMOS QUE EL EMAIL EXISTA.
	var userExist domain.User
	userExistResult := postgres.DB.First(&userExist, "email	=	?", body.Email)
	if !validations.DataBaseValidations(c, userExistResult, userExist.ID, "Invalid email or password...") {
		c.Abort()
		return
	}

	// VERIFICAR QUE LA CONSTRASEÑA SEA LA MISMA.
	errPasswordVerification := bcrypt.CompareHashAndPassword([]byte(userExist.Passsword), []byte(body.Password))
	if !validations.ErrorValidations(c, errPasswordVerification, "Invalid email or password...") {
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
	if !validations.ErrorValidations(c, errToken, "Failed to create token...") {
		c.Abort()
		return
	}

	// REGRESAR EL TOKEN.
	c.JSON(http.StatusOK, gin.H{"TOKEN-CREATED": tokenString})
}
