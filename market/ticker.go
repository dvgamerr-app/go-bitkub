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

type TickerResponse struct {
	Error  int      `json:"error"`
	Result []Ticker `json:"result"`
}

func GetTicker(sym string) ([]Ticker, error) {
	var result TickerResponse

	url := "/v3/market/ticker"
	if sym != "" {
		url = fmt.Sprintf("%s?sym=%s", url, sym)
	}

	if err := bitkub.FetchNonSecure("GET", url, nil, &result); err != nil {
		return nil, err
	}

	if result.Error != 0 {
		errMsg, exists := bitkub.ErrorCode[result.Error]
		if !exists {
			errMsg = "Unknown error"
		}
		return nil, fmt.Errorf("[error %d] %s", result.Error, errMsg)
	}

	return result.Result, nil
}
