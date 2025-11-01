package market

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBids(t *testing.T) {
	result, err := GetBids("btc_thb", 10)
	assert.Nil(t, err)
	assert.NotNil(t, result)
}
