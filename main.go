package main

import (
	"os"

	"github.com/leekchan/accounting"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/touno-io/go-bitkub/helper"
	"github.com/touno-io/go-bitkub/market"
)

var (
	symbolMoney string                = "₮"
	aNo         accounting.Accounting = accounting.Accounting{Precision: 2, Thousand: ",", Format: "%s%v"}
)

func init() {
	aNo.Symbol = symbolMoney

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(zerolog.TraceLevel)

	if err := helper.LoadDotEnv(); err != nil {
		log.Warn().Err(err)
	}

	if err := helper.CheckEnvVars("BTK_APIKEY", "BTK_SECRETKEY"); err != nil {
		log.Warn().Err(err)
	}

	// if os.Getenv("ENV") != "development" {
	// 	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	// }
}

func main() {
	// bal, err := QueryBalances()
	// helper.FatalError(err)
	// log.Trace().Interface("Balance", bal).Send()

	wal, err := market.GetWallet()
	helper.FatalError(err)
	log.Trace().Interface("wal", wal).Send()

	ord, err := market.GetMyOpenOrders("ankr")
	helper.FatalError(err)
	log.Trace().Interface("ord", ord).Send()

}
