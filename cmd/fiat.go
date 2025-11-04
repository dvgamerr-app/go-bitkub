package cmd

import (
	"strconv"

	"github.com/dvgamerr-app/go-bitkub/fiat"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var fiatCmd = &cobra.Command{
	Use:   "fiat",
	Short: "Fiat API commands",
	Long:  "Commands for fiat deposits, withdrawals, and bank accounts",
}

var fiatAccountsCmd = &cobra.Command{
	Use:   "accounts",
	Short: "Get bank accounts",
	Run: func(cmd *cobra.Command, args []string) {
		params := fiat.AccountsParams{}

		page, _ := cmd.Flags().GetInt("page")
		limit, _ := cmd.Flags().GetInt("limit")

		if page > 0 {
			params.Page = page
		}
		if limit > 0 {
			params.Limit = limit
		}

		accounts, err := fiat.GetAccounts(params)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to get accounts")
		}

		log.Info().
			Int("page", accounts.Pagination.Page).
			Int("last", accounts.Pagination.Last).
			Msg("Pagination")

		for _, acc := range accounts.Result {
			log.Info().
				Str("id", acc.ID).
				Str("bank", acc.Bank).
				Str("name", acc.Name).
				Int64("time", acc.Time).
				Msg("Account")
		}
	},
}

var fiatDepositHistoryCmd = &cobra.Command{
	Use:   "deposit-history",
	Short: "Get deposit history",
	Run: func(cmd *cobra.Command, args []string) {
		params := fiat.DepositHistoryParams{}

		page, _ := cmd.Flags().GetInt("page")
		limit, _ := cmd.Flags().GetInt("limit")

		if page > 0 {
			params.Page = page
		}
		if limit > 0 {
			params.Limit = limit
		}

		deposits, err := fiat.GetDepositHistory(params)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to get deposit history")
		}

		log.Info().
			Int("page", deposits.Pagination.Page).
			Int("last", deposits.Pagination.Last).
			Msg("Pagination")

		for _, dep := range deposits.Result {
			log.Info().
				Str("txn_id", dep.TxnID).
				Str("currency", dep.Currency).
				Float64("amount", dep.Amount).
				Str("status", dep.Status).
				Int64("time", dep.Time).
				Msg("Deposit")
		}
	},
}

var fiatWithdrawHistoryCmd = &cobra.Command{
	Use:   "withdraw-history",
	Short: "Get withdraw history",
	Run: func(cmd *cobra.Command, args []string) {
		params := fiat.WithdrawHistoryParams{}

		page, _ := cmd.Flags().GetInt("page")
		limit, _ := cmd.Flags().GetInt("limit")

		if page > 0 {
			params.Page = page
		}
		if limit > 0 {
			params.Limit = limit
		}

		withdraws, err := fiat.GetWithdrawHistory(params)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to get withdraw history")
		}

		log.Info().
			Int("page", withdraws.Pagination.Page).
			Int("last", withdraws.Pagination.Last).
			Msg("Pagination")

		for _, wth := range withdraws.Result {
			log.Info().
				Str("txn_id", wth.TxnID).
				Str("currency", wth.Currency).
				Str("amount", wth.Amount).
				Float64("fee", wth.Fee).
				Str("status", wth.Status).
				Int64("time", wth.Time).
				Msg("Withdraw")
		}
	},
}

var fiatWithdrawCmd = &cobra.Command{
	Use:   "withdraw [bank-account-id] [amount]",
	Short: "Create fiat withdrawal",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		amount, err := strconv.ParseFloat(args[1], 64)
		if err != nil {
			log.Fatal().Err(err).Msg("Invalid amount")
		}

		req := fiat.WithdrawRequest{
			ID:     args[0],
			Amount: amount,
		}

		result, err := fiat.Withdraw(req)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to create withdraw")
		}

		log.Info().
			Str("txn_id", result.Transaction).
			Str("account", result.Account).
			Str("currency", result.Currency).
			Float64("amount", result.Amount).
			Float64("fee", result.Fee).
			Float64("receive", result.Receive).
			Msg("Withdraw Created")
	},
}

func init() {
	rootCmd.AddCommand(fiatCmd)

	fiatCmd.AddCommand(fiatAccountsCmd)
	fiatCmd.AddCommand(fiatDepositHistoryCmd)
	fiatCmd.AddCommand(fiatWithdrawHistoryCmd)
	fiatCmd.AddCommand(fiatWithdrawCmd)

	fiatAccountsCmd.Flags().IntP("page", "p", 1, "Page number")
	fiatAccountsCmd.Flags().IntP("limit", "l", 20, "Limit per page")

	fiatDepositHistoryCmd.Flags().IntP("page", "p", 1, "Page number")
	fiatDepositHistoryCmd.Flags().IntP("limit", "l", 20, "Limit per page")

	fiatWithdrawHistoryCmd.Flags().IntP("page", "p", 1, "Page number")
	fiatWithdrawHistoryCmd.Flags().IntP("limit", "l", 20, "Limit per page")
}
