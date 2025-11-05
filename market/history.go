package market

import (
	"fmt"
	"time"

	"github.com/dvgamerr-app/go-bitkub/bitkub"
	"github.com/dvgamerr-app/go-bitkub/utils"
)

type HistoryRequest struct {
	Symbol     string
	Resolution string
	From       int64
	To         int64
}

type HistoryResponse struct {
	Status string    `json:"s"`
	Close  []float64 `json:"c"`
	Open   []float64 `json:"o"`
	High   []float64 `json:"h"`
	Low    []float64 `json:"l"`
	Volume []float64 `json:"v"`
	Time   []int64   `json:"t"`
}

func GetHistory(req HistoryRequest) (*HistoryResponse, error) {
	if req.Symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}
	if req.From == 0 && req.To == 0 {
		req.Resolution = "60"
	}
	if req.From == 0 {
		req.To = time.Now().Unix()
		req.From = req.To - 86400
	}
	if req.To == 0 {
		req.To = time.Now().Unix()
	}

	url := fmt.Sprintf("/tradingview/history?symbol=%s&resolution=%s&from=%d&to=%d",
		utils.UppercaseSymbol(req.Symbol), req.Resolution, req.From, req.To)

	var result HistoryResponse
	if err := bitkub.FetchNonSecure("GET", url, nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
