package contract

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func BuildTransactor(client *ethclient.Client) (*bind.TransactOpts, error) {
	pkeystr := os.Getenv("PRIVATE_KEY")

	if client == nil {
		fmt.Println("client is nil")
		return nil, errors.New("client is nil")
	}

	privateKey, err := crypto.HexToECDSA(pkeystr)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	transactor, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal("Could not create keyed transactor")
	}

	transactor.Nonce = big.NewInt(int64(nonce))
	transactor.Value = big.NewInt(0)     // in wei
	transactor.GasLimit = uint64(300000) // in units
	transactor.GasPrice = gasPrice

	return transactor, err
}
