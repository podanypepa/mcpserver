// Package tools provides MCP tool implementations
package tools

import (
	"context"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/mark3labs/mcp-go/mcp"
)

// EchoInput is the input structure for the Echo tool.
type EchoInput struct {
	Text      string `json:"text" jsonschema:"required,description=Text to echo back"`
	Uppercase bool   `json:"uppercase,omitempty" jsonschema:"description=Convert to uppercase?"`
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

// ReverseInput is the input structure for the Reverse tool.
type ReverseInput struct {
	Text string `json:"text" jsonschema:"required,description=Text to reverse"`
}

// ReverseOutput is the output structure for the Reverse tool.
type ReverseOutput struct {
	Text string `json:"text"`
}

// Reverse reverses the provided text.
func Reverse(_ context.Context, _ mcp.CallToolRequest, in ReverseInput) (ReverseOutput, error) {
	runes := []rune(in.Text)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return ReverseOutput{Text: string(runes)}, nil
}

// HashInput is the input structure for the Hash tool.
type HashInput struct {
	Text      string `json:"text" jsonschema:"required,description=Text to hash"`
	Algorithm string `json:"algorithm,omitempty" jsonschema:"description=Hash algorithm (md5 or sha256),enum=md5,enum=sha256"`
}

// HashOutput is the output structure for the Hash tool.
type HashOutput struct {
	Hash      string `json:"hash"`
	Algorithm string `json:"algorithm"`
}

// Hash generates a hash of the provided text.
func Hash(_ context.Context, _ mcp.CallToolRequest, in HashInput) (HashOutput, error) {
	algo := strings.ToLower(in.Algorithm)
	if algo == "" {
		algo = "sha256"
	}

	var hash string
	switch algo {
	case "md5":
		h := md5.Sum([]byte(in.Text))
		hash = hex.EncodeToString(h[:])
	case "sha256":
		h := sha256.Sum256([]byte(in.Text))
		hash = hex.EncodeToString(h[:])
	default:
		return HashOutput{}, fmt.Errorf("unsupported algorithm: %s (use md5 or sha256)", algo)
	}

	return HashOutput{Hash: hash, Algorithm: algo}, nil
}

// UUIDOutput is the output structure for the UUID tool.
type UUIDOutput struct {
	UUID string `json:"uuid"`
}

// GenerateUUID generates a new UUID v4.
func GenerateUUID(_ context.Context, _ mcp.CallToolRequest, _ struct{}) (UUIDOutput, error) {
	return UUIDOutput{UUID: uuid.New().String()}, nil
}

// TimestampInput is the input structure for the Timestamp tool.
type TimestampInput struct {
	Format string `json:"format,omitempty" jsonschema:"description=Time format (RFC3339, Unix, or UnixMilli)"`
}

// TimestampOutput is the output structure for the Timestamp tool.
type TimestampOutput struct {
	Timestamp string `json:"timestamp"`
	Format    string `json:"format"`
}

// GetTimestamp returns the current timestamp in various formats.
func GetTimestamp(_ context.Context, _ mcp.CallToolRequest, in TimestampInput) (TimestampOutput, error) {
	now := time.Now()
	format := strings.ToLower(in.Format)
	if format == "" {
		format = "rfc3339"
	}

	var ts string
	switch format {
	case "rfc3339":
		ts = now.Format(time.RFC3339)
	case "unix":
		ts = fmt.Sprintf("%d", now.Unix())
	case "unixmilli":
		ts = fmt.Sprintf("%d", now.UnixMilli())
	default:
		return TimestampOutput{}, fmt.Errorf("unsupported format: %s (use RFC3339, Unix, or UnixMilli)", in.Format)
	}

	return TimestampOutput{Timestamp: ts, Format: format}, nil
}
