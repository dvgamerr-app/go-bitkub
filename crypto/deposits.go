package crypto

import (
	"net/url"
	"strconv"

	"github.com/dvgamerr-app/go-bitkub/bitkub"
)

type GetDepositsParams struct {
	Page         int
	Limit        int
	Symbol       string
	Status       string
	CreatedStart string
	CreatedEnd   string
}

func GetDeposits(params GetDepositsParams) (*DepositsResponse, error) {
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
	if params.Status != "" {
		queryParams.Add("status", params.Status)
	}
	if params.CreatedStart != "" {
		queryParams.Add("created_start", params.CreatedStart)
	}
	if params.CreatedEnd != "" {
		queryParams.Add("created_end", params.CreatedEnd)
	}

	path := "/v4/crypto/deposits"
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

	data := DepositsResponse{}
	if err = stdJson.Unmarshal(byteData, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
