package market

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlaceAsk(t *testing.T) {
	// Note: This test requires BTC balance in the account
	result, err := PlaceAsk(PlaceAskRequest{
		Symbol: "btc_thb",
		Amount: 100,
		Rate:   100,
		Type:   "limit",
	})
	if err != nil {
		t.Skipf("Cannot place ask order (may need BTC balance): %v", err)
		return
	}
	assert.NotNil(t, result)
	assert.NotEmpty(t, result.ID)

	// Clean up: cancel the order
	_ = CancelOrder(CancelOrderRequest{
		Symbol: "btc_thb",
		ID:     result.ID,
		Side:   "sell",
	})
}
