package market

import (
	"fmt"

	"github.com/dvgamerr-app/go-bitkub/bitkub"
)

func GetTradingCredits() (float64, error) {
	var result bitkub.ResponseAPI

	if err := bitkub.FetchSecure("POST", "/api/v3/user/trading-credits", nil, &result); err != nil {
		return 0, err
	}

	if err := result.CheckResponseError(); err != nil {
		return 0, err
	}

	crd, ok := result.Result.(float64)
	if !ok {
		return 0, fmt.Errorf("can't parse Result %#v", result)
	}

	return crd, nil
}
