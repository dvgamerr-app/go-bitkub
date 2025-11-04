package cmd

import (
	"github.com/dvgamerr-app/go-bitkub/crypto"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var cryptoCmd = &cobra.Command{
	Use:   "crypto",
	Short: "Crypto API commands",
	Long:  "Commands for crypto deposits, withdrawals, and addresses",
}

var cryptoCoinsCmd = &cobra.Command{
	Use:   "coins",
	Short: "Get coin information",
	Run: func(cmd *cobra.Command, args []string) {
		params := crypto.Coins{}

		symbol, _ := cmd.Flags().GetString("symbol")
		network, _ := cmd.Flags().GetString("network")

		if symbol != "" {
			params.Symbol = symbol
		}
		if network != "" {
			params.Network = network
		}

		coins, err := crypto.GetCoins(params)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to get coins")
		}

		for _, coin := range coins.Items {
			log.Info().
				Str("symbol", coin.Symbol).
				Str("name", coin.Name).
				Bool("deposit_enable", coin.DepositEnable).
				Bool("withdraw_enable", coin.WithdrawEnable).
				Msg("Coin")
		}
	},
}

var cryptoAddressesCmd = &cobra.Command{
	Use:   "addresses",
	Short: "Get deposit addresses",
	Run: func(cmd *cobra.Command, args []string) {
		params := crypto.Addresses{}

		page, _ := cmd.Flags().GetInt("page")
		limit, _ := cmd.Flags().GetInt("limit")

		if page > 0 {
			params.Page = page
		}
		if limit > 0 {
			params.Limit = limit
		}

		addresses, err := crypto.GetAddresses(params)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to get addresses")
		}

		log.Info().
			Int("page", addresses.Page).
			Int("total_page", addresses.TotalPage).
			Int("total_item", addresses.TotalItem).
			Msg("Pagination")

		for _, addr := range addresses.Items {
			log.Info().
				Str("symbol", addr.Symbol).
				Str("network", addr.Network).
				Str("address", addr.Address).
				Str("memo", addr.Memo).
				Str("created_at", addr.CreatedAt).
				Msg("Address")
		}
	},
}

var cryptoCreateAddressCmd = &cobra.Command{
	Use:   "create-address [symbol] [network]",
	Short: "Create new deposit address",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		req := crypto.CreateAddressRequest{
			Symbol:  args[0],
			Network: args[1],
		}

		addresses, err := crypto.CreateAddress(req)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to create address")
		}

		for _, addr := range addresses {
			log.Info().
				Str("symbol", addr.Symbol).
				Str("network", addr.Network).
				Str("address", addr.Address).
				Str("memo", addr.Memo).
				Msg("Address Created")
		}
	},
}

var cryptoDepositsCmd = &cobra.Command{
	Use:   "deposits",
	Short: "Get deposit history",
	Run: func(cmd *cobra.Command, args []string) {
		params := crypto.Deposits{}

		page, _ := cmd.Flags().GetInt("page")
		limit, _ := cmd.Flags().GetInt("limit")
		symbol, _ := cmd.Flags().GetString("symbol")

		if page > 0 {
			params.Page = page
		}
		if limit > 0 {
			params.Limit = limit
		}
		if symbol != "" {
			params.Symbol = symbol
		}

		deposits, err := crypto.GetDeposits(params)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to get deposits")
		}

		log.Info().
			Int("page", deposits.Page).
			Int("total_page", deposits.TotalPage).
			Int("total_item", deposits.TotalItem).
			Msg("Pagination")

		for _, dep := range deposits.Items {
			log.Info().
				Str("hash", dep.Hash).
				Str("symbol", dep.Symbol).
				Str("network", dep.Network).
				Str("amount", dep.Amount).
				Str("from_address", dep.FromAddress).
				Str("to_address", dep.ToAddress).
				Int("confirmations", dep.Confirmations).
				Str("status", dep.Status).
				Str("created_at", dep.CreatedAt).
				Msg("Deposit")
		}
	},
}

var cryptoWithdrawsCmd = &cobra.Command{
	Use:   "withdraws",
	Short: "Get withdraw history",
	Run: func(cmd *cobra.Command, args []string) {
		params := crypto.Withdraws{}

		page, _ := cmd.Flags().GetInt("page")
		limit, _ := cmd.Flags().GetInt("limit")
		symbol, _ := cmd.Flags().GetString("symbol")

		if page > 0 {
			params.Page = page
		}
		if limit > 0 {
			params.Limit = limit
		}
		if symbol != "" {
			params.Symbol = symbol
		}

		withdraws, err := crypto.GetWithdraws(params)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to get withdraws")
		}

		log.Info().
			Int("page", withdraws.Page).
			Int("total_page", withdraws.TotalPage).
			Int("total_item", withdraws.TotalItem).
			Msg("Pagination")

		for _, wth := range withdraws.Items {
			log.Info().
				Str("txn_id", wth.TxnID).
				Str("hash", wth.Hash).
				Str("symbol", wth.Symbol).
				Str("network", wth.Network).
				Str("amount", wth.Amount).
				Str("fee", wth.Fee).
				Str("address", wth.Address).
				Str("status", wth.Status).
				Str("created_at", wth.CreatedAt).
				Msg("Withdraw")
		}
	},
}

var cryptoWithdrawCmd = &cobra.Command{
	Use:   "withdraw [symbol] [amount] [address] [network]",
	Short: "Create crypto withdrawal",
	Args:  cobra.ExactArgs(4),
	Run: func(cmd *cobra.Command, args []string) {
		req := crypto.CreateWithdrawRequest{
			Symbol:  args[0],
			Amount:  args[1],
			Address: args[2],
			Network: args[3],
		}

		memo, _ := cmd.Flags().GetString("memo")
		if memo != "" {
			req.Memo = memo
		}

		result, err := crypto.CreateWithdraw(req)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to create withdraw")
		}

		log.Info().
			Str("txn_id", result.TxnID).
			Str("symbol", result.Symbol).
			Str("amount", result.Amount).
			Str("fee", result.Fee).
			Str("address", result.Address).
			Msg("Withdraw Created")
	},
}

var cryptoCompensationsCmd = &cobra.Command{
	Use:   "compensations",
	Short: "Get compensation history",
	Run: func(cmd *cobra.Command, args []string) {
		params := crypto.Compensations{}

		page, _ := cmd.Flags().GetInt("page")
		limit, _ := cmd.Flags().GetInt("limit")
		symbol, _ := cmd.Flags().GetString("symbol")

		if page > 0 {
			params.Page = page
		}
		if limit > 0 {
			params.Limit = limit
		}
		if symbol != "" {
			params.Symbol = symbol
		}

		comps, err := crypto.GetCompensations(params)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to get compensations")
		}

		log.Info().
			Int("page", comps.Page).
			Int("total_page", comps.TotalPage).
			Int("total_item", comps.TotalItem).
			Msg("Pagination")

		for _, comp := range comps.Items {
			log.Info().
				Str("txn_id", comp.TxnID).
				Str("symbol", comp.Symbol).
				Str("amount", comp.Amount).
				Str("type", comp.Type).
				Str("status", comp.Status).
				Str("created_at", comp.CreatedAt).
				Msg("Compensation")
		}
	},
}

func init() {
	rootCmd.AddCommand(cryptoCmd)

	cryptoCmd.AddCommand(cryptoCoinsCmd)
	cryptoCmd.AddCommand(cryptoAddressesCmd)
	cryptoCmd.AddCommand(cryptoCreateAddressCmd)
	cryptoCmd.AddCommand(cryptoDepositsCmd)
	cryptoCmd.AddCommand(cryptoWithdrawsCmd)
	cryptoCmd.AddCommand(cryptoWithdrawCmd)
	cryptoCmd.AddCommand(cryptoCompensationsCmd)

	cryptoCoinsCmd.Flags().StringP("symbol", "s", "", "Filter by symbol")
	cryptoCoinsCmd.Flags().StringP("network", "n", "", "Filter by network")

	cryptoAddressesCmd.Flags().IntP("page", "p", 1, "Page number")
	cryptoAddressesCmd.Flags().IntP("limit", "l", 20, "Limit per page")

	cryptoDepositsCmd.Flags().IntP("page", "p", 1, "Page number")
	cryptoDepositsCmd.Flags().IntP("limit", "l", 20, "Limit per page")
	cryptoDepositsCmd.Flags().StringP("symbol", "s", "", "Filter by symbol")

	cryptoWithdrawsCmd.Flags().IntP("page", "p", 1, "Page number")
	cryptoWithdrawsCmd.Flags().IntP("limit", "l", 20, "Limit per page")
	cryptoWithdrawsCmd.Flags().StringP("symbol", "s", "", "Filter by symbol")

	cryptoWithdrawCmd.Flags().String("memo", "", "Memo/Tag for withdrawal")

	cryptoCompensationsCmd.Flags().IntP("page", "p", 1, "Page number")
	cryptoCompensationsCmd.Flags().IntP("limit", "l", 20, "Limit per page")
	cryptoCompensationsCmd.Flags().StringP("symbol", "s", "", "Filter by symbol")
}
