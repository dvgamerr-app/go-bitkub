package main

import (
	"github.com/dvgamerr-app/go-bitkub/market"
)

func QueryWallet() (*market.BitkubWallet, error) {
	data, err := market.GetWallet()
	if err != nil {
		return nil, err
	}

	for ccy, coin := range *data {
		if coin == 0.0 {
			delete(*data, ccy)
			continue
		}
	}

	return data, nil
}

func QueryCoins() ([]string, error) {
	data, err := market.GetWallet()
	if err != nil {
		return nil, err
	}

	keys := make([]string, 0, len(*data))
	for k := range *data {
		keys = append(keys, k)
	}
	return keys, nil
}
