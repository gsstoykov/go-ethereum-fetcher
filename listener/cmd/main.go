package main

import (
	"context"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gsstoykov/go-ethereum-fetcher/contract/repository"
	"github.com/gsstoykov/go-ethereum-fetcher/listener"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Error loading .env file: %v\n", err)
		os.Exit(1)
	}

	// Get database connection string and initialize database
	connstr := os.Getenv("DB_CONNECTION_STRING")
	if connstr == "" {
		fmt.Printf("Environment variable DB_CONNECTION_STRING is not set\n")
		os.Exit(1)
	}

	db, err := gorm.Open(postgres.Open(connstr), &gorm.Config{})
	if err != nil {
		fmt.Printf("Could not connect to database: %v\n", err)
		os.Exit(1)
	}

	// Get WebSocket node URL and initialize Ethereum client
	ethurl := os.Getenv("WS_NODE_URL")
	if ethurl == "" {
		fmt.Printf("Environment variable WS_NODE_URL is not set\n")
		os.Exit(1)
	}

	client, err := ethclient.Dial(ethurl)
	if err != nil {
		fmt.Printf("Could not connect to Ethereum node: %v\n", err)
		os.Exit(1)
	}

	// Initialize person repository
	personRepository := repository.NewPersonRepository(db)

	// Create a context with cancellation
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Ensure the context is cancelled when main exits

	// Start the listener
	if err := listener.SubPIC(ctx, client, personRepository); err != nil {
		fmt.Printf("Error in listener: %v\n", err)
		os.Exit(1)
	}
}
