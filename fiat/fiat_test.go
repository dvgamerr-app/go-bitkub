package fiat

import (
	"testing"

	"github.com/dvgamerr-app/go-bitkub/bitkub"
	"github.com/dvgamerr-app/go-bitkub/utils"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func init() {
	zerolog.SetGlobalLevel(zerolog.Disabled)
	utils.LoadDotEnv("../.env")
	bitkub.Initlizer()
}

func TestGetAccounts(t *testing.T) {
	params := AccountsParams{
		Page:  1,
		Limit: 10,
	}
	result, err := GetAccounts(params)
	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestGetDepositHistory(t *testing.T) {
	params := DepositHistoryParams{
		Page:  1,
		Limit: 10,
	}
	result, err := GetDepositHistory(params)
	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestWithdraw(t *testing.T) {
	t.Skip("requires real bank account and withdrawal permission")
	request := WithdrawRequest{
		ID:     "account123",
		Amount: 1000,
	}
	result, err := Withdraw(request)
	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestGetWithdrawHistory(t *testing.T) {
	params := WithdrawHistoryParams{
		Page:  1,
		Limit: 10,
	}
	result, err := GetWithdrawHistory(params)
	assert.Nil(t, err)
	assert.NotNil(t, result)
}
