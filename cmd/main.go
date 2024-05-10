package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gsstoykov/go-ethereum-fetcher/handlers"
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

	router := gin.Default()
	router.GET("users", handlers.GetUsers)

	server := &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	db, err := gorm.Open(postgres.Open(connstr), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to db: %v", err)
		panic(err)
	}

	fmt.Println(db)

	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
