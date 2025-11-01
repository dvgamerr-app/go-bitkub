package market

import (
	"github.com/dvgamerr-app/go-bitkub/bitkub"
)

// PlaceAskRequest represents the request body for placing a sell order
type PlaceAskRequest struct {
	Sym      string  `json:"sym"`                 // The symbol (e.g. btc_thb)
	Amt      float64 `json:"amt"`                 // Amount you want to sell
	Rat      float64 `json:"rat"`                 // Rate you want for the order (0 for market order)
	Typ      string  `json:"typ"`                 // Order type: limit or market
	ClientID string  `json:"client_id,omitempty"` // Your id for reference (optional)
	PostOnly bool    `json:"post_only,omitempty"` // Post-only flag (optional)
}

// PlaceAskResult represents the result from placing a sell order
type PlaceAskResult struct {
	ID       string  `json:"id"`  // Order id
	Typ      string  `json:"typ"` // Order type
	Amt      float64 `json:"amt"` // Selling amount
	Rat      float64 `json:"rat"` // Rate
	Fee      float64 `json:"fee"` // Fee
	Cre      float64 `json:"cre"` // Fee credit used
	Rec      float64 `json:"rec"` // Amount to receive
	Ts       string  `json:"ts"`  // Timestamp
	ClientID string  `json:"ci"`  // Input id for reference
}

// PlaceAskResponse represents the response from /api/v3/market/place-ask endpoint
type PlaceAskResponse struct {
	Error  int            `json:"error"`
	Result PlaceAskResult `json:"result"`
}

// PlaceAsk creates a sell order
// POST /api/v3/market/place-ask
func PlaceAsk(req PlaceAskRequest) (*PlaceAskResult, error) {
	var response bitkub.ResponseAPI

	if err := bitkub.FetchSecure("POST", "/v3/market/place-ask", req, &response); err != nil {
		return nil, err
	}

	byteData, err := stdJson.Marshal(response.Result)
	if err != nil {
		return nil, err
	}

	var result PlaceAskResult
	if err = stdJson.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
