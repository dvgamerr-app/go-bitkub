package main

import (
	"github.com/dvgamerr-app/go-bitkub/market"
)

type BitkubBalances struct {
	Total     float64
	Available float64
	Coins     map[string]market.Balance
}

func QueryBalances() (*BitkubBalances, error) {
	var (
		data BitkubBalances
		err  error
	)
	data.Coins, err = market.GetBalances()
	if err != nil {
		return nil, err
	}

	for ccy, coin := range data.Coins {
		if coin.Available == 0 && coin.Reserved == 0 {
			delete(data.Coins, ccy)
			continue
		}

		rate := 1.0
		if ccy != "THB" {
			ticker, err := market.GetMarketTicker(ccy)
			if err != nil {
				return nil, err
			}
			rate = ticker.Last
		}
		data.Total += (coin.Available + coin.Reserved) * rate
		data.Available += coin.Available * rate
	}

	return &data, nil
}
