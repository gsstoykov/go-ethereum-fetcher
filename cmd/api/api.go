package api

import (
	"net/http"
	"time"

	"gorm.io/gorm"
)

type EthereumFetcher struct {
	port   string
	server *http.Server
	db     *gorm.DB
}

func NewEthereumFetcher(port string, db *gorm.DB) *EthereumFetcher {
	router := NewRouter(db)
	return &EthereumFetcher{
		port: port,
		server: &http.Server{
			Addr:           ":" + port,
			Handler:        router.gin_router,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
		db: db,
	}
}

func (ef EthereumFetcher) Listen() error {
	err := ef.server.ListenAndServe()
	if err != nil {
		panic(err)
	}
	return nil
}
