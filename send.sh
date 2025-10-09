#!/bin/sh

curl -sS -X POST http://127.0.0.1:8080/mcp/messages \
  -H 'Authorization: Bearer secret123' \
  -H 'Content-Type: application/json' \
  -H 'Mcp-Protocol-Version: 2025-06-18' \
  --data '{"jsonrpc":"2.0","id":1,"method":"initialize","params":{"protocolVersion":"2025-06-18","capabilities":{},"clientInfo":{"name":"curl","version":"0.1"}}}'

curl -sS -X POST http://127.0.0.1:8080/mcp/messages \
  -H 'Authorization: Bearer secret123' \
  -H 'Content-Type: application/json' \
  -H 'Mcp-Protocol-Version: 2025-06-18' \
  --data '{"jsonrpc":"2.0","method":"initialized"}'

curl -sS -X POST http://127.0.0.1:8080/mcp/messages \
  -H 'Authorization: Bearer secret123' \
  -H 'Content-Type: application/json' \
  --data '{"jsonrpc":"2.0","id":2,"method":"tools/list"}'

curl -sS -X POST http://127.0.0.1:8080/mcp/messages \
  -H "Authorization: Bearer secret123" \
  -H "Content-Type: application/json" \
  -H "Mcp-Protocol-Version: 2025-06-18" \
  --data '{"jsonrpc":"2.0","id":3,"method":"tools/call","params":{"name":"echo","arguments":{"text":"Ahoj, pepp","uppercase":true}}}'
