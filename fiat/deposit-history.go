package fiat

import (
	"fmt"

	"github.com/dvgamerr-app/go-bitkub/bitkub"
)

type DepositHistory struct {
	TxnID    string  `json:"txn_id"`
	Currency string  `json:"currency"`
	Amount   float64 `json:"amount"`
	Status   string  `json:"status"`
	Time     int64   `json:"time"`
}

type DepositHistoryPagination struct {
	Page int `json:"page"`
	Last int `json:"last"`
}

type DepositHistoryResponse struct {
	Error      int                      `json:"error"`
	Result     []DepositHistory         `json:"result"`
	Pagination DepositHistoryPagination `json:"pagination"`
}

type DepositHistoryParams struct {
	P   int
	Lmt int
}

func GetDepositHistory(params DepositHistoryParams) (*DepositHistoryResponse, error) {
	var response bitkub.ResponseAPI

	url := "/v3/fiat/deposit-history?"

	if params.P > 0 {
		url = fmt.Sprintf("%sp=%d&", url, params.P)
	}
	if params.Lmt > 0 {
		url = fmt.Sprintf("%slmt=%d&", url, params.Lmt)
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

	var result DepositHistoryResponse
	if err = stdJson.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
