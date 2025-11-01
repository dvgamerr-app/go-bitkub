package fiat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetWithdrawHistory(t *testing.T) {
	params := WithdrawHistoryParams{
		P:   1,
		Lmt: 10,
	}

	result, err := GetWithdrawHistory(params)

	assert.Nil(t, err)
	assert.NotNil(t, result)
}
