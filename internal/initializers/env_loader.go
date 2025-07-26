package initializers

import (
	"fmt"
	"log"
	"os"
	// "github.com/joho/godotenv"
)

func EnvLoader() {
	// === LOCAL ===
	// curDir, err := os.Getwd()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// loaderr := godotenv.Load(curDir + "/.env")
	// if loaderr != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	// === LOCAL ===
	// =============

	// === CON DOCKER ===
	DB_CONNECTION_STRING := os.Getenv("DB_CONNECTION_STRING")
	SECRET := os.Getenv("SECRET")
	if DB_CONNECTION_STRING == "" || SECRET == "" {
		log.Fatal("Error loading .env file")
	}
	// === DEV ===
	fmt.Println("=== TODO BIEN CON .env ===")
	// === CON DOCKER ===
}
