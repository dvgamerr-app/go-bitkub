package market

import (
	"fmt"

	"github.com/dvgamerr-app/go-bitkub/bitkub"
	"github.com/rs/zerolog/log"
)

type Ticker struct {
	Symbol        string  `json:"symbol"`
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
	var resultMap map[string]any

	url := "/api/market/ticker"
	if symbol != "" {
		url = fmt.Sprintf("%s?sym=%s", url, symbol)
	}

	if err := bitkub.FetchNonSecure("GET", url, nil, &resultMap); err != nil {
		return nil, err
	}

	log.Debug().Interface("resultMap", resultMap).Msg("Ticker response")

	result := make([]Ticker, 0, len(resultMap))
	for sym, value := range resultMap {
		tickerData, ok := value.(map[string]any)
		if !ok {
			log.Debug().Str("symbol", sym).Interface("value", value).Msg("Skipping non-map value")
			continue
		}

		ticker := Ticker{Symbol: sym}
		if id, ok := tickerData["id"].(float64); ok {
			ticker.ID = int(id)
		}
		if last, ok := tickerData["last"].(float64); ok {
			ticker.Last = last
		}
		if lowestAsk, ok := tickerData["lowestAsk"].(float64); ok {
			ticker.LowestAsk = lowestAsk
		}
		if highestBid, ok := tickerData["highestBid"].(float64); ok {
			ticker.HighestBid = highestBid
		}
		if percentChange, ok := tickerData["percentChange"].(float64); ok {
			ticker.PercentChange = percentChange
		}
		if baseVolume, ok := tickerData["baseVolume"].(float64); ok {
			ticker.BaseVolume = baseVolume
		}
		if quoteVolume, ok := tickerData["quoteVolume"].(float64); ok {
			ticker.QuoteVolume = quoteVolume
		}
		if isFrozen, ok := tickerData["isFrozen"].(float64); ok {
			ticker.IsFrozen = int(isFrozen)
		}
		if high24hr, ok := tickerData["high24hr"].(float64); ok {
			ticker.High24hr = high24hr
		}
		if low24hr, ok := tickerData["low24hr"].(float64); ok {
			ticker.Low24hr = low24hr
		}
		if change, ok := tickerData["change"].(float64); ok {
			ticker.Change = change
		}
		if prevClose, ok := tickerData["prevClose"].(float64); ok {
			ticker.PrevClose = prevClose
		}
		if prevOpen, ok := tickerData["prevOpen"].(float64); ok {
			ticker.PrevOpen = prevOpen
		}

		result = append(result, ticker)
	}

	return result, nil
}
