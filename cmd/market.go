package cmd

import (
	"strconv"

	"github.com/dvgamerr-app/go-bitkub/market"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var marketCmd = &cobra.Command{
	Use:   "market",
	Short: "Market API commands",
	Long:  "Commands for trading, orders, balances, and market data",
}

var marketSymbolsCmd = &cobra.Command{
	Use:   "symbols",
	Short: "Get all trading symbols",
	Run: func(cmd *cobra.Command, args []string) {
		symbols, err := market.GetSymbols()
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to get symbols")
		}
		for _, s := range symbols {
			if s.Status != "active" {
				continue
			}

			log.Info().
				Str("symbol", s.Symbol).
				Str("name", s.Name).
				Str("status", s.Status).
				Msg("Symbol")
		}
	},
}

var marketTickerCmd = &cobra.Command{
	Use:   "ticker [symbol]",
	Short: "Get ticker information",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		symbol := ""
		if len(args) > 0 {
			symbol = args[0]
		}
		tickers, err := market.GetTicker(symbol)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to get ticker")
		}
		for _, t := range tickers {
			output(map[string]any{
				"symbol":  t.Symbol,
				"last":    t.Last,
				"high24h": t.High24hr,
				"low24h":  t.Low24hr,
				"volume":  t.BaseVolume,
				"change":  t.PercentChange,
			})
		}
	},
}

var marketTradesCmd = &cobra.Command{
	Use:   "trades [symbol]",
	Short: "Get recent trades",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		limit, _ := cmd.Flags().GetInt("limit")
		trades, err := market.GetTrades(args[0], limit)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to get trades")
		}
		for _, t := range trades {
			log.Info().
				Interface("trade", t).
				Msg("Trade")
		}
	},
}

var marketDepthCmd = &cobra.Command{
	Use:   "depth [symbol]",
	Short: "Get order book depth",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		limit, _ := cmd.Flags().GetInt("limit")
		depth, err := market.GetDepth(args[0], limit)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to get depth")
		}
		log.Info().Int("bids", len(depth.Bids)).Int("asks", len(depth.Asks)).Msg("Order Book")

		log.Info().Msg("Top 5 Bids:")
		for i, bid := range depth.Bids {
			if i >= 5 {
				break
			}
			log.Info().Float64("price", bid[0]).Float64("volume", bid[1]).Msg("  Bid")
		}

		log.Info().Msg("Top 5 Asks:")
		for i, ask := range depth.Asks {
			if i >= 5 {
				break
			}
			log.Info().Float64("price", ask[0]).Float64("volume", ask[1]).Msg("  Ask")
		}
	},
}

var marketAsksCmd = &cobra.Command{
	Use:   "asks [symbol]",
	Short: "Get ask orders",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		limit, _ := cmd.Flags().GetInt("limit")
		asks, err := market.GetAsks(args[0], limit)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to get asks")
		}
		for _, ask := range asks {
			log.Info().
				Str("order_id", ask.OrderID).
				Str("price", ask.Price).
				Str("size", ask.Size).
				Str("volume", ask.Volume).
				Int64("timestamp", ask.Timestamp).
				Msg("Ask")
		}
	},
}

var marketBidsCmd = &cobra.Command{
	Use:   "bids [symbol]",
	Short: "Get bid orders",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		limit, _ := cmd.Flags().GetInt("limit")
		bids, err := market.GetBids(args[0], limit)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to get bids")
		}
		for _, bid := range bids {
			log.Info().
				Str("order_id", bid.OrderID).
				Str("price", bid.Price).
				Str("size", bid.Size).
				Str("volume", bid.Volume).
				Int64("timestamp", bid.Timestamp).
				Msg("Bid")
		}
	},
}

var marketBalancesCmd = &cobra.Command{
	Use:   "balances",
	Short: "Get account balances",
	Run: func(cmd *cobra.Command, args []string) {
		balances, err := market.GetBalances()
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to get balances")
		}
		for coin, balance := range balances {
			if balance.Available > 0 || balance.Reserved > 0 {
				log.Info().
					Str("coin", coin).
					Float64("available", balance.Available).
					Float64("reserved", balance.Reserved).
					Msg("Balance")
			}
		}
	},
}

var marketWalletCmd = &cobra.Command{
	Use:   "wallet",
	Short: "Get wallet information",
	Run: func(cmd *cobra.Command, args []string) {
		wallet, err := market.GetWallet()
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to get wallet")
		}
		for coin, balance := range *wallet {
			if balance > 0 {
				log.Info().Str("coin", coin).Float64("balance", balance).Msg("Wallet")
			}
		}
	},
}

var marketOpenOrdersCmd = &cobra.Command{
	Use:   "open-orders [symbol]",
	Short: "Get open orders",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		symbol := ""
		if len(args) > 0 {
			symbol = args[0]
		}
		orders, err := market.GetMyOpenOrders(symbol)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to get open orders")
		}
		for _, order := range orders {
			log.Info().
				Str("id", order.ID).
				Str("side", order.Side).
				Str("type", order.Type).
				Str("rate", order.Rate).
				Str("amount", order.Amount).
				Int64("timestamp", order.Timestamp).
				Msg("Open Order")
		}
	},
}

var marketOrderHistoryCmd = &cobra.Command{
	Use:   "order-history [symbol]",
	Short: "Get order history",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		params := market.MyOrderHistoryParams{
			Page:  "1",
			Limit: "20",
		}
		if len(args) > 0 {
			params.Symbol = args[0]
		}

		page, _ := cmd.Flags().GetInt("page")
		limit, _ := cmd.Flags().GetInt("limit")
		start, _ := cmd.Flags().GetInt64("start")
		end, _ := cmd.Flags().GetInt64("end")

		if page > 0 {
			params.Page = strconv.Itoa(page)
		}
		if limit > 0 {
			params.Limit = strconv.Itoa(limit)
		}
		if start > 0 {
			params.Start = strconv.FormatInt(start, 10)
		}
		if end > 0 {
			params.End = strconv.FormatInt(end, 10)
		}

		history, err := market.GetMyOrderHistory(params)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to get order history")
		}

		log.Info().
			Int("page", history.Pagination.Page).
			Int("last", history.Pagination.Last).
			Msg("Pagination")

		for _, order := range history.Result {
			log.Info().
				Str("txn_id", order.TxnID).
				Str("order_id", order.OrderID).
				Str("side", order.Side).
				Str("type", order.Type).
				Str("rate", order.Rate).
				Str("amount", order.Amount).
				Str("fee", order.Fee).
				Str("credit", order.Credit).
				Int64("timestamp", order.Timestamp).
				Msg("Order")
		}
	},
}

var marketOrderInfoCmd = &cobra.Command{
	Use:   "order-info [symbol] [order-id] [side]",
	Short: "Get order information",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		info, err := market.GetOrderInfo(args[0], args[1], args[2])
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to get order info")
		}
		log.Info().
			Str("id", info.ID).
			Float64("rate", info.Rate).
			Float64("fee", info.Fee).
			Float64("credit", info.Credit).
			Float64("amount", info.Amount).
			Float64("filled", info.Filled).
			Str("status", info.Status).
			Str("parent", info.Parent).
			Str("client_id", info.ClientID).
			Msg("Order Info")
	},
}

var marketPlaceBidCmd = &cobra.Command{
	Use:   "place-bid [symbol] [amount] [rate]",
	Short: "Place a bid order (buy)",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		var amount, rate float64
		var err error

		amount, err = strconv.ParseFloat(args[1], 64)
		if err != nil {
			log.Fatal().Err(err).Msg("Invalid amount")
		}
		rate, err = strconv.ParseFloat(args[2], 64)
		if err != nil {
			log.Fatal().Err(err).Msg("Invalid rate")
		}

		req := market.PlaceBidRequest{
			Symbol: args[0],
			Amount: amount,
			Rate:   rate,
			Type:   "limit",
		}

		result, err := market.PlaceBid(req)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to place bid")
		}

		log.Info().
			Str("id", result.ID).
			Str("type", result.Type).
			Float64("amount", result.Amount).
			Float64("rate", result.Rate).
			Float64("fee", result.Fee).
			Float64("credit", result.Credit).
			Str("timestamp", result.Timestamp).
			Msg("Bid Placed")
	},
}

var marketPlaceAskCmd = &cobra.Command{
	Use:   "place-ask [symbol] [amount] [rate]",
	Short: "Place an ask order (sell)",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		var amount, rate float64
		var err error

		amount, err = strconv.ParseFloat(args[1], 64)
		if err != nil {
			log.Fatal().Err(err).Msg("Invalid amount")
		}
		rate, err = strconv.ParseFloat(args[2], 64)
		if err != nil {
			log.Fatal().Err(err).Msg("Invalid rate")
		}

		req := market.PlaceAskRequest{
			Symbol: args[0],
			Amount: amount,
			Rate:   rate,
			Type:   "limit",
		}

		result, err := market.PlaceAsk(req)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to place ask")
		}

		log.Info().
			Str("id", result.ID).
			Str("type", result.Type).
			Float64("amount", result.Amount).
			Float64("rate", result.Rate).
			Float64("fee", result.Fee).
			Float64("credit", result.Credit).
			Str("timestamp", result.Timestamp).
			Msg("Ask Placed")
	},
}

var marketCancelOrderCmd = &cobra.Command{
	Use:   "cancel [symbol] [order-id] [side]",
	Short: "Cancel an order",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		req := market.CancelOrderRequest{
			Symbol: args[0],
			ID:     args[1],
			Side:   args[2],
		}

		if err := market.CancelOrder(req); err != nil {
			log.Fatal().Err(err).Msg("Failed to cancel order")
		}

		log.Info().
			Str("symbol", args[0]).
			Str("order_id", args[1]).
			Str("side", args[2]).
			Msg("Order Cancelled")
	},
}

var marketLimitsCmd = &cobra.Command{
	Use:   "limits",
	Short: "Get user trading limits",
	Run: func(cmd *cobra.Command, args []string) {
		limits, err := market.GetUserLimits()
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to get limits")
		}

		log.Info().
			Float64("deposit", limits.Limits.Crypto.Deposit).
			Float64("withdraw", limits.Limits.Crypto.Withdraw).
			Msg("Crypto Limits")

		log.Info().
			Float64("deposit", limits.Limits.Fiat.Deposit).
			Float64("withdraw", limits.Limits.Fiat.Withdraw).
			Msg("Fiat Limits")
	},
}

var marketTradingCreditsCmd = &cobra.Command{
	Use:   "credits",
	Short: "Get trading credits",
	Run: func(cmd *cobra.Command, args []string) {
		credits, err := market.GetTradingCredits()
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to get trading credits")
		}
		log.Info().Float64("credits", credits).Msg("Trading Credits")
	},
}

var marketWSTokenCmd = &cobra.Command{
	Use:   "wstoken",
	Short: "Get WebSocket token",
	Run: func(cmd *cobra.Command, args []string) {
		token, err := market.GetWSToken()
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to get WS token")
		}
		log.Info().Str("token", token).Msg("WebSocket Token")
	},
}

var marketHistoryCmd = &cobra.Command{
	Use:   "history [symbol]",
	Short: "Get historical data for TradingView chart",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		resolution, _ := cmd.Flags().GetString("resolution")
		from, _ := cmd.Flags().GetInt64("from")
		to, _ := cmd.Flags().GetInt64("to")

		result, err := market.GetHistory(market.HistoryRequest{
			Symbol:     args[0],
			Resolution: resolution,
			From:       from,
			To:         to,
		})
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to get history")
		}

		output(map[string]any{
			"status": result.Status,
			"bars":   len(result.Close),
		})

		for i := 0; i < len(result.Time); i++ {
			output(map[string]any{
				"time":   result.Time[i],
				"open":   result.Open[i],
				"high":   result.High[i],
				"low":    result.Low[i],
				"close":  result.Close[i],
				"volume": result.Volume[i],
			})
		}
	},
}

func init() {
	rootCmd.AddCommand(marketCmd)

	marketCmd.AddCommand(marketSymbolsCmd)
	marketCmd.AddCommand(marketTickerCmd)
	marketCmd.AddCommand(marketTradesCmd)
	marketCmd.AddCommand(marketDepthCmd)
	marketCmd.AddCommand(marketAsksCmd)
	marketCmd.AddCommand(marketBidsCmd)
	marketCmd.AddCommand(marketBalancesCmd)
	marketCmd.AddCommand(marketWalletCmd)
	marketCmd.AddCommand(marketOpenOrdersCmd)
	marketCmd.AddCommand(marketOrderHistoryCmd)
	marketCmd.AddCommand(marketOrderInfoCmd)
	marketCmd.AddCommand(marketPlaceBidCmd)
	marketCmd.AddCommand(marketPlaceAskCmd)
	marketCmd.AddCommand(marketCancelOrderCmd)
	marketCmd.AddCommand(marketLimitsCmd)
	marketCmd.AddCommand(marketTradingCreditsCmd)
	marketCmd.AddCommand(marketWSTokenCmd)
	marketCmd.AddCommand(marketHistoryCmd)

	marketTradesCmd.Flags().IntP("limit", "l", 10, "Limit number of results")
	marketDepthCmd.Flags().IntP("limit", "l", 10, "Limit number of results")
	marketAsksCmd.Flags().IntP("limit", "l", 10, "Limit number of results")
	marketBidsCmd.Flags().IntP("limit", "l", 10, "Limit number of results")
	marketOrderHistoryCmd.Flags().IntP("page", "p", 1, "Page number")
	marketOrderHistoryCmd.Flags().IntP("limit", "l", 20, "Limit per page")
	marketOrderHistoryCmd.Flags().Int64("start", 0, "Start timestamp")
	marketOrderHistoryCmd.Flags().Int64("end", 0, "End timestamp")
	marketHistoryCmd.Flags().StringP("resolution", "r", "1D", "Chart resolution (1, 5, 15, 60, 240, 1D)")
	marketHistoryCmd.Flags().Int64("from", 0, "From timestamp (default: 24h ago)")
	marketHistoryCmd.Flags().Int64("to", 0, "To timestamp (default: now)")
}
