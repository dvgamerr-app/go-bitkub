package user

import (
	"github.com/dvgamerr-app/go-bitkub/bitkub"
)

type CryptoLimit struct {
	Deposit  float64 `json:"deposit"`
	Withdraw float64 `json:"withdraw"`
}

type FiatLimit struct {
	Deposit  float64 `json:"deposit"`
	Withdraw float64 `json:"withdraw"`
}

type Limits struct {
	Crypto CryptoLimit `json:"crypto"`
	Fiat   FiatLimit   `json:"fiat"`
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

type Usage struct {
	Crypto CryptoUsage `json:"crypto"`
	Fiat   FiatUsage   `json:"fiat"`
}

type UserLimits struct {
	Limits Limits  `json:"limits"`
	Usage  Usage   `json:"usage"`
	Rate   float64 `json:"rate"`
}

type UserLimitsResponse struct {
	Error  int        `json:"error"`
	Result UserLimits `json:"result"`
}

func GetUserLimits() (*UserLimits, error) {
	var response bitkub.ResponseAPI

	if err := bitkub.FetchSecure("POST", "/v3/user/limits", nil, &response); err != nil {
		return nil, err
	}

	if err := response.CheckResponseError(); err != nil {
		return nil, err
	}

	byteData, err := stdJson.Marshal(response.Result)
	if err != nil {
		return nil, err
	}

	var result UserLimits
	if err = stdJson.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
