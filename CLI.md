# Bitkub CLI

A command-line interface for interacting with Bitkub API, built with Go, Cobra, and Zerolog.

## Features

- 🎯 **Market Commands**: Trading, orders, balances, and market data
- 💰 **Crypto Commands**: Deposits, withdrawals, and addresses management
- 💵 **Fiat Commands**: Bank accounts, deposits, and withdrawals
- 👤 **User Commands**: User information, limits, and trading credits
- 📊 **Beautiful Logging**: Clean and readable output with zerolog

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
go build -o bitkub ./cmd/bitkub

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
bitkub-cli --key YOUR_KEY --secret YOUR_SECRET [command]
```

### 3. Short Flags

```bash
bitkub-cli -k YOUR_KEY -s YOUR_SECRET [command]
```

## Usage

### General Commands

```bash
# Show help
bitkub --help

# Enable debug mode
bitkub --debug [command]
```

### Market Commands

```bash
# Get all trading symbols
bitkub market symbols

# Get ticker information
bitkub market ticker               # All symbols
bitkub market ticker THB_BTC       # Specific symbol

# Get market depth
bitkub market depth THB_BTC --limit 10

# Get recent trades
bitkub market trades THB_BTC --limit 20

# Get account balances
bitkub market balances

# Get wallet information
bitkub market wallet

# Get open orders
bitkub market open-orders
bitkub market open-orders THB_BTC

# Get order history
bitkub market order-history --page 1 --limit 20
bitkub market order-history THB_BTC

# Get order information
bitkub market order-info THB_BTC ORDER_ID buy

# Place orders
bitkub market place-bid THB_BTC 0.001 1000000
bitkub market place-ask THB_BTC 0.001 1200000

# Cancel order
bitkub market cancel THB_BTC ORDER_ID buy

# Get user limits
bitkub market limits

# Get trading credits
bitkub market credits

# Get WebSocket token
bitkub market wstoken
```

### Crypto Commands

```bash
# Get coin information
bitkub crypto coins
bitkub crypto coins --symbol BTC
bitkub crypto coins --symbol BTC --network BTC

# Get deposit addresses
bitkub crypto addresses --page 1 --limit 20

# Create new deposit address
bitkub crypto create-address BTC BTC

# Get deposit history
bitkub crypto deposits --page 1 --limit 20
bitkub crypto deposits --symbol BTC

# Get withdraw history
bitkub crypto withdraws --page 1 --limit 20
bitkub crypto withdraws --symbol BTC

# Create withdrawal
bitkub crypto withdraw BTC 0.001 ADDRESS NETWORK --memo MEMO

# Get compensation history
bitkub crypto compensations --page 1 --limit 20
```

### Fiat Commands

```bash
# Get bank accounts
bitkub fiat accounts --page 1 --limit 20

# Get deposit history
bitkub fiat deposit-history --page 1 --limit 20

# Get withdraw history
bitkub fiat withdraw-history --page 1 --limit 20

# Create withdrawal
bitkub fiat withdraw BANK_ACCOUNT_ID 1000
```

### User Commands

```bash
# Get user limits
bitkub user limits

# Get trading credits
bitkub user credits

# Get coin convert history
bitkub user coin-convert-history --page 1 --limit 20
```

## Examples

### Check BTC price

```bash
bitkub market ticker THB_BTC
```

Output:
```
12:00AM INF Ticker change=2.5 high24h=1250000 last=1200000 low24h=1180000 symbol=THB_BTC volume=150.5
```

### Get your balance

```bash
bitkub -k YOUR_KEY -s YOUR_SECRET market balances
```

Output:
```
12:00AM INF Balance available=1000000 coin=THB reserved=0
12:00AM INF Balance available=0.5 coin=BTC reserved=0
```

### Place a buy order

```bash
bitkub -k YOUR_KEY -s YOUR_SECRET market place-bid THB_BTC 0.001 1200000
```

Output:
```
12:00AM INF Bid Placed amount=0.001 credit=0 fee=3 id=12345 rate=1200000 timestamp=1699000000 type=limit
```

## Debug Mode

Enable debug logging to see detailed request/response information:

```bash
bitkub --debug market ticker
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
- **Bitkub API**: v3 and v4 endpoints

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
