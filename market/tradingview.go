package market

import (
	"fmt"

	"github.com/dvgamerr-app/go-bitkub/bitkub"
)

type TradingViewHistory struct {
	C []float64 `json:"c"`
	H []float64 `json:"h"`
	L []float64 `json:"l"`
	O []float64 `json:"o"`
	S string    `json:"s"`
	T []int64   `json:"t"`
	V []float64 `json:"v"`
}

type TradingViewHistoryParams struct {
	Symbol     string
	Resolution string
	From       int64
	To         int64
}

func GetTradingViewHistory(params TradingViewHistoryParams) (*TradingViewHistory, error) {
	var result TradingViewHistory

	url := fmt.Sprintf("/tradingview/history?symbol=%s&resolution=%s&from=%d&to=%d",
		params.Symbol, params.Resolution, params.From, params.To)

	if err := bitkub.FetchNonSecure("GET", url, nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
