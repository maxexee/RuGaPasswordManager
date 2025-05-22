package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	postgres "github.com/maxexee/rugaPasswordManager/infrastructure/db"
	"github.com/maxexee/rugaPasswordManager/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

// REGISTRO -
func SignUp(c *gin.Context) {

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

	// HASHEA LA CONTRASEÑA RECIBIDA.
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	// CREAR EL USUARIO.
	user := domain.User{Email: body.Email, Passsword: string(hash)}

	// LO GUARDA EN LA BASE DE DATOS.
	result := postgres.DB.Create(&user)

	// SI RETORNA UN ERROR EL GUARDAR EN LA BASE DE DATOS, NOS MUESTRA...
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user.",
		})
		fmt.Println(result.Error)
		return
	}

	// CREACION DE LA RAIZ UNA VEZ EL USUARIO SEA CREADO.

	// RETORNA UN 200
	c.JSON(http.StatusOK, gin.H{"great-user": "User created successfully..."})

}
