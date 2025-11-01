package fiat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithdraw(t *testing.T) {
	t.Skip("requires real bank account and withdrawal permission")

	request := WithdrawRequest{
		ID:  "account123",
		Amt: 1000,
	}

	result, err := Withdraw(request)

	assert.Nil(t, err)
	assert.NotNil(t, result)
}
