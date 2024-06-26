package api

import (
	"net/http"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
)

type EthereumFetcher struct {
	server *http.Server
}

func NewEthereumFetcher(db *gorm.DB, client *ethclient.Client) *EthereumFetcher {
	hm := NewHandleManager(db, client)
	return &EthereumFetcher{
		server: &http.Server{
			Addr:           ":" + os.Getenv("API_PORT"),
			Handler:        hm.InitRouter(),
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
	}
}

func (ef EthereumFetcher) Listen() error {
	err := ef.server.ListenAndServe()
	if err != nil {
		panic(err)
	}
	return nil
}
