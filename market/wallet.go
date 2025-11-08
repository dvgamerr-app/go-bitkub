package market

import (
	"github.com/dvgamerr-app/go-bitkub/bitkub"
)

type BitkubWallet map[string]float64

func GetWallet() (*BitkubWallet, error) {
	var result bitkub.ResponseAPI

	if err := bitkub.FetchSecure("POST", "/api/v3/market/wallet", nil, &result); err != nil {
		return nil, err
	}

	if err := result.CheckResponseError(); err != nil {
		return nil, err
	}

	return bitkub.DecodeResult[BitkubWallet](result.Result)
}
