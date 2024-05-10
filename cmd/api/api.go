package api

import (
	"context"
	"net/http"
	"time"
)

type EthereumFetcher struct {
	port   string
	server *http.Server
}

func NewEthereumFetcher(port string, ctx context.Context, ctxcf context.CancelFunc) *EthereumFetcher {
	router := NewRouter()
	return &EthereumFetcher{
		port: port,
		server: &http.Server{
			Addr:           ":" + port,
			Handler:        router.gin_router,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
	}
}
