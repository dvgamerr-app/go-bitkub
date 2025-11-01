package bitkub

import (
	"os"

	"github.com/dvgamerr-app/go-bitkub/helper"
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

type ResponseAPIV4 struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func Initlizer(key ...string) error {
	if len(key) >= 2 {
		apiKey = key[0]
		secretKey = key[1]
	}

	if apiKey == "" || secretKey == "" {
		if err := helper.CheckEnvVars("BTK_APIKEY", "BTK_SECRETKEY"); err != nil {
			return err
		}

		apiKey = os.Getenv("BTK_APIKEY")
		secretKey = os.Getenv("BTK_SECRETKEY")
	}
	return nil
}
