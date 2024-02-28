package market

import (
	"encoding/binary"
	"math"

	"github.com/touno-io/go-bitkub/bitkub"
)

type Limits struct {
	Crypto CryptoLimits `json:"crypto"`
	Fiat   FiatLimits   `json:"fiat"`
}

type CryptoLimits struct {
	Deposit  float64 `json:"deposit"`
	Withdraw float64 `json:"withdraw"`
}

type FiatLimits struct {
	Deposit  float64 `json:"deposit"`
	Withdraw float64 `json:"withdraw"`
}

type Usage struct {
	Crypto CryptoUsage `json:"crypto"`
	Fiat   FiatUsage   `json:"fiat"`
}

type CryptoUsage struct {
	Deposit               float64 `json:"deposit"`
	Withdraw              float64 `json:"withdraw"`
	DepositPercentage     float64 `json:"deposit_percentage"`
	WithdrawPercentage    float64 `json:"withdraw_percentage"`
	DepositThbEquivalent  float64 `json:"deposit_thb_equivalent"`
	WithdrawThbEquivalent float64 `json:"withdraw_thb_equivalent"`
}

type FiatUsage struct {
	Deposit            float64 `json:"deposit"`
	Withdraw           float64 `json:"withdraw"`
	DepositPercentage  float64 `json:"deposit_percentage"`
	WithdrawPercentage float64 `json:"withdraw_percentage"`
}

type Data struct {
	Limits Limits `json:"limits"`
	Usage  Usage  `json:"usage"`
	Rate   int    `json:"rate"`
}

func GetUserLimits() (float64, error) {
	var result bitkub.ResponseAPI

	if err := bitkub.FetchSecure("POST", "/api/v3/user/limits", nil, &result); err != nil {
		return 0, err
	}

	byteData, err := stdJson.Marshal(result.Result)
	if err != nil {
		return 0, err
	}

	return math.Float64frombits(binary.LittleEndian.Uint64(byteData)), nil
}
