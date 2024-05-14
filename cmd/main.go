package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gsstoykov/go-ethereum-fetcher/cmd/api"
	"github.com/gsstoykov/go-ethereum-fetcher/model"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env file: %v", err)
		panic(err)
	}

	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8080" // Default port if not provided
	}

	connstr := os.Getenv("DB_CONNECTION_STRING")
	if connstr == "" {
		log.Fatalf("bad db connection string: %v", err)
		panic(err)
	}

	ethurl := os.Getenv("ETH_NODE_URL")
	if ethurl == "" {
		log.Fatalf("bad ethurl string: %v", err)
		panic(err)
	}

	client, err := ethclient.Dial(ethurl)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	txHash := common.HexToHash("0xb659a9e20c44392ea0fe8f717a85f31dfa85346dd4632adc86a82fba273ff245")
	tx, _, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}

	// Print transaction details
	fmt.Printf("Transaction: %v\n", tx)

	db, err := gorm.Open(postgres.Open(connstr), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to db: %v", err)
		panic(err)
	}

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Transaction{})

	ef := api.NewEthereumFetcher(port, db, client)
	ef.Listen()
}
