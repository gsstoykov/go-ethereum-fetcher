package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
)

// EthereumFetcher represents the main server for the Ethereum fetcher.
type EthereumFetcher struct {
	server *http.Server
}

// NewEthereumFetcher initializes and returns a new EthereumFetcher.
func NewEthereumFetcher(db *gorm.DB, client *ethclient.Client, ethurl string) *EthereumFetcher {
	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8080" // Default to port 8080 if API_PORT is not set
	}

	hm := NewHandleManager(db, client, ethurl)
	router := hm.InitRouter()

	server := &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return &EthereumFetcher{server: server}
}

// Listen starts the HTTP server and implements graceful shutdown.
func (ef *EthereumFetcher) Listen() error {
	// Channel to listen for interrupt signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// Start the server in a separate goroutine
	go func() {
		if err := ef.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on %s: %v", ef.server.Addr, err)
		}
	}()

	// Wait for interrupt signal
	<-stop

	// Create a deadline to wait for the server to shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("Shutting down server...")
	if err := ef.server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")
	return nil
}
