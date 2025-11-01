package fiat

import (
	"fmt"

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
	bitkub.GetError
	Result     []BankAccount      `json:"result"`
	Pagination AccountsPagination `json:"pagination"`
}

type AccountsParams struct {
	Page  int
	Limit int
}

func GetAccounts(params AccountsParams) (*AccountsResponse, error) {
	var response bitkub.ResponseAPI

	url := "/v3/fiat/accounts?"

	if params.Page > 0 {
		url = fmt.Sprintf("%sp=%d&", url, params.Page)
	}
	if params.Limit > 0 {
		url = fmt.Sprintf("%slmt=%d&", url, params.Limit)
	}

	if url[len(url)-1] == '&' || url[len(url)-1] == '?' {
		url = url[:len(url)-1]
	}

	if err := bitkub.FetchSecure("POST", url, nil, &response); err != nil {
		return nil, err
	}

	if err := response.CheckResponseError(); err != nil {
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
