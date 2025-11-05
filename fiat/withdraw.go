package fiat

import (
	"github.com/dvgamerr-app/go-bitkub/bitkub"
)

// WithdrawRequest represents the request body for fiat withdrawal
type WithdrawRequest struct {
	ID     string  `json:"id"`  // Bank account id
	Amount float64 `json:"amt"` // Amount you want to withdraw
}

// WithdrawResult represents the result from fiat withdrawal
type WithdrawResult struct {
	Transaction string  `json:"txn"` // Local transaction id
	Account     string  `json:"acc"` // Bank account id
	Currency    string  `json:"cur"` // Currency
	Amount      float64 `json:"amt"` // Withdraw amount
	Fee         float64 `json:"fee"` // Withdraw fee
	Receive     float64 `json:"rec"` // Amount to receive
	Timestamp   int64   `json:"ts"`  // Timestamp
}

// WithdrawResponse represents the response from /api/v3/fiat/withdraw endpoint
type WithdrawResponse struct {
	bitkub.Error
	Result WithdrawResult `json:"result"`
}

// Withdraw makes a withdrawal to an approved bank account
// POST /api/v3/fiat/withdraw
func Withdraw(req WithdrawRequest) (*WithdrawResult, error) {
	var response bitkub.ResponseAPI

	if err := bitkub.FetchSecure("POST", "/v3/fiat/withdraw", req, &response); err != nil {
		return nil, err
	}

	if err := response.CheckResponseError(); err != nil {
		return nil, err
	}

	byteData, err := stdJson.Marshal(response.Result)
	if err != nil {
		return nil, err
	}

	var result WithdrawResult
	if err = stdJson.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
