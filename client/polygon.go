package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

var rpcURL = "https://polygon-rpc.com/"

var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

type RPCRequest struct {
	JSONRPC string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params,omitempty"`
	ID      int           `json:"id"`
}

type RPCResponse struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      int             `json:"id"`
	Result  json.RawMessage `json:"result"`
	Error   *RPCError       `json:"error,omitempty"`
}

type RPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func sendRPCRequest(ctx context.Context, method string, params []interface{}) (*RPCResponse, error) {
	reqBody, err := json.Marshal(RPCRequest{
		JSONRPC: "2.0",
		Method:  method,
		Params:  params,
		ID:      2,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, rpcURL, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("RPC request failed with status %d: %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	var rpcResp RPCResponse
	if err := json.Unmarshal(body, &rpcResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if rpcResp.Error != nil {
		return nil, fmt.Errorf("RPC error %d: %s", rpcResp.Error.Code, rpcResp.Error.Message)
	}

	return &rpcResp, nil
}

func GetBlockNumber(ctx context.Context) (string, error) {
	rpcResp, err := sendRPCRequest(ctx, "eth_blockNumber", nil)
	if err != nil {
		return "", err
	}

	var blockNumber string
	if err := json.Unmarshal(rpcResp.Result, &blockNumber); err != nil {
		return "", fmt.Errorf("failed to unmarshal block number: %w", err)
	}
	return blockNumber, nil
}

func GetBlockByNumber(ctx context.Context, blockNumber string) (json.RawMessage, error) {
	rpcResp, err := sendRPCRequest(ctx, "eth_getBlockByNumber", []interface{}{blockNumber, true})
	if err != nil {
		return nil, err
	}
	return rpcResp.Result, nil
}
