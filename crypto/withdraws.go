package crypto

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/dvgamerr-app/go-bitkub/bitkub"
)

type Withdraws struct {
	Pagination
	DateRange
	Symbol string
	Status string
}

func GetWithdraws(params Withdraws) (*WithdrawsResponse, error) {
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

	path := "/api/v4/crypto/withdraws"
	if len(queryParams) > 0 {
		path += "?" + queryParams.Encode()
	}

	if err := bitkub.FetchSecureV4("GET", path, nil, &result); err != nil {
		return nil, err
	}

	return bitkub.DecodeResult[WithdrawsResponse](result.Data)
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

	var result bitkub.ResponseAPIV4

	if err := bitkub.FetchSecureV4("POST", "/api/v4/crypto/withdraws", req, &result); err != nil {
		return nil, err
	}

	return bitkub.DecodeResult[CreateWithdrawResponse](result.Data)
}
