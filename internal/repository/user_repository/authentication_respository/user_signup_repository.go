package authenticationrespository

import (
	postgres "github.com/maxexee/rugaPasswordManager/infrastructure/db"
	"github.com/maxexee/rugaPasswordManager/internal/domain"
)

func SignUpRepository(user domain.User) (bool, error) {

	userCreationResult := postgres.DB.Create(&user)
	if userCreationResult.Error != nil {
		return false, userCreationResult.Error
	}

	return true, nil
}
