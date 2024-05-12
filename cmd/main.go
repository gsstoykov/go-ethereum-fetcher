package main

import (
	"log"
	"os"

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

	db, err := gorm.Open(postgres.Open(connstr), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to db: %v", err)
		panic(err)
	}

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Transaction{})

	ef := api.NewEthereumFetcher(port, db)
	ef.Listen()
}
