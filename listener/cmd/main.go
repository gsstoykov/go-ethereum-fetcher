package main

import (
	"context"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gsstoykov/go-ethereum-fetcher/contract/repository"
	"github.com/gsstoykov/go-ethereum-fetcher/listener"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Get database connection string and initialize database
	connstr := os.Getenv("DB_CONNECTION_STRING")
	if connstr == "" {
		log.Fatalf("Environment variable DB_CONNECTION_STRING is not set")
	}

	db, err := gorm.Open(postgres.Open(connstr), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	// Get WebSocket node URL and initialize Ethereum client
	ethurl := os.Getenv("WS_NODE_URL")
	if ethurl == "" {
		log.Fatalf("Environment variable WS_NODE_URL is not set")
	}

	client, err := ethclient.Dial(ethurl)
	if err != nil {
		log.Fatalf("Could not connect to Ethereum node: %v", err)
	}

	// Initialize person repository
	personRepository := repository.NewPersonRepository(db)

	// Create a context with cancellation
	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel() // Ensure the context is cancelled when main exits
		log.Println("Shutting down gracefully...")
	}()

	// Start the listener
	if err := listener.SubPIC(ctx, client, personRepository); err != nil {
		log.Fatalf("Error in listener: %v", err)
	}
}
