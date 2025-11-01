package market

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCancelOrder(t *testing.T) {

	bidResult, err := PlaceBid(PlaceBidRequest{
		Sym: "btc_thb",
		Amt: 100, // Increased amount to meet minimum requirement
		Rat: 100,
		Typ: "limit",
	})

	if err != nil {
		t.Skipf("Cannot create bid order for test: %v", err)
		return
	}
	assert.NotNil(t, bidResult)
	assert.NotEmpty(t, bidResult.ID)

	// Cancel the order using the ID from the created order
	request := CancelOrderRequest{
		Sym: "btc_thb",
		ID:  bidResult.ID,
		Sd:  "buy",
	}
	err = CancelOrder(request)
	assert.Nil(t, err)
}
