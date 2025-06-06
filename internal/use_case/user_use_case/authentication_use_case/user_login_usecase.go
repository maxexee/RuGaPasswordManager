package authenticationusecase

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/maxexee/rugaPasswordManager/internal/dto"
	authenticationrespository "github.com/maxexee/rugaPasswordManager/internal/repository/user_repository/authentication_respository"
)

func LogInUseCase(user *dto.LogInDTO) (bool, string, error) {
	// ===========================================================================================
	// ===========================================================================================
	// =========================================== VALIDACIONES ==================================
	//
	ok, userExist, userError := authenticationrespository.LogInRepository(*user)
	if !ok {
		return false, "", userError
	}

	// ===========================================================================================
	// ===========================================================================================
	// =========================================== TOKEN =========================================
	// DEFINICION DEL TOKEN.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userExist.ID,
		"exp": time.Now().Add(time.Minute).Unix(),
	})

	//
	tokeString, tokenError := token.SignedString([]byte(os.Getenv("SECRET")))
	if tokenError != nil {

	}

	return true, tokeString, nil
}
