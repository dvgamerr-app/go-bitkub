package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetWithdraws(t *testing.T) {
	params := Withdraws{
		Pagination: Pagination{
			Page:  1,
			Limit: 10,
		},
	}

	result, err := GetWithdraws(params)

	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestGetWithdrawsWithFilters(t *testing.T) {
	params := Withdraws{
		Pagination: Pagination{
			Page:  1,
			Limit: 5,
		},
		Status: "complete",
	}

	result, err := GetWithdraws(params)

	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestGetWithdrawsBySymbol(t *testing.T) {
	params := Withdraws{
		Pagination: Pagination{
			Page:  1,
			Limit: 5,
		},
		Symbol: "THB",
	}

	result, err := GetWithdraws(params)

	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestCreateWithdraw(t *testing.T) {
	t.Skip("requires is_withdraw permission and real transaction")

	req := CreateWithdrawRequest{
		Symbol:  "RDNT",
		Amount:  "2.00000000",
		Address: "0xDaCd17d1E77604aaFB6e47F5Ffa1F7E35F83fDa7",
		Network: "ARB",
	}

	result, err := CreateWithdraw(req)

	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestCreateWithdrawValidation(t *testing.T) {
	// Test missing symbol
	req := CreateWithdrawRequest{
		Amount:  "2.00000000",
		Address: "0xDaCd17d1E77604aaFB6e47F5Ffa1F7E35F83fDa7",
		Network: "ARB",
	}
	_, err := CreateWithdraw(req)
	assert.NotNil(t, err)

	// Test missing amount
	req = CreateWithdrawRequest{
		Symbol:  "RDNT",
		Address: "0xDaCd17d1E77604aaFB6e47F5Ffa1F7E35F83fDa7",
		Network: "ARB",
	}
	_, err = CreateWithdraw(req)
	assert.NotNil(t, err)

	// Test missing address
	req = CreateWithdrawRequest{
		Symbol:  "RDNT",
		Amount:  "2.00000000",
		Network: "ARB",
	}
	_, err = CreateWithdraw(req)
	assert.NotNil(t, err)

	// Test missing network
	req = CreateWithdrawRequest{
		Symbol:  "RDNT",
		Amount:  "2.00000000",
		Address: "0xDaCd17d1E77604aaFB6e47F5Ffa1F7E35F83fDa7",
	}
	_, err = CreateWithdraw(req)
	assert.NotNil(t, err)
}
