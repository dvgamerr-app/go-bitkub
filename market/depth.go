package market

import (
	"fmt"

	"github.com/dvgamerr-app/go-bitkub/bitkub"
	"github.com/dvgamerr-app/go-bitkub/utils"
)

type DepthEntry [2]float64

type DepthResult struct {
	Asks []DepthEntry `json:"asks"`
	Bids []DepthEntry `json:"bids"`
}

func GetDepth(symbol string, limit int) (*DepthResult, error) {
	var result DepthResult

	url := fmt.Sprintf("/api/market/depth?sym=%s", utils.UppercaseSymbol(symbol))
	if limit > 0 {
		url = fmt.Sprintf("%s&lmt=%d", url, limit)
	}

	if err := bitkub.FetchNonSecure("GET", url, nil, &result); err != nil {
		return nil, err
	}

	if len(result.Asks) == 0 && len(result.Bids) == 0 {
		return nil, fmt.Errorf("invalid symbol or no data available for %s", symbol)
	}

	return &result, nil
}
