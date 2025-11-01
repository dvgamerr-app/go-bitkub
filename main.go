package main

import (
	"os"

	"github.com/alexflint/go-arg"
	"github.com/dvgamerr-app/go-bitkub/bitkub"
	"github.com/dvgamerr-app/go-bitkub/helper"
	"github.com/dvgamerr-app/go-bitkub/market"
	"github.com/leekchan/accounting"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	aNo accounting.Accounting = accounting.Accounting{Precision: 2, Thousand: ",", Format: "%s%v"}
)

var cli struct {
	Key    string `arg:"--key,-K" help:"optimization level"`
	Secret string `arg:"--secret, -S" help:"optimization level"`
}

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(zerolog.TraceLevel)

	if err := helper.LoadDotEnv(); err != nil {
		log.Warn().Err(err).Msg("Failed to load .env file")
	}
}

func main() {
	arg.MustParse(&cli)
	bitkub.Initlizer(cli.Key, cli.Secret)

	log.Info().Msg("📊 Fetching balances from all wallets...")
	balances, err := QueryBalances()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to fetch balances")
	}

	log.Info().
		Str("total", aNo.FormatMoney(balances.Total)).
		Str("available", aNo.FormatMoney(balances.Available)).
		Msg("💰 Total Balance Summary")

	log.Info().Msg("📋 Balance Details:")
	for ccy, balance := range balances.Coins {
		if ccy == "THB" {
			log.Info().
				Str("coin", ccy).
				Str("available", aNo.FormatMoney(balance.Available)).
				Str("reserved", aNo.FormatMoney(balance.Reserved)).
				Msg("  ")
		} else {
			log.Info().
				Str("coin", ccy).
				Float64("available", balance.Available).
				Float64("reserved", balance.Reserved).
				Msg("  ")
		}
	}

	log.Info().Msg("📝 Fetching open orders...")
	ordersByCoin := make(map[string][]market.OpenOrder)

	for ccy := range balances.Coins {
		if ccy == "THB" {
			continue
		}

		orders, err := market.GetMyOpenOrders(ccy)
		if err != nil {
			log.Warn().Err(err).Str("coin", ccy).Msg("Failed to fetch orders")
			continue
		}

		if len(orders) > 0 {
			ordersByCoin[ccy] = orders
		}
	}

	if len(ordersByCoin) == 0 {
		log.Info().Msg("✅ No open orders")
	} else {
		log.Info().Int("coins", len(ordersByCoin)).Msg("🔄 Open Orders by Coin:")

		for ccy, orders := range ordersByCoin {
			log.Info().
				Str("coin", ccy).
				Int("count", len(orders)).
				Msg("  ")

			buyOrders := 0
			sellOrders := 0
			limitOrders := 0
			marketOrders := 0

			for _, order := range orders {
				switch order.Side {
				case "buy":
					buyOrders++
				case "sell":
					sellOrders++
				}

				switch order.Type {
				case "limit":
					limitOrders++
				case "market":
					marketOrders++
				}

				log.Debug().
					Str("id", order.ID).
					Str("side", order.Side).
					Str("type", order.Type).
					Str("rate", order.Rate).
					Str("amount", order.Amount).
					Msg("    ")
			}

			log.Info().
				Str("coin", ccy).
				Int("buy", buyOrders).
				Int("sell", sellOrders).
				Int("limit", limitOrders).
				Int("market", marketOrders).
				Msg("    Summary")
		}
	}

	log.Info().Msg("✨ Done!")
}
