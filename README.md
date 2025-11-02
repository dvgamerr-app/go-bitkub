# Bitkub Go SDK

[![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.18-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)

Go SDK for [Bitkub](https://www.bitkub.com/) Cryptocurrency Exchange API - Complete implementation with full V3 & V4 API support.

![](./docs/example.png)

## ⚠️ Important Updates

This SDK implements the latest Bitkub API V3 specification (November 2025) with:
- ✅ All deprecated endpoints removed
- ✅ Using V3 endpoints exclusively
- ✅ Keyset-based pagination (page-based removed)
- ✅ Simplified function names (removed V3 suffix)

## 🚀 Features

### Market API (V3)
- ✅ Non-secure endpoints (Market data, server status)
- ✅ Secure endpoints (Trading, user info, fiat operations)
- ✅ WebSocket token support
- ✅ Full order management (place, cancel, history)
- ✅ Wallet & balance operations
- ✅ TradingView chart data

### Crypto API (V4) ✨
- 📋 List crypto addresses with pagination
- ➕ Generate new crypto addresses
- 💰 View deposit history
- 💸 View withdrawal history
- 🚀 Create withdrawals to trusted addresses
- 🪙 Get available coins and networks
- 🎁 View compensations history

### Core Features
- ✅ Type-safe API responses
- ✅ Proper error handling
- ✅ Connection pooling & optimization
- ✅ HMAC SHA256 signature authentication
- ✅ Rate limit awareness

## 📦 Installation

```bash
go get github.com/dvgamerr-app/go-bitkub
```

## 🔧 Quick Start

### Initialize

```go
package main

import (
    "log"
    "github.com/dvgamerr-app/go-bitkub"
    "github.com/dvgamerr-app/go-bitkub/bitkub"
)

func main() {
    // Initialize with API credentials
    bitkub.Initlizer("YOUR_API_KEY", "YOUR_SECRET_KEY")
    // Or use environment variables BTK_APIKEY and BTK_SECRETKEY
    bitkub.Initlizer()
    
    // Get wallet balance
    wallet, err := Wallet()
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Wallet: %+v", wallet)

    // Get detailed balances
    balances, err := Balances()
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Balances: %+v", balances)
    
}
```

## 📚 API Coverage

### Non-Secure Endpoints (V3)

```go
// Get system status
status, err := bitkub.GetStatus()

// Get server time
timestamp, err := bitkub.GetServerTimeV3()

// Get all symbols
symbols, err := market.GetSymbols()

// Get ticker data
tickers, err := market.GetTickerV3("btc_thb")

#### Get Market Depth
```go
depth, err := market.GetDepth("btc_thb", 10)
```

#### Get Recent Trades
```go
trades, err := market.GetTrades("btc_thb", 10)
```

#### Get Order Books
```go
// Get buy orders (bids)
bids, err := market.GetBids("btc_thb", 10)

// Get sell orders (asks)
asks, err := market.GetAsks("btc_thb", 10)

// Get recent trades
trades, err := market.GetTradesV3("btc_thb", 10)

// Get TradingView history
params := market.TradingViewHistoryParams{
    Symbol:     "BTC_THB",
    Resolution: "60",
    From:       1633424427,
    To:         1633427427,
}
history, err := market.GetTradingViewHistory(params)
```

### Trading Endpoints (V3 - Secure)

```go
// Get wallet balances
wallet, err := market.GetWallet()

// Get detailed balances
balances, err := market.GetBalances()

// Place buy order
bidReq := market.PlaceBidRequest{
    Sym:      "btc_thb",
    Amt:      1000,
    Rat:      2500000,
    Typ:      "limit",
    ClientID: "my-order-1",
}
bidResult, err := market.PlaceBid(bidReq)

// Place sell order
askReq := market.PlaceAskRequest{
    Sym: "btc_thb",
    Amt: 0.001,
    Rat: 2600000,
    Typ: "limit",
}
askResult, err := market.PlaceAsk(askReq)

// Cancel order
cancelReq := market.CancelOrderRequest{
    Sym: "btc_thb",
    ID:  "12345",
    Sd:  "buy",
}
err = market.CancelOrder(cancelReq)

// Get open orders
orders, err := market.GetMyOpenOrders("btc_thb")

// Get order history (with keyset pagination)
historyParams := market.MyOrderHistoryParams{
    Sym:            "BTC_THB",
    Lmt:            "10",
    PaginationType: "keyset",
}
orderHistory, err := market.GetMyOrderHistory(historyParams)

// Get order info
orderInfo, err := market.GetOrderInfo("btc_thb", "12345", "buy")

// Get WebSocket token
token, err := market.GetWSToken()
```

### User Endpoints (V3 - Secure)

```go
import "github.com/dvgamerr-app/go-bitkub/user"

// Get trading credits
credits, err := user.GetTradingCredits()

// Get user limits
limits, err := user.GetUserLimits()

// Get coin convert history
convertParams := user.CoinConvertHistoryParams{
    P:      1,
    Lmt:    100,
    Status: "success",
}
convertHistory, err := user.GetCoinConvertHistory(convertParams)
```

### Fiat Endpoints (V3 - Secure)

```go
import "github.com/dvgamerr-app/go-bitkub/fiat"

// Get bank accounts
accountsParams := fiat.AccountsParams{
    P:   1,
    Lmt: 10,
}
accounts, err := fiat.GetAccounts(accountsParams)

// Withdraw fiat
withdrawReq := fiat.WithdrawRequest{
    ID:  "bank-account-id",
    Amt: 1000.0,
}
withdrawResult, err := fiat.Withdraw(withdrawReq)

// Get deposit history
depositParams := fiat.DepositHistoryParams{
    P:   1,
    Lmt: 10,
}
deposits, err := fiat.GetDepositHistory(depositParams)

// Get withdrawal history
withdrawParams := fiat.WithdrawHistoryParams{
    P:   1,
    Lmt: 10,
}
withdrawals, err := fiat.GetWithdrawHistory(withdrawParams)
```

### Crypto API (V4) Examples

```go
### Crypto API (V4) Examples

```go
// List crypto addresses with pagination
addresses, err := crypto.GetAddresses(crypto.GetAddressesParams{
    PaginationParams: crypto.PaginationParams{
        Page:  1,
        Limit: 10,
    },
    SymbolNetworkParams: crypto.SymbolNetworkParams{
        Symbol:  "ATOM",
        Network: "ATOM",
    },
})

// Get deposit history with filters
deposits, err := crypto.GetDeposits(crypto.GetDepositsParams{
    PaginationParams: crypto.PaginationParams{
        Page:  1,
        Limit: 10,
    },
    Symbol: "BTC",
    Status: "complete",
})

// Get available coins
coins, err := crypto.GetCoins(crypto.GetCoinsParams{
    Symbol: "BTC",
})

// Withdraw crypto
withdrawReq := crypto.WithdrawRequest{
    Symbol:  "BTC",
    Network: "BTC",
    Address: "bc1q...",
    Amount:  0.001,
    Memo:    "",
}
txn, err := crypto.Withdraw(withdrawReq)
```

For detailed crypto API documentation, see [crypto/README.md](crypto/README.md)

## 📖 Documentation

- [API Migration Guide](docs/API_MIGRATION_GUIDE.md) - Complete V3 migration guide
- [Crypto API Guide](crypto/README.md) - Crypto V4 endpoints documentation
- [Official Bitkub V3 API](https://github.com/bitkub/bitkub-official-api-docs/blob/master/restful-api.md)
- [Official Bitkub V4 API](https://github.com/bitkub/bitkub-official-api-docs/blob/master/restful-api-v4.md)

## 🔒 Security Best Practices

1. **Never commit API keys** to version control
2. **Use environment variables** for credentials:
   ```bash
   export BTK_APIKEY="your_api_key"
   export BTK_SECRETKEY="your_secret_key"
   ```
3. **Use IP whitelist** in Bitkub API settings
4. **Implement rate limiting** in your application
5. **Monitor API usage** regularly

## ⚡ Rate Limits

| Endpoint Type | Rate Limit |
|---------------|------------|
| Market Data V3 (ticker, trades, etc.) | 100 req/sec |
| Trading Operations | 150-200 req/sec |
| Fiat/User Operations | 20 req/sec |
| Crypto V4 Operations | 250 req/10sec |

See [official documentation](https://github.com/bitkub/bitkub-official-api-docs/blob/master/restful-api.md#rate-limits) for complete information.

## 🧪 Testing

### Run Validation Tests (No API credentials needed)
```bash
go test ./crypto/... -v -run "Validation"
```

### Run Integration Tests (Requires API credentials)
```bash
# Set your credentials
export BTK_APIKEY="your_api_key"
export BTK_SECRETKEY="your_secret_key"

# Run tests
go test ./crypto/... -v
go test ./market/... -v -run "Test"
```

## 📁 Project Structure

```
go-bitkub/
├── bitkub/          # Core API client and authentication
│   ├── bitkub.go    # Initialization and configuration
│   ├── fetch.go     # HTTP client with v3 & v4 support
│   ├── error.go     # Error code mappings
│   ├── status.go    # System status endpoints
│   └── types.go     # Common type definitions
├── market/          # Market API (V3) endpoints
│   ├── symbols.go
│   ├── ticker_v3.go
│   ├── depth_v3.go
│   ├── trades_v3.go
│   ├── place-bid.go
│   ├── place-ask.go
│   ├── cancel-order.go
│   ├── my-order-history.go
│   ├── order-info.go
│   ├── balances.go
│   ├── wallet.go
│   ├── wstoken.go
│   └── ...
├── user/            # User API (V3) endpoints
│   ├── trading-credits.go
│   ├── limits.go
│   └── coin-convert-history.go
├── fiat/            # Fiat API (V3) endpoints
│   ├── accounts.go
│   ├── withdraw.go
│   ├── deposit-history.go
│   └── withdraw-history.go
├── crypto/          # Crypto API (V4) endpoints
│   ├── addresses.go
│   ├── deposits.go
│   ├── withdraws.go
│   ├── coins.go
│   ├── compensations.go
│   └── types.go
├── helper/          # Utility functions
└── docs/            # Documentation
    └── API_MIGRATION_GUIDE.md
```

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 📄 License

MIT License

## 🔗 Links

- [Bitkub Official Website](https://www.bitkub.com/)
- [Bitkub V3 API Documentation](https://github.com/bitkub/bitkub-official-api-docs/blob/master/restful-api.md)
- [Bitkub V4 API Documentation](https://github.com/bitkub/bitkub-official-api-docs/blob/master/restful-api-v4.md)
- [Bitkub Support](https://support.bitkub.com/)

## ⚠️ Disclaimer

This is an unofficial SDK. Use at your own risk. Always test thoroughly before using in production.

---

**Note**: This SDK implements the Bitkub API specification as of November 2025. API specifications are subject to change by Bitkub.

