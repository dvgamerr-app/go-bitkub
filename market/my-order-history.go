package market

import (
	"fmt"

	"github.com/dvgamerr-app/go-bitkub/bitkub"
)

type OpenOrder struct {
	ID        string `json:"id"`
	Side      string `json:"side"`
	Type      string `json:"type"`
	Rate      string `json:"rate"`
	Fee       string `json:"fee"`
	Credit    string `json:"credit"`
	Amount    string `json:"amount"`
	Receive   string `json:"receive"`
	ParentID  string `json:"parent_id"`
	SuperID   string `json:"super_id"`
	ClientID  string `json:"client_id"`
	Timestamp int64  `json:"ts"`
}

type OrderHistory struct {
	TxnID         string `json:"txn_id"`
	OrderID       string `json:"order_id"`
	ParentOrderID string `json:"parent_order_id"`
	SuperOrderID  string `json:"super_order_id"`
	ClientID      string `json:"client_id"`
	TakenByMe     bool   `json:"taken_by_me"`
	IsMaker       bool   `json:"is_maker"`
	Side          string `json:"side"`
	Type          string `json:"type"`
	Rate          string `json:"rate"`
	Fee           string `json:"fee"`
	Credit        string `json:"credit"`
	Amount        string `json:"amount"`
	Timestamp     int64  `json:"ts"`
	OrderClosedAt int64  `json:"order_closed_at"`
}

type Pagination struct {
	Page    int    `json:"page,omitempty"`
	Last    int    `json:"last,omitempty"`
	Next    *int   `json:"next,omitempty"`
	Prev    *int   `json:"prev,omitempty"`
	Cursor  string `json:"cursor,omitempty"`
	HasNext bool   `json:"has_next,omitempty"`
}

type MyOrderHistoryParams struct {
	Symbol         string
	Page           string
	Limit          string
	Cursor         string
	Start          string
	End            string
	PaginationType string
}

type MyOrderHistoryResponse struct {
	bitkub.Error
	Result     []OrderHistory `json:"result"`
	Pagination Pagination     `json:"pagination"`
}

func GetMyOpenOrders(symbol string) ([]OpenOrder, error) {
	var response bitkub.ResponseAPI

	url := fmt.Sprintf("/api/v3/market/my-open-orders?sym=%s", symbol)

	if err := bitkub.FetchSecure("GET", url, nil, &response); err != nil {
		return nil, err
	}

	if err := response.CheckResponseError(); err != nil {
		return nil, err
	}

	byteData, err := stdJson.Marshal(response.Result)
	if err != nil {
		return nil, err
	}

	var result []OpenOrder
	if err = stdJson.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func GetMyOrderHistory(params MyOrderHistoryParams) (*MyOrderHistoryResponse, error) {
	var response bitkub.ResponseAPI

	url := fmt.Sprintf("/api/v3/market/my-order-history?sym=%s", params.Symbol)

	if params.Page != "" {
		url = fmt.Sprintf("%s&p=%s", url, params.Page)
	}
	if params.Limit != "" {
		url = fmt.Sprintf("%s&lmt=%s", url, params.Limit)
	}
	if params.Cursor != "" {
		url = fmt.Sprintf("%s&cursor=%s", url, params.Cursor)
	}
	if params.Start != "" {
		url = fmt.Sprintf("%s&start=%s", url, params.Start)
	}
	if params.End != "" {
		url = fmt.Sprintf("%s&end=%s", url, params.End)
	}
	if params.PaginationType != "" {
		url = fmt.Sprintf("%s&pagination_type=%s", url, params.PaginationType)
	}

	if err := bitkub.FetchSecure("GET", url, nil, &response); err != nil {
		return nil, err
	}

	if err := response.CheckResponseError(); err != nil {
		return nil, err
	}

	byteData, err := stdJson.Marshal(response)
	if err != nil {
		return nil, err
	}

	var result MyOrderHistoryResponse
	if err = stdJson.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
