package bitkub

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
)

var stdJson = jsoniter.ConfigCompatibleWithStandardLibrary

func DecodeResult[T any](result any) (*T, error) {
	byteData, err := stdJson.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("marshal result: %w", err)
	}

	var data T
	if err = stdJson.Unmarshal(byteData, &data); err != nil {
		return nil, fmt.Errorf("unmarshal result: %w", err)
	}

	return &data, nil
}

type Error struct {
	Error int `json:"error"`
}

func (e *Error) GetError() int {
	return e.Error
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
