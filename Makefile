.PHONY: help build run test clean docker-build docker-run

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: ## Build the application
	@echo "🔨 Building..."
	@go build -o bin/mcpserver .

run: ## Run the application
	@echo "🚀 Starting server..."
	@go run main.go

test: ## Run tests
	@echo "🧪 Running tests..."
	@go test -v ./...

test-coverage: ## Run tests with coverage
	@echo "🧪 Running tests with coverage..."
	@go test -v -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "📊 Coverage report: coverage.html"

lint: ## Run linter
	@echo "🔍 Running linter..."
	@if command -v golangci-lint > /dev/null; then \
		golangci-lint run; \
	else \
		echo "⚠️  golangci-lint not installed. Run: brew install golangci-lint"; \
	fi

fmt: ## Format code
	@echo "✨ Formatting code..."
	@go fmt ./...
	@go vet ./...

clean: ## Clean build artifacts
	@echo "🧹 Cleaning..."
	@rm -rf bin/
	@rm -f coverage.out coverage.html

deps: ## Download dependencies
	@echo "📦 Downloading dependencies..."
	@go mod download
	@go mod tidy

docker-build: ## Build Docker image
	@echo "🐳 Building Docker image..."
	@docker build -t mcpserver:latest .

docker-run: ## Run Docker container
	@echo "🐳 Running Docker container..."
	@docker run -p 8080:8080 -e MCP_TOKEN=secret123 mcpserver:latest

install: ## Install the binary
	@echo "📦 Installing..."
	@go install .

all: clean deps fmt test build ## Clean, download deps, format, test, and build
