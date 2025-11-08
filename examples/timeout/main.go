package main

import (
	"context"
	"fmt"
	"time"

	"github.com/dvgamerr-app/go-bitkub/bitkub"
	"github.com/dvgamerr-app/go-bitkub/stream"
)

func main() {
	// Configure HTTP client timeouts
	bitkub.SetConfig(bitkub.Config{
		APITimeout:      60 * time.Second,  // Increase API timeout for slow connections
		ServerTimeout:   10 * time.Second,  // Server time request timeout
		MaxIdleConns:    200,               // More idle connections for high-frequency trading
		IdleConnTimeout: 120 * time.Second, // Keep connections alive longer
	})

	fmt.Printf("HTTP Client Config: %+v\n\n", bitkub.GetConfig())

	// Configure WebSocket stream with custom buffer
	s := stream.New(&stream.StreamConfig{
		ReconnectInterval: 3 * time.Second,
		MaxReconnect:      5,
		PingInterval:      20 * time.Second,
		ReadTimeout:       30 * time.Second,
		MessageBuffer:     500, // Higher buffer for high-frequency data
	})

	if err := s.ConnectMarket("market.trade.thb_btc"); err != nil {
		panic(err)
	}
	defer s.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	fmt.Println("⏱️  Collecting data for 30 seconds...")
	fmt.Println()

	messageCount := 0
	for {
		select {
		case msg, ok := <-s.Messages():
			if !ok {
				fmt.Println("Channel closed")
				return
			}

			if msg.Error != nil {
				fmt.Printf("Error: %v\n", msg.Error)
				continue
			}

			messageCount++
			fmt.Printf("[%d] %s\n", messageCount, msg.Type)

		case <-ctx.Done():
			fmt.Printf("\n✅ Collected %d messages in 30 seconds\n", messageCount)
			return
		}
	}
}
