package main

import (
	"log"

	"github.com/faruqii/EAI/internal/apps"
	"github.com/joho/godotenv"
)

func main() {
	// load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Call the Start function from the apps package
	apps.Start()
}
