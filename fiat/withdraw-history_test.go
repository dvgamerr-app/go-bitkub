package fiat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetWithdrawHistory(t *testing.T) {
	params := WithdrawHistoryParams{
		Page:  1,
		Limit: 10,
	}

	result, err := GetWithdrawHistory(params)

	assert.Nil(t, err)
	assert.NotNil(t, result)
}
