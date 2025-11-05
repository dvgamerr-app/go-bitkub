package cmd

import (
	"fmt"
	"os"

	"github.com/dvgamerr-app/go-bitkub/bitkub"
	"github.com/dvgamerr-app/go-bitkub/utils"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	apiKey      string
	secretKey   string
	debug       bool
	format      string
	isFirstLine bool = true
)

func output(data any, isLastLine ...bool) {
	switch format {
	case "json":
		if isFirstLine {
			fmt.Print("[")
			isFirstLine = false
		} else {
			fmt.Print(",")
		}

		jsonStr, err := stdJson.MarshalToString(data)
		if err != nil {
			log.Error().Err(err).Msg("Failed to marshal JSON")
			return
		}

		if len(isLastLine) > 0 && isLastLine[0] {
			fmt.Print(jsonStr)
			fmt.Print("]")
			isFirstLine = true
		} else {
			fmt.Print(jsonStr)
		}
	case "jsonl":
		jsonStr, err := stdJson.MarshalToString(data)
		if err != nil {
			log.Error().Err(err).Msg("Failed to marshal JSON")
			return
		}
		fmt.Println(jsonStr)
	default:
		switch v := data.(type) {
		case map[string]any:
			event := log.Info()
			for key, val := range v {
				event = event.Interface(key, val)
			}
			event.Send()
		default:
			log.Info().Interface("data", data).Send()
		}
	}
}

var rootCmd = &cobra.Command{
	Use:   "bitkub",
	Short: "Bitkub API CLI Tool",
	Long:  "A command-line interface for interacting with Bitkub API",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if format == "json" || format == "jsonl" {
			zerolog.SetGlobalLevel(zerolog.Disabled)
		} else {
			log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
			if debug {
				zerolog.SetGlobalLevel(zerolog.DebugLevel)
			} else {
				zerolog.SetGlobalLevel(zerolog.InfoLevel)
			}
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
	rootCmd.PersistentFlags().StringVar(&secretKey, "secret", os.Getenv("API_SECRET"), "API Secret")
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "Enable debug mode")
	rootCmd.PersistentFlags().StringVarP(&format, "format", "f", "text", "Output format (text, json, jsonl)")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal().Err(err).Msg("Failed to execute command")
	}
}
