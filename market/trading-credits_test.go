package market

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTradingCredits(t *testing.T) {
	_, err := GetTradingCredits()
	assert.Equal(t, err, nil)
}
