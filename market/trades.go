package market

import (
	"fmt"

	"github.com/dvgamerr-app/go-bitkub/bitkub"
)

type Trade [4]any

type TradesResponse struct {
	bitkub.Error
	Result []Trade `json:"result"`
}

func GetTrades(symbol string, limit int) ([]Trade, error) {
	var result TradesResponse

	url := fmt.Sprintf("/api/v3/market/trades?sym=%s", symbol)
	if limit > 0 {
		url = fmt.Sprintf("%s&lmt=%d", url, limit)
	}

	if err := bitkub.FetchNonSecure("GET", url, nil, &result); err != nil {
		return nil, err
	}

	return result.Result, nil
}
