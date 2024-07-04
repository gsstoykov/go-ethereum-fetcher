package main

import (
	"log"
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
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env file: %v", err)
		panic(err)
	}

	connstr := os.Getenv("DB_CONNECTION_STRING")
	db, err := gorm.Open(postgres.Open(connstr), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to db: %v", err)
		panic(err)
	}

	ethurl := os.Getenv("ETH_NODE_URL")
	client, err := ethclient.Dial(ethurl)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	db.AutoMigrate(&model.User{}, &model.Transaction{}, &cmodel.Person{})

	ef := api.NewEthereumFetcher(db, client)
	ef.Listen()
}
