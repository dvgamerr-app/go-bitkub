package market

import (
	"fmt"

	"github.com/dvgamerr-app/go-bitkub/bitkub"
)

// Trade represents a recent trade [timestamp, price, size, side]
type Trade [4]interface{}

// TradesResponse represents the response from /api/v3/market/trades endpoint
type TradesResponse struct {
	Error  int     `json:"error"`
	Result []Trade `json:"result"`
}

// GetTrades lists recent trades
// GET /api/v3/market/trades
// sym: The symbol (e.g. btc_thb)
// lmt: No. of limit to query recent trades (optional)
func GetTrades(sym string, lmt int) ([]Trade, error) {
	var result TradesResponse

	url := fmt.Sprintf("/v3/market/trades?sym=%s", sym)
	if lmt > 0 {
		url = fmt.Sprintf("%s&lmt=%d", url, lmt)
	}

	if err := bitkub.FetchNonSecure("GET", url, nil, &result); err != nil {
		return nil, err
	}

	return result.Result, nil
}
