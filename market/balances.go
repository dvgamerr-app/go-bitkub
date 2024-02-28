package market

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/touno-io/go-bitkub/bitkub"
)

var stdJson = jsoniter.ConfigCompatibleWithStandardLibrary

type Balance struct {
	Available float64 `json:"available"`
	Reserved  float64 `json:"reserved"`
}

type BitkubBalances struct {
	Total     float64
	Available float64
	Coins     map[string]Balance
}

func GetBalances() {
	var result bitkub.ResponseAPI

	// sugar.Debugf("POST /v3/market/balances")
	if err := bitkub.FetchSecure("POST", "/v3/market/balances", nil, &result); err != nil {
		// sugar.Error(err)
	}

	byteData, err := stdJson.Marshal(result.Result)
	if err != nil {
		// sugar.Errorln("Error marshaling:", err)
	}

	data := BitkubBalances{
		Total:     0,
		Available: 0,
		Coins:     map[string]Balance{},
	}

	if err = stdJson.Unmarshal(byteData, &data.Coins); err != nil {
		// sugar.Errorln("Error unmarshaling:", err)
	}
}
