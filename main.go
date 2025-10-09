// Package main
package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// EchoInput is the input structure for the Echo tool.
type EchoInput struct {
	Text      string `json:"text" jsonschema:"Text to echo back"`
	Uppercase bool   `json:"uppercase,omitempty" jsonschema:"Uppercase output?"`
}

// EchoOutput is the output structure for the Echo tool.
type EchoOutput struct {
	Text string `json:"text"`
}

// Echo is a simple tool that echoes back the provided text, optionally uppercased.
func Echo(_ context.Context, _ mcp.CallToolRequest, in EchoInput) (EchoOutput, error) {
	out := in.Text
	if in.Uppercase {
		out = strings.ToUpper(out)
	}
	return EchoOutput{Text: out}, nil
}

func main() {
	var (
		addr  = flag.String("addr", ":8080", "HTTP listen address")
		token = flag.String("token", "", "Bearer token required for access (recommended)")
		path  = flag.String("path", "/mcp", "Base path for MCP endpoints (/mcp)")
	)
	flag.Parse()

	srv := server.NewMCPServer("nomodo-mcp-go-http", "0.2.0")

	echoTool := mcp.NewTool(
		"echo",
		mcp.WithDescription("Echo the provided text"),
		mcp.WithInputSchema[EchoInput](),
		mcp.WithOutputSchema[EchoOutput](),
		mcp.WithString("text", mcp.Required(), mcp.Description("Text to echo")),
		mcp.WithBoolean("uppercase", mcp.Description("Uppercase output?")),
	)
	srv.AddTool(echoTool, mcp.NewStructuredToolHandler(Echo))

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

	log.Printf("[mcp] streamable HTTP on %s (base %s)", *addr, *path)
	if *token != "" {
		log.Printf("[mcp] auth: Bearer token required")
	}

	if err := http.ListenAndServe(*addr, mux); err != nil {
		log.Fatal(err)
	}
}
