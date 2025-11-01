package bitkub

// Common types and structures used across the SDK

// PaginationInfo represents common pagination information
type PaginationInfo struct {
	Page    int    `json:"page,omitempty"`
	Last    int    `json:"last,omitempty"`
	Next    *int   `json:"next,omitempty"`
	Prev    *int   `json:"prev,omitempty"`
	Cursor  string `json:"cursor,omitempty"`
	HasNext bool   `json:"has_next,omitempty"`
}

// OrderSide represents the side of an order
type OrderSide string

const (
	OrderSideBuy  OrderSide = "buy"
	OrderSideSell OrderSide = "sell"
)

// OrderType represents the type of an order
type OrderType string

const (
	OrderTypeLimit  OrderType = "limit"
	OrderTypeMarket OrderType = "market"
)

// OrderStatus represents the status of an order
type OrderStatus string

const (
	OrderStatusFilled   OrderStatus = "filled"
	OrderStatusUnfilled OrderStatus = "unfilled"
	OrderStatusCanceled OrderStatus = "cancelled"
)

// TransactionStatus represents the status of a transaction
type TransactionStatus string

const (
	TransactionStatusComplete TransactionStatus = "complete"
	TransactionStatusPending  TransactionStatus = "pending"
	TransactionStatusFailed   TransactionStatus = "failed"
)
