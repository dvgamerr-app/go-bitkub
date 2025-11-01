package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTradingCredits(t *testing.T) {
	result, err := GetTradingCredits()
	assert.Nil(t, err)
	assert.NotNil(t, result)
}
