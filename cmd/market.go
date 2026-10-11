package cmd

import (
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
			symbol = args[0]
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
		symbol := args[0]
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
		symbol := args[0]
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

func newOrderBookCmd(side string, short string, fetch func(symbol string, limit int) ([]market.AskOrder, error)) *cobra.Command {
	return &cobra.Command{
		Use:   side + " [symbol]",
		Short: short,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			limit, _ := cmd.Flags().GetInt("limit")
			orders, err := fetch(args[0], limit)
			if err != nil {
				log.Fatal().Err(err).Msg("Failed to get " + side)
			}
			for i, order := range orders {
				output(map[string]any{
					"order_id":  order.OrderID,
					"price":     order.Price,
					"size":      order.Size,
					"volume":    order.Volume,
					"timestamp": order.Timestamp,
				}, i == len(orders)-1)
			}
		},
	}
}

var marketAsksCmd = newOrderBookCmd("asks", "Get ask orders", market.GetAsks)

var marketBidsCmd = newOrderBookCmd("bids", "Get bid orders",
	func(symbol string, limit int) ([]market.AskOrder, error) {
		bids, err := market.GetBids(symbol, limit)
		orders := make([]market.AskOrder, len(bids))
		for i, bid := range bids {
			orders[i] = market.AskOrder(bid)
		}
		return orders, err
	})

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
