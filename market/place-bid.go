package market

import (
	"github.com/dvgamerr-app/go-bitkub/bitkub"
)

// PlaceBidRequest represents the request body for placing a buy order
type PlaceBidRequest struct {
	Sym      string  `json:"sym"`                 // The symbol (e.g. btc_thb)
	Amt      float64 `json:"amt"`                 // Amount you want to spend
	Rat      float64 `json:"rat"`                 // Rate you want for the order (0 for market order)
	Typ      string  `json:"typ"`                 // Order type: limit or market
	ClientID string  `json:"client_id,omitempty"` // Your id for reference (optional)
	PostOnly bool    `json:"post_only,omitempty"` // Post-only flag (optional)
}

// PlaceBidResult represents the result from placing a buy order
type PlaceBidResult struct {
	ID       string  `json:"id"`  // Order id
	Typ      string  `json:"typ"` // Order type
	Amt      float64 `json:"amt"` // Spending amount
	Rat      float64 `json:"rat"` // Rate
	Fee      float64 `json:"fee"` // Fee
	Cre      float64 `json:"cre"` // Fee credit used
	Rec      float64 `json:"rec"` // Amount to receive
	Ts       string  `json:"ts"`  // Timestamp
	ClientID string  `json:"ci"`  // Input id for reference
}

// PlaceBidResponse represents the response from /api/v3/market/place-bid endpoint
type PlaceBidResponse struct {
	Error  int            `json:"error"`
	Result PlaceBidResult `json:"result"`
}

// PlaceBid creates a buy order
// POST /api/v3/market/place-bid
func PlaceBid(req PlaceBidRequest) (*PlaceBidResult, error) {
	var response bitkub.ResponseAPI

	if err := bitkub.FetchSecure("POST", "/v3/market/place-bid", req, &response); err != nil {
		return nil, err
	}

	byteData, err := stdJson.Marshal(response.Result)
	if err != nil {
		return nil, err
	}

	var result PlaceBidResult
	if err = stdJson.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
