package market

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
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
	if req.Resolution == "" {
		return nil, fmt.Errorf("resolution is required")
	}
	if req.From == 0 {
		req.To = time.Now().Unix()
		req.From = req.To - 86400
	}
	if req.To == 0 {
		req.To = time.Now().Unix()
	}

	url := fmt.Sprintf("https://api.bitkub.com/tradingview/history?symbol=%s&resolution=%s&from=%d&to=%d",
		req.Symbol, req.Resolution, req.From, req.To)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result HistoryResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
