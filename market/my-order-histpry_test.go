package market

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMyOpenOrders(t *testing.T) {
	_, err := GetMyOpenOrders("btc")
	assert.Equal(t, err, nil)
}
