package market

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSymbols(t *testing.T) {
	result, err := GetSymbols()
	assert.Nil(t, err)
	assert.NotNil(t, result)
}
