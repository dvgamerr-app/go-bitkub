package market

import (
	"fmt"

	"github.com/dvgamerr-app/go-bitkub/bitkub"
)

type Trade [4]interface{}

type TradesResponse struct {
	Error  int     `json:"error"`
	Result []Trade `json:"result"`
}

func GetTrades(sym string, lmt int) ([]Trade, error) {
	var result TradesResponse

	url := fmt.Sprintf("/v3/market/trades?sym=%s", sym)
	if lmt > 0 {
		url = fmt.Sprintf("%s&lmt=%d", url, lmt)
	}

	if err := bitkub.FetchNonSecure("GET", url, nil, &result); err != nil {
		return nil, err
	}

	if result.Error != 0 {
		errMsg, exists := bitkub.ErrorCode[result.Error]
		if !exists {
			errMsg = "Unknown error"
		}
		return nil, fmt.Errorf("[error %d] %s", result.Error, errMsg)
	}

	return result.Result, nil
}
