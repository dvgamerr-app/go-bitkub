package market

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetWallet(t *testing.T) {
	result, err := GetWallet()
	assert.Nil(t, err)
	assert.NotNil(t, result)
}
