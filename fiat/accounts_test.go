package fiat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAccounts(t *testing.T) {
	params := AccountsParams{
		P:   1,
		Lmt: 10,
	}

	result, err := GetAccounts(params)

	assert.Nil(t, err)
	assert.NotNil(t, result)
}
