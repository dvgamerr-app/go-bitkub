package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCompensations(t *testing.T) {
	params := Compensations{
		Pagination: Pagination{
			Page:  1,
			Limit: 10,
		},
	}

	result, err := GetCompensations(params)

	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestGetCompensationsByType(t *testing.T) {
	params := Compensations{
		Pagination: Pagination{
			Page:  1,
			Limit: 5,
		},
		Type: "COMPENSATE",
	}

	result, err := GetCompensations(params)

	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestGetCompensationsByStatus(t *testing.T) {
	params := Compensations{
		Pagination: Pagination{
			Page:  1,
			Limit: 5,
		},
		Status: "COMPLETED",
	}

	result, err := GetCompensations(params)

	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestGetCompensationsBySymbol(t *testing.T) {
	params := Compensations{
		Pagination: Pagination{
			Page:  1,
			Limit: 5,
		},
		Symbol: "XRP",
	}

	result, err := GetCompensations(params)

	assert.Nil(t, err)
	assert.NotNil(t, result)
}
