package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tunechi28/blockchain-client/client"
)

// BlockNumberResponse defines the response for GET /block/number.
// swagger:model BlockNumberResponse
type BlockNumberResponse struct {
	// The latest block number in hexadecimal format.
	// example: 0x41ce50d
	BlockNumber string `json:"blockNumber"`
}

// ErrorResponse defines the error response model.
// swagger:model ErrorResponse
type ErrorResponse struct {
	// Error message describing the issue.
	// example: HTTP request failed: Post "https://polygon-rpc.com/": dial tcp: lookup polygon-rpc.com: no such host
	Error string `json:"error"`
}

// BlockDetails represents detailed block information returned by GET /block/{number}.
// swagger:model BlockDetails
type BlockDetails struct {
	// The block difficulty.
	// example: 0x7
	Difficulty string `json:"difficulty" example:"0x7"`

	// Extra data associated with the block.
	// example: 0xd58301090083626f7286676f312e3133856c696e75780000000000000000000020e630202e8ff344c8b9504088f0d59c81fad4fecf54e352319be3085fe5bf6d12e11a34625810f586e80b8388e131b04f4fb964b777d757cba48851880bc1bf00
	ExtraData string `json:"extraData" example:"0xd58301090083626f7286676f312e3133856c696e75780000000000000000000020e630202e8ff344c8b9504088f0d59c81fad4fecf54e352319be3085fe5bf6d12e11a34625810f586e80b8388e131b04f4fb964b777d757cba48851880bc1bf00"`

	// The gas limit for the block.
	// example: 0x992f4a
	GasLimit string `json:"gasLimit" example:"0x992f4a"`

	// The gas used in the block.
	// example: 0x0
	GasUsed string `json:"gasUsed" example:"0x0"`

	// The hash of the block.
	// example: 0x30b128813e745454f46978d5073cd9e9474ae78df969ae4928992470fa9379dc
	Hash string `json:"hash" example:"0x30b128813e745454f46978d5073cd9e9474ae78df969ae4928992470fa9379dc"`

	// The bloom filter for the block logs.
	// example: 0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000
	LogsBloom string `json:"logsBloom" example:"0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"`

	// The address of the miner who mined the block.
	// example: 0x0000000000000000000000000000000000000000
	Miner string `json:"miner" example:"0x0000000000000000000000000000000000000000"`

	// The mix hash used in the proof-of-work algorithm.
	// example: 0x0000000000000000000000000000000000000000000000000000000000000000
	MixHash string `json:"mixHash" example:"0x0000000000000000000000000000000000000000000000000000000000000000"`

	// The nonce used in the mining process.
	// example: 0x0000000000000000
	Nonce string `json:"nonce" example:"0x0000000000000000"`

	// The block number.
	// example: 0x4
	Number string `json:"number" example:"0x4"`

	// The hash of the parent block.
	// example: 0x449afa8d9ea4c39c65a508414aa1b1110a225e2d8d4a86d40e5f8ae23103186a
	ParentHash string `json:"parentHash" example:"0x449afa8d9ea4c39c65a508414aa1b1110a225e2d8d4a86d40e5f8ae23103186a"`

	// The root hash of the receipts trie.
	// example: 0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421
	ReceiptsRoot string `json:"receiptsRoot" example:"0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421"`

	// The SHA3 hash of the uncles data.
	// example: 0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347
	Sha3Uncles string `json:"sha3Uncles" example:"0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"`

	// The size of the block in bytes.
	// example: 0x25e
	Size string `json:"size" example:"0x25e"`

	// The root hash of the state trie.
	// example: 0x654f28d19b44239d1012f27038f1f71b3d4465dc415a382fb2b7009cba1527c8
	StateRoot string `json:"stateRoot" example:"0x654f28d19b44239d1012f27038f1f71b3d4465dc415a382fb2b7009cba1527c8"`

	// The timestamp of the block.
	// example: 0x5ed28a0e
	Timestamp string `json:"timestamp" example:"0x5ed28a0e"`

	// The total difficulty of the chain until this block.
	// example: 0x1d
	TotalDifficulty string `json:"totalDifficulty" example:"0x1d"`

	// List of transactions included in the block.
	// example: []
	Transactions []string `json:"transactions" example:"[]"`

	// The root hash of the transactions trie.
	// example: 0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421
	TransactionsRoot string `json:"transactionsRoot" example:"0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421"`

	// List of uncle hashes.
	// example: []
	Uncles []string `json:"uncles" example:"[]"`
}


// BlockResponse defines the response for GET /block/{number}.
// swagger:model BlockResponse
type BlockResponse struct {
	// Detailed block information.
	Block BlockDetails `json:"block"`
}

// Handler holds the blockchain client and provides HTTP handlers.
type BlockchainClient interface {
	GetBlockNumber(ctx context.Context) (string, error)
	GetBlockByNumber(ctx context.Context, blockNumber string) (json.RawMessage, error)
}

// RealClient wraps the production client.
type RealClient struct{}

func (rc *RealClient) GetBlockNumber(ctx context.Context) (string, error) {
	return client.GetBlockNumber(ctx)
}

func (rc *RealClient) GetBlockByNumber(ctx context.Context, blockNumber string) (json.RawMessage, error) {
	return client.GetBlockByNumber(ctx, blockNumber)
}

type Handler struct {
	BC BlockchainClient
}

// GetBlockNumberHandler handles the API request to fetch the latest block number.
//
// @Summary Get Latest Block Number
// @Description Retrieves the latest block number from the Polygon RPC endpoint.
// @Tags Blockchain
// @Produce json
// @Success 200 {object} BlockNumberResponse "Example: {\"blockNumber\": \"0x41ce50d\"}"
// @Failure 500 {object} ErrorResponse "Example: {\"error\": \"HTTP request failed: Post \\\"https://polygon-rpc.com/\\\": dial tcp: lookup polygon-rpc.com: no such host\"}"
// @Router /block/number [get]
func (h *Handler) GetBlockNumberHandler(c *gin.Context) {
	ctx := c.Request.Context()
	blockNumber, err := h.BC.GetBlockNumber(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"blockNumber": blockNumber})
}

// GetBlockByNumberHandler handles the API request to fetch a block by number.
//
// @Summary Get Block by Number
// @Description Retrieves block details by block number.
// @Tags Blockchain
// @Produce json
// @Param number path string true "Block Number"
// @Success 200 {object} BlockResponse "Example: {\"block\": {\"difficulty\": \"0x7\", \"extraData\": \"0xd58301090083626f7286676f312e3133856c696e75780000000000000000000020e630202e8ff344c8b9504088f0d59c81fad4fecf54e352319be3085fe5bf6d12e11a34625810f586e80b8388e131b04f4fb964b777d757cba48851880bc1bf00\", \"gasLimit\": \"0x992f4a\", \"gasUsed\": \"0x0\", \"hash\": \"0x30b128813e745454f46978d5073cd9e9474ae78df969ae4928992470fa9379dc\", \"number\": \"0x4\"}}"
// @Failure 500 {object} ErrorResponse "Example: {\"error\": \"HTTP request failed: Post \\\"https://polygon-rpc.com/\\\": dial tcp: lookup polygon-rpc.com: no such host\"}"
// @Router /block/{number} [get]
func (h *Handler) GetBlockByNumberHandler(c *gin.Context) {
	blockNumber := c.Param("number")
	ctx := c.Request.Context()
	block, err := h.BC.GetBlockByNumber(ctx, blockNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var blockData interface{}
	if err := json.Unmarshal(block, &blockData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse block data"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"block": blockData})
}
