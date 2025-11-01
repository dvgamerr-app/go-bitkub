# Bitkub Crypto API v4

This package provides Go client implementation for Bitkub API v4 crypto endpoints.

## Features

- ✅ List crypto addresses (GET /api/v4/crypto/addresses)
- ✅ Generate new crypto address (POST /api/v4/crypto/addresses)
- ✅ List deposit history (GET /api/v4/crypto/deposits)
- ✅ List withdrawal history (GET /api/v4/crypto/withdraws)
- ✅ Create withdrawal (POST /api/v4/crypto/withdraws)
- ✅ Get available coins (GET /api/v4/crypto/coins)
- ✅ List compensations history (GET /api/v4/crypto/compensations)

## Installation

```bash
go get github.com/dvgamerr-app/go-bitkub
```

## Usage

### Initialize

```go
import (
    "github.com/dvgamerr-app/go-bitkub/bitkub"
    "github.com/dvgamerr-app/go-bitkub/crypto"
)

func main() {
    // Initialize with API credentials
    bitkub.Initlizer("YOUR_API_KEY", "YOUR_SECRET_KEY")
    
    // Or use environment variables BTK_APIKEY and BTK_SECRETKEY
    bitkub.Initlizer()
}
```

### Common Parameter Types

The library uses reusable parameter structs for consistent API patterns:

- `PaginationParams` - For page and limit parameters
- `DateRangeParams` - For created_start and created_end filters
- `SymbolNetworkParams` - For symbol and network filters

### Get Crypto Addresses

```go
addresses, err := crypto.GetAddresses(crypto.GetAddressesParams{
    PaginationParams: crypto.PaginationParams{
        Page:  1,
        Limit: 10,
    },
    SymbolNetworkParams: crypto.SymbolNetworkParams{
        Symbol:  "ATOM",
        Network: "ATOM",
    },
    Memo: "", // Optional
})
if err != nil {
    log.Fatal(err)
}

for _, addr := range addresses.Items {
    fmt.Printf("Address: %s, Memo: %s\n", addr.Address, addr.Memo)
}
```

### Create New Address

**Required Permission: `is_deposit`**

```go
newAddr, err := crypto.CreateAddress(crypto.CreateAddressRequest{
    Symbol:  "ATOM",
    Network: "ATOM",
})
if err != nil {
    log.Fatal(err)
}

fmt.Printf("New address: %s\n", newAddr[0].Address)
```

### Get Deposit History

```go
deposits, err := crypto.GetDeposits(crypto.GetDepositsParams{
    PaginationParams: crypto.PaginationParams{
        Page:  1,
        Limit: 10,
    },
    DateRangeParams: crypto.DateRangeParams{
        CreatedStart: "2025-01-01T00:00:00.000Z",
        CreatedEnd:   "2025-01-31T23:59:59.999Z",
    },
    Symbol: "BTC",
    Status: "complete",
})
if err != nil {
    log.Fatal(err)
}

for _, deposit := range deposits.Items {
    fmt.Printf("Hash: %s, Amount: %s %s, Status: %s\n", 
        deposit.Hash, deposit.Amount, deposit.Symbol, deposit.Status)
}
```

### Get Withdrawal History

```go
withdraws, err := crypto.GetWithdraws(crypto.GetWithdrawsParams{
    PaginationParams: crypto.PaginationParams{
        Page:  1,
        Limit: 10,
    },
    DateRangeParams: crypto.DateRangeParams{
        CreatedStart: "2025-01-01T00:00:00.000Z",
        CreatedEnd:   "2025-01-31T23:59:59.999Z",
    },
    Symbol: "ETH",
    Status: "complete",
})
if err != nil {
    log.Fatal(err)
}

for _, withdraw := range withdraws.Items {
    fmt.Printf("TxnID: %s, Amount: %s %s, Fee: %s, Status: %s\n", 
        withdraw.TxnID, withdraw.Amount, withdraw.Symbol, withdraw.Fee, withdraw.Status)
}
```

### Create Withdrawal

**Required Permission: `is_withdraw`**

```go
withdraw, err := crypto.CreateWithdraw(crypto.CreateWithdrawRequest{
    Symbol:  "RDNT",
    Amount:  "2.00000000",
    Address: "0xDaCd17d1E77604aaFB6e47F5Ffa1F7E35F83fDa7",
    Network: "ARB",
    Memo:    "", // Optional
})
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Withdrawal created: TxnID=%s, Amount=%s %s, Fee=%s\n",
    withdraw.TxnID, withdraw.Amount, withdraw.Symbol, withdraw.Fee)
```

### Get Available Coins

```go
// Get all coins
coins, err := crypto.GetCoins(crypto.GetCoinsParams{})

// Or filter by symbol and network
coins, err := crypto.GetCoins(crypto.GetCoinsParams{
    Symbol:  "BTC",
    Network: "BTC",
})
if err != nil {
    log.Fatal(err)
}

for _, coin := range coins.Items {
    fmt.Printf("Coin: %s (%s)\n", coin.Name, coin.Symbol)
    for _, network := range coin.Networks {
        fmt.Printf("  Network: %s, Min: %s, Fee: %s\n",
            network.Network, network.WithdrawMin, network.WithdrawFee)
    }
}
```

### Get Compensations History

```go
compensations, err := crypto.GetCompensations(crypto.GetCompensationsParams{
    PaginationParams: crypto.PaginationParams{
        Page:  1,
        Limit: 10,
    },
    Symbol: "XRP",
    Type:   "COMPENSATE", // or "DECOMPENSATE"
    Status: "COMPLETED",  // or "PENDING"
})
if err != nil {
    log.Fatal(err)
}

for _, comp := range compensations.Items {
    fmt.Printf("TxnID: %s, Type: %s, Amount: %s %s, Status: %s\n",
        comp.TxnID, comp.Type, comp.Amount, comp.Symbol, comp.Status)
}
```

## API Reference

### Transaction Status

#### Deposit Status
- `pending` - Pending confirmation
- `rejected` - Rejected
- `complete` - Completed

#### Withdrawal Status
- `pending` - Pending
- `processing` - Processing
- `reported` - Reported
- `rejected` - Rejected
- `complete` - Completed

#### Compensation Status
- `PENDING` - Pending
- `COMPLETED` - Completed

### Compensation Types
- `COMPENSATE` - Compensation added
- `DECOMPENSATE` - Compensation deducted

## Error Handling

The API returns error codes as documented in the [Bitkub API documentation](https://github.com/bitkub/bitkub-official-api-docs/blob/master/restful-api-v4.md#error-codes).

Common error codes:
- `A1000-CW` - Unauthorized Access
- `A1001-CW` - Permission denied
- `V1003-CW` - Invalid signature
- `B1000-CW` - User account is suspended
- `B1001-CW` - Network is disabled
- `B1003-CW` - Insufficient balance
- `B1012-CW` - Address is not trusted
- `B1013-CW` - Withdrawal is frozen
- `B1016-CW` - Deposit is frozen

## Rate Limits

Rate limit: **250 requests per 10 seconds** per user for all `/api/v4/crypto/*` endpoints.

When rate limit is exceeded, HTTP 429 (Too Many Requests) is returned and the request will be blocked for 30 seconds.

## Notes

1. **Coin/Network Symbols**: Use official symbols from https://www.bitkub.com/fee/cryptocurrency
   - Terra Classic (LUNC) uses symbol: `LUNA`
   - Terra 2.0 (LUNA) uses symbol: `LUNA2`

2. **Timestamp Format**: Use ISO 8601 format (e.g., `2025-01-11T10:00:00.000Z`)

3. **Pagination**: Default page = 1, default limit = 100, max limit = 200

4. **Permissions**: 
   - Creating addresses requires `is_deposit` permission
   - Creating withdrawals requires `is_withdraw` permission

## Testing

Run tests with:

```bash
go test ./crypto/...
```

Note: Most tests are skipped by default as they require valid API credentials.

## License

MIT License
