package main

import (
	"log"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"

	"github.com/gsstoykov/go-ethereum-fetcher/contract"
)

func main() {
	// Load environment variables from .env file.
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get Ethereum node URL from environment variables.
	ethurl := os.Getenv("ETH_NODE_URL")
	if ethurl == "" {
		log.Fatal("ETH_NODE_URL environment variable not set")
	}

	// Connect to the Ethereum client.
	client, err := ethclient.Dial(ethurl)
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum client: %v", err)
	}

	// Build the transactor using the Ethereum client.
	transactor, err := contract.BuildTransactor(client)
	if err != nil {
		log.Fatalf("Failed to build transactor: %v", err)
	}

	// Deploy the SimplePersonInfoContract using the transactor and client.
	address, tx, _, err := contract.DeploySimplePersonInfoContract(transactor, client)
	if err != nil {
		log.Fatalf("Failed to deploy contract: %v", err)
	}

	// Print the deployed contract address and transaction hash.
	log.Printf("Contract deployed at address: %s", address.Hex())
	log.Printf("Transaction hash: %s", tx.Hash().Hex())
}
