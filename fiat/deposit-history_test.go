package fiat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDepositHistory(t *testing.T) {
	params := DepositHistoryParams{
		Page:  1,
		Limit: 10,
	}

	result, err := GetDepositHistory(params)

	assert.Nil(t, err)
	assert.NotNil(t, result)
}
