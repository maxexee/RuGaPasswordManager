package initializers

import (
	postgres "github.com/maxexee/rugaPasswordManager/infrastructure/db"
	"github.com/maxexee/rugaPasswordManager/internal/domain"
)

func DB_migratetion() {
	postgres.DB.AutoMigrate(&domain.User{}, &domain.Section{}, &domain.Password{})
}
