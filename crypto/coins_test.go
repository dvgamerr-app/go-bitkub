package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCoins(t *testing.T) {
	params := Coins{}

	result, err := GetCoins(params)

	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestGetCoinsWithSymbol(t *testing.T) {
	params := Coins{
		Symbol: "KUB",
	}

	result, err := GetCoins(params)

	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestGetCoinsMultipleSymbols(t *testing.T) {
	symbols := []string{"ETH", "USDT", "BNB"}

	for _, symbol := range symbols {
		params := Coins{
			Symbol: symbol,
		}

		result, err := GetCoins(params)

		assert.Nil(t, err)
		assert.NotNil(t, result)
	}
}
