package crypto

import (
	"fmt"
	"net/url"
)

type Addresses struct {
	Pagination
	SymbolNetwork
	Memo string
}

func GetAddresses(params Addresses) (*AddressesResponse, error) {
	queryParams := url.Values{}
	addPagination(queryParams, params.Pagination)
	if params.Symbol != "" {
		queryParams.Set("symbol", params.Symbol)
	}
	if params.Network != "" {
		queryParams.Set("network", params.Network)
	}
	if params.Memo != "" {
		queryParams.Set("memo", params.Memo)
	}

	return fetchV4[AddressesResponse]("GET", pathWithQuery("/api/v4/crypto/addresses", queryParams), nil)
}

func CreateAddress(req CreateAddressRequest) ([]Address, error) {
	if req.Symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}
	if req.Network == "" {
		return nil, fmt.Errorf("network is required")
	}

	data, err := fetchV4[[]Address]("POST", "/api/v4/crypto/addresses", req)
	if err != nil {
		return nil, err
	}

	return *data, nil
}
