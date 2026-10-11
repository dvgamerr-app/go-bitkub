# Changelog

All notable changes to this project will be documented in this file.

## [v1.4.0] - 2026-10-11

### Fixed
- CLI public commands (market data and streams) no longer exit when `BTK_APIKEY` and `BTK_SECRET` are unset
- `market.GetTicker(symbol)` and `bitkub market ticker <symbol>` return the requested symbol instead of an empty result
- `market.GetDepth` and `bitkub market depth` use the V3 depth endpoint instead of the unresponsive legacy one

### Added
- `bitkub.ErrMissingCredentials`: secure requests without API credentials fail before any network call with a clear message
- Version-free release archive names so `releases/latest/download/go-bitkub_<Os>_<Arch>.tar.gz` always points at the newest build
- Windows release archives are published as `.zip`
- GitHub Pages landing site and terminal screenshots

### Documentation
- New English and Thai READMEs; the previous README moved to `docs/reference.md`
- Install guide covers prebuilt binaries; removed the nonexistent `-k` flag from examples

## [v1.1.0] - 2025-11-07

### Added
- WebSocket stream commands for trade and ticker data
- Testing workflow for automated testing on main branch

### Changed
- Updated README badges for Go version and added Discord link
- Consolidated server time fetching logic and removed redundant code
- Updated symbol references in tests from btc_thb to kub_thb

### Removed
- Unused test functions for address and withdraw creation
- Logging of message details in tests

### Documentation
- Added streaming command examples to CLI and README

## [v1.0.2] - 2025-11-06

### Fixed
- Increased timeout for API server time requests and added retry logic
- Normalized symbol casing in API requests for consistency

## [v1.0.1] - 2025-11-06

### Added
- Update Go module proxy step in release workflow
- Versioning information in CLI and GoReleaser configuration

### Changed
- Optimized HTTP client initialization and added retry logic for server time fetch

### Fixed
- Updated API key and secret flag defaults to empty strings

## [v1.0.0] - 2025-11-05

### Added
- Complete API V3 implementation
- Market API endpoints (asks, bids, trades, order management)
- Crypto API V4 support (deposits, withdrawals, addresses, coins)
- Fiat account and transaction management (deposit/withdrawal history)
- User limits and coin convert history management
- WebSocket streaming implementation for real-time data
- Connection pooling for HTTP client
- Comprehensive error handling with detailed error codes
- Unit tests for all major modules
- CLI commands for market, crypto, fiat, and user operations
- Release workflow and GoReleaser configuration
- Project documentation (README, CLI.md, AGENTS.md)

### Changed
- Replaced interface{} with any for improved type safety
- Consolidated error handling using GetError type
- Renamed functions for consistency (place-bid → buy, place-ask → sell)
- Updated import paths to use dvgamerr-app/go-bitkub
- Improved output formatting with support for JSON, JSONL, and table formats
- Optimized request handling and API parameter structures

### Fixed
- API endpoints to include '/api' prefix for consistency
- Command flags for consistency and improved usage clarity
- Header formatting in Quick Start section
- Signature encoding issues
- Balance retrieval and wallet query methods

### Documentation
- Comprehensive README with usage examples
- CLI documentation with command references
- WebSocket streaming examples
- Project structure and guidelines (AGENTS.md)

