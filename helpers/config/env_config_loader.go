package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func ExportENV() {
	env := os.Getenv("ENV")
	if env == "" {
		fmt.Println("ENV is not set. Defaulting to 'local'")
		env = "local"
	}
	// Construct the filename based on the ENV variable
	envFile := fmt.Sprintf("config/%s.env", env)

	// Load the environment variables from the file
	if err := godotenv.Load(envFile); err != nil {
		fmt.Printf("Error loading %s file\n", envFile)
		os.Exit(1)
	}
}
