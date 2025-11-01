package market

import (
	"testing"

	"github.com/dvgamerr-app/go-bitkub/bitkub"
	"github.com/dvgamerr-app/go-bitkub/helper"
	"github.com/rs/zerolog"
)

func init() {
	// Disable zerolog during tests
	zerolog.SetGlobalLevel(zerolog.Disabled)

	helper.LoadDotEnv("../.env")
	bitkub.Initlizer()
}

func TestGetWallet(t *testing.T) {
	result, err := GetWallet()
	if err != nil {
		t.Error("API credentials may be invalid")
		return
	}

	// Assertions
	if result == nil {
		t.Fatal("❌ GetWallet failed: result is nil")
	}

	// Validate wallet structure
	assetsWithBalance := 0
	for symbol, balance := range *result {
		if symbol == "" {
			t.Error("Symbol is empty")
		}
		if balance < 0 {
			t.Errorf("Balance should not be negative for %s", symbol)
		}
		if balance > 0 {
			assetsWithBalance++
		}
	}
}
