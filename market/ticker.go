package market

import (
	"fmt"

	"github.com/touno-io/go-bitkub/bitkub"
)

type Ticker struct {
	ID            int     `json:"id"`
	Last          float64 `json:"last"`
	LowestAsk     float64 `json:"lowestAsk"`
	HighestBid    float64 `json:"highestBid"`
	PercentChange float64 `json:"percentChange"`
	BaseVolume    float64 `json:"baseVolume"`
	QuoteVolume   float64 `json:"quoteVolume"`
	IsFrozen      int     `json:"isFrozen"`
	High24hr      float64 `json:"high24hr"`
	Low24hr       float64 `json:"low24hr"`
	Change        float64 `json:"change"`
	PrevClose     float64 `json:"prevClose"`
	PrevOpen      float64 `json:"prevOpen"`
}

func GetMarketTicker(symbol string) Ticker {
	var res map[string]Ticker
	url := fmt.Sprintf("/market/ticker?sym=%s", symbol)
	// sugar.Debugf("GET %s", url)
	if err := bitkub.FetchNonSecure("GET", url, nil, &res); err != nil {
		// sugar.Errorln(err)
	}

	// sugar.Debugf("Response: %#v\n", res[symbol])
	return res[symbol]
}
