package initializers

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvLoader() {
	curDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	loaderr := godotenv.Load(curDir + "/.env")
	if loaderr != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println("=== TODO BIEN CON .env ===")
}
