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

### Option 2: Download a prebuilt binary

Each release ships archives for Linux and macOS (amd64 and arm64) and Windows (amd64). The `releases/latest/download` links below always point at the newest release.

```bash
# Linux x86_64 (other targets: Linux_arm64, Darwin_x86_64, Darwin_arm64)
curl -fsSL https://github.com/dvgamerr-app/go-bitkub/releases/latest/download/go-bitkub_Linux_x86_64.tar.gz | tar -xz bitkub
sudo mv bitkub /usr/local/bin/
```

```powershell
# Windows PowerShell
Invoke-WebRequest https://github.com/dvgamerr-app/go-bitkub/releases/latest/download/go-bitkub_Windows_x86_64.zip -OutFile bitkub.zip
Expand-Archive bitkub.zip -DestinationPath "$env:USERPROFILE\bitkub"
$env:PATH += ";$env:USERPROFILE\bitkub"
```

Verify downloads against `checksums.txt` from the same release.

### Option 3: Build from source

```bash
git clone https://github.com/dvgamerr-app/go-bitkub.git
cd go-bitkub
go build -o bitkub ./cmd/bitkub
```

### Option 4: Install locally

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

Market data commands (`market symbols`, `ticker`, `trades`, `depth`, `bids`, `asks`, `history` and both `stream` commands) need no configuration. Account, trading, crypto, fiat and user commands need a Bitkub API key and secret; without them they stop with `missing API credentials`.

### Method 1: Environment Variables (Recommended)

Export the variables, or create a `.env` file in the directory you run `bitkub` from:
```env
BTK_APIKEY=your_api_key_here
BTK_SECRET=your_api_secret_here
```

### Method 2: Command Line Flags

```bash
bitkub --key YOUR_KEY --secret YOUR_SECRET market balances
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
bitkub --key YOUR_KEY --secret YOUR_SECRET market balances          # Your balances
bitkub --key YOUR_KEY --secret YOUR_SECRET market wallet            # Wallet info
bitkub --key YOUR_KEY --secret YOUR_SECRET market open-orders       # Open orders
bitkub --key YOUR_KEY --secret YOUR_SECRET market order-history     # Order history

# Trading (requires API keys)
bitkub --key YOUR_KEY --secret YOUR_SECRET market buy BTC_THB 0.001 3500000   # Buy
bitkub --key YOUR_KEY --secret YOUR_SECRET market sell BTC_THB 0.001 3600000   # Sell
bitkub --key YOUR_KEY --secret YOUR_SECRET market cancel BTC_THB ORDER_ID buy       # Cancel

# Crypto operations (requires API keys)
bitkub --key YOUR_KEY --secret YOUR_SECRET crypto addresses             # Deposit addresses
bitkub --key YOUR_KEY --secret YOUR_SECRET crypto deposits              # Deposit history
bitkub --key YOUR_KEY --secret YOUR_SECRET crypto withdraws             # Withdraw history

# Fiat operations (requires API keys)
bitkub --key YOUR_KEY --secret YOUR_SECRET fiat accounts                # Bank accounts
bitkub --key YOUR_KEY --secret YOUR_SECRET fiat deposit-history         # Deposits
bitkub --key YOUR_KEY --secret YOUR_SECRET fiat withdraw-history        # Withdraws

# User information (requires API keys)
bitkub --key YOUR_KEY --secret YOUR_SECRET user limits                  # Trading limits
bitkub --key YOUR_KEY --secret YOUR_SECRET user credits                 # Trading credits
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
