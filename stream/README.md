# Bitkub WebSocket Stream

Real-time WebSocket streaming client for Bitkub exchange data.

## Features

- ✅ Market Trade Stream
- ✅ Market Ticker Stream
- ✅ Live Order Book Stream
- ✅ Multiple Streams Subscription
- ✅ Auto Reconnect
- ✅ Graceful Shutdown
- ✅ Thread-Safe
- ✅ Context Support

## Installation

```bash
go get github.com/dvgamerr-app/go-bitkub/stream
```

## Quick Start

### Market Stream (Trade + Ticker)

```go
package main

import (
	"fmt"
	"github.com/dvgamerr-app/go-bitkub/stream"
)

func main() {
	s := stream.New(nil)
	
	if err := s.ConnectMarket("market.trade.thb_btc", "market.ticker.thb_btc"); err != nil {
		panic(err)
	}
	defer s.Close()

	for msg := range s.Messages() {
		if msg.Error != nil {
			fmt.Printf("Error: %v\n", msg.Error)
			continue
		}
		
		fmt.Printf("[%s] %+v\n", msg.Type, msg.Data)
	}
}
```

### Order Book Stream

```go
package main

import (
	"fmt"
	"github.com/dvgamerr-app/go-bitkub/stream"
)

func main() {
	s := stream.New(nil)
	
	// Connect to THB_BTC orderbook (symbol ID = 1)
	if err := s.ConnectOrderBook(1); err != nil {
		panic(err)
	}
	defer s.Close()

	for msg := range s.Messages() {
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
```

## Configuration

```go
config := &stream.StreamConfig{
	ReconnectInterval: 5 * time.Second,  // Wait time before reconnect
	MaxReconnect:      10,                // Max reconnect attempts (0 = unlimited)
	PingInterval:      30 * time.Second,  // Ping interval
	ReadTimeout:       60 * time.Second,  // Read timeout
}

s := stream.New(config)
```

## Stream Types

### Market Streams

Format: `market.<type>.<symbol>`

**Available Types:**
- `trade` - Real-time trade data
- `ticker` - Real-time ticker updates

**Example Symbols:**
- `thb_btc` - Bitcoin
- `thb_eth` - Ethereum
- `thb_usdt` - Tether

**Multiple Streams:**
```go
s.ConnectMarket(
	"market.trade.thb_btc",
	"market.ticker.thb_btc",
	"market.trade.thb_eth",
	"market.ticker.thb_eth",
)
```

### Order Book Stream

Format: `orderbook/<symbol_id>`

**Example Symbol IDs:**
- `1` - THB_BTC
- `2` - THB_ETH
- See [Bitkub Symbols API](https://github.com/bitkub/bitkub-official-api-docs/blob/master/restful-api.md#get-apimarketsymbols) for all symbol IDs

## Message Structure

```go
type Message struct {
	Data      interface{}  // Raw message data
	Type      string       // Message type (stream name or event)
	Error     error        // Error if occurred
	Timestamp time.Time    // Message timestamp
}
```

## Advanced Usage

### With Context and Timeout

```go
package main

import (
	"context"
	"fmt"
	"time"
	"github.com/dvgamerr-app/go-bitkub/stream"
)

func main() {
	s := stream.New(nil)
	if err := s.ConnectMarket("market.trade.thb_btc"); err != nil {
		panic(err)
	}
	defer s.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	for {
		select {
		case msg := <-s.Messages():
			if msg.Error != nil {
				fmt.Printf("Error: %v\n", msg.Error)
				continue
			}
			fmt.Println("Message:", msg.Type)

		case <-ctx.Done():
			fmt.Println("Timeout reached")
			return
		}
	}
}
```

### With Graceful Shutdown

```go
package main

import (
	"fmt"
	"os"
	"os/signal"
	"github.com/dvgamerr-app/go-bitkub/stream"
)

func main() {
	s := stream.New(nil)
	if err := s.ConnectMarket("market.trade.thb_btc"); err != nil {
		panic(err)
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	for {
		select {
		case msg := <-s.Messages():
			if msg.Error != nil {
				fmt.Printf("Error: %v\n", msg.Error)
				continue
			}
			fmt.Println("Message:", msg.Type)

		case <-interrupt:
			fmt.Println("Shutting down...")
			s.Close()
			return
		}
	}
}
```

## Error Handling

The stream automatically handles:
- Connection errors with auto-reconnect
- Ping/Pong timeout detection
- Graceful shutdown
- Concurrent write protection

Errors are sent through the message channel:

```go
for msg := range s.Messages() {
	if msg.Error != nil {
		// Handle error
		fmt.Printf("Error: %v\n", msg.Error)
		continue
	}
	
	// Handle message
}
```

## Examples

See the `examples/` directory for complete working examples:

- `examples/market/` - Market stream example
- `examples/orderbook/` - Order book stream example
- `examples/timeout/` - Timeout with context example

## Testing

```bash
# Run all tests
go test -v ./stream

# Run specific test
go test -v ./stream -run TestMarketTrade

# Run with timeout
go test -v ./stream -run TestMarketTrade -timeout 35s
```

## Reference

- [Bitkub WebSocket API Documentation](https://github.com/bitkub/bitkub-official-api-docs/blob/master/websocket-api.md)
- [Bitkub Official API Docs](https://github.com/bitkub/bitkub-official-api-docs)
