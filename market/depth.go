package market

import (
	"fmt"

	"github.com/dvgamerr-app/go-bitkub/bitkub"
)

// DepthEntry represents a depth entry [price, size]
type DepthEntry [2]float64

// DepthResult represents the depth information
type DepthResult struct {
	Asks []DepthEntry `json:"asks"`
	Bids []DepthEntry `json:"bids"`
}

// DepthResponse represents the response from /api/v3/market/depth endpoint
type DepthResponse struct {
	Error  int         `json:"error"`
	Result DepthResult `json:"result"`
}

// GetDepth gets depth information
// GET /api/v3/market/depth
// sym: The symbol (e.g. btc_thb)
// lmt: Depth size (optional)
func GetDepth(sym string, lmt int) (*DepthResult, error) {
	var result DepthResponse

	url := fmt.Sprintf("/v3/market/depth?sym=%s", sym)
	if lmt > 0 {
		url = fmt.Sprintf("%s&lmt=%d", url, lmt)
	}

	if err := bitkub.FetchNonSecure("GET", url, nil, &result); err != nil {
		return nil, err
	}

	return &result.Result, nil
}
