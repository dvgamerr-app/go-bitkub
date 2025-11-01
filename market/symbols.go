package market

import (
	"fmt"

	"github.com/dvgamerr-app/go-bitkub/bitkub"
)

type Symbol struct {
	BaseAsset             string  `json:"base_asset"`
	BaseAssetScale        int     `json:"base_asset_scale"`
	BuyPriceGapAsPercent  float64 `json:"buy_price_gap_as_percent"`
	CreatedAt             string  `json:"created_at"`
	Description           string  `json:"description"`
	FreezeBuy             bool    `json:"freeze_buy"`
	FreezeCancel          bool    `json:"freeze_cancel"`
	FreezeSell            bool    `json:"freeze_sell"`
	MarketSegment         string  `json:"market_segment"`
	MinQuoteSize          float64 `json:"min_quote_size"`
	ModifiedAt            string  `json:"modified_at"`
	Name                  string  `json:"name"`
	PairingID             int     `json:"pairing_id"`
	PriceScale            int     `json:"price_scale"`
	PriceStep             string  `json:"price_step"`
	QuantityScale         int     `json:"quantity_scale"`
	QuantityStep          string  `json:"quantity_step"`
	QuoteAsset            string  `json:"quote_asset"`
	QuoteAssetScale       int     `json:"quote_asset_scale"`
	SellPriceGapAsPercent float64 `json:"sell_price_gap_as_percent"`
	Status                string  `json:"status"`
	Symbol                string  `json:"symbol"`
	Source                string  `json:"source"`
}

type SymbolsResponse struct {
	Error  int      `json:"error"`
	Result []Symbol `json:"result"`
}

func GetSymbols() ([]Symbol, error) {
	var result SymbolsResponse

	if err := bitkub.FetchNonSecure("GET", "/v3/market/symbols", nil, &result); err != nil {
		return nil, err
	}

	if result.Error != 0 {
		errMsg, exists := bitkub.ErrorCode[result.Error]
		if !exists {
			errMsg = "Unknown error"
		}
		return nil, fmt.Errorf("[error %d] %s", result.Error, errMsg)
	}

	return result.Result, nil
}
