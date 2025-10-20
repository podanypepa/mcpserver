# API Documentation

## MCP Protocol Endpoints

The server implements the Model Context Protocol (MCP) and exposes endpoints at the configured base path (default: `/mcp`).

### Base URL

```
http://localhost:8080/mcp
```

### Endpoints

- `POST /mcp/messages` - JSON-RPC 2.0 messages endpoint
- `POST /mcp/stream` - Streaming endpoint (SSE)

## Authentication

All requests require a Bearer token in the Authorization header:

```
Authorization: Bearer YOUR_TOKEN_HERE
```

Set the token using the `-token` flag or `MCP_TOKEN` environment variable.

## Protocol Version

Include the protocol version header:

```
Mcp-Protocol-Version: 2025-06-18
```

## JSON-RPC Methods

### initialize

Initialize the MCP connection.

**Request:**
```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "method": "initialize",
  "params": {
    "protocolVersion": "2025-06-18",
    "capabilities": {},
    "clientInfo": {
      "name": "your-client",
      "version": "1.0.0"
    }
  }
}
```

**Response:**
```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "result": {
    "protocolVersion": "2025-06-18",
    "capabilities": {...},
    "serverInfo": {
      "name": "mcp-go-utilities",
      "version": "1.0.0"
    }
  }
}
```

### tools/list

List all available tools.

**Request:**
```json
{
  "jsonrpc": "2.0",
  "id": 2,
  "method": "tools/list"
}
```

**Response:**
```json
{
  "jsonrpc": "2.0",
  "id": 2,
  "result": {
    "tools": [
      {
        "name": "echo",
        "description": "Echo the provided text, optionally in uppercase",
        "inputSchema": {...}
      },
      {
        "name": "reverse",
        "description": "Reverse the provided text",
        "inputSchema": {...}
      },
      ...
    ]
  }
}
```

### tools/call

Call a specific tool.

**Request:**
```json
{
  "jsonrpc": "2.0",
  "id": 3,
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

**Response:**
```json
{
  "jsonrpc": "2.0",
  "id": 3,
  "result": {
    "content": [
      {
        "type": "text",
        "text": "{\"text\":\"HELLO, WORLD!\"}"
      }
    ]
  }
}
```

## Available Tools

### 1. echo

Echo text back, optionally in uppercase.

**Arguments:**
- `text` (string, required): Text to echo
- `uppercase` (boolean, optional): Convert to uppercase

**Example:**
```json
{
  "name": "echo",
  "arguments": {
    "text": "Hello",
    "uppercase": true
  }
}
```

**Output:**
```json
{
  "text": "HELLO"
}
```

### 2. reverse

Reverse any text string (Unicode-safe).

**Arguments:**
- `text` (string, required): Text to reverse

**Example:**
```json
{
  "name": "reverse",
  "arguments": {
    "text": "Hello, World!"
  }
}
```

**Output:**
```json
{
  "text": "!dlroW ,olleH"
}
```

### 3. hash

Generate MD5 or SHA256 hash.

**Arguments:**
- `text` (string, required): Text to hash
- `algorithm` (string, optional): Hash algorithm - "md5" or "sha256" (default: "sha256")

**Example:**
```json
{
  "name": "hash",
  "arguments": {
    "text": "Hello",
    "algorithm": "sha256"
  }
}
```

**Output:**
```json
{
  "hash": "185f8db32271fe25f561a6fc938b2e264306ec304eda518007d1764826381969",
  "algorithm": "sha256"
}
```

### 4. uuid

Generate a new UUID v4.

**Arguments:** None

**Example:**
```json
{
  "name": "uuid",
  "arguments": {}
}
```

**Output:**
```json
{
  "uuid": "f47ac10b-58cc-4372-a567-0e02b2c3d479"
}
```

### 5. timestamp

Get current timestamp in various formats.

**Arguments:**
- `format` (string, optional): Format - "RFC3339", "Unix", or "UnixMilli" (default: "RFC3339")

**Example:**
```json
{
  "name": "timestamp",
  "arguments": {
    "format": "RFC3339"
  }
}
```

**Output:**
```json
{
  "timestamp": "2025-10-20T16:00:00Z",
  "format": "rfc3339"
}
```

## Error Responses

Errors follow the JSON-RPC 2.0 specification:

```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "error": {
    "code": -32600,
    "message": "Invalid Request",
    "data": "additional error information"
  }
}
```

## HTTP Status Codes

- `200 OK` - Successful request
- `401 Unauthorized` - Missing or invalid bearer token
- `400 Bad Request` - Invalid JSON or malformed request
- `500 Internal Server Error` - Server error

## Rate Limiting

Currently, no rate limiting is implemented. Consider adding rate limiting in production environments.

## CORS

CORS is not enabled by default. Configure your reverse proxy (nginx, Apache, etc.) to handle CORS if needed.
