package market

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTicker(t *testing.T) {
	_, err := GetTicker("btc_thb")
	assert.Equal(t, err, nil)
}
