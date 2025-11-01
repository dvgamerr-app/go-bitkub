package fiat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDepositHistory(t *testing.T) {
	params := DepositHistoryParams{
		P:   1,
		Lmt: 10,
	}

	result, err := GetDepositHistory(params)

	assert.Nil(t, err)
	assert.NotNil(t, result)
}
