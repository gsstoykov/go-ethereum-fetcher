package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	crepo "github.com/gsstoykov/go-ethereum-fetcher/contract/repository"
	"github.com/gsstoykov/go-ethereum-fetcher/listener/ws"
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

	ethurl := os.Getenv("WS_NODE_URL")
	client, err := ethclient.Dial(ethurl)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	fmt.Println(db)
	fmt.Println(client)

	personRepository := crepo.NewPersonRepository(db)

	ctx, _ := context.WithCancel(context.Background())
	var el ws.EventListener
	el.Subscirbe(ctx, client, personRepository)

	time.Sleep(time.Second * 100)
}
