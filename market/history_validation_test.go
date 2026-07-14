package market

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHistoryValidationOffline(t *testing.T) {
	_, err := GetHistory(HistoryRequest{})
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "symbol is required")
}
