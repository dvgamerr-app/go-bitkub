package user

import (
	"fmt"

	"github.com/dvgamerr-app/go-bitkub/bitkub"
)

// CoinConvertHistory represents a coin convert history item
type CoinConvertHistory struct {
	TransactionID      string `json:"transaction_id"`
	Status             string `json:"status"`
	Amount             string `json:"amount"`
	FromCurrency       string `json:"from_currency"`
	TradingFeeReceived string `json:"trading_fee_received"`
	Timestamp          int64  `json:"timestamp"`
}

// CoinConvertPagination represents pagination information for coin convert history
type CoinConvertPagination struct {
	Page int  `json:"page"`
	Last int  `json:"last"`
	Next *int `json:"next,omitempty"`
}

// CoinConvertHistoryResponse represents the response from /api/v3/user/coin-convert-history endpoint
type CoinConvertHistoryResponse struct {
	Error      int                   `json:"error"`
	Result     []CoinConvertHistory  `json:"result"`
	Pagination CoinConvertPagination `json:"pagination"`
}

// CoinConvertHistoryParams represents parameters for coin convert history request
type CoinConvertHistoryParams struct {
	P      int    // Page default = 1 (optional)
	Lmt    int    // Limit default = 100 (optional)
	Sort   int    // Sort [1, -1] default = 1 (optional)
	Status string // Status [success, fail, all] (default = all) (optional)
	Sym    string // The symbol (e.g. KUB) (optional)
	Start  int64  // Start timestamp (optional)
	End    int64  // End timestamp (optional)
}

// GetCoinConvertHistory lists all coin convert histories (paginated)
// GET /api/v3/user/coin-convert-history
func GetCoinConvertHistory(params CoinConvertHistoryParams) (*CoinConvertHistoryResponse, error) {
	var response bitkub.ResponseAPI

	url := "/v3/user/coin-convert-history?"

	if params.P > 0 {
		url = fmt.Sprintf("%sp=%d&", url, params.P)
	}
	if params.Lmt > 0 {
		url = fmt.Sprintf("%slmt=%d&", url, params.Lmt)
	}
	if params.Sort != 0 {
		url = fmt.Sprintf("%ssort=%d&", url, params.Sort)
	}
	if params.Status != "" {
		url = fmt.Sprintf("%sstatus=%s&", url, params.Status)
	}
	if params.Sym != "" {
		url = fmt.Sprintf("%ssym=%s&", url, params.Sym)
	}
	if params.Start > 0 {
		url = fmt.Sprintf("%sstart=%d&", url, params.Start)
	}
	if params.End > 0 {
		url = fmt.Sprintf("%send=%d&", url, params.End)
	}

	// Remove trailing '&' or '?'
	if url[len(url)-1] == '&' || url[len(url)-1] == '?' {
		url = url[:len(url)-1]
	}

	if err := bitkub.FetchSecure("GET", url, nil, &response); err != nil {
		return nil, err
	}

	byteData, err := stdJson.Marshal(response)
	if err != nil {
		return nil, err
	}

	var result CoinConvertHistoryResponse
	if err = stdJson.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
