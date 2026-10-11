package fiat

import (
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
	bitkub.Error
	Result     []DepositHistory         `json:"result"`
	Pagination DepositHistoryPagination `json:"pagination"`
}

type DepositHistoryParams struct {
	Page  int
	Limit int
}

func GetDepositHistory(params DepositHistoryParams) (*DepositHistoryResponse, error) {
	return fetchPage[DepositHistoryResponse]("/api/v3/fiat/deposit-history", params.Page, params.Limit)
}
