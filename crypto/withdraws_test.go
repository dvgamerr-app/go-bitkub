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

func TestGetWithdraws(t *testing.T) {
	params := Withdraws{
		Pagination: Pagination{
			Page:  1,
			Limit: 10,
		},
	}

	result, err := GetWithdraws(params)
	if err != nil {
		t.Error("API credentials may be invalid")
		return
	}

	// Assertions
	if result == nil {
		t.Fatal("❌ GetWithdraws failed: result is nil")
	}
	if result.Page != params.Page {
		t.Errorf("Expected page %d, got %d", params.Page, result.Page)
	}
	if result.TotalItem < 0 {
		t.Errorf("TotalItem should not be negative, got %d", result.TotalItem)
	}

	// Validate items structure
	for _, withdraw := range result.Items {
		if withdraw.TxnID == "" {
			t.Error("Withdraw TxnID is empty")
		}
		if withdraw.Symbol == "" {
			t.Error("Withdraw symbol is empty")
		}
		if withdraw.Network == "" {
			t.Error("Withdraw network is empty")
		}
		if withdraw.Status == "" {
			t.Error("Withdraw status is empty")
		}
		if withdraw.Address == "" {
			t.Error("Withdraw address is empty")
		}
	}

}

func TestGetWithdrawsWithFilters(t *testing.T) {
	params := Withdraws{
		Pagination: Pagination{
			Page:  1,
			Limit: 5,
		},
		Status: "complete",
	}

	result, err := GetWithdraws(params)
	if err != nil {
		t.Fatalf("❌ GetWithdraws with filters failed: %v", err)
	}
	// Assertions
	if result == nil {
		t.Fatal("Result is nil")
	}
	for _, withdraw := range result.Items {
		if withdraw.Status != params.Status {
			t.Errorf("Expected status %s, got %s", params.Status, withdraw.Status)
		}
		if withdraw.Amount == "" {
			t.Error("Amount is empty")
		}
		if withdraw.Fee == "" {
			t.Error("Fee is empty")
		}
	}

}

func TestGetWithdrawsBySymbol(t *testing.T) {
	params := Withdraws{
		Pagination: Pagination{
			Page:  1,
			Limit: 5,
		},
		Symbol: "THB",
	}

	result, err := GetWithdraws(params)
	if err != nil {
		t.Fatalf("❌ GetWithdraws by symbol failed: %v", err)
	}
	// Assertions
	if result == nil {
		t.Fatal("Result is nil")
	}
	for _, withdraw := range result.Items {
		if withdraw.Symbol != params.Symbol {
			t.Errorf("Expected symbol %s, got %s", params.Symbol, withdraw.Symbol)
		}
	}
}

func TestCreateWithdraw(t *testing.T) {
	t.Skip("requires is_withdraw permission and real transaction")

	req := CreateWithdrawRequest{
		Symbol:  "RDNT",
		Amount:  "2.00000000",
		Address: "0xDaCd17d1E77604aaFB6e47F5Ffa1F7E35F83fDa7",
		Network: "ARB",
	}

	result, err := CreateWithdraw(req)
	if err != nil {
		t.Fatalf("❌ CreateWithdraw failed: %v", err)
	}

	// Assertions
	if result.TxnID == "" {
		t.Error("TxnID is empty")
	}
	if result.Symbol != req.Symbol {
		t.Errorf("Expected symbol %s, got %s", req.Symbol, result.Symbol)
	}
	if result.Network != req.Network {
		t.Errorf("Expected network %s, got %s", req.Network, result.Network)
	}
	if result.Address != req.Address {
		t.Errorf("Expected address %s, got %s", req.Address, result.Address)
	}
	if result.Amount == "" {
		t.Error("Amount is empty")
	}
	if result.Fee == "" {
		t.Error("Fee is empty")
	}

	t.Logf("✅ CreateWithdraw passed: TxnID %s created", result.TxnID)
}

func TestCreateWithdrawValidation(t *testing.T) {
	// Test missing symbol
	req := CreateWithdrawRequest{
		Amount:  "2.00000000",
		Address: "0xDaCd17d1E77604aaFB6e47F5Ffa1F7E35F83fDa7",
		Network: "ARB",
	}
	_, err := CreateWithdraw(req)
	if err == nil {
		t.Error("Expected error for missing symbol")
	}

	// Test missing amount
	req = CreateWithdrawRequest{
		Symbol:  "RDNT",
		Address: "0xDaCd17d1E77604aaFB6e47F5Ffa1F7E35F83fDa7",
		Network: "ARB",
	}
	_, err = CreateWithdraw(req)
	if err == nil {
		t.Error("Expected error for missing amount")
	}

	// Test missing address
	req = CreateWithdrawRequest{
		Symbol:  "RDNT",
		Amount:  "2.00000000",
		Network: "ARB",
	}
	_, err = CreateWithdraw(req)
	if err == nil {
		t.Error("Expected error for missing address")
	}

	// Test missing network
	req = CreateWithdrawRequest{
		Symbol:  "RDNT",
		Amount:  "2.00000000",
		Address: "0xDaCd17d1E77604aaFB6e47F5Ffa1F7E35F83fDa7",
	}
	_, err = CreateWithdraw(req)
	if err == nil {
		t.Error("Expected error for missing network")
	}
}
