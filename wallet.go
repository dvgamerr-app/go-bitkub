package main

import (
	"github.com/touno-io/go-bitkub/market"
)

type BitkubWallet map[string]float64

func QueryWallet() (*BitkubWallet, error) {
	var (
		data BitkubWallet
		err  error
	)
	data, err = market.GetWallet()
	if err != nil {
		return nil, err
	}

	for ccy, coin := range data {
		if coin == 0.0 {
			delete(data, ccy)
			continue
		}
	}

	return &data, nil
}
