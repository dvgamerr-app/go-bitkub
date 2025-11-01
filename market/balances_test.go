package market

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBalances(t *testing.T) {
	result, err := GetBalances()
	assert.Nil(t, err)
	assert.NotNil(t, result)
}
