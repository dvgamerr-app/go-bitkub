# Quick Install Guide

## Installation

### Option 1: Install via `go install` (Recommended)

```bash
go install github.com/dvgamerr-app/go-bitkub/cmd/bitkub@latest
```

The `bitkub` command will be installed in `$GOPATH/bin` (usually `~/go/bin` on Linux/Mac or `%USERPROFILE%\go\bin` on Windows).

**Make sure `$GOPATH/bin` is in your PATH:**

```bash
# Linux/Mac - Add to ~/.bashrc or ~/.zshrc
export PATH=$PATH:$(go env GOPATH)/bin

# Windows PowerShell - Add to $PROFILE
$env:PATH += ";$(go env GOPATH)\bin"
```

### Option 2: Build from source

```bash
git clone https://github.com/dvgamerr-app/go-bitkub.git
cd go-bitkub
go build -o bitkub ./cmd/bitkub
```

### Option 3: Install locally

```bash
git clone https://github.com/dvgamerr-app/go-bitkub.git
cd go-bitkub
go install ./cmd/bitkub
```

## Verify Installation

```bash
bitkub --help
```

## Quick Test

```bash
# Get all market symbols
bitkub market symbols

# Get Bitcoin ticker
bitkub market ticker BTC_THB

# Get order book depth
bitkub market depth BTC_THB --limit 5
```

## Configuration

### Method 1: Environment Variables (Recommended)

Create `.env` file:
```env
API_KEY=your_api_key_here
API_SECRET=your_api_secret_here
```

### Method 2: Command Line Flags

```bash
bitkub -k YOUR_KEY -s YOUR_SECRET market balances
```

### Method 3: Short Flags

```bash
bitkub -k YOUR_KEY -s YOUR_SECRET market balances
```

## Common Commands

```bash
# Market data (no authentication required)
bitkub market symbols                    # List all trading pairs
bitkub market ticker                     # Get all tickers
bitkub market ticker BTC_THB             # Get specific ticker
bitkub market depth BTC_THB --limit 10   # Order book
bitkub market trades BTC_THB --limit 20  # Recent trades

# Account operations (requires API keys)
bitkub -k KEY -s SECRET market balances          # Your balances
bitkub -k KEY -s SECRET market wallet            # Wallet info
bitkub -k KEY -s SECRET market open-orders       # Open orders
bitkub -k KEY -s SECRET market order-history     # Order history

# Trading (requires API keys)
bitkub -k KEY -s SECRET market buy BTC_THB 0.001 3500000   # Buy
bitkub -k KEY -s SECRET market sell BTC_THB 0.001 3600000   # Sell
bitkub -k KEY -s SECRET market cancel BTC_THB ORDER_ID buy       # Cancel

# Crypto operations (requires API keys)
bitkub -k KEY -s SECRET crypto addresses             # Deposit addresses
bitkub -k KEY -s SECRET crypto deposits              # Deposit history
bitkub -k KEY -s SECRET crypto withdraws             # Withdraw history

# Fiat operations (requires API keys)
bitkub -k KEY -s SECRET fiat accounts                # Bank accounts
bitkub -k KEY -s SECRET fiat deposit-history         # Deposits
bitkub -k KEY -s SECRET fiat withdraw-history        # Withdraws

# User information (requires API keys)
bitkub -k KEY -s SECRET user limits                  # Trading limits
bitkub -k KEY -s SECRET user credits                 # Trading credits
```

## Debug Mode

Enable detailed logging:

```bash
bitkub --debug market ticker
```

## Update

```bash
go install github.com/dvgamerr-app/go-bitkub/cmd/bitkub@latest
```

## Uninstall

```bash
rm $(which bitkub)
# or on Windows
Remove-Item (Get-Command bitkub).Source
```

## Help

```bash
bitkub --help                    # General help
bitkub market --help             # Market commands help
bitkub market ticker --help      # Specific command help
```

## Documentation

- [Full CLI Documentation](CLI.md)
- [API Documentation](README.md)

## Support

- GitHub Issues: https://github.com/dvgamerr-app/go-bitkub/issues
- Bitkub API Docs: https://github.com/bitkub/bitkub-official-api-docs
