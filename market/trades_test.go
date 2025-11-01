package market

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTrades(t *testing.T) {
	result, err := GetTrades("btc_thb", 10)
	assert.Nil(t, err)
	assert.NotNil(t, result)
}
