package market

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAsks(t *testing.T) {
	result, err := GetAsks("btc_thb", 10)
	assert.Nil(t, err)
	assert.NotNil(t, result)
}
