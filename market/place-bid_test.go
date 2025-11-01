package market

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlaceBid(t *testing.T) {
	request := PlaceBidRequest{
		Sym: "btc_thb",
		Amt: 100,
		Rat: 100,
		Typ: "limit",
	}
	result, err := PlaceBid(request)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.NotEmpty(t, result.ID)

	if err != nil {
		t.Skipf("Cannot place ask order (may need BTC balance): %v", err)
		return
	}

	// Clean up: cancel the order
	_ = CancelOrder(CancelOrderRequest{
		Sym: "btc_thb",
		ID:  result.ID,
		Sd:  "buy",
	})

}
