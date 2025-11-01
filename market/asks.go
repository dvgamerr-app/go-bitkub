package market

import (
	"fmt"

	"github.com/dvgamerr-app/go-bitkub/bitkub"
)

// AskOrder represents an open sell order
type AskOrder struct {
	OrderID   string `json:"order_id"`
	Price     string `json:"price"`
	Side      string `json:"side"`
	Size      string `json:"size"`
	Timestamp int64  `json:"timestamp"`
	Volume    string `json:"volume"`
}

// AsksResponse represents the response from /api/v3/market/asks endpoint
type AsksResponse struct {
	Error  int        `json:"error"`
	Result []AskOrder `json:"result"`
}

// GetAsks lists open sell orders
// GET /api/v3/market/asks
// sym: The symbol (e.g. btc_thb)
// lmt: No. of limit to query open sell orders (optional)
func GetAsks(sym string, lmt int) ([]AskOrder, error) {
	var result AsksResponse

	url := fmt.Sprintf("/v3/market/asks?sym=%s", sym)
	if lmt > 0 {
		url = fmt.Sprintf("%s&lmt=%d", url, lmt)
	}

	if err := bitkub.FetchNonSecure("GET", url, nil, &result); err != nil {
		return nil, err
	}

	return result.Result, nil
}
