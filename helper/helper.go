package helper

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// PrettyPrintJSON takes an interface containing the data and prints it as formatted JSON
func PrettyPrintJSON(data interface{}) error {
	// Marshal the data with indentation
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	// Print the formatted JSON
	fmt.Println(string(jsonData))
	return nil
}

// Load environment variables from .env
func LoadDotEnv() error {
	if _, err := os.Stat(".env"); err == nil {
		if err = godotenv.Load(); err != nil {
			return err
		}
	}
	return nil
}

// checkEnvVars checks that all specified environment variables are set and not empty.
func CheckEnvVars(envs ...string) error {
	var missingEnvVars []string
	for _, env := range envs {
		if value := os.Getenv(env); value == "" {
			missingEnvVars = append(missingEnvVars, env)
		}
	}

	if len(missingEnvVars) > 0 {
		return fmt.Errorf("missing required environment variables: %s", strings.Join(missingEnvVars, ", "))
	}
	return nil
}
