package market

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetOrderInfo(t *testing.T) {
	// Try to create an order first to get a valid order ID
	bidResult, err := PlaceBid(PlaceBidRequest{
		Sym: "btc_thb",
		Amt: 100,
		Rat: 100,
		Typ: "limit",
	})
	if err != nil {
		t.Skipf("Cannot create bid order for test: %v", err)
		return
	}

	// Now test GetOrderInfo with the created order ID
	result, err := GetOrderInfo("btc_thb", bidResult.ID, "buy")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, bidResult.ID, result.ID)

	// Clean up: cancel the order
	_ = CancelOrder(CancelOrderRequest{
		Sym: "btc_thb",
		ID:  bidResult.ID,
		Sd:  "buy",
	})
}
