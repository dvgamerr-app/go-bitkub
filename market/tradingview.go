package market

import (
	"fmt"

	"github.com/dvgamerr-app/go-bitkub/bitkub"
)

// TradingViewHistory represents historical data for TradingView chart
type TradingViewHistory struct {
	C []float64 `json:"c"` // Close prices
	H []float64 `json:"h"` // High prices
	L []float64 `json:"l"` // Low prices
	O []float64 `json:"o"` // Open prices
	S string    `json:"s"` // Status
	T []int64   `json:"t"` // Timestamps
	V []float64 `json:"v"` // Volumes
}

// TradingViewHistoryParams represents parameters for historical data request
type TradingViewHistoryParams struct {
	Symbol     string // The symbol (e.g. BTC_THB)
	Resolution string // Chart resolution (1, 5, 15, 60, 240, 1D)
	From       int64  // Timestamp of the starting time
	To         int64  // Timestamp of the ending time
}

// GetTradingViewHistory gets historical data for TradingView chart
// GET /tradingview/history
func GetTradingViewHistory(params TradingViewHistoryParams) (*TradingViewHistory, error) {
	var result TradingViewHistory

	url := fmt.Sprintf("/tradingview/history?symbol=%s&resolution=%s&from=%d&to=%d",
		params.Symbol, params.Resolution, params.From, params.To)

	if err := bitkub.FetchNonSecure("GET", url, nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
