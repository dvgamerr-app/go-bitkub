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

	if err := bitkub.FetchNonSecure("GET", "/api/market/ticker", nil, &resultMap); err != nil {
		return nil, err
	}

	wanted := ""
	if symbol != "" {
		wanted = utils.NormalizeSymbol(symbol)
	}

	result := make([]Ticker, 0, len(resultMap))
	for sym, rawData := range resultMap {
		if wanted != "" && sym != wanted {
			continue
		}
		var ticker Ticker
		if err := json.Unmarshal(rawData, &ticker); err != nil {
			continue
		}
		ticker.Symbol = sym
		result = append(result, ticker)
	}

	if wanted != "" && len(result) == 0 {
		return nil, fmt.Errorf("invalid symbol or no data available for %s", symbol)
	}

	return result, nil
}
