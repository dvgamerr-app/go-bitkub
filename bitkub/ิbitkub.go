package bitkub

import (
	"os"

	"github.com/touno-io/go-bitkub/helper"
)

const BASE_URL = "https://api.bitkub.com/api"

var (
	apiKey    string
	secretKey string
)

type ResponseAPI struct {
	Error   int    `json:"error"`
	Message string `json:"message"`
	Result  any    `json:"result"`
}

func Initlizer(key string, secret string) error {
	apiKey = key
	secretKey = secret

	if apiKey == "" || secretKey == "" {
		if err := helper.CheckEnvVars("BTK_APIKEY", "BTK_SECRETKEY"); err != nil {
			return err
		}

		apiKey = os.Getenv("BTK_APIKEY")
		secretKey = os.Getenv("BTK_SECRETKEY")
	}
	return nil
}
