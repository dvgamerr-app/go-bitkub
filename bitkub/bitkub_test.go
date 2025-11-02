package bitkub

import (
	"testing"

	"github.com/dvgamerr-app/go-bitkub/utils"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func init() {
	// Disable zerolog during tests
	zerolog.SetGlobalLevel(zerolog.Disabled)

	utils.LoadDotEnv("../.env")
	Initlizer()
}
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
