package market

import (
	"fmt"

	"github.com/dvgamerr-app/go-bitkub/bitkub"
)

type DepthEntry [2]float64

type DepthResult struct {
	Asks []DepthEntry `json:"asks"`
	Bids []DepthEntry `json:"bids"`
}

type DepthResponse struct {
	bitkub.GetError
	Result DepthResult `json:"result"`
}

func GetDepth(symbol string, limit int) (*DepthResult, error) {
	var result DepthResponse

	url := fmt.Sprintf("/v3/market/depth?sym=%s", symbol)
	if limit > 0 {
		url = fmt.Sprintf("%s&lmt=%d", url, limit)
	}

	if err := bitkub.FetchNonSecure("GET", url, nil, &result); err != nil {
		return nil, err
	}

	return &result.Result, nil
}
