package main

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	cmodel "github.com/gsstoykov/go-ethereum-fetcher/contract/model"
	"github.com/gsstoykov/go-ethereum-fetcher/fetcher/cmd/api"
	"github.com/gsstoykov/go-ethereum-fetcher/fetcher/model"
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

	// Get DB connection string and initialize database
	connstr := os.Getenv("DB_CONNECTION_STRING")
	if connstr == "" {
		fmt.Printf("Environment variable %s not set\n", "DB_CONNECTION_STRING")
		os.Exit(1)
	}

	db, err := gorm.Open(postgres.Open(connstr), &gorm.Config{})
	if err != nil {
		fmt.Printf("Could not connect to database: %v\n", err)
		os.Exit(1)
	}

	// Get Ethereum node URL and initialize Ethereum client
	ethurl := os.Getenv("ETH_NODE_URL")
	if ethurl == "" {
		fmt.Printf("Environment variable %s not set\n", "ETH_NODE_URL")
		os.Exit(1)
	}

	client, err := ethclient.Dial(ethurl)
	if err != nil {
		fmt.Printf("Could not connect to Ethereum node: %v\n", err)
		os.Exit(1)
	}

	// Run database migrations
	if err := db.AutoMigrate(&model.User{}, &model.Transaction{}, &cmodel.Person{}); err != nil {
		fmt.Printf("Database migration failed: %v\n", err)
		os.Exit(1)
	}

	// Initialize and start EthereumFetcher
	ef := api.NewEthereumFetcher(db, client, ethurl)
	if err := ef.Listen(); err != nil {
		fmt.Printf("Server failed: %v\n", err)
		os.Exit(1)
	}
}
