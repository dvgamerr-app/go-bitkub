package market

import (
	"encoding/json"
	"fmt"

	"github.com/dvgamerr-app/go-bitkub/bitkub"
	"github.com/dvgamerr-app/go-bitkub/utils"
)

type Ticker struct {
	Symbol        string  `json:"-"`
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

func GetTicker(symbol string) ([]Ticker, error) {
	var resultMap map[string]json.RawMessage

	url := "/api/market/ticker"
	if symbol != "" {
		url = fmt.Sprintf("%s?sym=%s", url, utils.NormalizeSymbol(symbol))
	}

	if err := bitkub.FetchNonSecure("GET", url, nil, &resultMap); err != nil {
		return nil, err
	}

	result := make([]Ticker, 0, len(resultMap))
	for sym, rawData := range resultMap {
		var ticker Ticker
		if err := json.Unmarshal(rawData, &ticker); err != nil {
			continue
		}
		ticker.Symbol = sym
		result = append(result, ticker)
	}

	return result, nil
}
