package market

import (
	"fmt"

	"github.com/dvgamerr-app/go-bitkub/bitkub"
)

// Ticker represents ticker information from V3 API
type Ticker struct {
	Symbol        string `json:"symbol"`
	BaseVolume    string `json:"base_volume"`
	High24hr      string `json:"high_24_hr"`
	HighestBid    string `json:"highest_bid"`
	Last          string `json:"last"`
	Low24hr       string `json:"low_24_hr"`
	LowestAsk     string `json:"lowest_ask"`
	PercentChange string `json:"percent_change"`
	QuoteVolume   string `json:"quote_volume"`
}

// TickerResponse represents the response from /api/v3/market/ticker endpoint
type TickerResponse struct {
	Error  int      `json:"error"`
	Result []Ticker `json:"result"`
}

// GetTicker gets ticker information (V3 API)
// GET /api/v3/market/ticker
// sym: The symbol (e.g. btc_thb) - optional, returns all if not specified
func GetTicker(sym string) ([]Ticker, error) {
	var result TickerResponse

	url := "/v3/market/ticker"
	if sym != "" {
		url = fmt.Sprintf("%s?sym=%s", url, sym)
	}

	if err := bitkub.FetchNonSecure("GET", url, nil, &result); err != nil {
		return nil, err
	}

	return result.Result, nil
}
