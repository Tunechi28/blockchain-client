package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
)

type mockBlockchainClient struct {
	BlockNumber string
	Block       json.RawMessage
	Err         error
}

func (f *mockBlockchainClient) GetBlockNumber(ctx context.Context) (string, error) {
	return f.BlockNumber, f.Err
}

func (f *mockBlockchainClient) GetBlockByNumber(ctx context.Context, blockNumber string) (json.RawMessage, error) {
	return f.Block, f.Err
}

func TestHandler_GetBlockNumberHandler_Success(t *testing.T) {
	mockClient := &mockBlockchainClient{BlockNumber: "0x10d4f"}
	h := &Handler{BC: mockClient}

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/block/number", h.GetBlockNumberHandler)

	req, _ := http.NewRequest("GET", "/block/number", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status 200, got %d", w.Code)
	}

	var response map[string]string
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if response["blockNumber"] != "0x10d4f" {
		t.Errorf("Expected block number %s, got %s", "0x10d4f", response["blockNumber"])
	}
}

func TestHandler_GetBlockNumberHandler_Error(t *testing.T) {
	mockClient := &mockBlockchainClient{Err: errors.New("client error")}
	h := &Handler{BC: mockClient}

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/block/number", h.GetBlockNumberHandler)

	req, _ := http.NewRequest("GET", "/block/number", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Fatalf("Expected status 500, got %d", w.Code)
	}
}

func TestHandler_GetBlockByNumberHandler_Success(t *testing.T) {
	blockData := map[string]interface{}{
		"number": "0x10d4f",
		"hash":   "0xabc123",
	}
	blockJSON, _ := json.Marshal(blockData)
	mockClient := &mockBlockchainClient{Block: blockJSON}
	h := &Handler{BC: mockClient}

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/block/:number", h.GetBlockByNumberHandler)

	req, _ := http.NewRequest("GET", "/block/0x10d4f", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status 200, got %d", w.Code)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	block, ok := response["block"].(map[string]interface{})
	if !ok {
		t.Fatalf("Expected block data to be a map, got %T", response["block"])
	}

	if block["number"] != "0x10d4f" {
		t.Errorf("Expected block number %s, got %v", "0x10d4f", block["number"])
	}
}

func TestHandler_GetBlockByNumberHandler_Error(t *testing.T) {
	mockClient := &mockBlockchainClient{Err: errors.New("client error")}
	h := &Handler{BC: mockClient}

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/block/:number", h.GetBlockByNumberHandler)

	req, _ := http.NewRequest("GET", "/block/0x10d4f", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Fatalf("Expected status 500, got %d", w.Code)
	}
}
