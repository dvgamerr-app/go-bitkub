package bitkub

import (
	"fmt"
	"os"

	util "github.com/dvgamerr-app/go-bitkub/Util"
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

// CheckResponseError checks if the response has an error and returns an error if it does
func (r *ResponseAPI) CheckResponseError() error {
	if r.Error != 0 {
		errMsg, exists := ErrorCode[r.Error]
		if !exists {
			errMsg = "Unknown error"
		}
		if r.Message != "" {
			return fmt.Errorf("[error %d] %s: %s", r.Error, errMsg, r.Message)
		}
		return fmt.Errorf("[error %d] %s", r.Error, errMsg)
	}
	return nil
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
		if err := util.CheckEnvVars("BTK_APIKEY", "BTK_SECRETKEY"); err != nil {
			return err
		}

		apiKey = os.Getenv("BTK_APIKEY")
		secretKey = os.Getenv("BTK_SECRETKEY")
	}
	return nil
}
