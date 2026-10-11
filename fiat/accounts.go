package fiat

import (
	"github.com/dvgamerr-app/go-bitkub/bitkub"
)

type BankAccount struct {
	ID   string `json:"id"`
	Bank string `json:"bank"`
	Name string `json:"name"`
	Time int64  `json:"time"`
}

type AccountsPagination struct {
	Page int `json:"page"`
	Last int `json:"last"`
}

type AccountsResponse struct {
	bitkub.Error
	Result     []BankAccount      `json:"result"`
	Pagination AccountsPagination `json:"pagination"`
}

type AccountsParams struct {
	Page  int
	Limit int
}

func GetAccounts(params AccountsParams) (*AccountsResponse, error) {
	return fetchPage[AccountsResponse]("/api/v3/fiat/accounts", params.Page, params.Limit)
}
