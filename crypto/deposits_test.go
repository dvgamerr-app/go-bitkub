package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDeposits(t *testing.T) {
	params := Deposits{
		Pagination: Pagination{
			Page:  1,
			Limit: 10,
		},
	}

	result, err := GetDeposits(params)

	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestGetDepositsWithFilters(t *testing.T) {
	params := Deposits{
		Pagination: Pagination{
			Page:  1,
			Limit: 5,
		},
		Status: "complete",
	}

	result, err := GetDeposits(params)

	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestGetDepositsBySymbol(t *testing.T) {
	params := Deposits{
		Pagination: Pagination{
			Page:  1,
			Limit: 5,
		},
		Symbol: "BTC",
	}

	result, err := GetDeposits(params)

	assert.Nil(t, err)
	assert.NotNil(t, result)
}
