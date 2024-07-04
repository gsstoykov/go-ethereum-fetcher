package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"

	"github.com/gsstoykov/go-ethereum-fetcher/contract"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("error loading .env file: %v", err)
		panic(err)
	}

	ethurl := os.Getenv("ETH_NODE_URL")
	client, err := ethclient.Dial(ethurl)
	if err != nil {
		log.Fatal(err)
	}

	transactor, err := contract.BuildTransactor(client)
	if err != nil {
		log.Fatal(err)
	}

	address, tx, instance, err := contract.DeploySimplePersonInfoContract(transactor, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(address.Hex())
	fmt.Println(tx.Hash().Hex())

	_ = instance
}
