package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	postgres "github.com/maxexee/rugaPasswordManager/infrastructure/db"
	"github.com/maxexee/rugaPasswordManager/internal/domain"
	validations "github.com/maxexee/rugaPasswordManager/internal/handler/validations"
	"golang.org/x/crypto/bcrypt"
)

// REGISTRO -
func SignUp(c *gin.Context) {

	// BODY DE LA PETICION HTTP.
	var body struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=8"`
	}
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
	// HASHEA LA CONTRASEÑA RECIBIDA.
	passwordHash, errHashContraseña := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if !validations.ErrorValidations(c, errHashContraseña, "Failed to hash password...") {
		c.Abort()
		return
	}

	// CREAR EL USUARIO.
	userCreation := domain.User{
		Email:     strings.ToLower(body.Email),
		Passsword: string(passwordHash),
	}

	// GUARDADO DEL USUARIO EN LA BASE DE DATOS.
	userCreationResult := postgres.DB.Create(&userCreation)
	if !validations.DataBaseValidations(c, userCreationResult, userCreation.ID, "Failed to create user...") {
		c.Abort()
		return
	}

	// CREACION DE SECCIONES UNA VEZ EL USUARIO SEA CREADO (EN PROCESO).
	/*
		.
		.
		.
	*/

	// RETORNA UN 200
	c.JSON(http.StatusOK, gin.H{"USER-CREATED": "User created successfully..."})
}
