package client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

var testTimeout = 5 * time.Second

func TestGetBlockNumber_Success(t *testing.T) {
	mockResponse := `{"jsonrpc":"2.0","id":1,"result":"0x10d4f"}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(mockResponse))
	}))
	defer server.Close()
	rpcURL = server.URL

	ctx, cancel := context.WithTimeout(context.Background(), testTimeout)
	defer cancel()

	blockNumber, err := GetBlockNumber(ctx)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := "0x10d4f"
	if blockNumber != expected {
		t.Errorf("Expected block number %s, got %s", expected, blockNumber)
	}
}

func TestGetBlockNumber_HTTPError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}))
	defer server.Close()
	rpcURL = server.URL

	ctx, cancel := context.WithTimeout(context.Background(), testTimeout)
	defer cancel()

	_, err := GetBlockNumber(ctx)
	if err == nil || !strings.Contains(err.Error(), "RPC request failed with status") {
		t.Errorf("Expected HTTP error, got %v", err)
	}
}

func TestGetBlockNumber_MalformedResponse(t *testing.T) {
	mockResponse := `invalid json`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(mockResponse))
	}))
	defer server.Close()
	rpcURL = server.URL

	ctx, cancel := context.WithTimeout(context.Background(), testTimeout)
	defer cancel()

	_, err := GetBlockNumber(ctx)
	if err == nil || !strings.Contains(err.Error(), "failed to unmarshal response") {
		t.Errorf("Expected unmarshalling error, got %v", err)
	}
}

func TestGetBlockByNumber_Success(t *testing.T) {
	block := map[string]interface{}{
		"number": "0x10d4f",
		"hash":   "0xabc123",
	}
	blockJSON, _ := json.Marshal(block)
	mockResponse := `{"jsonrpc":"2.0","id":1,"result":` + string(blockJSON) + `}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(mockResponse))
	}))
	defer server.Close()
	rpcURL = server.URL

	ctx, cancel := context.WithTimeout(context.Background(), testTimeout)
	defer cancel()

	result, err := GetBlockByNumber(ctx, "0x10d4f")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	var blockData map[string]interface{}
	if err := json.Unmarshal(result, &blockData); err != nil {
		t.Fatalf("Failed to unmarshal block data: %v", err)
	}
	if blockData["number"] != "0x10d4f" {
		t.Errorf("Expected block number %s, got %v", "0x10d4f", blockData["number"])
	}
}

func TestGetBlockByNumber_RPCError(t *testing.T) {
	mockResponse := `{"jsonrpc":"2.0","id":1,"error":{"code":-32000,"message":"Test error"}}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(mockResponse))
	}))
	defer server.Close()
	rpcURL = server.URL

	ctx, cancel := context.WithTimeout(context.Background(), testTimeout)
	defer cancel()

	_, err := GetBlockByNumber(ctx, "0x10d4f")
	if err == nil || !strings.Contains(err.Error(), "RPC error") {
		t.Errorf("Expected RPC error, got %v", err)
	}
}
