package authenticationusecase

import (
	"github.com/maxexee/rugaPasswordManager/internal/domain"
	"github.com/maxexee/rugaPasswordManager/internal/dto"
	authenticationrespository "github.com/maxexee/rugaPasswordManager/internal/repository/user_repository/authentication_respository"
	"golang.org/x/crypto/bcrypt"
)

func SignUpUseCase(body dto.SignUpDTO) (bool, error) {
	// ===========================================================================================
	// ===========================================================================================
	// =========================================== CONTRASEÑA ====================================
	// HASHEA LA CONTRASEÑA RECIBIDA.
	passwordHash, errHashContraseña := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if errHashContraseña != nil {
		return false, errHashContraseña
	}

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== USUARO ========================================
	// CREAR EL USUARIO.
	userCreation := domain.User{
		Email:     body.Email,
		Passsword: string(passwordHash),
	}

	ok, userCreationError := authenticationrespository.SignUpRepository(userCreation)
	if !ok {
		return false, userCreationError
	}

	// CREACION DE SECCIONES UNA VEZ EL USUARIO SEA CREADO (EN PROCESO).
	/*
		.
		.
		.
	*/

	return true, nil
}
