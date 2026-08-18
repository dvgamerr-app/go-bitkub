package crypto

import (
	"net/url"
)

type Deposits struct {
	Pagination
	DateRange
	Symbol string
	Status string
}

func GetDeposits(params Deposits) (*DepositsResponse, error) {
	queryParams := url.Values{}
	addPagination(queryParams, params.Pagination)
	if params.Symbol != "" {
		queryParams.Set("symbol", params.Symbol)
	}
	if params.Status != "" {
		queryParams.Set("status", params.Status)
	}
	addDateRange(queryParams, params.DateRange)

	return fetchV4[DepositsResponse]("GET", pathWithQuery("/api/v4/crypto/deposits", queryParams), nil)
}
