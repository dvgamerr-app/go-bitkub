package market

import (
	"github.com/dvgamerr-app/go-bitkub/bitkub"
)

type Balance struct {
	Available float64 `json:"available"`
	Reserved  float64 `json:"reserved"`
}

func GetBalances() (map[string]Balance, error) {
	var result bitkub.ResponseAPI

	if err := bitkub.FetchSecure("POST", "/api/v3/market/balances", nil, &result); err != nil {
		return nil, err
	}

	if err := result.CheckResponseError(); err != nil {
		return nil, err
	}

	data, err := bitkub.DecodeResult[map[string]Balance](result.Result)
	if err != nil {
		return nil, err
	}

	return *data, nil
}
