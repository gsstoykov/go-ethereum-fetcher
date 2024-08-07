package main

import (
	"log"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	cmodel "github.com/gsstoykov/go-ethereum-fetcher/contract/model"
	"github.com/gsstoykov/go-ethereum-fetcher/fetcher/cmd/api"
	"github.com/gsstoykov/go-ethereum-fetcher/fetcher/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	// Get DB connection string and initialize database
	connstr := os.Getenv("DB_CONNECTION_STRING")
	if connstr == "" {
		log.Fatalf("Environment variable %s not set", "DB_CONNECTION_STRING")
	}

	db, err := gorm.Open(postgres.Open(connstr), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  logger.Silent,
				IgnoreRecordNotFoundError: true,
				ParameterizedQueries:      true,
			},
		),
	})
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	// Run database migrations
	if err := db.AutoMigrate(&model.User{}, &model.Transaction{}, &cmodel.Person{}); err != nil {
		log.Fatalf("Database migration failed: %v", err)
	}

	// Get Ethereum node URL and initialize Ethereum client
	ethurl := os.Getenv("ETH_NODE_URL")
	if ethurl == "" {
		log.Fatalf("Environment variable %s not set", "ETH_NODE_URL")
	}

	client, err := ethclient.Dial(ethurl)
	if err != nil {
		log.Fatalf("Could not connect to Ethereum node: %v", err)
	}

	// Initialize and start EthereumFetcher
	ef := api.NewEthereumFetcher(db, client, ethurl)
	if err := ef.Listen(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
