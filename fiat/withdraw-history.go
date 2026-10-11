package fiat

import (
	"github.com/dvgamerr-app/go-bitkub/bitkub"
)

// WithdrawHistory represents a fiat withdrawal history item
type WithdrawHistory struct {
	TxnID    string  `json:"txn_id"`
	Currency string  `json:"currency"`
	Amount   string  `json:"amount"`
	Fee      float64 `json:"fee"`
	Status   string  `json:"status"`
	Time     int64   `json:"time"`
}

// WithdrawHistoryPagination represents pagination information
type WithdrawHistoryPagination struct {
	Page int `json:"page"`
	Last int `json:"last"`
}

// WithdrawHistoryResponse represents the response from /api/v3/fiat/withdraw-history endpoint
type WithdrawHistoryResponse struct {
	bitkub.Error
	Result     []WithdrawHistory         `json:"result"`
	Pagination WithdrawHistoryPagination `json:"pagination"`
}

// WithdrawHistoryParams represents parameters for withdraw history request
type WithdrawHistoryParams struct {
	Page  int // Page (optional)
	Limit int // Limit (optional)
}

func GetWithdrawHistory(params WithdrawHistoryParams) (*WithdrawHistoryResponse, error) {
	return fetchPage[WithdrawHistoryResponse]("/api/v3/fiat/withdraw-history", params.Page, params.Limit)
}
