package user

import (
	"testing"

	"github.com/dvgamerr-app/go-bitkub/bitkub"
	"github.com/dvgamerr-app/go-bitkub/utils"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func init() {
	zerolog.SetGlobalLevel(zerolog.Disabled)
	utils.LoadDotEnv("../.env")
	bitkub.Initlizer()
}

func TestGetCoinConvertHistory(t *testing.T) {
	result, err := GetCoinConvertHistory(CoinHistoryParams{
		Page:  1,
		Limit: 10,
	})
	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestGetUserLimits(t *testing.T) {
	result, err := GetUserLimits()
	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestGetTradingCredits(t *testing.T) {
	result, err := GetTradingCredits()
	assert.Nil(t, err)
	assert.NotNil(t, result)
}
