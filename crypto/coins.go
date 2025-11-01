package crypto

import (
	"net/url"

	"github.com/dvgamerr-app/go-bitkub/bitkub"
)

type Coins = SymbolNetwork

func GetCoins(params Coins) (*CoinsResponse, error) {
	var result bitkub.ResponseAPIV4

	queryParams := url.Values{}
	if params.Symbol != "" {
		queryParams.Add("symbol", params.Symbol)
	}
	if params.Network != "" {
		queryParams.Add("network", params.Network)
	}

	path := "/v4/crypto/coins"
	if len(queryParams) > 0 {
		path += "?" + queryParams.Encode()
	}

	if err := bitkub.FetchSecureV4("GET", path, nil, &result); err != nil {
		return nil, err
	}

	byteData, err := stdJson.Marshal(result.Data)
	if err != nil {
		return nil, err
	}

	data := CoinsResponse{}
	if err = stdJson.Unmarshal(byteData, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
