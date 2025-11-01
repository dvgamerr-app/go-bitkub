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

func TestGetCompensations(t *testing.T) {
	params := GetCompensationsParams{
		Page:  1,
		Limit: 10,
	}

	result, err := GetCompensations(params)
	if err != nil {
		t.Error("API credentials may be invalid")
		return
	}

	// Assertions
	if result == nil {
		t.Fatal("❌ GetCompensations failed: result is nil")
	}
	if result.Page != params.Page {
		t.Errorf("Expected page %d, got %d", params.Page, result.Page)
	}
	if result.TotalItem < 0 {
		t.Errorf("TotalItem should not be negative, got %d", result.TotalItem)
	}

	// Validate items structure if exists
	for _, comp := range result.Items {
		if comp.TxnID == "" {
			t.Error("Compensation TxnID is empty")
		}
		if comp.Symbol == "" {
			t.Error("Compensation symbol is empty")
		}
		if comp.Type == "" {
			t.Error("Compensation type is empty")
		}
		if comp.Status == "" {
			t.Error("Compensation status is empty")
		}
	}
}

func TestGetCompensationsByType(t *testing.T) {
	params := GetCompensationsParams{
		Page:  1,
		Limit: 5,
		Type:  "COMPENSATE",
	}

	result, err := GetCompensations(params)
	if err != nil {
		t.Fatalf("❌ GetCompensations (COMPENSATE) failed: %v", err)
	}

	// Assertions
	if result == nil {
		t.Fatal("Result is nil")
	}
	for _, comp := range result.Items {
		if comp.Type != params.Type {
			t.Errorf("Expected type %s, got %s", params.Type, comp.Type)
		}
		if comp.Amount == "" {
			t.Error("Amount is empty")
		}
		if comp.Status == "" {
			t.Error("Status is empty")
		}
	}
}

func TestGetCompensationsByStatus(t *testing.T) {
	params := GetCompensationsParams{
		Page:   1,
		Limit:  5,
		Status: "COMPLETED",
	}

	result, err := GetCompensations(params)
	if err != nil {
		t.Fatalf("❌ GetCompensations (COMPLETED) failed: %v", err)
	}

	// Assertions
	if result == nil {
		t.Fatal("Result is nil")
	}
	for _, comp := range result.Items {
		if comp.Status != params.Status {
			t.Errorf("Expected status %s, got %s", params.Status, comp.Status)
		}
	}
}

func TestGetCompensationsBySymbol(t *testing.T) {
	params := GetCompensationsParams{
		Page:   1,
		Limit:  5,
		Symbol: "XRP",
	}

	result, err := GetCompensations(params)
	if err != nil {
		t.Fatalf("❌ GetCompensations by symbol failed: %v", err)
	}

	// Assertions
	if result == nil {
		t.Fatal("Result is nil")
	}
	for _, comp := range result.Items {
		if comp.Symbol != params.Symbol {
			t.Errorf("Expected symbol %s, got %s", params.Symbol, comp.Symbol)
		}
		if comp.Amount == "" {
			t.Error("Amount is empty")
		}
		if comp.Type == "" {
			t.Error("Type is empty")
		}
	}
}
