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

// NewEthereumGateway creates a new instance of EthereumGateway.
func NewEthereumGateway(ethNodeURL string) IEthereumGateway {
	return &EthereumGateway{
		ethNodeURL: ethNodeURL,
		client:     &http.Client{},
	}
}

// GetByTransactionHash retrieves a transaction by its hash and returns a model.Transaction.
func (eg *EthereumGateway) GetByTransactionHash(txHashString string) (*model.Transaction, error) {
	// Fetch transaction details and receipt
	tx, receipt, err := eg.fetchTransactionData(txHashString)
	if err != nil {
		return nil, err
	}

	// Build and return the transaction model
	transaction := eg.buildTransactionModel(tx, receipt)
	return transaction, nil
}

// fetchTransactionData retrieves transaction and receipt data by hash.
func (eg *EthereumGateway) fetchTransactionData(txHashString string) (map[string]interface{}, map[string]interface{}, error) {
	tx, err := eg.getTransactionByHash(txHashString)
	if err != nil {
		return nil, nil, err
	}

	receipt, err := eg.getTransactionReceipt(txHashString)
	if err != nil {
		return nil, nil, err
	}

	return tx, receipt, nil
}

// getTransactionByHash retrieves a transaction by its hash using JSON-RPC.
func (eg *EthereumGateway) getTransactionByHash(hash string) (map[string]interface{}, error) {
	var response struct {
		Result map[string]interface{} `json:"result"`
	}

	payload := map[string]interface{}{
		"id":      1,
		"jsonrpc": "2.0",
		"method":  "eth_getTransactionByHash",
		"params":  []interface{}{hash},
	}

	requestBody, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", eg.ethNodeURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := eg.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code while fetching transaction: %d", resp.StatusCode)
	}

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&response); err != nil {
		return nil, err
	}

	return response.Result, nil
}

// getTransactionReceipt retrieves a transaction receipt by its hash using JSON-RPC.
func (eg *EthereumGateway) getTransactionReceipt(hash string) (map[string]interface{}, error) {
	var response struct {
		Result map[string]interface{} `json:"result"`
	}

	payload := map[string]interface{}{
		"id":      1,
		"jsonrpc": "2.0",
		"method":  "eth_getTransactionReceipt",
		"params":  []interface{}{hash},
	}

	requestBody, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", eg.ethNodeURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := eg.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code while fetching transaction receipt: %d", resp.StatusCode)
	}

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&response); err != nil {
		return nil, err
	}

	return response.Result, nil
}

// buildTransactionModel creates a model.Transaction instance from the transaction and receipt data.
func (eg *EthereumGateway) buildTransactionModel(tx map[string]interface{}, receipt map[string]interface{}) *model.Transaction {
	to := ""
	if tx["to"] != nil {
		to = tx["to"].(string)
	}

	// Extract from address directly from transaction
	from := tx["from"].(string)

	// Convert block number and status to integers
	blockNumber, _ := strconv.ParseUint(receipt["blockNumber"].(string), 0, 64)
	status, _ := strconv.ParseUint(receipt["status"].(string), 0, 64)

	// Handle nil contract address
	contractAddress := ""
	if receipt["contractAddress"] != nil {
		contractAddress = receipt["contractAddress"].(string)
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
	}
}
