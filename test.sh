#!/bin/bash
# Integration test script - tests all available tools

set -e

BASE_URL="${BASE_URL:-http://127.0.0.1:8080}"
TOKEN="${MCP_TOKEN:-secret123}"

echo "🧪 Running integration tests..."
echo "📍 Server: $BASE_URL"
echo ""

# Helper function to make MCP requests
mcp_call() {
    local method=$1
    local params=$2
    local id=${3:-1}
    
    curl -sS -X POST "$BASE_URL/mcp/messages" \
        -H "Authorization: Bearer $TOKEN" \
        -H "Content-Type: application/json" \
        -H "Mcp-Protocol-Version: 2025-06-18" \
        --data "{\"jsonrpc\":\"2.0\",\"id\":$id,\"method\":\"$method\",\"params\":$params}"
}

# Test 1: Initialize
echo "1️⃣  Testing initialize..."
response=$(mcp_call "initialize" '{"protocolVersion":"2025-06-18","capabilities":{},"clientInfo":{"name":"test","version":"1.0"}}')
if echo "$response" | grep -q "result"; then
    echo "✅ Initialize successful"
else
    echo "❌ Initialize failed: $response"
    exit 1
fi

# Test 2: List tools
echo ""
echo "2️⃣  Testing tools/list..."
response=$(mcp_call "tools/list" '{}' 2)
if echo "$response" | grep -q "echo" && echo "$response" | grep -q "reverse" && echo "$response" | grep -q "hash"; then
    echo "✅ Tools list successful (found echo, reverse, hash)"
else
    echo "❌ Tools list failed: $response"
    exit 1
fi

# Test 3: Echo tool
echo ""
echo "3️⃣  Testing echo tool..."
response=$(mcp_call "tools/call" '{"name":"echo","arguments":{"text":"Hello","uppercase":true}}' 3)
if echo "$response" | grep -q "HELLO"; then
    echo "✅ Echo tool successful"
else
    echo "❌ Echo tool failed: $response"
    exit 1
fi

# Test 4: Reverse tool
echo ""
echo "4️⃣  Testing reverse tool..."
response=$(mcp_call "tools/call" '{"name":"reverse","arguments":{"text":"hello"}}' 4)
if echo "$response" | grep -q "olleh"; then
    echo "✅ Reverse tool successful"
else
    echo "❌ Reverse tool failed: $response"
    exit 1
fi

# Test 5: Hash tool (SHA256)
echo ""
echo "5️⃣  Testing hash tool (SHA256)..."
response=$(mcp_call "tools/call" '{"name":"hash","arguments":{"text":"hello","algorithm":"sha256"}}' 5)
if echo "$response" | grep -q "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824"; then
    echo "✅ Hash tool (SHA256) successful"
else
    echo "❌ Hash tool failed: $response"
    exit 1
fi

# Test 6: Hash tool (MD5)
echo ""
echo "6️⃣  Testing hash tool (MD5)..."
response=$(mcp_call "tools/call" '{"name":"hash","arguments":{"text":"hello","algorithm":"md5"}}' 6)
if echo "$response" | grep -q "5d41402abc4b2a76b9719d911017c592"; then
    echo "✅ Hash tool (MD5) successful"
else
    echo "❌ Hash tool failed: $response"
    exit 1
fi

# Test 7: UUID tool
echo ""
echo "7️⃣  Testing uuid tool..."
response=$(mcp_call "tools/call" '{"name":"uuid","arguments":{}}' 7)
if echo "$response" | grep -qE "[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"; then
    echo "✅ UUID tool successful"
else
    echo "❌ UUID tool failed: $response"
    exit 1
fi

# Test 8: Timestamp tool
echo ""
echo "8️⃣  Testing timestamp tool..."
response=$(mcp_call "tools/call" '{"name":"timestamp","arguments":{"format":"RFC3339"}}' 8)
if echo "$response" | grep -qE "[0-9]{4}-[0-9]{2}-[0-9]{2}T"; then
    echo "✅ Timestamp tool successful"
else
    echo "❌ Timestamp tool failed: $response"
    exit 1
fi

echo ""
echo "🎉 All integration tests passed!"
