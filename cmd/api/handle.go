package api

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/gsstoykov/go-ethereum-fetcher/handlers"
	"github.com/gsstoykov/go-ethereum-fetcher/repository"
	"gorm.io/gorm"
)

type HandleManager struct {
	router *gin.Engine
	client *ethclient.Client
	db     *gorm.DB
}

func NewHandleManager(db *gorm.DB, client *ethclient.Client) *HandleManager {
	return &HandleManager{
		router: gin.Default(),
		client: client,
		db:     db,
	}
}

func (hm *HandleManager) InitRouter() *gin.Engine {
	// model handlers
	userHandler := handlers.NewUserHandler(repository.NewUserRepository(hm.db))
	transactionHandler := handlers.NewTransactionHandler(repository.NewTransactionRepository(hm.db))
	// user routes
	hm.router.GET("users", userHandler.FetchUsers)
	hm.router.POST("user", userHandler.CreateUser)
	// transaction routes
	hm.router.GET("transactions", transactionHandler.FetchTransactions)
	hm.router.POST("transaction", transactionHandler.CreateTransaction)
	return hm.router
}
