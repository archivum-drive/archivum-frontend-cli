# Archivum Frontend CLI Makefile

.PHONY: build run clean test deps help

# Default target
all: build

# Build the application
build:
	@echo "Building Archivum CLI..."
	@mkdir -p bin
	@go build -o bin/archivum cmd/archivum/main.go
	@echo "Build complete: bin/archivum"

# Run the application directly
run:
	@echo "Running Archivum CLI..."
	@go run cmd/archivum/main.go

# Run the compiled binary
run-bin: build
	@echo "Running compiled binary..."
	@./bin/archivum

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf bin/*
	@echo "Clean complete"

# Install/update dependencies
deps:
	@echo "Installing dependencies..."
	@go mod tidy
	@go mod download
	@echo "Dependencies updated"

# Run tests (when we have them)
test:
	@echo "Running tests..."
	@go test ./...

# Format code
fmt:
	@echo "Formatting code..."
	@go fmt ./...

# Lint code (requires golangci-lint)
lint:
	@echo "Linting code..."
	@golangci-lint run

# Show help
help:
	@echo "Available commands:"
	@echo "  build      - Build the application"
	@echo "  run        - Run the application with go run"
	@echo "  run-bin    - Build and run the compiled binary"
	@echo "  clean      - Remove build artifacts"
	@echo "  deps       - Install/update dependencies"
	@echo "  test       - Run tests"
	@echo "  fmt        - Format code"
	@echo "  lint       - Lint code (requires golangci-lint)"
	@echo "  help       - Show this help message"
