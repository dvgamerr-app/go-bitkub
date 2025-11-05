package market

import (
	"fmt"

	"github.com/dvgamerr-app/go-bitkub/bitkub"
)

// OrderHistoryItem represents an order history item in order info
type OrderHistoryItem struct {
	Amount    float64 `json:"amount"`
	Credit    float64 `json:"credit"`
	Fee       float64 `json:"fee"`
	ID        string  `json:"id"`
	Rate      float64 `json:"rate"`
	Timestamp int64   `json:"timestamp"`
	TxnID     string  `json:"txn_id"`
}

// OrderInfo represents detailed information about an order
type OrderInfo struct {
	ID            string             `json:"id"`             // Order id
	First         string             `json:"first"`          // First order id
	Parent        string             `json:"parent"`         // Parent order id
	Last          string             `json:"last"`           // Last order id
	ClientID      string             `json:"client_id"`      // Your id for reference
	PostOnly      bool               `json:"post_only"`      // Post only flag
	Amount        float64            `json:"amount"`         // Order amount
	Rate          float64            `json:"rate"`           // Order rate
	Fee           float64            `json:"fee"`            // Order fee
	Credit        float64            `json:"credit"`         // Order fee credit used
	Filled        float64            `json:"filled"`         // Filled amount
	Total         float64            `json:"total"`          // Total amount
	Status        string             `json:"status"`         // Order status: filled, unfilled, cancelled
	PartialFilled bool               `json:"partial_filled"` // Partially filled flag
	Remaining     float64            `json:"remaining"`      // Remaining amount
	History       []OrderHistoryItem `json:"history"`        // Order history
}

// OrderInfoResponse represents the response from /api/v3/market/order-info endpoint
type OrderInfoResponse struct {
	bitkub.Error
	Result OrderInfo `json:"result"`
}

// GetOrderInfo gets information regarding the specified order
// GET /api/v3/market/order-info
// symbol: The symbol (e.g. btc_thb)
// id: Order id
// side: Order side: buy or sell
func GetOrderInfo(symbol, id, side string) (*OrderInfo, error) {
	var response bitkub.ResponseAPI

	url := fmt.Sprintf("/v3/market/order-info?sym=%s&id=%s&sd=%s", symbol, id, side)

	if err := bitkub.FetchSecure("GET", url, nil, &response); err != nil {
		return nil, err
	}

	if err := response.CheckResponseError(); err != nil {
		return nil, err
	}

	byteData, err := stdJson.Marshal(response.Result)
	if err != nil {
		return nil, err
	}

	var result OrderInfo
	if err = stdJson.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
