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

func TestGetUserLimits(t *testing.T) {
	result, err := GetUserLimits()
	if err != nil {
		t.Error("API credentials may be invalid")
		return
	}

	// Assertions
	if result == nil {
		t.Fatal("❌ GetUserLimits failed: result is nil")
	}
	if result.Rate < 0 {
		t.Error("Rate level should not be negative")
	}

	// Validate crypto limits
	if result.Limits.Crypto.Deposit < 0 {
		t.Error("Crypto deposit limit should not be negative")
	}
	if result.Limits.Crypto.Withdraw < 0 {
		t.Error("Crypto withdraw limit should not be negative")
	}

	// Validate fiat limits
	if result.Limits.Fiat.Deposit < 0 {
		t.Error("Fiat deposit limit should not be negative")
	}
	if result.Limits.Fiat.Withdraw < 0 {
		t.Error("Fiat withdraw limit should not be negative")
	}

	// Validate usage percentages
	if result.Usage.Crypto.DepositPercentage < 0 || result.Usage.Crypto.DepositPercentage > 100 {
		t.Errorf("Crypto deposit percentage should be 0-100, got %.2f", result.Usage.Crypto.DepositPercentage)
	}
	if result.Usage.Crypto.WithdrawPercentage < 0 || result.Usage.Crypto.WithdrawPercentage > 100 {
		t.Errorf("Crypto withdraw percentage should be 0-100, got %.2f", result.Usage.Crypto.WithdrawPercentage)
	}
}
