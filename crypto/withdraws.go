package crypto

import (
	"fmt"
	"net/url"
)

type Withdraws struct {
	Pagination
	DateRange
	Symbol string
	Status string
}

func GetWithdraws(params Withdraws) (*WithdrawsResponse, error) {
	queryParams := url.Values{}
	addPagination(queryParams, params.Pagination)
	if params.Symbol != "" {
		queryParams.Set("symbol", params.Symbol)
	}
	if params.Status != "" {
		queryParams.Set("status", params.Status)
	}
	addDateRange(queryParams, params.DateRange)

	return fetchV4[WithdrawsResponse]("GET", pathWithQuery("/api/v4/crypto/withdraws", queryParams), nil)
}

func CreateWithdraw(req CreateWithdrawRequest) (*CreateWithdrawResponse, error) {
	if req.Symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}
	if req.Amount == "" {
		return nil, fmt.Errorf("amount is required")
	}
	if req.Address == "" {
		return nil, fmt.Errorf("address is required")
	}
	if req.Network == "" {
		return nil, fmt.Errorf("network is required")
	}

	return fetchV4[CreateWithdrawResponse]("POST", "/api/v4/crypto/withdraws", req)
}
