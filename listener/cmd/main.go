package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	time.Sleep(600 * time.Second)
	err := godotenv.Load("../.env")
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

	fmt.Println(db)
	fmt.Println(client)
}
