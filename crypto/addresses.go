package crypto

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/dvgamerr-app/go-bitkub/bitkub"
)

type Addresses struct {
	Pagination
	SymbolNetwork
	Memo string
}

func GetAddresses(params Addresses) (*AddressesResponse, error) {
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
	if params.Network != "" {
		queryParams.Add("network", params.Network)
	}
	if params.Memo != "" {
		queryParams.Add("memo", params.Memo)
	}

	path := "/api/v4/crypto/addresses"
	if len(queryParams) > 0 {
		path += "?" + queryParams.Encode()
	}

	if err := bitkub.FetchSecureV4("GET", path, nil, &result); err != nil {
		return nil, err
	}

	return bitkub.DecodeResult[AddressesResponse](result.Data)
}

func CreateAddress(req CreateAddressRequest) ([]Address, error) {
	if req.Symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}
	if req.Network == "" {
		return nil, fmt.Errorf("network is required")
	}

	var result bitkub.ResponseAPIV4

	if err := bitkub.FetchSecureV4("POST", "/api/v4/crypto/addresses", req, &result); err != nil {
		return nil, err
	}

	data, err := bitkub.DecodeResult[[]Address](result.Data)
	if err != nil {
		return nil, err
	}

	return *data, nil
}
