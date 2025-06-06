package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maxexee/rugaPasswordManager/internal/dto"
	validations "github.com/maxexee/rugaPasswordManager/internal/handler/validations"
	authenticationusecase "github.com/maxexee/rugaPasswordManager/internal/use_case/user_use_case/authentication_use_case"
)

func Login(c *gin.Context) {
	// DTO CON LOS CAMPOS DEL BODY.
	var body dto.LogInDTO

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== VALIDACIONES ==================================
	// OBTENCION Y VALIDACION DEL BODY.
	if !validations.BodyValidation(c, &body) {
		c.Abort()
		return
	}

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== LOGIN =========================================
	//
	ok, token, userLoginResultError := authenticationusecase.LogInUseCase(&body)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"STATUS": "Email o contraseña inválida...",
			"ERROR":  userLoginResultError,
		})
		c.Abort()
		return
	}
	fmt.Println(token)

	// REGRESAR EL TOKEN.
	c.JSON(http.StatusOK, gin.H{"TOKEN-CREATED": token})
}
