package main

import (
	"os"

	"github.com/alexflint/go-arg"
	"github.com/dvgamerr-app/go-bitkub/bitkub"
	"github.com/dvgamerr-app/go-bitkub/crypto"
	"github.com/dvgamerr-app/go-bitkub/helper"
	_ "github.com/dvgamerr-app/go-bitkub/market"
	"github.com/leekchan/accounting"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
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
}

func main() {
	arg.MustParse(&cli)
	bitkub.Initlizer(cli.Key, cli.Secret)

	addresses, err := crypto.GetAddresses(crypto.GetAddressesParams{
		Page:  1,
		Limit: 10,
	})
	helper.FatalError(err)
	log.Info().Interface("addresses", addresses).Send()

}
