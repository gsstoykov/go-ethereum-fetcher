package main

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"

	"github.com/gsstoykov/go-ethereum-fetcher/contract"
)

func main() {
	// Load environment variables from .env file.
	err := godotenv.Load("../../.env")
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		os.Exit(1)
	}

	// Get Ethereum node URL from environment variables.
	ethurl := os.Getenv("ETH_NODE_URL")
	if ethurl == "" {
		fmt.Println("ETH_NODE_URL environment variable not set")
		os.Exit(1)
	}

	// Connect to the Ethereum client.
	client, err := ethclient.Dial(ethurl)
	if err != nil {
		fmt.Println("Failed to connect to Ethereum client:", err)
		os.Exit(1)
	}

	// Build the transactor using the Ethereum client.
	transactor, err := contract.BuildTransactor(client)
	if err != nil {
		fmt.Println("Failed to build transactor:", err)
		os.Exit(1)
	}

	// Deploy the SimplePersonInfoContract using the transactor and client.
	address, tx, _, err := contract.DeploySimplePersonInfoContract(transactor, client)
	if err != nil {
		fmt.Println("Failed to deploy contract:", err)
		os.Exit(1)
	}

	// Print the deployed contract address and transaction hash.
	fmt.Println("Contract deployed at address:", address.Hex())
	fmt.Println("Transaction hash:", tx.Hash().Hex())
}
