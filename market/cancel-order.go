package market

import (
	"github.com/dvgamerr-app/go-bitkub/bitkub"
)

// CancelOrderRequest represents the request body for canceling an order
type CancelOrderRequest struct {
	Sym string `json:"sym"` // The symbol (e.g. btc_thb)
	ID  string `json:"id"`  // Order id you wish to cancel
	Sd  string `json:"sd"`  // Order side: buy or sell
}

// CancelOrder cancels an open order
// POST /api/v3/market/cancel-order
func CancelOrder(req CancelOrderRequest) error {
	var response bitkub.ResponseAPI

	if err := bitkub.FetchSecure("POST", "/v3/market/cancel-order", req, &response); err != nil {
		return err
	}

	return nil
}
