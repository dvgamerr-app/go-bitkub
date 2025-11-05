package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func PrettyPrintJSON(data any) error {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(jsonData))
	return nil
}

func LoadDotEnv(filenames ...string) error {
	if len(filenames) == 0 {
		if _, err := os.Stat(".env"); err != nil {
			return err
		}
	} else {
		for _, f := range filenames {
			if _, err := os.Stat(f); err != nil {
				return err
			}
		}
	}

	if err := godotenv.Load(filenames...); err != nil {
		return err
	}

	return nil
}

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
