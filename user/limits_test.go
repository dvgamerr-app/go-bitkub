package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserLimits(t *testing.T) {
	result, err := GetUserLimits()
	assert.Nil(t, err)
	assert.NotNil(t, result)
}
