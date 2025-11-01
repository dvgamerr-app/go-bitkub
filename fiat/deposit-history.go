package fiat

import (
	"fmt"

	"github.com/dvgamerr-app/go-bitkub/bitkub"
)

// DepositHistory represents a fiat deposit history item
type DepositHistory struct {
	TxnID    string  `json:"txn_id"`
	Currency string  `json:"currency"`
	Amount   float64 `json:"amount"`
	Status   string  `json:"status"`
	Time     int64   `json:"time"`
}

// DepositHistoryPagination represents pagination information
type DepositHistoryPagination struct {
	Page int `json:"page"`
	Last int `json:"last"`
}

// DepositHistoryResponse represents the response from /api/v3/fiat/deposit-history endpoint
type DepositHistoryResponse struct {
	Error      int                      `json:"error"`
	Result     []DepositHistory         `json:"result"`
	Pagination DepositHistoryPagination `json:"pagination"`
}

// DepositHistoryParams represents parameters for deposit history request
type DepositHistoryParams struct {
	P   int // Page (optional)
	Lmt int // Limit (optional)
}

// GetDepositHistory lists fiat deposit history
// POST /api/v3/fiat/deposit-history
func GetDepositHistory(params DepositHistoryParams) (*DepositHistoryResponse, error) {
	var response bitkub.ResponseAPI

	url := "/v3/fiat/deposit-history?"

	if params.P > 0 {
		url = fmt.Sprintf("%sp=%d&", url, params.P)
	}
	if params.Lmt > 0 {
		url = fmt.Sprintf("%slmt=%d&", url, params.Lmt)
	}

	// Remove trailing '&' or '?'
	if url[len(url)-1] == '&' || url[len(url)-1] == '?' {
		url = url[:len(url)-1]
	}

	if err := bitkub.FetchSecure("POST", url, nil, &response); err != nil {
		return nil, err
	}

	byteData, err := stdJson.Marshal(response)
	if err != nil {
		return nil, err
	}

	var result DepositHistoryResponse
	if err = stdJson.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
