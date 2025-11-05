package market

import (
	"fmt"

	"github.com/dvgamerr-app/go-bitkub/bitkub"
)

type BidOrder struct {
	OrderID   string `json:"order_id"`
	Price     string `json:"price"`
	Side      string `json:"side"`
	Size      string `json:"size"`
	Timestamp int64  `json:"timestamp"`
	Volume    string `json:"volume"`
}

type BidsResponse struct {
	bitkub.Error
	Result []BidOrder `json:"result"`
}

func GetBids(symbol string, limit int) ([]BidOrder, error) {
	var result BidsResponse

	url := fmt.Sprintf("/v3/market/bids?sym=%s", symbol)
	if limit > 0 {
		url = fmt.Sprintf("%s&lmt=%d", url, limit)
	}

	if err := bitkub.FetchNonSecure("GET", url, nil, &result); err != nil {
		return nil, err
	}

	return result.Result, nil
}
