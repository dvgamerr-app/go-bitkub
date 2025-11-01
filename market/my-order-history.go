package market

import (
	"fmt"

	"github.com/dvgamerr-app/go-bitkub/bitkub"
)

// OpenOrder represents an open order
type OpenOrder struct {
	ID       string `json:"id"`        // Order id
	Side     string `json:"side"`      // Order side: buy or sell
	Type     string `json:"type"`      // Order type: limit or market
	Rate     string `json:"rate"`      // Rate
	Fee      string `json:"fee"`       // Fee
	Credit   string `json:"credit"`    // Credit used
	Amount   string `json:"amount"`    // Amount (THB for buy, crypto for sell)
	Receive  string `json:"receive"`   // Amount to receive
	ParentID string `json:"parent_id"` // Parent order id
	SuperID  string `json:"super_id"`  // Super parent order id
	ClientID string `json:"client_id"` // Client id
	Ts       int64  `json:"ts"`        // Timestamp
}

// OrderHistory represents an order history item
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
	Ts            int64  `json:"ts"`
	OrderClosedAt int64  `json:"order_closed_at"`
}

// Pagination represents pagination information
type Pagination struct {
	Page    int    `json:"page,omitempty"`
	Last    int    `json:"last,omitempty"`
	Next    *int   `json:"next,omitempty"`
	Prev    *int   `json:"prev,omitempty"`
	Cursor  string `json:"cursor,omitempty"`
	HasNext bool   `json:"has_next,omitempty"`
}

// MyOrderHistoryParams represents parameters for order history request
type MyOrderHistoryParams struct {
	Sym            string // The trading symbol (e.g. BTC_THB) - required
	P              string // Page number for page-based pagination (optional)
	Lmt            string // Limit per page, default: 10, min: 1 (optional)
	Cursor         string // Base64 encoded cursor for keyset pagination (optional)
	Start          string // Start timestamp (optional)
	End            string // End timestamp (optional)
	PaginationType string // Pagination type: "page" or "keyset", default: "page" (optional)
}

// MyOrderHistoryResponse represents the response from /api/v3/market/my-order-history endpoint
type MyOrderHistoryResponse struct {
	Error      int            `json:"error"`
	Result     []OrderHistory `json:"result"`
	Pagination Pagination     `json:"pagination"`
}

// GetMyOpenOrders lists all open orders of the given symbol
// GET /api/v3/market/my-open-orders
// sym: The symbol (e.g. btc_thb or BTC_THB)
func GetMyOpenOrders(sym string) ([]OpenOrder, error) {
	var response bitkub.ResponseAPI

	url := fmt.Sprintf("/v3/market/my-open-orders?sym=%s", sym)

	if err := bitkub.FetchSecure("GET", url, nil, &response); err != nil {
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

// GetMyOrderHistory lists all orders that have already matched
// GET /api/v3/market/my-order-history
// Uses keyset-based pagination for better performance
func GetMyOrderHistory(params MyOrderHistoryParams) (*MyOrderHistoryResponse, error) {
	var response bitkub.ResponseAPI

	url := fmt.Sprintf("/v3/market/my-order-history?sym=%s", params.Sym)

	if params.P != "" {
		url = fmt.Sprintf("%s&p=%s", url, params.P)
	}
	if params.Lmt != "" {
		url = fmt.Sprintf("%s&lmt=%s", url, params.Lmt)
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
