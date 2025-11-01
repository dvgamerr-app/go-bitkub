package crypto

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

func TestGetCoins(t *testing.T) {
	params := GetCoinsParams{}

	result, err := GetCoins(params)
	if err != nil {
		t.Skip("Skipping - API credentials may be invalid")
		return
	}

	// Assertions
	if result == nil {
		t.Fatal("❌ result is nil")
	}
	if len(result.Items) == 0 {
		t.Fatal("❌ no coins returned")
	}

	// Validate first coin structure
	coin := result.Items[0]
	if coin.Symbol == "" {
		t.Error("First coin symbol is empty")
	}
	if coin.Name == "" {
		t.Error("First coin name is empty")
	}
	if len(coin.Networks) > 0 {
		network := coin.Networks[0]
		if network.Network == "" {
			t.Error("First network name is empty")
		}
		if network.Decimal < 0 {
			t.Error("Network decimal should not be negative")
		}
	}
}

func TestGetCoinsWithSymbol(t *testing.T) {
	params := GetCoinsParams{
		Symbol: "KUB",
	}

	result, err := GetCoins(params)
	if err != nil {
		t.Fatalf("❌ GetCoins (%s) failed: %v", err, params.Symbol)
	}

	// Assertions
	if result == nil {
		t.Fatal("Result is nil")
	}
	if len(result.Items) == 0 {
		t.Fatalf("%s not found in results", params.Symbol)
	}

	coin := result.Items[0]
	if coin.Symbol != params.Symbol {
		t.Errorf("Expected symbol %s, got %s", params.Symbol, coin.Symbol)
	}
	if coin.Name == "" {
		t.Error("Coin name is empty")
	}
	if len(coin.Networks) == 0 {
		t.Errorf("No networks available for %s", params.Symbol)
	}

	// Validate network structure
	for _, network := range coin.Networks {
		if network.Network == "" {
			t.Error("Network name is empty")
		}
		if network.WithdrawMin == "" {
			t.Error("WithdrawMin is empty")
		}
		if network.WithdrawFee == "" {
			t.Error("WithdrawFee is empty")
		}
		if network.Decimal < 0 {
			t.Error("Decimal should not be negative")
		}
	}
}

func TestGetCoinsMultipleSymbols(t *testing.T) {
	symbols := []string{"ETH", "USDT", "BNB"}
	successCount := 0
	totalNetworks := 0

	for _, symbol := range symbols {
		params := GetCoinsParams{
			Symbol: symbol,
		}

		result, err := GetCoins(params)
		if err != nil {
			t.Errorf("Failed to get coin %s: %v", symbol, err)
			continue
		}

		if len(result.Items) == 0 {
			t.Errorf("No results for symbol %s", symbol)
			continue
		}

		coin := result.Items[0]
		if coin.Symbol != symbol {
			t.Errorf("Expected symbol %s, got %s", symbol, coin.Symbol)
		}
		if coin.Name == "" {
			t.Errorf("Coin %s has empty name", symbol)
		}

		successCount++
		totalNetworks += len(coin.Networks)
	}

	if successCount != len(symbols) {
		t.Fatalf("❌ GetCoinsMultipleSymbols failed: only %d/%d symbols succeeded", successCount, len(symbols))
	}
}
