package user

import (
	"fmt"

	"github.com/dvgamerr-app/go-bitkub/bitkub"
)

type CoinConvertHistory struct {
	TransactionID      string `json:"transaction_id"`
	Status             string `json:"status"`
	Amount             string `json:"amount"`
	FromCurrency       string `json:"from_currency"`
	TradingFeeReceived string `json:"trading_fee_received"`
	Timestamp          int64  `json:"timestamp"`
}

type CoinConvertPagination struct {
	Page int  `json:"page"`
	Last int  `json:"last"`
	Next *int `json:"next,omitempty"`
}

type CoinConvertHistoryResponse struct {
	Error      int                   `json:"error"`
	Result     []CoinConvertHistory  `json:"result"`
	Pagination CoinConvertPagination `json:"pagination"`
}

type CoinConvertHistoryParams struct {
	P      int
	Lmt    int
	Sort   int
	Status string
	Sym    string
	Start  int64
	End    int64
}

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

	if url[len(url)-1] == '&' || url[len(url)-1] == '?' {
		url = url[:len(url)-1]
	}

	if err := bitkub.FetchSecure("GET", url, nil, &response); err != nil {
		return nil, err
	}

	if err := response.CheckResponseError(); err != nil {
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
