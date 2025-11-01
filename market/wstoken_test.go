package market

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetWSToken(t *testing.T) {
	result, err := GetWSToken()
	assert.Nil(t, err)
	assert.NotNil(t, result)
}
