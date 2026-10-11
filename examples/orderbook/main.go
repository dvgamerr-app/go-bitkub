package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/dvgamerr-app/go-bitkub/stream"
)

var messageTitles = map[string]string{
	"bidschanged":   "📈 Bids Changed",
	"askschanged":   "📉 Asks Changed",
	"tradeschanged": "💱 Trades Changed",
	"global.ticker": "🌐 Global Ticker",
}

func printMessage(msg stream.Message) {
	if msg.Error != nil {
		fmt.Printf("❌ Error: %v\n", msg.Error)
		return
	}

	if title, ok := messageTitles[msg.Type]; ok {
		fmt.Println(title)
	}

	data, ok := msg.Data.(map[string]any)
	if !ok {
		return
	}
	if event, ok := data["event"].(string); ok {
		fmt.Printf("   Event: %s\n", event)
	}
	if pairingID, ok := data["pairing_id"].(float64); ok {
		fmt.Printf("   Pairing ID: %.0f\n", pairingID)
	}
}

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
			printMessage(msg)
		case <-interrupt:
			fmt.Println("\n👋 Closing connection...")
			s.Close()
			return
		}
	}
}
