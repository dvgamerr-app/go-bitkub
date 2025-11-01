# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

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
