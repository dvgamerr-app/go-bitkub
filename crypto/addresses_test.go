package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAddresses(t *testing.T) {
	params := Addresses{
		Pagination: Pagination{
			Page:  1,
			Limit: 10,
		},
	}

	result, err := GetAddresses(params)

	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestGetAddressesWithFilter(t *testing.T) {
	params := Addresses{
		Pagination: Pagination{
			Page:  1,
			Limit: 5,
		},
		SymbolNetwork: SymbolNetwork{
			Symbol: "KUB",
		},
	}

	result, err := GetAddresses(params)

	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestCreateAddress(t *testing.T) {
	req := CreateAddressRequest{
		Symbol:  "ETH",
		Network: "ETH",
	}

	result, err := CreateAddress(req)

	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestValidationCreateAddress(t *testing.T) {
	// Test missing symbol
	req := CreateAddressRequest{
		Network: "KUB",
	}
	_, err := CreateAddress(req)
	assert.NotNil(t, err)

	// Test missing network
	req = CreateAddressRequest{
		Symbol: "KUB",
	}
	_, err = CreateAddress(req)
	assert.NotNil(t, err)
}
