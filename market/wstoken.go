package market

import (
	"github.com/dvgamerr-app/go-bitkub/bitkub"
)

type WSTokenResult struct {
	Token string `json:"token"`
}

func GetWSToken() (string, error) {
	var response bitkub.ResponseAPI

	if err := bitkub.FetchSecure("POST", "/api/v3/market/wstoken", nil, &response); err != nil {
		return "", err
	}

	if err := response.CheckResponseError(); err != nil {
		return "", err
	}

	if token, ok := response.Result.(string); ok {
		return token, nil
	}

	result, err := bitkub.DecodeResult[WSTokenResult](response.Result)
	if err != nil {
		return "", err
	}

	return result.Token, nil
}
