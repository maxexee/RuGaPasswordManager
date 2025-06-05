package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maxexee/rugaPasswordManager/internal/dto"
	validations "github.com/maxexee/rugaPasswordManager/internal/handler/validations"
	authenticationusecase "github.com/maxexee/rugaPasswordManager/internal/use_case/user_use_case/authentication_use_case"
)

// REGISTRO -
func SignUp(c *gin.Context) {

	// DTO CON LOS CAMPOS DEL BODY.
	var body dto.SignUpDTO

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
	// =========================================== CONTRASEÑA ====================================
	// LLAMADA AL USER CASE DE SIGN UP, DONDE LE MANDAMOS EL BODY.
	ok, userCreationResult := authenticationusecase.SignUpUseCase(&body)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"STATUS": "Error al crear el Usuario...",
			"ERROR":  userCreationResult.Error(),
		})
		c.Abort()
		return
	}

	// SI TODO SALE BIEN, RETORNA UN 200
	c.JSON(http.StatusOK, gin.H{"USER-CREATED": "User created successfully..."})
}
