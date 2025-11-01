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

func TestGetDeposits(t *testing.T) {
	params := Deposits{
		Pagination: Pagination{
			Page:  1,
			Limit: 10,
		},
	}

	result, err := GetDeposits(params)
	if err != nil {
		t.Error("API credentials may be invalid")
		return
	}

	// Assertions
	if result == nil {
		t.Fatal("❌ GetDeposits failed: result is nil")
	}
	if result.Page != params.Page {
		t.Errorf("Expected page %d, got %d", params.Page, result.Page)
	}
	if result.TotalItem < 0 {
		t.Errorf("TotalItem should not be negative, got %d", result.TotalItem)
	}

	// Validate items structure
	for _, deposit := range result.Items {
		if deposit.Hash == "" {
			t.Error("Deposit hash is empty")
		}
		if deposit.Symbol == "" {
			t.Error("Deposit symbol is empty")
		}
		if deposit.Status == "" {
			t.Error("Deposit status is empty")
		}
		if deposit.Confirmations < 0 {
			t.Error("Confirmations should not be negative")
		}
	}
}

func TestGetDepositsWithFilters(t *testing.T) {
	params := Deposits{
		Pagination: Pagination{
			Page:  1,
			Limit: 5,
		},
		Status: "complete",
	}

	result, err := GetDeposits(params)
	if err != nil {
		t.Fatalf("❌ GetDeposits with filters failed: %v", err)
	}

	// Assertions
	if result == nil {
		t.Fatal("Result is nil")
	}
	for _, deposit := range result.Items {
		if deposit.Status != params.Status {
			t.Errorf("Expected status %s, got %s", params.Status, deposit.Status)
		}
		if deposit.Amount == "" {
			t.Error("Amount is empty")
		}
		if deposit.Symbol == "" {
			t.Error("Symbol is empty")
		}
	}
}

func TestGetDepositsBySymbol(t *testing.T) {
	params := Deposits{
		Pagination: Pagination{
			Page:  1,
			Limit: 5,
		},
		Symbol: "BTC",
	}

	result, err := GetDeposits(params)
	if err != nil {
		t.Fatalf("❌ GetDeposits by symbol failed: %v", err)
	}

	// Assertions
	if result == nil {
		t.Fatal("Result is nil")
	}
	for _, deposit := range result.Items {
		if deposit.Symbol != params.Symbol {
			t.Errorf("Expected symbol %s, got %s", params.Symbol, deposit.Symbol)
		}
	}
}
