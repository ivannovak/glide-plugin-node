.PHONY: build install test clean build-all help

# Default target
all: build

# Build the plugin binary for the current platform
build:
	@echo "Building glide-plugin-node..."
	@go build -o glide-plugin-node cmd/glide-plugin-node/main.go

# Install the plugin to GOPATH/bin
install:
	@echo "Installing glide-plugin-node..."
	@go install ./cmd/glide-plugin-node

# Run tests
test:
	@echo "Running tests..."
	@go test ./...

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	@go test -cover ./...

# Clean build artifacts
clean:
	@echo "Cleaning..."
	@rm -f glide-plugin-node
	@rm -rf dist/

# Build for all platforms
build-all:
	@echo "Building for all platforms..."
	@mkdir -p dist
	@echo "Building for macOS (amd64)..."
	@GOOS=darwin GOARCH=amd64 go build -o dist/glide-plugin-node-darwin-amd64 cmd/glide-plugin-node/main.go
	@echo "Building for macOS (arm64)..."
	@GOOS=darwin GOARCH=arm64 go build -o dist/glide-plugin-node-darwin-arm64 cmd/glide-plugin-node/main.go
	@echo "Building for Linux (amd64)..."
	@GOOS=linux GOARCH=amd64 go build -o dist/glide-plugin-node-linux-amd64 cmd/glide-plugin-node/main.go
	@echo "Building for Windows (amd64)..."
	@GOOS=windows GOARCH=amd64 go build -o dist/glide-plugin-node-windows-amd64.exe cmd/glide-plugin-node/main.go
	@echo "Build complete! Binaries available in dist/"

# Tidy dependencies
tidy:
	@echo "Tidying dependencies..."
	@go mod tidy

# Show help
help:
	@echo "Available targets:"
	@echo "  build         - Build plugin binary for current platform"
	@echo "  install       - Install plugin to GOPATH/bin"
	@echo "  test          - Run tests"
	@echo "  test-coverage - Run tests with coverage"
	@echo "  clean         - Remove build artifacts"
	@echo "  build-all     - Build for all platforms (macOS, Linux, Windows)"
	@echo "  tidy          - Tidy Go dependencies"
	@echo "  help          - Show this help message"
