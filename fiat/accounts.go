package fiat

import (
	"fmt"

	"github.com/dvgamerr-app/go-bitkub/bitkub"
)

// BankAccount represents a bank account
type BankAccount struct {
	ID   string `json:"id"`
	Bank string `json:"bank"`
	Name string `json:"name"`
	Time int64  `json:"time"`
}

// AccountsPagination represents pagination information
type AccountsPagination struct {
	Page int `json:"page"`
	Last int `json:"last"`
}

// AccountsResponse represents the response from /api/v3/fiat/accounts endpoint
type AccountsResponse struct {
	Error      int                `json:"error"`
	Result     []BankAccount      `json:"result"`
	Pagination AccountsPagination `json:"pagination"`
}

// AccountsParams represents parameters for accounts request
type AccountsParams struct {
	P   int // Page (optional)
	Lmt int // Limit (optional)
}

// GetAccounts lists all approved bank accounts
// POST /api/v3/fiat/accounts
func GetAccounts(params AccountsParams) (*AccountsResponse, error) {
	var response bitkub.ResponseAPI

	url := "/v3/fiat/accounts?"

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

	var result AccountsResponse
	if err = stdJson.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
