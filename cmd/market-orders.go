package cmd

import (
	"strconv"

	"github.com/dvgamerr-app/go-bitkub/market"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var marketOpenOrdersCmd = &cobra.Command{
	Use:   "open-orders [symbol]",
	Short: "Get open orders",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		symbol := args[0]
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
			Symbol: args[0],
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
		symbol := args[0]
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
