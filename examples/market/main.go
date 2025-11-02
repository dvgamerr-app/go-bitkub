package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/dvgamerr-app/go-bitkub/stream"
)

func main() {
	config := &stream.StreamConfig{
		ReconnectInterval: 5 * time.Second,
		MaxReconnect:      10,
		PingInterval:      30 * time.Second,
		ReadTimeout:       60 * time.Second,
	}

	s := stream.New(config)

	if err := s.ConnectMarket("market.trade.thb_btc", "market.ticker.thb_btc"); err != nil {
		panic(err)
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	fmt.Println("🚀 Connected to Bitkub WebSocket")
	fmt.Println("Press Ctrl+C to stop...")
	fmt.Println()

	messageCount := 0
	for {
		select {
		case msg := <-s.Messages():
			if msg.Error != nil {
				fmt.Printf("❌ Error: %v\n", msg.Error)
				continue
			}

			if msg.Type == "reconnected" {
				fmt.Println("🔄 Reconnected")
				continue
			}

			messageCount++
			fmt.Printf("[%d] %s: %+v\n", messageCount, msg.Type, msg.Data)

		case <-interrupt:
			fmt.Println("\n👋 Closing connection...")
			s.Close()
			return
		}
	}
}
