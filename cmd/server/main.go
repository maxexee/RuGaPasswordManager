package main

import (
	postgres "github.com/maxexee/rugaPasswordManager/infrastructure/db"
	"github.com/maxexee/rugaPasswordManager/internal/initializers"
	"github.com/maxexee/rugaPasswordManager/router"
)

func init() {
	initializers.EnvLoader()
	postgres.DbPostgresConnection()
	initializers.DB_migratetion()
}

func main() {
	r := router.InitRoutes()
	r.Run()
}
