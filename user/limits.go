package user

import (
	"github.com/dvgamerr-app/go-bitkub/bitkub"
)

// CryptoLimit represents crypto deposit/withdraw limitations
type CryptoLimit struct {
	Deposit  float64 `json:"deposit"`  // BTC value equivalent
	Withdraw float64 `json:"withdraw"` // BTC value equivalent
}

// FiatLimit represents fiat deposit/withdraw limitations
type FiatLimit struct {
	Deposit  float64 `json:"deposit"`  // THB value equivalent
	Withdraw float64 `json:"withdraw"` // THB value equivalent
}

// Limits represents limitations by KYC level
type Limits struct {
	Crypto CryptoLimit `json:"crypto"`
	Fiat   FiatLimit   `json:"fiat"`
}

// CryptoUsage represents today's crypto usage
type CryptoUsage struct {
	Deposit               float64 `json:"deposit"`  // BTC value equivalent
	Withdraw              float64 `json:"withdraw"` // BTC value equivalent
	DepositPercentage     float64 `json:"deposit_percentage"`
	WithdrawPercentage    float64 `json:"withdraw_percentage"`
	DepositThbEquivalent  float64 `json:"deposit_thb_equivalent"`  // THB value equivalent
	WithdrawThbEquivalent float64 `json:"withdraw_thb_equivalent"` // THB value equivalent
}

// FiatUsage represents today's fiat usage
type FiatUsage struct {
	Deposit            float64 `json:"deposit"`  // THB value equivalent
	Withdraw           float64 `json:"withdraw"` // THB value equivalent
	DepositPercentage  float64 `json:"deposit_percentage"`
	WithdrawPercentage float64 `json:"withdraw_percentage"`
}

// Usage represents today's usage
type Usage struct {
	Crypto CryptoUsage `json:"crypto"`
	Fiat   FiatUsage   `json:"fiat"`
}

// UserLimits represents the complete user limits response
type UserLimits struct {
	Limits Limits  `json:"limits"`
	Usage  Usage   `json:"usage"`
	Rate   float64 `json:"rate"` // Current THB rate used to calculate
}

// UserLimitsResponse represents the response from /api/v3/user/limits endpoint
type UserLimitsResponse struct {
	Error  int        `json:"error"`
	Result UserLimits `json:"result"`
}

// GetUserLimits checks deposit/withdraw limitations and usage
// POST /api/v3/user/limits
func GetUserLimits() (*UserLimits, error) {
	var response bitkub.ResponseAPI

	if err := bitkub.FetchSecure("POST", "/v3/user/limits", nil, &response); err != nil {
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
