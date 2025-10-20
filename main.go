// Package main provides an MCP (Model Context Protocol) server with various utility tools.
package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/podanypepa/mcpserver/tools"
)

func main() {
	var (
		addr  = flag.String("addr", getEnv("MCP_ADDR", ":8080"), "HTTP listen address")
		token = flag.String("token", getEnv("MCP_TOKEN", ""), "Bearer token required for access (recommended)")
		path  = flag.String("path", getEnv("MCP_PATH", "/mcp"), "Base path for MCP endpoints")
	)
	flag.Parse()

	srv := server.NewMCPServer("mcp-go-utilities", "1.0.0")

	// Register Echo tool
	echoTool := mcp.NewTool(
		"echo",
		mcp.WithDescription("Echo the provided text, optionally in uppercase"),
		mcp.WithInputSchema[tools.EchoInput](),
		mcp.WithOutputSchema[tools.EchoOutput](),
		mcp.WithString("text", mcp.Required(), mcp.Description("Text to echo")),
		mcp.WithBoolean("uppercase", mcp.Description("Convert to uppercase?")),
	)
	srv.AddTool(echoTool, mcp.NewStructuredToolHandler(tools.Echo))

	// Register Reverse tool
	reverseTool := mcp.NewTool(
		"reverse",
		mcp.WithDescription("Reverse the provided text"),
		mcp.WithInputSchema[tools.ReverseInput](),
		mcp.WithOutputSchema[tools.ReverseOutput](),
		mcp.WithString("text", mcp.Required(), mcp.Description("Text to reverse")),
	)
	srv.AddTool(reverseTool, mcp.NewStructuredToolHandler(tools.Reverse))

	// Register Hash tool
	hashTool := mcp.NewTool(
		"hash",
		mcp.WithDescription("Generate hash of the provided text (MD5 or SHA256)"),
		mcp.WithInputSchema[tools.HashInput](),
		mcp.WithOutputSchema[tools.HashOutput](),
		mcp.WithString("text", mcp.Required(), mcp.Description("Text to hash")),
		mcp.WithString("algorithm", mcp.Description("Hash algorithm: md5 or sha256 (default)")),
	)
	srv.AddTool(hashTool, mcp.NewStructuredToolHandler(tools.Hash))

	// Register UUID tool
	uuidTool := mcp.NewTool(
		"uuid",
		mcp.WithDescription("Generate a new UUID v4"),
		mcp.WithOutputSchema[tools.UUIDOutput](),
	)
	srv.AddTool(uuidTool, mcp.NewStructuredToolHandler(tools.GenerateUUID))

	// Register Timestamp tool
	timestampTool := mcp.NewTool(
		"timestamp",
		mcp.WithDescription("Get current timestamp in various formats (RFC3339, Unix, UnixMilli)"),
		mcp.WithInputSchema[tools.TimestampInput](),
		mcp.WithOutputSchema[tools.TimestampOutput](),
		mcp.WithString("format", mcp.Description("Format: RFC3339 (default), Unix, or UnixMilli")),
	)
	srv.AddTool(timestampTool, mcp.NewStructuredToolHandler(tools.GetTimestamp))

	httpSrv := server.NewStreamableHTTPServer(
		srv,
		server.WithEndpointPath(*path), // provides /mcp/messages and /mcp/stream
		server.WithStateLess(true),     // ← stateless for simplicity
	)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if *token != "" {
			want := "Bearer " + *token
			if r.Header.Get("Authorization") != want {
				w.WriteHeader(http.StatusUnauthorized)
				_, _ = w.Write([]byte("unauthorized"))
				return
			}
		}
		httpSrv.ServeHTTP(w, r)
	})

	mux := http.NewServeMux()
	mux.Handle(*path+"/", http.StripPrefix(*path, handler))

	log.Printf("🚀 MCP Utilities Server starting...")
	log.Printf("📍 Address: %s", *addr)
	log.Printf("🔗 Base path: %s", *path)
	log.Printf("🔧 Tools: echo, reverse, hash, uuid, timestamp")
	if *token != "" {
		log.Printf("🔐 Auth: Bearer token required")
	} else {
		log.Printf("⚠️  Warning: No authentication token set (use -token or MCP_TOKEN)")
	}

	if err := http.ListenAndServe(*addr, mux); err != nil {
		log.Fatal(err)
	}
}

// getEnv returns the value of an environment variable or a default value if not set.
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
