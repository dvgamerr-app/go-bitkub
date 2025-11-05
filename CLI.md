# go run . CLI

A command-line interface for interacting with go run . API, built with Go, Cobra, and Zerolog.

## Features

- 🎯 **Market Commands**: Trading, orders, balances, historical data, and market data
- 💰 **Crypto Commands**: Deposits, withdrawals, and addresses management
- 💵 **Fiat Commands**: Bank accounts, deposits, and withdrawals
- 👤 **User Commands**: User information, limits, and trading credits
- 📊 **Output Formats**: JSON, JSONL, and text format support
- 🪵 **Beautiful Logging**: Clean and readable output with zerolog

## Installation

### Install via go install (Recommended)

```bash
go install github.com/dvgamerr-app/go-bitkub/cmd/bitkub@latest
```

After installation, the `bitkub` command will be available in your `$GOPATH/bin` directory. Make sure this directory is in your PATH.

### Build from source

```bash
# Clone the repository
git clone https://github.com/dvgamerr-app/go-bitkub.git
cd go-bitkub

# Build
go build -o go run . ./cmd/bitkub

# Or install locally
go install ./cmd/bitkub
```

### Run directly (Development)

```bash
go run ./cmd/bitkub/main.go [command]
```

## Configuration

Set your API credentials using one of these methods:

### 1. Environment Variables

Create a `.env` file:

```env
API_KEY=your_api_key
API_SECRET=your_api_secret
```

### 2. Command Line Flags

```bash
go run . --key YOUR_KEY --secret YOUR_SECRET [command]
```

### 3. Short Flags

```bash
go run . -k YOUR_KEY --secret YOUR_SECRET [command]
```

## Usage

### General Commands

```bash
# Show help
go run . --help

# Enable debug mode
go run . --debug [command]

# Output formats
go run . --format json [command]   # JSON output
go run . --format jsonl [command]  # JSONL output (one JSON per line)
go run . --format text [command]   # Text output (default)
```

### Market Commands

```bash
# Get all trading symbols
go run . market symbols

# Get ticker information
go run . market ticker               # All symbols
go run . market ticker btc_thb       # Specific symbol

# Get market depth
go run . market depth btc_thb --limit 10

# Get recent trades
go run . market trades btc_thb --limit 20

# Get historical data (TradingView)
go run . market history btc_thb                                    # Last 24h with 1D resolution
go run . market history btc_thb -r 1                     # 1 minute candles
go run . market history btc_thb -r 60 --from 1234567890  # Custom timeframe

# Get account balances
go run . market balances

# Get wallet information
go run . market wallet

# Get open orders
go run . market open-orders btc_thb

# Get order history
go run . market order-history btc_thb
go run . market order-history btc_thb --page 1 --limit 20

# Get order information
go run . market order-info btc_thb 446958802

# Place orders
go run . market buy kub_thb 10 1
go run . market sell kub_thb 10 1

# Cancel order
go run . market cancel btc_thb 446958802

# Get user limits
go run . market limits

# Get trading credits
go run . market credits

# Get WebSocket token
go run . market wstoken
```

### Crypto Commands

```bash
# Get coin information
go run . crypto coins
go run . crypto coins -s btc
go run . crypto coins -s btc
go run . crypto coins -s btc --network btc
go run . crypto coins -s btc -n btc

# Get deposit addresses
go run . crypto addresses --page 1 --limit 20

# Create new deposit address
go run . crypto create-address btc btc

# Get deposit history
go run . crypto deposits --page 1 --limit 20
go run . crypto deposits -s btc
go run . crypto deposits -s btc

# Get withdraw history
go run . crypto withdraws --page 1 --limit 20
go run . crypto withdraws -s btc
go run . crypto withdraws -s btc

# Create withdrawal
go run . crypto withdraw btc 0.001 ADDRESS NETWORK --memo MEMO

# Get compensation history
go run . crypto compensations --page 1 --limit 20
go run . crypto compensations -s btc
```

### Fiat Commands

```bash
# Get bank accounts
go run . fiat accounts --page 1 --limit 20

# Get deposit history
go run . fiat deposit-history --page 1 --limit 20

# Get withdraw history
go run . fiat withdraw-history --page 1 --limit 20

# Create withdrawal
go run . fiat withdraw BANK_ACCOUNT_ID 1000
```

### User Commands

```bash
# Get user limits
go run . user limits

# Get trading credits
go run . user credits

# Get coin convert history
go run . user coin-convert-history --page 1 --limit 20
```

## Examples

### Check BTC price

```bash
go run . market ticker btc_thb
```

Output:
```
12:00AM INF Ticker change=2.5 high24h=1250000 last=1200000 low24h=1180000 symbol=btc_thb volume=150.5
```

### Get your balance

```bash
go run . -k YOUR_KEY --secret YOUR_SECRET market balances
```

Output:
```
12:00AM INF Balance available=1000000 coin=THB reserved=0
12:00AM INF Balance available=0.5 coin=BTC reserved=0
```

### Place a buy order

```bash
go run . -k YOUR_KEY --secret YOUR_SECRET market buy btc_thb 0.001 1200000
```

Output:
```
12:00AM INF Bid Placed amount=0.001 credit=0 fee=3 id=12345 rate=1200000 timestamp=1699000000 type=limit
```

## Debug Mode

Enable debug logging to see detailed request/response information:

```bash
go run . --debug market ticker
```

## Features by Module

### Market Module (17 commands)
- ✅ Symbols, Ticker, Trades, Depth
- ✅ Asks, Bids, Balances, Wallet
- ✅ Open Orders, Order History, Order Info
- ✅ Place Bid/Ask, Cancel Order
- ✅ User Limits, Trading Credits, WS Token

### Crypto Module (7 commands)
- ✅ Coins, Addresses, Create Address
- ✅ Deposits, Withdraws, Create Withdraw
- ✅ Compensations

### Fiat Module (4 commands)
- ✅ Accounts, Deposit History
- ✅ Withdraw History, Withdraw

### User Module (3 commands)
- ✅ Limits, Trading Credits
- ✅ Coin Convert History

## Tech Stack

- **Go**: Programming language
- **Cobra**: CLI framework
- **Zerolog**: Structured logging
- **go run . API**: v3 and v4 endpoints

## Development

### Project Structure

```
go-bitkub/
├── cmd/
│   ├── bitkub/      # CLI entry point (go install target)
│   │   └── main.go
│   ├── root.go      # Root command
│   ├── market.go    # Market commands
│   ├── crypto.go    # Crypto commands
│   ├── fiat.go      # Fiat commands
│   └── user.go      # User commands
├── bitkub/          # Core API client
├── market/          # Market API
├── crypto/          # Crypto API
├── fiat/            # Fiat API
├── user/            # User API
└── main.go          # Library entry point
```

### Adding New Commands

1. Open the relevant command file (e.g., `cmd/market.go`)
2. Add your command following the existing pattern
3. Register it in the `init()` function
4. Build and test

## Contributing

Feel free to submit issues and pull requests!

## License

See LICENSE file for details.
