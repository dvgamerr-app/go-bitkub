package market

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTicker(t *testing.T) {
	result, err := GetTicker("btc_thb")
	assert.Nil(t, err)
	assert.NotNil(t, result)
}
