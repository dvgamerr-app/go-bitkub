package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueryWallet(t *testing.T) {
	result, err := QueryWallet()
	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestQueryCoins(t *testing.T) {
	result, err := QueryCoins()
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.IsType(t, []string{}, result)
}
