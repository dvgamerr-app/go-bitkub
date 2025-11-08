package user

import (
	"github.com/dvgamerr-app/go-bitkub/bitkub"
)

type TradingCreditsResponse struct {
	bitkub.Error
	Result float64 `json:"result"`
}

func GetTradingCredits() (float64, error) {
	var response bitkub.ResponseAPI

	if err := bitkub.FetchSecure("POST", "/api/v3/user/trading-credits", nil, &response); err != nil {
		return 0, err
	}

	if err := response.CheckResponseError(); err != nil {
		return 0, err
	}

	result, err := bitkub.DecodeResult[float64](response.Result)
	if err != nil {
		return 0, err
	}

	return *result, nil
}
