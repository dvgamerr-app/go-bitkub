# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [2.0.0] - 2025-11-01

### Added - Complete API V3 Implementation

#### Non-Secure Endpoints
- `bitkub.GetStatus()` - Get endpoint status
- `bitkub.GetServerTime()` - Get server timestamp in milliseconds (V3)
- `market.GetSymbols()` - List all available symbols
- `market.GetTicker()` - Get ticker information
- `market.GetBids()` - List open buy orders
- `market.GetAsks()` - List open sell orders
- `market.GetDepth()` - Get market depth
- `market.GetTrades()` - List recent trades
- `market.GetTradingViewHistory()` - Get TradingView historical data

#### Trading Endpoints (Secure)
- `market.PlaceBid()` - Create buy order with post-only support
- `market.PlaceAsk()` - Create sell order with post-only support
- `market.CancelOrder()` - Cancel an open order
- `market.GetMyOpenOrders()` - List all open orders
- `market.GetMyOrderHistory()` - Get order history with keyset pagination
- `market.GetOrderInfo()` - Get detailed order information
- `market.GetWSToken()` - Get WebSocket authentication token
- `market.GetWallet()` - Get user available balances
- `market.GetBalances()` - Get detailed balances (available + reserved)

#### User Endpoints (Secure)
- `user.GetTradingCredits()` - Check trading credit balance
- `user.GetUserLimits()` - Check deposit/withdraw limitations and usage
- `user.GetCoinConvertHistory()` - List coin convert history with pagination

#### Fiat Endpoints (Secure)
- `fiat.GetAccounts()` - List all approved bank accounts
- `fiat.Withdraw()` - Make fiat withdrawal
- `fiat.GetDepositHistory()` - List fiat deposit history
- `fiat.GetWithdrawHistory()` - List fiat withdrawal history

### Enhanced Features

#### Type Safety
- Added comprehensive struct types for all API requests and responses
- Added common types: `OrderSide`, `OrderType`, `OrderStatus`, `TransactionStatus`
- Proper JSON marshaling/unmarshaling throughout

#### Pagination Support
- Keyset-based pagination for order history
- Cursor encoding/decoding utilities
- `Pagination` struct with keyset fields

#### Documentation
- Updated API guide with current endpoints
- Comprehensive README with usage examples
- Clean and modern documentation structure

#### Error Handling
- Improved error messages with context
- HTTP status code validation
- API error code mapping

### Fixed
- Connection pooling optimization in HTTP client
- Proper timestamp handling (milliseconds only)
- Query parameter encoding for GET requests
- Signature generation for both GET and POST methods

## [1.0.0] - Previous Version

### Added
- **Bitkub API v4 Support**: Complete implementation of Bitkub Crypto API v4 endpoints
- New `crypto` package with the following features:
  - `GetAddresses()` - List all crypto addresses with pagination
  - `CreateAddress()` - Generate new crypto addresses (requires `is_deposit` permission)
  - `GetDeposits()` - List crypto deposit history with filters
  - `GetWithdraws()` - List crypto withdrawal history with filters
  - `CreateWithdraw()` - Make withdrawals to trusted addresses (requires `is_withdraw` permission)
  - `GetCoins()` - Get available coins and networks information
  - `GetCompensations()` - List crypto compensations history
- Enhanced signature generation to support both v3 and v4 API formats
- New `FetchSecureV4()` function for v4 API endpoints
- `ResponseAPIV4` type for v4 API responses
- Comprehensive type definitions for all v4 API request and response structures
- Complete test suite for crypto package
- Detailed documentation in `crypto/README.md`
- Example code demonstrating all crypto API features in `examples/crypto_v4_example.go`

### Changed
- Updated `bitkub/fetch.go` to include request body in signature calculation for v4 API
- Enhanced main README.md with v4 API information and examples
- Updated `main.go` with examples of using both v3 and v4 APIs

### Technical Details
- All v4 endpoints follow the same pattern as existing v3 endpoints
- Maintains backward compatibility with existing v3 API code
- Follows Go project structure standards with proper package organization
- All code passes Go linting and compilation checks
- Rate limiting: 250 requests per 10 seconds for v4 crypto endpoints

## [Previous Versions]

### Market API v3 (Existing)
- Get wallet balances
- Get user limits
- Get trading credits
- Place orders
- View order history
- Get ticker information
