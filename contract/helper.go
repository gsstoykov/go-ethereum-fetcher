package contract

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// BuildTransactor creates a new TransactOpts using the provided Ethereum client and private key from the environment variables.
func BuildTransactor(client *ethclient.Client) (*bind.TransactOpts, error) {
	// Get the private key from environment variables.
	pkeystr := os.Getenv("PRIVATE_KEY")
	if pkeystr == "" {
		return nil, errors.New("PRIVATE_KEY environment variable not set")
	}

	// Ensure the client is not nil.
	if client == nil {
		return nil, errors.New("Ethereum client is nil")
	}

	// Convert the private key from hex string to ECDSA.
	privateKey, err := crypto.HexToECDSA(pkeystr)
	if err != nil {
		return nil, fmt.Errorf("failed to convert hex string to ECDSA: %v", err)
	}

	// Get the public key from the private key.
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("error casting public key to ECDSA")
	}

	// Get the address of the public key.
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// Get the nonce for the address.
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to get nonce: %v", err)
	}

	// Get the chain ID.
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get chain ID: %v", err)
	}

	// Suggest the gas price.
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to suggest gas price: %v", err)
	}

	// Create a keyed transactor with the chain ID.
	transactor, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, fmt.Errorf("could not create keyed transactor: %v", err)
	}

	// Set the transaction options.
	transactor.Nonce = big.NewInt(int64(nonce))
	transactor.Value = big.NewInt(0)     // in wei
	transactor.GasLimit = uint64(300000) // in units
	transactor.GasPrice = gasPrice

	return transactor, nil
}
