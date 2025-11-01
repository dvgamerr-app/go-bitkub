package market

import (
	"github.com/dvgamerr-app/go-bitkub/bitkub"
)

// WSTokenResult represents the websocket token result
type WSTokenResult struct {
	Token string `json:"token"`
}

// GetWSToken gets websocket token for authenticated websocket connections
// POST /api/v3/market/wstoken
func GetWSToken() (string, error) {
	var response bitkub.ResponseAPI

	if err := bitkub.FetchSecure("POST", "/v3/market/wstoken", nil, &response); err != nil {
		return "", err
	}

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
