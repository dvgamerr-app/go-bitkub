package bitkub

// GetError represents the error response from Bitkub API
type GetError struct {
	Error int `json:"error"`
}

type PaginationInfo struct {
	Page    int    `json:"page,omitempty"`
	Last    int    `json:"last,omitempty"`
	Next    *int   `json:"next,omitempty"`
	Prev    *int   `json:"prev,omitempty"`
	Cursor  string `json:"cursor,omitempty"`
	HasNext bool   `json:"has_next,omitempty"`
}

type OrderSide string

const (
	OrderSideBuy  OrderSide = "buy"
	OrderSideSell OrderSide = "sell"
)

type OrderType string

const (
	OrderTypeLimit  OrderType = "limit"
	OrderTypeMarket OrderType = "market"
)

type OrderStatus string

const (
	OrderStatusFilled   OrderStatus = "filled"
	OrderStatusUnfilled OrderStatus = "unfilled"
	OrderStatusCanceled OrderStatus = "cancelled"
)

type TransactionStatus string

const (
	TransactionStatusComplete TransactionStatus = "complete"
	TransactionStatusPending  TransactionStatus = "pending"
	TransactionStatusFailed   TransactionStatus = "failed"
)
