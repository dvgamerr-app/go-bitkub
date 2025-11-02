package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/dvgamerr-app/go-bitkub/stream"
)

func main() {
	s := stream.New(nil)

	if err := s.ConnectOrderBook(1); err != nil {
		panic(err)
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	fmt.Println("🚀 Connected to Bitkub Order Book (THB_BTC)")
	fmt.Println("Press Ctrl+C to stop...")
	fmt.Println()

	for {
		select {
		case msg := <-s.Messages():
			if msg.Error != nil {
				fmt.Printf("❌ Error: %v\n", msg.Error)
				continue
			}

			switch msg.Type {
			case "bidschanged":
				fmt.Println("📈 Bids Changed")
			case "askschanged":
				fmt.Println("📉 Asks Changed")
			case "tradeschanged":
				fmt.Println("💱 Trades Changed")
			case "global.ticker":
				fmt.Println("🌐 Global Ticker")
			}

			if data, ok := msg.Data.(map[string]interface{}); ok {
				if event, ok := data["event"].(string); ok {
					fmt.Printf("   Event: %s\n", event)
				}
				if pairingID, ok := data["pairing_id"].(float64); ok {
					fmt.Printf("   Pairing ID: %.0f\n", pairingID)
				}
			}

		case <-interrupt:
			fmt.Println("\n👋 Closing connection...")
			s.Close()
			return
		}
	}
}
