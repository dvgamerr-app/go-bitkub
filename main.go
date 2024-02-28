package main

import (
	"github.com/leekchan/accounting"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/touno-io/go-bitkub/helper"
)

var (
	symbolMoney string                = "₮"
	aNo         accounting.Accounting = accounting.Accounting{Precision: 2, Thousand: ",", Format: "%s%v"}
)

func init() {
	aNo.Symbol = symbolMoney
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	err := helper.LoadDotEnv()
	if err != nil {
		log.Error().Err(err)
	}

	err = helper.CheckEnvVars("BTK_APIKEY", "BTK_SECRETKEY")
	if err != nil {
		log.Error().Err(err)
	}
}
func main() {
	bal, err := QueryBalances()
	if err != nil {
		log.Error().Err(err)
	}
	log.Info().Msgf("%+v", bal)
}
