# Repository notes

## Project conventions

- This repository is a Go SDK and Cobra CLI for Bitkub V3/V4 APIs.
- Follow `AGENTS.md`: keep code self-explanatory, do not add explanatory comments or inline documentation, and prefer modifying existing files over adding files.
- Keep public endpoint signatures and exported response types backward compatible.
- Public endpoints must not depend on API credentials, server-time synchronization, or HMAC signing.
- Secure endpoints must synchronize server time and set `X-BTK-TIMESTAMP`, `X-BTK-APIKEY`, and `X-BTK-SIGN`.
- Centralize Crypto V4 response decoding with `fetchV4` and query composition with the helpers in `crypto/types.go`.
- Run `gofmt` only on changed Go files. Before handoff, run package tests, `go test ./...`, `go vet ./...`, `go build ./...`, and `git diff --check`.
- Integration-tagged REST tests call live Bitkub endpoints and require credentials. The untagged `stream` tests also open live WebSocket connections.

## Optimization record — 2026-08-02

- Non-secure HTTP requests no longer call the server-time endpoint or calculate an unused HMAC signature. They now perform only the requested public API round trip.
- JSON decoding now targets the caller-provided response directly instead of passing a redundant pointer to the interface value.
- HTTP transport failures safely handle a nil response instead of dereferencing `resp.StatusCode` and panicking.
- Crypto V4 endpoints share one generic request/decode path and common pagination, date-range, and URL query helpers.
- Added offline core regression tests for public request behavior, secure authentication headers, and transport-error handling.
- Extended existing Crypto validation tests to cover shared query generation.
- Updated `README.md` with the optimized public-request behavior, accurate network requirements, development checks, and the current repository tree.

## Verification record — 2026-08-02

- Baseline working tree: clean.
- Baseline `go test ./...`: pass.
- Baseline `go vet ./...`: pass.
- Baseline `go build ./...`: blocked by VCS stamping because Go's Git subprocess does not inherit the command-level `safe.directory`; verify with `go build -buildvcs=false ./...` in this environment.
- `golangci-lint`: not installed in this environment; do not install or mutate dependencies solely for verification.
- Final `gofmt -d` on changed Go files: pass with no diff.
- Final `go test -count=1 ./bitkub ./crypto`: pass.
- Final `go test -count=1 ./...`: pass, including the live WebSocket smoke tests.
- Final `go vet ./...`: pass.
- Final `go build -buildvcs=false ./...`: pass; standard `go build ./...` remains blocked only by the VCS-stamping limitation above.
- Final `git diff --check`: pass.
- The touched production Go files decreased from 528 to 477 lines while adding shared behavior and preserving endpoint signatures.
