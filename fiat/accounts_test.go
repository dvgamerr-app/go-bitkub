package fiat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAccounts(t *testing.T) {
	params := AccountsParams{
		Page:  1,
		Limit: 10,
	}

	result, err := GetAccounts(params)

	assert.Nil(t, err)
	assert.NotNil(t, result)
}
