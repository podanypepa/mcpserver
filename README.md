# 🚀 MCP Go Utilities Server

[![Go CI](https://github.com/podanypepa/mcpserver/actions/workflows/ci.yml/badge.svg)](https://github.com/podanypepa/mcpserver/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/podanypepa/mcpserver)](https://goreportcard.com/report/github.com/podanypepa/mcpserver)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A production-ready implementation of a [Model Context Protocol (MCP)](https://modelcontextprotocol.io/) server written in Go. This server provides various utility tools accessible through the MCP protocol.

## ✨ Features

- 🔧 **5 Useful Tools**:
  - `echo` - Echo text with optional uppercase conversion
  - `reverse` - Reverse any text string (supports Unicode)
  - `hash` - Generate MD5 or SHA256 hashes
  - `uuid` - Generate UUID v4 identifiers
  - `timestamp` - Get current timestamps in various formats

- 🔐 **Bearer Token Authentication** - Secure your server with API tokens
- 🐳 **Docker Support** - Easy containerization
- 🧪 **Full Test Coverage** - Comprehensive unit tests
- 🏗️ **Clean Architecture** - Modular and maintainable code structure
- ⚙️ **Configurable** - Via command-line flags or environment variables

## 📋 Prerequisites

- Go 1.22+ or Docker

## 🚀 Quick Start

### Using Go

```bash
# Clone the repository
git clone https://github.com/podanypepa/mcpserver.git
cd mcpserver

# Install dependencies
go mod download

# Run the server
go run main.go
```

### Using Make

```bash
make help           # Show all available commands
make run            # Run the server
make test           # Run tests
make build          # Build binary
```

### Using Docker

```bash
# Build the image
docker build -t mcpserver .

# Run the container
docker run -p 8080:8080 -e MCP_TOKEN=secret123 mcpserver
```

## ⚙️ Configuration

Configure the server using command-line flags or environment variables:

| Flag | Environment Variable | Default | Description |
|------|---------------------|---------|-------------|
| `-addr` | `MCP_ADDR` | `:8080` | HTTP listen address |
| `-token` | `MCP_TOKEN` | `""` | Bearer token for authentication |
| `-path` | `MCP_PATH` | `/mcp` | Base path for MCP endpoints |

### Examples

```bash
# With custom port and token
go run main.go -addr=":3000" -token="my-secret-token"

# Using environment variables
export MCP_TOKEN="my-secret-token"
export MCP_ADDR=":3000"
go run main.go
```

## 🔧 Available Tools

### 1. Echo
Echoes text back, optionally in uppercase.

```json
{
  "method": "tools/call",
  "params": {
    "name": "echo",
    "arguments": {
      "text": "Hello, World!",
      "uppercase": true
    }
  }
}
```

### 2. Reverse
Reverses any text string (Unicode-safe).

```json
{
  "method": "tools/call",
  "params": {
    "name": "reverse",
    "arguments": {
      "text": "Hello, World!"
    }
  }
}
```

### 3. Hash
Generates MD5 or SHA256 hash.

```json
{
  "method": "tools/call",
  "params": {
    "name": "hash",
    "arguments": {
      "text": "Hello, World!",
      "algorithm": "sha256"
    }
  }
}
```

### 4. UUID
Generates a new UUID v4.

```json
{
  "method": "tools/call",
  "params": {
    "name": "uuid",
    "arguments": {}
  }
}
```

### 5. Timestamp
Returns current timestamp in various formats.

```json
{
  "method": "tools/call",
  "params": {
    "name": "timestamp",
    "arguments": {
      "format": "RFC3339"
    }
  }
}
```

Supported formats: `RFC3339`, `Unix`, `UnixMilli`

## 🧪 Testing

```bash
# Run tests
make test

# Run tests with coverage
make test-coverage

# View coverage in browser
open coverage.html
```

## 🛠️ Development

```bash
# Format code
make fmt

# Run linter
make lint

# Install golangci-lint (if needed)
brew install golangci-lint

# Build binary
make build
```

## 📝 Example Client Usage

The `send.sh` script demonstrates how to interact with the server:

```bash
# Make sure the server is running
go run main.go -token="secret123"

# In another terminal, run the example script
./send.sh
```

You can also use `curl` directly:

```bash
curl -X POST http://127.0.0.1:8080/mcp/messages \
  -H 'Authorization: Bearer secret123' \
  -H 'Content-Type: application/json' \
  -H 'Mcp-Protocol-Version: 2025-06-18' \
  --data '{
    "jsonrpc":"2.0",
    "id":1,
    "method":"tools/call",
    "params":{
      "name":"echo",
      "arguments":{"text":"Hello","uppercase":true}
    }
  }'
```

## 📦 Project Structure

```
.
├── main.go              # Application entry point
├── tools/
│   ├── tools.go         # Tool implementations
│   └── tools_test.go    # Tool tests
├── Makefile             # Build automation
├── Dockerfile           # Container definition
├── .github/
│   └── workflows/
│       └── ci.yml       # CI/CD pipeline
└── README.md            # This file
```

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 🙏 Acknowledgments

- Built with [mcp-go](https://github.com/mark3labs/mcp-go) by Mark3 Labs
- Inspired by the [Model Context Protocol](https://modelcontextprotocol.io/)

## 📧 Contact

Your Name - [@podanypepa](https://github.com/podanypepa)

Project Link: [https://github.com/podanypepa/mcpserver](https://github.com/podanypepa/mcpserver)
