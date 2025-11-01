package main

import (
	"fmt"
	"strconv"

	"github.com/dvgamerr-app/go-bitkub/market"
	"github.com/rs/zerolog/log"
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
			ticker, err := market.GetTicker(fmt.Sprintf("%s_THB", ccy))
			if err != nil {
				log.Warn().Err(err).Str("coin", ccy).Msg("Failed to fetch ticker")
				continue
			}

			if len(ticker) == 0 {
				log.Warn().Str("coin", ccy).Msg("No ticker data available")
				continue
			}

			// Parse last price from string
			lastPrice, err := strconv.ParseFloat(ticker[0].Last, 64)
			if err != nil {
				log.Warn().Err(err).Str("coin", ccy).Msg("Failed to parse ticker price")
				continue
			}
			rate = lastPrice
		}
		data.Total += (coin.Available + coin.Reserved) * rate
		data.Available += coin.Available * rate
	}

	return &data, nil
}
