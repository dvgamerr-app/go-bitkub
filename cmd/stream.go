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
			printed, stop := handleStreamMessage(msg, continuous, filterType)
			if printed {
				count++
			}
			if stop || limitReached(printed, continuous, limit, count) {
				return
			}
		case <-interrupt:
			s.Close()
			return
		}
	}
}

func handleStreamMessage(msg stream.Message, continuous bool, filterType string) (printed bool, stop bool) {
	if msg.Error != nil {
		log.Error().Err(msg.Error).Msg("stream error")
		return false, !continuous
	}
	if filterType != "" && msg.Type != filterType {
		return false, false
	}
	output(map[string]any{
		"type": msg.Type,
		"data": msg.Data,
		"ts":   msg.Timestamp.UnixMilli(),
	})
	return true, false
}

func newStreamCmd(kind string, short string) *cobra.Command {
	return &cobra.Command{
		Use:   kind + " [symbol]",
		Short: short,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			count, _ := cmd.Flags().GetInt("count")
			tail, _ := cmd.Flags().GetBool("tail")
			streamName := "market." + kind + "." + args[0]
			s := streamConnect(streamName)
			defer func() {
				if !tail {
					s.Close()
				}
			}()
			tailLoop(s, count, tail, streamName)
		},
	}
}

var (
	streamTradeCmd  = newStreamCmd("trade", "Stream trade data")
	streamTickerCmd = newStreamCmd("ticker", "Stream ticker data")
)

func init() {
	rootCmd.AddCommand(streamCmd)
	streamCmd.AddCommand(streamTradeCmd)
	streamCmd.AddCommand(streamTickerCmd)

	for _, c := range []*cobra.Command{streamTradeCmd, streamTickerCmd} {
		c.Flags().IntP("count", "n", 5, "Number of messages to show")
		c.Flags().BoolP("tail", "t", false, "Continue streaming until interrupted")
	}
}

func limitReached(printed bool, continuous bool, limit int, count int) bool {
	return printed && !continuous && limit > 0 && count >= limit
}
