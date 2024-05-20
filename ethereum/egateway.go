package egateway

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gsstoykov/go-ethereum-fetcher/model"
)

type IEthereumGateway interface {
	GetByTransactionHash(txHashString string) (*model.Transaction, error)
}

type EthereumGateway struct {
	client *ethclient.Client
}

func NewEthereumGateway(client *ethclient.Client) IEthereumGateway {
	return &EthereumGateway{client: client}
}

func (eg *EthereumGateway) GetByTransactionHash(txHashString string) (*model.Transaction, error) {
	// Retrieve transaction by hash
	txHash := common.HexToHash(txHashString)
	tx, _, err := eg.client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve transaction: %v", err)
	}

	// Retrieve transaction receipt
	receipt, err := eg.client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve transaction receipt: %v", err)
	}

	// Derive the sender (from) address
	chainID, err := eg.client.NetworkID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get network ID: %v", err)
	}
	signer := types.NewEIP155Signer(chainID)
	from, err := types.Sender(signer, tx)
	if err != nil {
		return nil, fmt.Errorf("failed to derive sender from transaction: %v", err)
	}

	// Get the recipient (to) address
	to := ""
	if tx.To() != nil {
		to = tx.To().Hex()
	}

	// Marshal the transaction and receipt to JSON
	txJSON, err := tx.MarshalJSON()
	if err != nil {
		return nil, fmt.Errorf("failed to marshal transaction to JSON: %v", err)
	}
	receiptJSON, err := receipt.MarshalJSON()
	if err != nil {
		return nil, fmt.Errorf("failed to marshal receipt to JSON: %v", err)
	}

	// Print the JSON representations
	fmt.Println("Transaction JSON:", string(txJSON))
	fmt.Println("Receipt JSON:", string(receiptJSON))

	// Create the transaction model
	transaction := &model.Transaction{
		TransactionHash:   tx.Hash().Hex(),
		TransactionStatus: int(receipt.Status),
		BlockHash:         receipt.BlockHash.Hex(),
		BlockNumber:       int(receipt.BlockNumber.Uint64()),
		From:              from.Hex(),
		To:                to,
		ContractAddress:   receipt.ContractAddress.Hex(),
		LogsCount:         len(receipt.Logs),
		Input:             common.Bytes2Hex(tx.Data()),
		Value:             tx.Value().String(),
	}

	// Print transaction details
	fmt.Printf("Transaction: %v\n", transaction)
	return transaction, nil
}
