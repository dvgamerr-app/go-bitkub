package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCoinConvertHistory(t *testing.T) {
	result, err := GetCoinConvertHistory(CoinHistoryParams{
		Page:  1,
		Limit: 10,
	})
	assert.Nil(t, err)
	assert.NotNil(t, result)
}
