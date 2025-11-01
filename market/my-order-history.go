package market

import (
	"fmt"
	"strings"

	"github.com/dvgamerr-app/go-bitkub/bitkub"
)

type Order struct {
	ID        string `json:"id"`
	Hash      string `json:"hash"`
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

// `sym` string The symbol (e.g. btc_thb)
// `p` int Page (optional)
// `lmt` int Limit (optional)
// `start` int Start timestamp (optional)
// `end` int End timestamp (optional)

func GetMyOpenOrders(sym string) ([]Order, error) {
	var result bitkub.ResponseAPI

	if err := bitkub.FetchSecure("GET", fmt.Sprintf("/v3/market/my-order-history?sym=%s_THB", strings.ToUpper(sym)), nil, &result); err != nil {
		return nil, err
	}

	byteData, err := stdJson.Marshal(result.Result)
	if err != nil {
		return nil, err
	}

	data := []Order{}

	if err = stdJson.Unmarshal(byteData, &data); err != nil {
		return nil, err
	}
	return data, nil
}
