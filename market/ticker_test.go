package market

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMarketTicker(t *testing.T) {
	_, err := GetMarketTicker("btc")
	assert.Equal(t, err, nil)
}
