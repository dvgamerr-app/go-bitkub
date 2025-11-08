package stream

import (
	"fmt"
	"testing"
	"time"
)

func TestMarketTicker(t *testing.T) {
	stream := New(nil)

	if err := stream.ConnectMarket("market.ticker.thb_btc"); err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer stream.Close()

	timeout := time.After(5 * time.Second)
	messageCount := 0

	for {
		select {
		case msg := <-stream.Messages():
			if msg.Error != nil {
				t.Logf("Error: %v", msg.Error)
				continue
			}

			messageCount++

			if messageCount >= 5 {
				return
			}

		case <-timeout:
			return
		}
	}
}

func TestMultipleStreams(t *testing.T) {
	stream := New(nil)

	if err := stream.ConnectMarket(
		"market.trade.thb_btc",
		"market.ticker.thb_btc",
		"market.trade.thb_eth",
	); err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer stream.Close()

	timeout := time.After(5 * time.Second)
	messageCount := 0

	for {
		select {
		case msg := <-stream.Messages():
			if msg.Error != nil {
				t.Logf("Error: %v", msg.Error)
				continue
			}

			messageCount++

			if messageCount >= 20 {
				return
			}

		case <-timeout:
			return
		}
	}
}

func TestOrderBook(t *testing.T) {
	stream := New(nil)

	if err := stream.ConnectOrderBook(1); err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer stream.Close()

	timeout := time.After(5 * time.Second)
	messageCount := 0

	for {
		select {
		case msg := <-stream.Messages():
			if msg.Error != nil {
				t.Logf("Error: %v", msg.Error)
				continue
			}

			messageCount++

			if messageCount >= 10 {
				return
			}

		case <-timeout:
			return
		}
	}
}

func ExampleStream_market() {
	stream := New(&StreamConfig{
		ReconnectInterval: 5 * time.Second,
		MaxReconnect:      10,
	})

	if err := stream.ConnectMarket("market.trade.thb_btc", "market.ticker.thb_btc"); err != nil {
		panic(err)
	}
	defer stream.Close()

	for msg := range stream.Messages() {
		if msg.Error != nil {
			fmt.Printf("Error: %v\n", msg.Error)
			continue
		}

		fmt.Printf("Type: %s, Data: %+v\n", msg.Type, msg.Data)
	}
}

func ExampleStream_orderbook() {
	stream := New(nil)

	if err := stream.ConnectOrderBook(1); err != nil {
		panic(err)
	}
	defer stream.Close()

	for msg := range stream.Messages() {
		if msg.Error != nil {
			fmt.Printf("Error: %v\n", msg.Error)
			continue
		}

		switch msg.Type {
		case "bidschanged":
			fmt.Println("Bids changed:", msg.Data)
		case "askschanged":
			fmt.Println("Asks changed:", msg.Data)
		case "tradeschanged":
			fmt.Println("Trades changed:", msg.Data)
		case "global.ticker":
			fmt.Println("Global ticker:", msg.Data)
		}
	}
}
