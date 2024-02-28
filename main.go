package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/leekchan/accounting"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	symbolMoney string                = "₮"
	aNo         accounting.Accounting = accounting.Accounting{Precision: 2, Thousand: ",", Format: "%s%v"}
)

// checkEnvVars checks that all specified environment variables are set and not empty.
func checkEnvVars(envs ...string) {
	for _, v := range envs {
		if os.Getenv(v) == "" {
			log.Fatal().Msgf("Error: %s environment variable is not set\n", v)
			os.Exit(1)
		}
	}
}

// Load environment variables from .env
func loadEnv() {
	if _, err := os.Stat(".env"); err == nil {
		if err = godotenv.Load(); err != nil {
			log.Fatal().Err(err)
		}
	}
}

func init() {
	aNo.Symbol = symbolMoney
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	loadEnv()
	checkEnvVars("BTK_APIKEY", "BTK_SECRETKEY")
}
func main() {
	bal, err := QueryBalances()
	if err != nil {
		log.Error().Err(err)
	}
	log.Info().Msgf("%+v", bal)
}
