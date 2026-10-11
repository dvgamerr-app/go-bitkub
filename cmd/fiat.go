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
		runPaged(cmd, "Failed to get accounts", func(page int, limit int) (int, int, []fiat.BankAccount, error) {
			res, err := fiat.GetAccounts(fiat.AccountsParams{Page: page, Limit: limit})
			if err != nil {
				return 0, 0, nil, err
			}
			return res.Pagination.Page, res.Pagination.Last, res.Result, nil
		}, func(acc fiat.BankAccount) {
			log.Info().
				Str("id", acc.ID).
				Str("bank", acc.Bank).
				Str("name", acc.Name).
				Int64("time", acc.Time).
				Msg("Account")
		})
	},
}

var fiatDepositHistoryCmd = &cobra.Command{
	Use:   "deposit-history",
	Short: "Get deposit history",
	Run: func(cmd *cobra.Command, args []string) {
		runPaged(cmd, "Failed to get deposit history", func(page int, limit int) (int, int, []fiat.DepositHistory, error) {
			res, err := fiat.GetDepositHistory(fiat.DepositHistoryParams{Page: page, Limit: limit})
			if err != nil {
				return 0, 0, nil, err
			}
			return res.Pagination.Page, res.Pagination.Last, res.Result, nil
		}, func(dep fiat.DepositHistory) {
			log.Info().
				Str("txn_id", dep.TxnID).
				Str("currency", dep.Currency).
				Float64("amount", dep.Amount).
				Str("status", dep.Status).
				Int64("time", dep.Time).
				Msg("Deposit")
		})
	},
}

var fiatWithdrawHistoryCmd = &cobra.Command{
	Use:   "withdraw-history",
	Short: "Get withdraw history",
	Run: func(cmd *cobra.Command, args []string) {
		runPaged(cmd, "Failed to get withdraw history", func(page int, limit int) (int, int, []fiat.WithdrawHistory, error) {
			res, err := fiat.GetWithdrawHistory(fiat.WithdrawHistoryParams{Page: page, Limit: limit})
			if err != nil {
				return 0, 0, nil, err
			}
			return res.Pagination.Page, res.Pagination.Last, res.Result, nil
		}, func(wth fiat.WithdrawHistory) {
			log.Info().
				Str("txn_id", wth.TxnID).
				Str("currency", wth.Currency).
				Str("amount", wth.Amount).
				Float64("fee", wth.Fee).
				Str("status", wth.Status).
				Int64("time", wth.Time).
				Msg("Withdraw")
		})
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
