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
	bitkub.GetError
	Result []BidOrder `json:"result"`
}

func GetBids(sym string, lmt int) ([]BidOrder, error) {
	var result BidsResponse

	url := fmt.Sprintf("/v3/market/bids?sym=%s", sym)
	if lmt > 0 {
		url = fmt.Sprintf("%s&lmt=%d", url, lmt)
	}

	if err := bitkub.FetchNonSecure("GET", url, nil, &result); err != nil {
		return nil, err
	}

	return result.Result, nil
}
