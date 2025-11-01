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

func TestGetAddresses(t *testing.T) {
	params := GetAddressesParams{
		Page:  1,
		Limit: 10,
	}

	result, err := GetAddresses(params)
	if err != nil {
		t.Skip("Skipping - API credentials may be invalid")
		return
	}

	// Assertions
	if result == nil {
		t.Fatal("❌ GetAddresses failed: result is nil")
	}
	if result.Page != params.Page {
		t.Errorf("Expected page %d, got %d", params.Page, result.Page)
	}
	if result.TotalItem < 0 {
		t.Errorf("TotalItem should not be negative, got %d", result.TotalItem)
	}
	if len(result.Items) > 0 {
		addr := result.Items[0]
		if addr.Symbol == "" {
			t.Error("First address symbol is empty")
		}
		if addr.Address == "" {
			t.Error("First address is empty")
		}
	}

	t.Logf("✅ GetAddresses passed: %d items, page %d/%d", len(result.Items), result.Page, result.TotalPage)
}

func TestGetAddressesWithFilter(t *testing.T) {
	params := GetAddressesParams{
		Page:   1,
		Limit:  5,
		Symbol: "KUB",
	}

	result, err := GetAddresses(params)
	if err != nil {
		t.Fatalf("❌ GetAddresses with filter failed: %v", err)
	}

	// Assertions
	if result == nil {
		t.Fatal("Result is nil")
	}
	for _, addr := range result.Items {
		if addr.Symbol != params.Symbol {
			t.Errorf("Expected symbol %s, got %s", params.Symbol, addr.Symbol)
		}
		if addr.Address == "" {
			t.Error("Address is empty")
		}
	}
}

func TestCreateAddress(t *testing.T) {
	req := CreateAddressRequest{
		Symbol:  "ETH",
		Network: "ETH",
	}

	result, err := CreateAddress(req)
	if err != nil {
		t.Fatalf("❌ CreateAddress failed: %v", err)
	}

	// Assertions
	if len(result) == 0 {
		t.Fatal("Result is empty")
	}
	addr := result[0]
	if addr.Symbol != req.Symbol {
		t.Errorf("Expected symbol %s, got %s", req.Symbol, addr.Symbol)
	}
	if addr.Network != req.Network {
		t.Errorf("Expected network %s, got %s", req.Network, addr.Network)
	}
	if addr.Address == "" {
		t.Error("Address is empty")
	}
}

func TestValidationCreateAddress(t *testing.T) {
	// Test missing symbol
	req := CreateAddressRequest{
		Network: "KUB",
	}
	_, err := CreateAddress(req)
	if err == nil {
		t.Fatal("❌ Expected error for missing symbol")
	}

	// Test missing network
	req = CreateAddressRequest{
		Symbol: "KUB",
	}
	_, err = CreateAddress(req)
	if err == nil {
		t.Fatal("❌ Expected error for missing network")
	}
}
