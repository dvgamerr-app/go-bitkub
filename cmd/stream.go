package cmd

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/dvgamerr-app/go-bitkub/stream"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var streamCmd = &cobra.Command{
	Use:   "stream",
	Short: "WebSocket stream commands",
	Long:  "Commands for real-time trade and ticker data",
}

func streamConnect(streams ...string) *stream.Stream {
	s := stream.New(nil)
	if err := s.ConnectMarket(streams...); err != nil {
		log.Fatal().Err(err).Msg("Failed to connect stream")
	}
	log.Info().Msg("Bitkub.com market connected")
	return s
}

func tailLoop(s *stream.Stream, limit int, continuous bool, filterType string) {
	count := 0
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	for {
		select {
		case msg, ok := <-s.Messages():
			if !ok {
				return
			}
			if msg.Error != nil {
				log.Error().Err(msg.Error).Msg("stream error")
				if !continuous {
					return
				}
				continue
			}
			if filterType != "" && msg.Type != filterType {
				continue
			}
			output(map[string]any{
				"type": msg.Type,
				"data": msg.Data,
				"ts":   msg.Timestamp.UnixMilli(),
			})
			count++
			if !continuous && limit > 0 && count >= limit {
				return
			}
		case <-interrupt:
			s.Close()
			return
		}
	}
}

var streamTradeCmd = &cobra.Command{
	Use:   "trade [symbol]",
	Short: "Stream trade data",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		symbol := args[0]
		count, _ := cmd.Flags().GetInt("count")
		tail, _ := cmd.Flags().GetBool("tail")
		streamName := "market.trade." + symbol
		s := streamConnect(streamName)
		defer func() {
			if !tail {
				s.Close()
			}
		}()
		tailLoop(s, count, tail, streamName)
	},
}

var streamTickerCmd = &cobra.Command{
	Use:   "ticker [symbol]",
	Short: "Stream ticker data",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		symbol := args[0]
		count, _ := cmd.Flags().GetInt("count")
		tail, _ := cmd.Flags().GetBool("tail")
		streamName := "market.ticker." + symbol
		s := streamConnect(streamName)
		defer func() {
			if !tail {
				s.Close()
			}
		}()
		tailLoop(s, count, tail, streamName)
	},
}

func init() {
	rootCmd.AddCommand(streamCmd)
	streamCmd.AddCommand(streamTradeCmd)
	streamCmd.AddCommand(streamTickerCmd)

	for _, c := range []*cobra.Command{streamTradeCmd, streamTickerCmd} {
		c.Flags().IntP("count", "n", 5, "Number of messages to show")
		c.Flags().BoolP("tail", "t", false, "Continue streaming until interrupted")
	}
}
