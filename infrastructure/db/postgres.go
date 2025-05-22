package postgres

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// VARIABLE GLOBAL DE LA BASE DA DATOS PARA QUE OTROS ARCHIVOS PUEDAN ENTRAR A LA CONEXCIÓN.
var DB *gorm.DB

func DbPostgresConnection() {
	var err error
	dsn := os.Getenv("POSTGRES_DB_CONNECTION_STRING")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("=== Failed to connect to database...")
	}

	fmt.Println("=== DB Connected ===")
}
