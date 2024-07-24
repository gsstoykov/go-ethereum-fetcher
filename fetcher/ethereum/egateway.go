package egateway

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gsstoykov/go-ethereum-fetcher/fetcher/model"
)

// IEthereumGateway defines the interface for interacting with Ethereum transactions.
type IEthereumGateway interface {
	GetByTransactionHash(txHashString string) (*model.Transaction, error)
}

// EthereumGateway provides methods to interact with Ethereum.
type EthereumGateway struct {
	ethNodeURL string
	client     *http.Client
}

// NewEthereumGateway creates a new instance of EthereumGateway with the specified Ethereum node URL.
func NewEthereumGateway(ethNodeURL string) IEthereumGateway {
	return &EthereumGateway{
		ethNodeURL: ethNodeURL,
		client:     &http.Client{}, // Initialize HTTP client
	}
}

// GetByTransactionHash retrieves a transaction by its hash and returns a model.Transaction.
func (eg *EthereumGateway) GetByTransactionHash(txHashString string) (*model.Transaction, error) {
	// Fetch transaction and receipt data
	tx, receipt, err := eg.fetchTransactionData(txHashString)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch transaction data: %w", err)
	}

	// Build transaction model from fetched data
	transaction, err := eg.buildTransactionModel(tx, receipt)
	if err != nil {
		return nil, fmt.Errorf("failed to build transaction model: %w", err)
	}

	return transaction, nil
}

// fetchTransactionData retrieves both transaction and receipt data by hash.
func (eg *EthereumGateway) fetchTransactionData(txHashString string) (map[string]interface{}, map[string]interface{}, error) {
	// Get transaction data by hash
	tx, err := eg.getTransactionByHash(txHashString)
	if err != nil {
		return nil, nil, fmt.Errorf("error fetching transaction by hash: %w", err)
	}

	// Get transaction receipt by hash
	receipt, err := eg.getTransactionReceipt(txHashString)
	if err != nil {
		return nil, nil, fmt.Errorf("error fetching transaction receipt: %w", err)
	}

	return tx, receipt, nil
}

// getTransactionByHash retrieves a transaction by its hash using JSON-RPC.
func (eg *EthereumGateway) getTransactionByHash(hash string) (map[string]interface{}, error) {
	payload := map[string]interface{}{
		"id":      1,
		"jsonrpc": "2.0",
		"method":  "eth_getTransactionByHash",
		"params":  []interface{}{hash},
	}

	// Send JSON-RPC request
	response, err := eg.sendRequest(payload)
	if err != nil {
		return nil, fmt.Errorf("error sending request to get transaction: %w", err)
	}

	return response.Result, nil
}

// getTransactionReceipt retrieves a transaction receipt by its hash using JSON-RPC.
func (eg *EthereumGateway) getTransactionReceipt(hash string) (map[string]interface{}, error) {
	payload := map[string]interface{}{
		"id":      1,
		"jsonrpc": "2.0",
		"method":  "eth_getTransactionReceipt",
		"params":  []interface{}{hash},
	}

	// Send JSON-RPC request
	response, err := eg.sendRequest(payload)
	if err != nil {
		return nil, fmt.Errorf("error sending request to get transaction receipt: %w", err)
	}

	return response.Result, nil
}

// sendRequest sends a JSON-RPC request and returns the response.
func (eg *EthereumGateway) sendRequest(payload map[string]interface{}) (*jsonResponse, error) {
	// Marshal payload to JSON
	requestBody, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request payload: %w", err)
	}

	// Create HTTP POST request
	req, err := http.NewRequest("POST", eg.ethNodeURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("error creating HTTP request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Execute HTTP request
	resp, err := eg.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error executing HTTP request: %w", err)
	}
	defer resp.Body.Close()

	// Check for unexpected status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Decode response body
	var response jsonResponse
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&response); err != nil {
		return nil, fmt.Errorf("error decoding response body: %w", err)
	}

	return &response, nil
}

// jsonResponse defines the structure of the JSON-RPC response.
type jsonResponse struct {
	Result map[string]interface{} `json:"result"`
}

// buildTransactionModel creates a model.Transaction instance from the transaction and receipt data.
func (eg *EthereumGateway) buildTransactionModel(tx map[string]interface{}, receipt map[string]interface{}) (*model.Transaction, error) {
	// Extract 'to' address from transaction data
	to, ok := tx["to"].(string)
	if !ok {
		to = ""
	}

	// Extract 'from' address from transaction data
	from, ok := tx["from"].(string)
	if !ok {
		return nil, fmt.Errorf("invalid 'from' address in transaction data")
	}

	// Convert block number and status to integers
	blockNumber, err := strconv.ParseUint(receipt["blockNumber"].(string), 0, 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing block number: %w", err)
	}

	status, err := strconv.ParseUint(receipt["status"].(string), 0, 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing status: %w", err)
	}

	// Extract contract address from receipt data
	contractAddress, ok := receipt["contractAddress"].(string)
	if !ok {
		contractAddress = ""
	}

	return &model.Transaction{
		TransactionHash:   tx["hash"].(string),
		TransactionStatus: int(status),
		BlockHash:         receipt["blockHash"].(string),
		BlockNumber:       int(blockNumber),
		From:              from,
		To:                to,
		ContractAddress:   contractAddress,
		LogsCount:         len(receipt["logs"].([]interface{})),
		Input:             tx["input"].(string),
		Value:             tx["value"].(string),
	}, nil
}
