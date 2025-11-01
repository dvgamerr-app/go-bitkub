package market

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlaceAsk(t *testing.T) {
	// Note: This test requires BTC balance in the account
	result, err := PlaceAsk(PlaceAskRequest{
		Sym: "btc_thb",
		Amt: 100,
		Rat: 100,
		Typ: "limit",
	})
	if err != nil {
		t.Skipf("Cannot place ask order (may need BTC balance): %v", err)
		return
	}
	assert.NotNil(t, result)
	assert.NotEmpty(t, result.ID)

	// Clean up: cancel the order
	_ = CancelOrder(CancelOrderRequest{
		Sym: "btc_thb",
		ID:  result.ID,
		Sd:  "sell",
	})
}
