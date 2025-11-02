# AGENTS.md

## Project Overview
Go SDK for Bitkub API - supports cryptocurrency trading, wallet management, and market data retrieval

## Code Structure

```
go-bitkub/
├── bitkub/       # Core API client (authentication, fetch, error handling)
├── market/       # Trading API (orders, balances, ticker, depth)
├── crypto/       # Crypto operations (deposit, withdraw, addresses)
├── fiat/         # Fiat operations (bank accounts, deposit/withdraw)
├── user/         # User info (limits, trading credits, history)
└── utils/        # Utilities (dotenv, error handling)
```

## Guidelines

- No explanatory comments in code
- No inline documentation
- Always use or modify existing code before creating new files
- Keep code clean and self-explanatory 