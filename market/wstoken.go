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

	// API returns token as a string directly in result
	if token, ok := response.Result.(string); ok {
		return token, nil
	}

	// Fallback: try to unmarshal as object
	byteData, err := stdJson.Marshal(response.Result)
	if err != nil {
		return "", err
	}

	var result WSTokenResult
	if err = stdJson.Unmarshal(byteData, &result); err != nil {
		return "", err
	}

	return result.Token, nil
}
