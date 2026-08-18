package crypto

import (
	"net/url"
)

type Coins = SymbolNetwork

func GetCoins(params Coins) (*CoinsResponse, error) {
	queryParams := url.Values{}
	if params.Symbol != "" {
		queryParams.Set("symbol", params.Symbol)
	}
	if params.Network != "" {
		queryParams.Set("network", params.Network)
	}

	return fetchV4[CoinsResponse]("GET", pathWithQuery("/api/v4/crypto/coins", queryParams), nil)
}
