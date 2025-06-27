package passwordusecase

import (
	"strconv"
	"strings"

	"github.com/maxexee/rugaPasswordManager/internal/dto"
	passwordsrepository "github.com/maxexee/rugaPasswordManager/internal/repository/passwords_repository"
)

// VERDE...
func PasswordGetByIdUseCase(userId int, passwordIdStr string) (bool, *dto.PasswordDto, error) {
	// CONVERSION DE STRING A INT PARA EL ID DE LA CONTRASEÑA.
	passwordId, passwordIdError := strconv.Atoi(passwordIdStr)

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== VALIDACIONES ==================================
	if passwordIdError != nil {
		return false, nil, passwordIdError
	}

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== QUERY =========================================
	// CREACION DEL DTO DE ENVIO CON EL ID DEL USUARIO Y DE LA CONTRASEÑA.
	dtoSend := dto.PasswordDto{
		UserID: uint(userId),
		ID:     uint(passwordId),
	}

	// LLAMADO AL REPOSITORIO.
	ok, passwordReturn, passwordReturnError := passwordsrepository.PasswordGetByIdUseCase(&dtoSend)
	if !ok {
		return false, nil, passwordReturnError
	}

	// SI TODO SALE BIEN...
	return true, passwordReturn, nil
}

// VERDE...
func PasswordGetByNameUseCase(userId int, passwordNameStr string) (bool, *dto.PasswordDto, error) {
	// ===========================================================================================
	// ===========================================================================================
	// =========================================== QUERY =========================================
	// CREACION DEL DTO DE ENVIO CON EL ID DEL USUARIO Y DE LA CONTRASEÑA.
	dtoSend := dto.PasswordDto{
		UserID: uint(userId),
		Name:   passwordNameStr,
	}

	// LLAMADA AL REPOSITORIO.
	ok, passwordReturn, passwordReturnError := passwordsrepository.PasswordGetByNameRepository(&dtoSend)
	if !ok {
		return false, nil, passwordReturnError
	}

	// SI TODO SALE BIEN...
	return true, passwordReturn, nil
}

// VERDE...
func PasswordPostUseCase(userId int, password *dto.PasswordDto) (bool, *dto.PasswordDto, error) {
	// ===========================================================================================
	// ===========================================================================================
	// =========================================== QUERY =========================================
	// CREACION DEL DTO DE ENVIO CON EL ID DEL USUARIO Y DE LA CONTRASEÑA.
	dtoSend := dto.PasswordDto{
		Name:                    strings.ToUpper(strings.ReplaceAll(password.Name, " ", "-")),
		Description:             password.Description,
		Password:                password.Password,
		UserID:                  uint(userId),
		SectionParentIdPassword: password.SectionParentIdPassword,
	}

	// LLAMADO AL REPOSITORIO.
	ok, passwordReturn, passwordReturnError := passwordsrepository.PasswordPostRepository(&dtoSend)
	if !ok {
		return false, nil, passwordReturnError
	}

	// SI TODO SALE BIEN...
	return true, passwordReturn, nil
}

// VERDE...
func PasswordUpdateUseCase(userId int, passwordIdStr string, password *dto.PasswordDto) (bool, *dto.PasswordDto, error) {
	// CONVERSION DEL *passwordIdStr* A INT.
	passwordId, passwordIdError := strconv.Atoi(passwordIdStr)

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== VALIDACIONES ==================================
	if passwordIdError != nil {
		return false, nil, passwordIdError
	}

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== QUERY =========================================
	// CONSTRUCCION DEL DTO DE ENVIO.
	dtoSend := dto.PasswordDto{
		ID:                      uint(passwordId),
		UserID:                  uint(userId),
		Name:                    password.Name,
		Description:             password.Description,
		Password:                password.Password,
		SectionParentIdPassword: password.SectionParentIdPassword,
	}

	// LLAMADA AL REPOSITORIO.
	ok, passwordReturn, passwordReturnError := passwordsrepository.PasswordUpdateRepositoy(&dtoSend)
	if !ok {
		return false, nil, passwordReturnError
	}

	// SI TODO SALE BIEN...
	return true, passwordReturn, nil
}

// VERDE...
func PasswordDeleteUseCase(userId int, passwordIdStr string) (bool, error) {
	// CONVERSION DEL *passwordIdStr* A INT.
	passwordId, passwordIdError := strconv.Atoi(passwordIdStr)

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== VALIDACIONES ==================================
	if passwordIdError != nil {
		return false, passwordIdError
	}

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== QUERY =========================================
	// CONTRUCCION DEL DTO DE ENVIO.
	// (-- VALIDAR SI ES NECESARIO O ES MEJOR SOLO MANDAR LOS VALORES INT DE ARRIBA -- )
	dtoSend := dto.PasswordDto{
		ID:     uint(passwordId),
		UserID: uint(userId),
	}

	// LLAMADO AL REPOSITORIO.
	ok, passwordDeleteError := passwordsrepository.PasswordDeleteRepository(&dtoSend)
	if !ok {
		return false, passwordDeleteError
	}

	// SI TODO SALE BIEN...
	return true, nil
}
