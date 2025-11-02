package main

import (
	"context"
	"fmt"
	"time"

	"github.com/dvgamerr-app/go-bitkub/stream"
)

func main() {
	s := stream.New(&stream.StreamConfig{
		ReconnectInterval: 3 * time.Second,
		MaxReconnect:      5,
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
