package egateway

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
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
		return nil, err
	}

	// Print transaction details
	fmt.Printf("Transaction: %v\n", tx)
	return nil, nil
}
