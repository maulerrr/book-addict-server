package helpers

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func SetEnvironment() error {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("Couldn't specify working directory")
		return err
	}
	file, err := os.Open(wd + "\\server" + "\\.env")
	if err != nil {
		log.Fatal("Couldn't open .env file")
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "=")
		if len(parts) == 2 {
			os.Setenv(parts[0], parts[1])
			log.Printf("setting env var: %s = %s", parts[0], parts[1])
		}
	}

	return nil
}
