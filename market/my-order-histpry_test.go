package market

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMyOpenOrders(t *testing.T) {
	result, err := GetMyOpenOrders("btc_thb")
	assert.Nil(t, err)
	assert.NotNil(t, result)
}
