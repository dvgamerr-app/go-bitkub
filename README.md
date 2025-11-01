# go-bitkub

Go client library for Bitkub API (v3 & v4)

![](./docs/example.png)

## Features

### Market API (v3)
- Get wallet balances
- Get user limits
- Get trading credits
- Place orders
- View order history
- Get ticker information

### Crypto API (v4) ✨ NEW
- 📋 List crypto addresses with pagination
- ➕ Generate new crypto addresses
- 💰 View deposit history
- 💸 View withdrawal history
- 🚀 Create withdrawals to trusted addresses
- 🪙 Get available coins and networks
- 🎁 View compensations history

## Installation

```bash
go get github.com/dvgamerr-app/go-bitkub
```

## Quick Start

### Initialize

```go
import (
    "github.com/dvgamerr-app/go-bitkub/bitkub"
    "github.com/dvgamerr-app/go-bitkub/market"
    "github.com/dvgamerr-app/go-bitkub/crypto"
)

func main() {
    // Initialize with API credentials
    bitkub.Initlizer("YOUR_API_KEY", "YOUR_SECRET_KEY")
    
    // Or use environment variables BTK_APIKEY and BTK_SECRETKEY
    bitkub.Initlizer()
}
```

### Market API (v3) Examples

```go
// Get wallet balances
balances, err := market.GetBalances()
if err != nil {
    log.Fatal(err)
}

// Get user limits
limits, err := market.GetUserLimits()
if err != nil {
    log.Fatal(err)
}
```

### Crypto API (v4) Examples

```go
// List crypto addresses
addresses, err := crypto.GetAddresses(crypto.GetAddressesParams{
    Page:  1,
    Limit: 10,
})

// Get deposit history
deposits, err := crypto.GetDeposits(crypto.GetDepositsParams{
    Page:   1,
    Limit:  10,
    Symbol: "BTC",
})

// Get available coins
coins, err := crypto.GetCoins(crypto.GetCoinsParams{
    Symbol: "BTC",
})
```

For detailed crypto API documentation, see [crypto/README.md](crypto/README.md)

## Testing

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

## Project Structure

```
go-bitkub/
├── bitkub/          # Core API client and authentication
│   ├── bitkub.go    # Initialization and response types
│   ├── fetch.go     # HTTP client with v3 & v4 signature support
│   └── error.go     # Error code mappings
├── market/          # Market API (v3) endpoints
│   ├── balances.go
│   ├── limit.go
│   ├── wallet.go
│   └── ...
├── crypto/          # Crypto API (v4) endpoints ✨ NEW
│   ├── addresses.go
│   ├── deposits.go
│   ├── withdraws.go
│   ├── coins.go
│   ├── compensations.go
│   └── types.go
└── helper/          # Utility functions
```

## API Documentation

- [Bitkub API v3 Documentation](https://github.com/bitkub/bitkub-official-api-docs/blob/master/restful-api.md)
- [Bitkub API v4 Documentation](https://github.com/bitkub/bitkub-official-api-docs/blob/master/restful-api-v4.md)

## Rate Limits

- **v3 endpoints**: Varies by endpoint
- **v4 crypto endpoints**: 250 requests per 10 seconds

## License

MIT License

