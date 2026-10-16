# go-ens build helpers.
#
# golangci-lint v2.12.2 and abigen@v1.17.5 do not build on Go 1.27, so lint
# and generation pin the Go 1.25 toolchain; tests run with whatever Go is
# installed (CI covers 1.25 and stable).

GOTOOLCHAIN_PIN ?= go1.25.0
GO_PINNED := GOTOOLCHAIN=$(GOTOOLCHAIN_PIN) go
GOLANGCI := GOTOOLCHAIN=$(GOTOOLCHAIN_PIN) golangci-lint

.PHONY: all build test test-live test-sepolia test-devnet lint fmt generate check-generated tidy

all: build test lint

build:
	go build ./...

## Hermetic tests: unit plus the in-process simulated chain.
test:
	go test -short -race -count=1 ./...

## Live tests against mainnet (GO_ENS_TEST_RPC, public fallback).
test-live:
	go test -count=1 -run 'TestLive' ./...

## Live tests against Sepolia, including the write path when
## GO_ENS_TEST_SEPOLIA_KEY is set.
test-sepolia:
	go test -count=1 -run 'TestSepolia' ./...

lint:
	$(GOLANGCI) run ./...

fmt:
	$(GOLANGCI) fmt ./...

## Re-vendor ABIs from the pinned contracts-v2 revision and regenerate bindings.
## Pass VENDOR_LOCAL=<contracts-v2>/contracts/deployments to work offline.
generate:
	$(GO_PINNED) run ./internal/cmd/vendorabi -root . $(if $(VENDOR_LOCAL),-local $(VENDOR_LOCAL)) $(if $(VENDOR_LOCAL_V1),-local-v1 $(VENDOR_LOCAL_V1)) -bytecode
	cd contracts && GOTOOLCHAIN=$(GOTOOLCHAIN_PIN) go generate $$(go list ./v1/... ./v2/... | grep -v 'contracts/v2$$')
	chmod 644 contracts/*/*/contract.go

## Fail if the vendored ABIs or generated bindings are stale.
check-generated:
	$(GO_PINNED) run ./internal/cmd/vendorabi -root . -check
	cd contracts && GOTOOLCHAIN=$(GOTOOLCHAIN_PIN) go generate $$(go list ./v1/... ./v2/... | grep -v 'contracts/v2$$')
	chmod 644 contracts/*/*/contract.go
	git diff --exit-code -- contracts/

tidy:
	go mod tidy
	git diff --exit-code go.mod go.sum
