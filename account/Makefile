#Go application Makefile
APP_NAME ?= account
BINARY_NAME ?= bin/$(APP_NAME)
MAIN_PATH ?= ./cmd
VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
BUILD_TIME ?= $(shell date -u '+%Y-%m-%d_%H:%M:%S')
LDFLAGS ?= -ldflags "-X main.Version=$(VERSION) -X main.BuildTime=$(BUILD_TIME)"

# Go files to format and lint
GOFILES ?= $(shell find . -name "*.go" -type f -not -path "./vendor/*" -not -path "./.git/*")

# Default target
.DEFAULT_GOAL := build

# Build the application
build:
	@echo "Building $(APP_NAME)..."
	@mkdir -p bin
	go build $(LDFLAGS) -o $(BINARY_NAME) $(MAIN_PATH)
	
# Run the application
run: build
	@echo "Running $(APP_NAME)..."
	./$(BINARY_NAME)

# Run without building (if binary exists)
run-only:
	@echo "Running $(APP_NAME)..."
	./$(BINARY_NAME)

# Install dependencies
deps:
	@echo "Installing dependencies..."
	go mod download
	go mod tidy

# Update dependencies
deps-update:
	@echo "Updating dependencies..."
	go get -u ./...
	go mod tidy

# Format code
fmt:
	@echo "Formatting code..."
	gofmt -s -w $(GOFILES)
	gofumpt -l -w -s $(GOFILES)
	goimports -l -w $(GOFILES)

# Run linter
lint:
	@echo "Running linter..."
	golangci-lint run --timeout=10m -v ./...

# Run tests
test:
	@echo "Running tests..."
	go test -v ./...

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Run benchmarks
bench:
	@echo "Running benchmarks..."
	go test -bench=. ./...

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -rf bin/
	rm -f coverage.out coverage.html

# Install the application
install: build
	@echo "Installing $(APP_NAME)..."
	go install $(LDFLAGS) $(MAIN_PATH)

# Generate documentation
docs:
	@echo "Generating documentation..."
	godoc -http=:6060

# Show help
help:
	@echo "Available targets:"
	@printf "  %-16s - %s\n" "build" "Build the application"
	@printf "  %-16s - %s\n" "run" "Build and run the application"
	@printf "  %-16s - %s\n" "run-only" "Run the application (without building)"
	@printf "  %-16s - %s\n" "deps" "Install dependencies"
	@printf "  %-16s - %s\n" "deps-update" "Update dependencies"
	@printf "  %-16s - %s\n" "fmt" "Format code"
	@printf "  %-16s - %s\n" "lint" "Run linter"
	@printf "  %-16s - %s\n" "test" "Run tests"
	@printf "  %-16s - %s\n" "test-coverage" "Run tests with coverage report"
	@printf "  %-16s - %s\n" "bench" "Run benchmarks"
	@printf "  %-16s - %s\n" "clean" "Clean build artifacts"
	@printf "  %-16s - %s\n" "install" "Install the application"
	@printf "  %-16s - %s\n" "docs" "Generate documentation"
	@printf "  %-16s - %s\n" "help" "Show this help"

# Development workflow
dev: fmt lint test build

# CI/CD workflow
ci: deps fmt lint test build

.PHONY: build run run-only deps deps-update fmt lint test test-coverage bench clean install docs help dev ci