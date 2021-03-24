package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func loadFile() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func Get(variableName string) string {
	loadFile()

	return os.Getenv(variableName)
}
