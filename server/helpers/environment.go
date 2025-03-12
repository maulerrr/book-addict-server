package helpers

import (
	"log"
	"os"
	"path/filepath"

	godotenv "github.com/joho/godotenv"
)

func SetEnvironment() error {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("Couldn't specify working directory:", err)
		return err
	}

	envPath := filepath.Join(wd, "server", ".env")

	if _, err := os.Stat(envPath); os.IsNotExist(err) {
		log.Fatal("Couldn't find .env file at:", envPath)
		return err
	}

	if err := godotenv.Load(envPath); err != nil {
		log.Fatal("Couldn't load .env file:", err)
		return err
	}

	return nil
}
