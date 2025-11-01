package market

import (
	"fmt"

	"github.com/dvgamerr-app/go-bitkub/bitkub"
)

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

func GetTicker(symbol string) ([]Ticker, error) {
	var result []Ticker

	url := "/v3/market/ticker"
	if symbol != "" {
		url = fmt.Sprintf("%s?sym=%s", url, symbol)
	}

	if err := bitkub.FetchNonSecure("GET", url, nil, &result); err != nil {
		return nil, err
	}

	return result, nil
}
