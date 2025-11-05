package fiat

import (
	"fmt"

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

// GetWithdrawHistory lists fiat withdrawal history
// POST /api/v3/fiat/withdraw-history
func GetWithdrawHistory(params WithdrawHistoryParams) (*WithdrawHistoryResponse, error) {
	var response bitkub.ResponseAPI

	url := "/api/v3/fiat/withdraw-history?"

	if params.Page > 0 {
		url = fmt.Sprintf("%sp=%d&", url, params.Page)
	}
	if params.Limit > 0 {
		url = fmt.Sprintf("%slmt=%d&", url, params.Limit)
	}

	// Remove trailing '&' or '?'
	if url[len(url)-1] == '&' || url[len(url)-1] == '?' {
		url = url[:len(url)-1]
	}

	if err := bitkub.FetchSecure("POST", url, nil, &response); err != nil {
		return nil, err
	}

	if err := response.CheckResponseError(); err != nil {
		return nil, err
	}

	byteData, err := stdJson.Marshal(response)
	if err != nil {
		return nil, err
	}

	var result WithdrawHistoryResponse
	if err = stdJson.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
