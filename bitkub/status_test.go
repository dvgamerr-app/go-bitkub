package bitkub

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetStatus(t *testing.T) {

	result, err := GetStatus()
	assert.Equal(t, err, nil)
	assert.NotEqual(t, result, nil)
}

func TestGetServerTime(t *testing.T) {
	result, err := GetServerTime()
	assert.Equal(t, err, nil)
	assert.NotEqual(t, result, nil)
}
