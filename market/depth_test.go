package market

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDepth(t *testing.T) {
	result, err := GetDepth("btc_thb", 10)
	assert.Nil(t, err)
	assert.NotNil(t, result)
}
