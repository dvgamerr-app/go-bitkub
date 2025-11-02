package main

import (
	"testing"

	"github.com/dvgamerr-app/go-bitkub/bitkub"
	"github.com/dvgamerr-app/go-bitkub/utils"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func init() {
	// Disable zerolog during tests
	zerolog.SetGlobalLevel(zerolog.Disabled)

	utils.LoadDotEnv()
	bitkub.Initlizer()
}

func TestQueryBalances(t *testing.T) {
	result, err := QueryBalances()
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.NotNil(t, result.Coins)
	assert.GreaterOrEqual(t, result.Total, 0.0)
	assert.GreaterOrEqual(t, result.Available, 0.0)
}

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
