package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueryBalances(t *testing.T) {
	result, err := QueryBalances()
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.NotNil(t, result.Coins)
	assert.GreaterOrEqual(t, result.Total, 0.0)
	assert.GreaterOrEqual(t, result.Available, 0.0)
}
