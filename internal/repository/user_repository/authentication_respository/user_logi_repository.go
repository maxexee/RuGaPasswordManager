package authenticationrespository

import (
	postgres "github.com/maxexee/rugaPasswordManager/infrastructure/db"
	"github.com/maxexee/rugaPasswordManager/internal/domain"
	"github.com/maxexee/rugaPasswordManager/internal/dto"
	"golang.org/x/crypto/bcrypt"
)

func LogInRepository(user dto.LogInDTO) (bool, domain.User, error) {
	// CREACION DE UN OBJETO USER.
	var userExist domain.User

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== VALIDACIONES ==================================
	// VERFICAMOS QUE EL EMAIL EXISTA.
	userExistResult := postgres.DB.First(&userExist, "email	=	?", user.Email)
	if userExist.ID == 0 || userExistResult.Error != nil {
		return false, userExist, userExistResult.Error
	}

	// VERIFICAR QUE LA CONSTRASEÑA SEA LA MISMA.
	passwordVerificationError := bcrypt.CompareHashAndPassword([]byte(userExist.Passsword), []byte(user.Password))
	if passwordVerificationError != nil {
		return false, userExist, passwordVerificationError
	}

	return true, userExist, nil
}
