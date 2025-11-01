package market

import (
	"github.com/dvgamerr-app/go-bitkub/bitkub"
)

type CancelOrderRequest struct {
	Symbol string `json:"sym"`
	ID     string `json:"id"`
	Side   string `json:"sd"`
}

func CancelOrder(req CancelOrderRequest) error {
	var response bitkub.ResponseAPI

	if err := bitkub.FetchSecure("POST", "/v3/market/cancel-order", req, &response); err != nil {
		return err
	}

	if err := response.CheckResponseError(); err != nil {
		return err
	}

	return nil
}
