package tools

import (
	"context"
	"testing"

	"github.com/mark3labs/mcp-go/mcp"
)

func TestEcho(t *testing.T) {
	tests := []struct {
		name     string
		input    EchoInput
		expected string
	}{
		{
			name:     "simple echo",
			input:    EchoInput{Text: "hello"},
			expected: "hello",
		},
		{
			name:     "uppercase echo",
			input:    EchoInput{Text: "hello", Uppercase: true},
			expected: "HELLO",
		},
		{
			name:     "empty string",
			input:    EchoInput{Text: ""},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Echo(context.Background(), mcp.CallToolRequest{}, tt.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if result.Text != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, result.Text)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "simple word",
			input:    "hello",
			expected: "olleh",
		},
		{
			name:     "with spaces",
			input:    "hello world",
			expected: "dlrow olleh",
		},
		{
			name:     "unicode characters",
			input:    "Ahoj, světe!",
			expected: "!etěvs ,johA",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Reverse(context.Background(), mcp.CallToolRequest{}, ReverseInput{Text: tt.input})
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if result.Text != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, result.Text)
			}
		})
	}
}

func TestHash(t *testing.T) {
	tests := []struct {
		name      string
		input     HashInput
		wantHash  string
		wantAlgo  string
		wantError bool
	}{
		{
			name:     "md5 hash",
			input:    HashInput{Text: "hello", Algorithm: "md5"},
			wantHash: "5d41402abc4b2a76b9719d911017c592",
			wantAlgo: "md5",
		},
		{
			name:     "sha256 hash",
			input:    HashInput{Text: "hello", Algorithm: "sha256"},
			wantHash: "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824",
			wantAlgo: "sha256",
		},
		{
			name:     "default to sha256",
			input:    HashInput{Text: "hello"},
			wantHash: "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824",
			wantAlgo: "sha256",
		},
		{
			name:      "unsupported algorithm",
			input:     HashInput{Text: "hello", Algorithm: "sha512"},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Hash(context.Background(), mcp.CallToolRequest{}, tt.input)
			if tt.wantError {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if result.Hash != tt.wantHash {
				t.Errorf("expected hash %q, got %q", tt.wantHash, result.Hash)
			}
			if result.Algorithm != tt.wantAlgo {
				t.Errorf("expected algorithm %q, got %q", tt.wantAlgo, result.Algorithm)
			}
		})
	}
}

func TestGenerateUUID(t *testing.T) {
	result, err := GenerateUUID(context.Background(), mcp.CallToolRequest{}, struct{}{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.UUID) != 36 {
		t.Errorf("expected UUID length 36, got %d", len(result.UUID))
	}

	// Generate another UUID to ensure uniqueness
	result2, err := GenerateUUID(context.Background(), mcp.CallToolRequest{}, struct{}{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.UUID == result2.UUID {
		t.Error("expected different UUIDs, got same")
	}
}

func TestGetTimestamp(t *testing.T) {
	tests := []struct {
		name      string
		format    string
		wantError bool
	}{
		{
			name:   "RFC3339 format",
			format: "RFC3339",
		},
		{
			name:   "Unix format",
			format: "Unix",
		},
		{
			name:   "UnixMilli format",
			format: "UnixMilli",
		},
		{
			name:   "default format",
			format: "",
		},
		{
			name:      "unsupported format",
			format:    "invalid",
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := GetTimestamp(context.Background(), mcp.CallToolRequest{}, TimestampInput{Format: tt.format})
			if tt.wantError {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if result.Timestamp == "" {
				t.Error("expected non-empty timestamp")
			}
		})
	}
}
