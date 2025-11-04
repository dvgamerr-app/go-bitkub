package cmd

import (
	"os"

	"github.com/dvgamerr-app/go-bitkub/bitkub"
	"github.com/dvgamerr-app/go-bitkub/utils"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	apiKey    string
	secretKey string
	debug     bool
	format    string
)

var rootCmd = &cobra.Command{
	Use:   "bitkub",
	Short: "Bitkub API CLI Tool",
	Long:  "A command-line interface for interacting with Bitkub API",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if format == "json" {
			log.Logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
		} else {
			log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
		}

		if debug {
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
		} else {
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
		}

		if err := bitkub.Initlizer(apiKey, secretKey); err != nil {
			log.Fatal().Err(err).Msg("Failed to initialize Bitkub client")
		}
	},
}

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	if err := utils.LoadDotEnv(); err != nil {
		log.Debug().Err(err).Msg("No .env file loaded")
	}

	rootCmd.PersistentFlags().StringVarP(&apiKey, "key", "k", os.Getenv("API_KEY"), "API Key")
	rootCmd.PersistentFlags().StringVarP(&secretKey, "secret", "s", os.Getenv("API_SECRET"), "API Secret")
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "Enable debug mode")
	rootCmd.PersistentFlags().StringVarP(&format, "format", "f", "text", "Output format (text, json)")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal().Err(err).Msg("Failed to execute command")
	}
}
