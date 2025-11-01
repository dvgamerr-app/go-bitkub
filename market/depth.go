package market

import (
	"fmt"

	"github.com/dvgamerr-app/go-bitkub/bitkub"
)

type DepthEntry [2]float64

type DepthResult struct {
	Asks []DepthEntry `json:"asks"`
	Bids []DepthEntry `json:"bids"`
}

type DepthResponse struct {
	Error  int         `json:"error"`
	Result DepthResult `json:"result"`
}

func GetDepth(sym string, lmt int) (*DepthResult, error) {
	var result DepthResponse

	url := fmt.Sprintf("/v3/market/depth?sym=%s", sym)
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

	return &result.Result, nil
}
