package cmd

import (
	"strconv"

	"github.com/dvgamerr-app/go-bitkub/market"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

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

type placeOrderFunc func(symbol string, amount float64, rate float64) (*market.PlaceAskResult, error)

func newPlaceOrderCmd(use string, short string, failure string, place placeOrderFunc) *cobra.Command {
	return &cobra.Command{
		Use:   use + " [symbol] [amount] [rate]",
		Short: short,
		Args:  cobra.ExactArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			amount, err := strconv.ParseFloat(args[1], 64)
			if err != nil {
				log.Fatal().Err(err).Msg("Invalid amount")
			}
			rate, err := strconv.ParseFloat(args[2], 64)
			if err != nil {
				log.Fatal().Err(err).Msg("Invalid rate")
			}

			result, err := place(args[0], amount, rate)
			if err != nil {
				log.Fatal().Err(err).Msg(failure)
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
}

var marketPlaceBidCmd = newPlaceOrderCmd("buy", "Place a bid order (buy)", "Failed to place bid",
	func(symbol string, amount float64, rate float64) (*market.PlaceAskResult, error) {
		result, err := market.PlaceBid(market.PlaceBidRequest{Symbol: symbol, Amount: amount, Rate: rate, Type: "limit"})
		if err != nil {
			return nil, err
		}
		return (*market.PlaceAskResult)(result), nil
	})

var marketPlaceAskCmd = newPlaceOrderCmd("sell", "Place an ask order (sell)", "Failed to place ask",
	func(symbol string, amount float64, rate float64) (*market.PlaceAskResult, error) {
		return market.PlaceAsk(market.PlaceAskRequest{Symbol: symbol, Amount: amount, Rate: rate, Type: "limit"})
	})

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
			Symbol: args[0],
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
