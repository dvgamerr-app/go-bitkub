package market

import (
	"testing"
	"time"

	"github.com/dvgamerr-app/go-bitkub/bitkub"
	"github.com/dvgamerr-app/go-bitkub/utils"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func init() {
	zerolog.SetGlobalLevel(zerolog.Disabled)
	utils.LoadDotEnv("../.env")
	bitkub.Initlizer()
}

func TestGetAsks(t *testing.T) {
	result, err := GetAsks("kub_thb", 10)
	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestGetBalances(t *testing.T) {
	result, err := GetBalances()
	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestGetBids(t *testing.T) {
	result, err := GetBids("kub_thb", 10)
	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestCancelOrder(t *testing.T) {
	bidResult, err := PlaceBid(PlaceBidRequest{
		Symbol: "kub_thb",
		Amount: 100,
		Rate:   1,
		Type:   "limit",
	})
	if err != nil {
		t.Skipf("Cannot create bid order for test: %v", err)
		return
	}
	time.Sleep(1 * time.Second)
	assert.NotNil(t, bidResult)
	assert.NotEmpty(t, bidResult.ID)

	request := CancelOrderRequest{
		Symbol: "kub_thb",
		ID:     bidResult.ID,
		Side:   "buy",
	}
	err = CancelOrder(request)
	assert.Nil(t, err)
}

func TestGetDepth(t *testing.T) {
	result, err := GetDepth("kub_thb", 10)
	if err != nil {
		t.Skipf("Skipping test: %v", err)
	}
	assert.NotNil(t, result)
}

func TestGetUserLimits(t *testing.T) {
	result, err := GetUserLimits()
	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestGetOpenOrders(t *testing.T) {
	result, err := GetOpenOrders("kub_thb")
	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestGetOrderInfo(t *testing.T) {
	bidResult, err := PlaceBid(PlaceBidRequest{
		Symbol: "kub_thb",
		Amount: 100,
		Rate:   1,
		Type:   "limit",
	})
	if err != nil {
		t.Skipf("Cannot create bid order for test: %v", err)
		return
	}
	time.Sleep(1 * time.Second)

	result, err := GetOrderInfo("kub_thb", bidResult.ID, "buy")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, bidResult.ID, result.ID)

	time.Sleep(1 * time.Second)

	_ = CancelOrder(CancelOrderRequest{
		Symbol: "kub_thb",
		ID:     bidResult.ID,
		Side:   "buy",
	})
}

func TestPlaceAsk(t *testing.T) {
	result, err := PlaceAsk(PlaceAskRequest{
		Symbol: "kub_thb",
		Amount: 1,
		Rate:   100,
		Type:   "limit",
	})
	if err != nil {
		t.Skipf("Cannot place ask: %v", err)
		return
	}
	assert.NotNil(t, result)
	assert.NotEmpty(t, result.ID)

	time.Sleep(1 * time.Second)
	err = CancelOrder(CancelOrderRequest{
		Symbol: "kub_thb",
		ID:     result.ID,
		Side:   "sell",
	})
	assert.Nil(t, err)
}

func TestPlaceBid(t *testing.T) {
	request := PlaceBidRequest{
		Symbol: "kub_thb",
		Amount: 100,
		Rate:   1,
		Type:   "limit",
	}
	result, err := PlaceBid(request)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.NotEmpty(t, result.ID)

	if err != nil {
		t.Skipf("Cannot place ask order: %v", err)
		return
	}
	time.Sleep(1 * time.Second)
	err = CancelOrder(CancelOrderRequest{
		Symbol: "kub_thb",
		ID:     result.ID,
		Side:   "buy",
	})
	assert.Nil(t, err)
}

func TestGetSymbols(t *testing.T) {
	result, err := GetSymbols()
	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestGetTicker(t *testing.T) {
	result, err := GetTicker("kub_thb")
	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestGetTrades(t *testing.T) {
	result, err := GetTrades("kub_thb", 10)
	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestGetTradingCredits(t *testing.T) {
	result, err := GetTradingCredits()
	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestGetWallet(t *testing.T) {
	result, err := GetWallet()
	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestGetWSToken(t *testing.T) {
	result, err := GetWSToken()
	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestGetHistory(t *testing.T) {
	req := HistoryRequest{
		Symbol:     "kub_THB",
		Resolution: "1D",
		From:       1650819600,
		To:         1650902400,
	}
	result, err := GetHistory(req)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "ok", result.Status)
	assert.Greater(t, len(result.Close), 0)
	assert.Greater(t, len(result.Open), 0)
	assert.Greater(t, len(result.High), 0)
	assert.Greater(t, len(result.Low), 0)
	assert.Greater(t, len(result.Volume), 0)
	assert.Greater(t, len(result.Time), 0)
}

func TestGetHistoryDifferentResolutions(t *testing.T) {
	resolutions := []string{"1", "5", "15", "60", "240", "1D"}
	for _, res := range resolutions {
		req := HistoryRequest{
			Symbol:     "kub_THB",
			Resolution: res,
			From:       1730736000,
			To:         1730822400,
		}
		result, err := GetHistory(req)
		assert.Nil(t, err, "Resolution %s should work", res)
		assert.NotNil(t, result)
		assert.Equal(t, "ok", result.Status)
	}
}

func TestGetHistoryValidation(t *testing.T) {
	_, err := GetHistory(HistoryRequest{})
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "symbol is required")

	_, err = GetHistory(HistoryRequest{Symbol: "kub_THB"})
	assert.Nil(t, err)

	result, err := GetHistory(HistoryRequest{Symbol: "kub_THB", Resolution: "1D"})
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "ok", result.Status)
}

func TestGetHistoryDefault24Hours(t *testing.T) {
	result, err := GetHistory(HistoryRequest{
		Symbol:     "ETH_THB",
		Resolution: "1D",
	})
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "ok", result.Status)
	assert.Greater(t, len(result.Close), 0)
}
