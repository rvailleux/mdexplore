# mdexplore Makefile

# Version handling
VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
BUILD_TIME = $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
COMMIT = $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")

# Build flags
LDFLAGS = -ldflags "-X main.version=$(VERSION) -X main.buildTime=$(BUILD_TIME) -X main.commit=$(COMMIT)"

# Go settings
GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOINSTALL = $(GOCMD) install
GOTEST = $(GOCMD) test
GOGET = $(GOCMD) get

# Binary name
BINARY_NAME = mdexplore
BINARY_PATH = ./cmd/mdexplore

.PHONY: all build install clean test version

all: build

# Build the binary with version info
build:
	@echo "Building $(BINARY_NAME) version $(VERSION)..."
	$(GOBUILD) $(LDFLAGS) -o $(BINARY_NAME) $(BINARY_PATH)
	@echo "Build complete: ./$(BINARY_NAME)"

# Install to GOPATH/bin
install:
	@echo "Installing $(BINARY_NAME) version $(VERSION)..."
	$(GOINSTALL) $(LDFLAGS) $(BINARY_PATH)
	@echo "Installed to $$(go env GOPATH)/bin/$(BINARY_NAME)"

# Build and run
run: build
	./$(BINARY_NAME)

# Clean build artifacts
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	@echo "Cleaned build artifacts"

# Run all tests
test:
	$(GOTEST) ./...

# Run tests with coverage
test-coverage:
	$(GOTEST) -cover ./...

# Run tests verbosely
test-verbose:
	$(GOTEST) -v ./...

# Run benchmarks
benchmark:
	$(GOTEST) ./tests/benchmark -bench=.

# Display version info
version:
	@echo "Version: $(VERSION)"
	@echo "Build Time: $(BUILD_TIME)"
	@echo "Commit: $(COMMIT)"

# Quick build for development (no version info)
dev:
	$(GOBUILD) -o $(BINARY_NAME) $(BINARY_PATH)

# Build for release (with version bump)
release:
	$(MAKE) build VERSION=$(VERSION)
	@echo "Release build complete: $(VERSION)"

# Install locally for development
dev-install:
	$(GOBUILD) -o $$(go env GOPATH)/bin/$(BINARY_NAME) $(BINARY_PATH)
	@echo "Installed development build to $$(go env GOPATH)/bin/$(BINARY_NAME)"
