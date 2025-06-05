package authenticationrespository

import (
	postgres "github.com/maxexee/rugaPasswordManager/infrastructure/db"
	"github.com/maxexee/rugaPasswordManager/internal/domain"
	"github.com/maxexee/rugaPasswordManager/internal/dto"
)

func SignUpRepository(user dto.SignUpDTO) (bool, error) {

	// CREAR EL USUARIO.
	userCreation := domain.User{
		Email:     user.Email,
		Passsword: user.Password,
	}

	// GUARDADO DEL USUARIO EN LA BASE DE DATOS.
	userCreationResult := postgres.DB.Create(&userCreation)
	if userCreationResult.Error != nil {
		return false, userCreationResult.Error
	}

	// SI TODO SALE BIEN, RETORNA AL USE CASE "true" y "nil".
	return true, nil
}
