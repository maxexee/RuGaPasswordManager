package authenticationusecase

import (
	"github.com/maxexee/rugaPasswordManager/internal/dto"
	authenticationrespository "github.com/maxexee/rugaPasswordManager/internal/repository/user_repository/authentication_respository"
	"golang.org/x/crypto/bcrypt"
)

func SignUpUseCase(user *dto.SignUpDTO) (bool, error) {
	// ===========================================================================================
	// ===========================================================================================
	// =========================================== CONTRASEÑA ====================================
	// HASHEA LA CONTRASEÑA RECIBIDA.
	passwordHash, errHashContraseña := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if errHashContraseña != nil {
		return false, errHashContraseña
	}

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== USUARO ========================================
	// ACTUALIZACION DE LA CONTRASEÑA YA HASHEADA.
	user.Password = string(passwordHash)

	// CREACION DE SECCIONES UNA VEZ EL USUARIO SEA CREADO (EN PROCESO).
	/*
		.
		.
		.
	*/

	// SE LE MANDA EL EMAIL Y LA CONTRASEÑA YA HASHEADA AL REPOSITORIO.
	ok, userError := authenticationrespository.SignUpRepository(*user)
	if !ok {
		return false, userError
	}

	// SI TODO SALE BIEN, RETORNA AL HANDLER "true" y "nil".
	return true, nil
}
