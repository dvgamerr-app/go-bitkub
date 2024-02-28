package market

import (
	"fmt"
	"strings"

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

func GetMarketTicker(symbol string) (*Ticker, error) {
	var res map[string]Ticker

	sym := fmt.Sprintf("THB_%s", strings.ToUpper(symbol))
	url := fmt.Sprintf("/market/ticker?sym=%s", sym)
	// sugar.Debugf("GET %s", url)
	if err := bitkub.FetchNonSecure("GET", url, nil, &res); err != nil {
		// sugar.Errorln(err)
		return nil, err
	}
	data := res[sym]
	return &data, nil
}
