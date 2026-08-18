package crypto

import (
	"net/url"
)

type Compensations struct {
	Pagination
	DateRange
	Symbol string
	Type   string
	Status string
}

func GetCompensations(params Compensations) (*CompensationsResponse, error) {
	queryParams := url.Values{}
	addPagination(queryParams, params.Pagination)
	if params.Symbol != "" {
		queryParams.Set("symbol", params.Symbol)
	}
	if params.Type != "" {
		queryParams.Set("type", params.Type)
	}
	if params.Status != "" {
		queryParams.Set("status", params.Status)
	}
	addDateRange(queryParams, params.DateRange)

	return fetchV4[CompensationsResponse]("GET", pathWithQuery("/api/v4/crypto/compensations", queryParams), nil)
}
