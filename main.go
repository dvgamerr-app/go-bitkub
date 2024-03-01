package main

import (
	"os"

	"github.com/alexflint/go-arg"
	"github.com/leekchan/accounting"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/touno-io/go-bitkub/bitkub"
	"github.com/touno-io/go-bitkub/helper"
	"github.com/touno-io/go-bitkub/market"
)

var (
	symbolMoney string                = "₮"
	aNo         accounting.Accounting = accounting.Accounting{Precision: 2, Thousand: ",", Format: "%s%v"}
)

var cli struct {
	Key    string `arg:"--key,-K" help:"optimization level"`
	Secret string `arg:"--secret, -S" help:"optimization level"`
}

func init() {
	aNo.Symbol = symbolMoney

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(zerolog.TraceLevel)

	if err := helper.LoadDotEnv(); err != nil {
		log.Warn().Err(err)
	}
	// if os.Getenv("ENV") != "development" {
	// 	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	// }
}

func main() {
	arg.MustParse(&cli)
	bitkub.Initlizer(cli.Key, cli.Secret)

	// bal, err := QueryBalances()
	// helper.FatalError(err)
	// log.Trace().Interface("Balance", bal).Send()

	ord, err := market.GetUserLimits()
	helper.FatalError(err)
	log.Trace().Interface("ord", ord).Send()

}
