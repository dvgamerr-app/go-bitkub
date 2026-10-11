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

type DepthResponse struct {
	bitkub.Error
	Result DepthResult `json:"result"`
}

func GetDepth(symbol string, limit int) (*DepthResult, error) {
	var response DepthResponse

	url := fmt.Sprintf("/api/v3/market/depth?sym=%s", utils.UppercaseSymbol(symbol))
	if limit > 0 {
		url = fmt.Sprintf("%s&lmt=%d", url, limit)
	}

	if err := bitkub.FetchNonSecure("GET", url, nil, &response); err != nil {
		return nil, err
	}

	result := response.Result
	if len(result.Asks) == 0 && len(result.Bids) == 0 {
		return nil, fmt.Errorf("invalid symbol or no data available for %s", symbol)
	}

	return &result, nil
}
