package market

import (
	"os"
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

func TestGetBalances(t *testing.T) {
	if os.Getenv("SKIP_INTEGRATION_TESTS") != "" {
		t.Skip("Skipping integration test")
	}

	result, err := GetBalances()
	if err != nil {
		t.Error("API credentials may be invalid")
		return
	}

	// Assertions
	if result == nil {
		t.Fatal("❌ GetBalances failed: result is nil")
	}

	// Validate balance structure
	assetsWithBalance := 0
	for symbol, balance := range result {
		if symbol == "" {
			t.Error("Symbol is empty")
		}
		if balance.Available < 0 {
			t.Errorf("Available balance should not be negative for %s", symbol)
		}
		if balance.Reserved < 0 {
			t.Errorf("Reserved balance should not be negative for %s", symbol)
		}
		if balance.Available > 0 || balance.Reserved > 0 {
			assetsWithBalance++
		}
	}

}
