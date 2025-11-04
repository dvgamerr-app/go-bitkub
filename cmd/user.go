package cmd

import (
	"github.com/dvgamerr-app/go-bitkub/user"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "User API commands",
	Long:  "Commands for user information and limits",
}

var userLimitsCmd = &cobra.Command{
	Use:   "limits",
	Short: "Get user limits",
	Run: func(cmd *cobra.Command, args []string) {
		limits, err := user.GetUserLimits()
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to get limits")
		}

		log.Info().
			Float64("crypto_deposit", limits.Limits.Crypto.Deposit).
			Float64("crypto_withdraw", limits.Limits.Crypto.Withdraw).
			Float64("fiat_deposit", limits.Limits.Fiat.Deposit).
			Float64("fiat_withdraw", limits.Limits.Fiat.Withdraw).
			Msg("Limits")

		log.Info().
			Float64("crypto_deposit", limits.Usage.Crypto.Deposit).
			Float64("crypto_withdraw", limits.Usage.Crypto.Withdraw).
			Float64("fiat_deposit", limits.Usage.Fiat.Deposit).
			Float64("fiat_withdraw", limits.Usage.Fiat.Withdraw).
			Msg("Usage")
	},
}

var userTradingCreditsCmd = &cobra.Command{
	Use:   "credits",
	Short: "Get trading credits",
	Run: func(cmd *cobra.Command, args []string) {
		credits, err := user.GetTradingCredits()
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to get trading credits")
		}
		log.Info().Float64("credits", credits).Msg("Trading Credits")
	},
}

var userCoinConvertHistoryCmd = &cobra.Command{
	Use:   "coin-convert-history",
	Short: "Get coin convert history",
	Run: func(cmd *cobra.Command, args []string) {
		params := user.CoinHistoryParams{}

		page, _ := cmd.Flags().GetInt("page")
		limit, _ := cmd.Flags().GetInt("limit")

		if page > 0 {
			params.Page = page
		}
		if limit > 0 {
			params.Limit = limit
		}

		history, err := user.GetCoinConvertHistory(params)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to get coin convert history")
		}

		log.Info().
			Int("page", history.Pagination.Page).
			Int("last", history.Pagination.Last).
			Msg("Pagination")

		for _, item := range history.Result {
			log.Info().
				Str("transaction_id", item.TransactionID).
				Str("from_currency", item.FromCurrency).
				Str("amount", item.Amount).
				Str("trading_fee_received", item.TradingFeeReceived).
				Str("status", item.Status).
				Int64("timestamp", item.Timestamp).
				Msg("Convert")
		}
	},
}

func init() {
	rootCmd.AddCommand(userCmd)

	userCmd.AddCommand(userLimitsCmd)
	userCmd.AddCommand(userTradingCreditsCmd)
	userCmd.AddCommand(userCoinConvertHistoryCmd)

	userCoinConvertHistoryCmd.Flags().IntP("page", "p", 1, "Page number")
	userCoinConvertHistoryCmd.Flags().IntP("limit", "l", 20, "Limit per page")
}
