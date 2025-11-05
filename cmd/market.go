package cmd

import (
	"strconv"

	"github.com/dvgamerr-app/go-bitkub/market"
	"github.com/dvgamerr-app/go-bitkub/utils"
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
		activeSymbols := []market.Symbol{}
		for _, s := range symbols {
			if s.Status == "active" {
				activeSymbols = append(activeSymbols, s)
			}
		}
		for i, s := range activeSymbols {
			output(map[string]any{
				"symbol": s.Symbol,
				"name":   s.Name,
				"status": s.Status,
			}, i == len(activeSymbols)-1)
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
			symbol = utils.NormalizeSymbol(args[0])
		}
		tickers, err := market.GetTicker(symbol)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to get ticker")
		}
		for i, t := range tickers {
			output(map[string]any{
				"symbol":  t.Symbol,
				"last":    t.Last,
				"high24h": t.High24hr,
				"low24h":  t.Low24hr,
				"volume":  t.BaseVolume,
				"change":  t.PercentChange,
			}, i == len(tickers)-1)
		}
	},
}

var marketTradesCmd = &cobra.Command{
	Use:   "trades [symbol]",
	Short: "Get recent trades",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		limit, _ := cmd.Flags().GetInt("limit")
		symbol := utils.UppercaseSymbol(args[0])
		trades, err := market.GetTrades(symbol, limit)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to get trades")
		}
		for i, t := range trades {
			output(map[string]any{
				"timestamp": t[0],
				"rate":      t[1],
				"amount":    t[2],
				"side":      t[3],
			}, i == len(trades)-1)
		}
	},
}

var marketDepthCmd = &cobra.Command{
	Use:   "depth [symbol]",
	Short: "Get order book depth",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		limit, _ := cmd.Flags().GetInt("limit")
		symbol := utils.NormalizeSymbol(args[0])
		depth, err := market.GetDepth(symbol, limit)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to get depth")
		}

		totalItems := 1 + len(depth.Bids) + len(depth.Asks)
		if len(depth.Bids) > 5 {
			totalItems = 1 + 5 + len(depth.Asks)
		}
		if len(depth.Asks) > 5 {
			totalItems = 1 + len(depth.Bids) + 5
		}
		if len(depth.Bids) > 5 && len(depth.Asks) > 5 {
			totalItems = 1 + 5 + 5
		}
		itemCount := 0

		output(map[string]any{
			"bids": len(depth.Bids),
			"asks": len(depth.Asks),
		}, false)
		itemCount++

		for i, bid := range depth.Bids {
			if i >= 5 {
				break
			}
			itemCount++
			output(map[string]any{
				"type":   "bid",
				"price":  bid[0],
				"volume": bid[1],
			}, itemCount == totalItems)
		}

		for i, ask := range depth.Asks {
			if i >= 5 {
				break
			}
			itemCount++
			output(map[string]any{
				"type":   "ask",
				"price":  ask[0],
				"volume": ask[1],
			}, itemCount == totalItems)
		}
	},
}

var marketAsksCmd = &cobra.Command{
	Use:   "asks [symbol]",
	Short: "Get ask orders",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		limit, _ := cmd.Flags().GetInt("limit")
		symbol := utils.NormalizeSymbol(args[0])
		asks, err := market.GetAsks(symbol, limit)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to get asks")
		}
		for i, ask := range asks {
			output(map[string]any{
				"order_id":  ask.OrderID,
				"price":     ask.Price,
				"size":      ask.Size,
				"volume":    ask.Volume,
				"timestamp": ask.Timestamp,
			}, i == len(asks)-1)
		}
	},
}

var marketBidsCmd = &cobra.Command{
	Use:   "bids [symbol]",
	Short: "Get bid orders",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		limit, _ := cmd.Flags().GetInt("limit")
		symbol := utils.NormalizeSymbol(args[0])
		bids, err := market.GetBids(symbol, limit)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to get bids")
		}
		for i, bid := range bids {
			output(map[string]any{
				"order_id":  bid.OrderID,
				"price":     bid.Price,
				"size":      bid.Size,
				"volume":    bid.Volume,
				"timestamp": bid.Timestamp,
			}, i == len(bids)-1)
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
		filteredBalances := []struct {
			coin    string
			balance market.Balance
		}{}
		for coin, balance := range balances {
			if balance.Available > 0 || balance.Reserved > 0 {
				filteredBalances = append(filteredBalances, struct {
					coin    string
					balance market.Balance
				}{coin, balance})
			}
		}
		for i, item := range filteredBalances {
			output(map[string]any{
				"coin":      item.coin,
				"available": item.balance.Available,
				"reserved":  item.balance.Reserved,
			}, i == len(filteredBalances)-1)
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
		filteredWallet := []struct {
			coin    string
			balance float64
		}{}
		for coin, balance := range *wallet {
			if balance > 0 {
				filteredWallet = append(filteredWallet, struct {
					coin    string
					balance float64
				}{coin, balance})
			}
		}
		for i, item := range filteredWallet {
			output(map[string]any{
				"coin":    item.coin,
				"balance": item.balance,
			}, i == len(filteredWallet)-1)
		}
	},
}

var marketOpenOrdersCmd = &cobra.Command{
	Use:   "open-orders [symbol]",
	Short: "Get open orders",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		symbol := utils.UppercaseSymbol(args[0])
		orders, err := market.GetOpenOrders(symbol)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to get open orders")
		}

		if len(orders) == 0 {
			log.Info().Msg("Empty open order")
		}

		for i, order := range orders {
			output(map[string]any{
				"id":        order.ID,
				"side":      order.Side,
				"type":      order.Type,
				"rate":      order.Rate,
				"amount":    order.Amount,
				"timestamp": order.Timestamp,
			}, i == len(orders)-1)
		}
	},
}

var marketOrderHistoryCmd = &cobra.Command{
	Use:   "order-history <symbol>",
	Short: "Get order history",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		params := market.OrderHistoryParams{
			Symbol: utils.UppercaseSymbol(args[0]),
			Page:   "1",
			Limit:  "20",
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

		history, err := market.GetOrderHistory(params)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to get order history")
		}

		output(map[string]any{
			"page": history.Pagination.Page,
			"last": history.Pagination.Last,
		}, len(history.Result) == 0)

		for i, order := range history.Result {
			output(map[string]any{
				"txn_id":    order.TxnID,
				"order_id":  order.OrderID,
				"side":      order.Side,
				"type":      order.Type,
				"rate":      order.Rate,
				"amount":    order.Amount,
				"fee":       order.Fee,
				"credit":    order.Credit,
				"timestamp": order.Timestamp,
			}, i == len(history.Result)-1)
		}
	},
}

var marketOrderInfoCmd = &cobra.Command{
	Use:   "order-info [symbol] [order-id] [side]",
	Short: "Get order information",
	Args:  cobra.RangeArgs(2, 3),
	Run: func(cmd *cobra.Command, args []string) {
		symbol := utils.UppercaseSymbol(args[0])
		side := "buy"
		if len(args) > 2 {
			side = args[2]
		}
		info, err := market.GetOrderInfo(symbol, args[1], side)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to get order info")
		}
		output(map[string]any{
			"id":        info.ID,
			"rate":      info.Rate,
			"fee":       info.Fee,
			"credit":    info.Credit,
			"amount":    info.Amount,
			"filled":    info.Filled,
			"status":    info.Status,
			"parent":    info.Parent,
			"client_id": info.ClientID,
		})
	},
}

var marketPlaceBidCmd = &cobra.Command{
	Use:   "buy [symbol] [amount] [rate]",
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
			Symbol: utils.UppercaseSymbol(args[0]),
			Amount: amount,
			Rate:   rate,
			Type:   "limit",
		}

		result, err := market.PlaceBid(req)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to place bid")
		}

		output(map[string]any{
			"id":        result.ID,
			"type":      result.Type,
			"amount":    result.Amount,
			"rate":      result.Rate,
			"fee":       result.Fee,
			"credit":    result.Credit,
			"timestamp": result.Timestamp,
		})
	},
}

var marketPlaceAskCmd = &cobra.Command{
	Use:   "sell [symbol] [amount] [rate]",
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
			Symbol: utils.UppercaseSymbol(args[0]),
			Amount: amount,
			Rate:   rate,
			Type:   "limit",
		}

		result, err := market.PlaceAsk(req)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to place ask")
		}

		output(map[string]any{
			"id":        result.ID,
			"type":      result.Type,
			"amount":    result.Amount,
			"rate":      result.Rate,
			"fee":       result.Fee,
			"credit":    result.Credit,
			"timestamp": result.Timestamp,
		})
	},
}

var marketCancelOrderCmd = &cobra.Command{
	Use:   "cancel [symbol] [order-id] [side]",
	Short: "Cancel an order",
	Args:  cobra.RangeArgs(2, 3),
	Run: func(cmd *cobra.Command, args []string) {
		side := "buy"
		if len(args) > 2 {
			side = args[2]
		}

		req := market.CancelOrderRequest{
			Symbol: utils.UppercaseSymbol(args[0]),
			ID:     args[1],
			Side:   side,
		}

		if err := market.CancelOrder(req); err != nil {
			log.Fatal().Err(err).Msg("Failed to cancel order")
		}

		output(map[string]any{
			"symbol":   args[0],
			"order_id": args[1],
			"side":     side,
			"status":   "cancelled",
		})
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

		output(map[string]any{
			"crypto": map[string]any{
				"deposit":  limits.Limits.Crypto.Deposit,
				"withdraw": limits.Limits.Crypto.Withdraw,
			},
			"fiat": map[string]any{
				"deposit":  limits.Limits.Fiat.Deposit,
				"withdraw": limits.Limits.Fiat.Withdraw,
			},
		})
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
		output(map[string]any{
			"credits": credits,
		})
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
		output(map[string]any{
			"token": token,
		})
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

		log.Info().
			Str("status", result.Status).
			Int("bars", len(result.Close)).
			Msg("Symbol")

		for i := 0; i < len(result.Time); i++ {
			output(map[string]any{
				"time":   result.Time[i],
				"open":   result.Open[i],
				"high":   result.High[i],
				"low":    result.Low[i],
				"close":  result.Close[i],
				"volume": result.Volume[i],
			}, i == len(result.Time)-1)
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
