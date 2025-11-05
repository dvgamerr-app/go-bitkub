package crypto

import (
	"net/url"
	"strconv"

	"github.com/dvgamerr-app/go-bitkub/bitkub"
)

type Compensations struct {
	Pagination
	DateRange
	Symbol string
	Type   string
	Status string
}

func GetCompensations(params Compensations) (*CompensationsResponse, error) {
	var result bitkub.ResponseAPIV4

	queryParams := url.Values{}
	if params.Page > 0 {
		queryParams.Add("page", strconv.Itoa(params.Page))
	}
	if params.Limit > 0 {
		queryParams.Add("limit", strconv.Itoa(params.Limit))
	}
	if params.Symbol != "" {
		queryParams.Add("symbol", params.Symbol)
	}
	if params.Type != "" {
		queryParams.Add("type", params.Type)
	}
	if params.Status != "" {
		queryParams.Add("status", params.Status)
	}
	if params.CreatedStart != "" {
		queryParams.Add("created_start", params.CreatedStart)
	}
	if params.CreatedEnd != "" {
		queryParams.Add("created_end", params.CreatedEnd)
	}

	path := "/api/v4/crypto/compensations"
	if len(queryParams) > 0 {
		path += "?" + queryParams.Encode()
	}

	if err := bitkub.FetchSecureV4("GET", path, nil, &result); err != nil {
		return nil, err
	}

	byteData, err := stdJson.Marshal(result.Data)
	if err != nil {
		return nil, err
	}

	data := CompensationsResponse{}
	if err = stdJson.Unmarshal(byteData, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
