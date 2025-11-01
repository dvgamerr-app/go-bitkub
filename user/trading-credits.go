package user

import (
	"github.com/dvgamerr-app/go-bitkub/bitkub"
)

// TradingCreditsResponse represents the response from /api/v3/user/trading-credits endpoint
type TradingCreditsResponse struct {
	Error  int     `json:"error"`
	Result float64 `json:"result"`
}

// GetTradingCredits checks trading credit balance
// POST /api/v3/user/trading-credits
func GetTradingCredits() (float64, error) {
	var response bitkub.ResponseAPI

	if err := bitkub.FetchSecure("POST", "/v3/user/trading-credits", nil, &response); err != nil {
		return 0, err
	}

	byteData, err := stdJson.Marshal(response.Result)
	if err != nil {
		return 0, err
	}

	var result float64
	if err = stdJson.Unmarshal(byteData, &result); err != nil {
		return 0, err
	}

	return result, nil
}
