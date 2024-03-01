package market

import (
	"github.com/touno-io/go-bitkub/bitkub"
)

type BitkubWallet map[string]float64

func GetWallet() (*BitkubWallet, error) {
	var result bitkub.ResponseAPI

	if err := bitkub.FetchSecure("POST", "/v3/market/wallet", nil, &result); err != nil {
		return nil, err
	}

	byteData, err := stdJson.Marshal(result.Result)
	if err != nil {
		return nil, err
	}

	data := BitkubWallet{}

	if err = stdJson.Unmarshal(byteData, &data); err != nil {
		return nil, err
	}
	return &data, nil
}
