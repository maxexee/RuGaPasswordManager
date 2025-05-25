package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	postgres "github.com/maxexee/rugaPasswordManager/infrastructure/db"
	"github.com/maxexee/rugaPasswordManager/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

// REGISTRO -
func SignUp(c *gin.Context) {

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
			"STATUS": "Failed to read body",
			"ERROR":  c.Errors.Errors(),
		})
		c.Abort()
		return
	}

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== CONTRASEÑA ====================================
	// HASHEA LA CONTRASEÑA RECIBIDA.
	hash, errHashContraseña := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if errHashContraseña != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"STATUS": "Failed to hash password",
			"ERROR":  errHashContraseña.Error(),
		})
		c.Abort()
		return
	}

	// CREAR EL USUARIO.
	userCreation := domain.User{
		Email:     body.Email,
		Passsword: string(hash),
	}

	// LO GUARDA EN LA BASE DE DATOS.
	userCreationResult := postgres.DB.Create(&userCreation)
	// SI RETORNA UN ERROR EL GUARDAR EN LA BASE DE DATOS, NOS MUESTRA...
	if userCreationResult.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"STATUS": "Failed to create user.",
			"ERROR":  userCreationResult.Error.Error(),
		})
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
