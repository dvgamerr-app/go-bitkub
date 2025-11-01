package user

import (
	"testing"
)

func TestGetCoinConvertHistory(t *testing.T) {
	params := CoinConvertHistoryParams{
		P:   1,
		Lmt: 10,
	}
	GetCoinConvertHistory(params)
}
