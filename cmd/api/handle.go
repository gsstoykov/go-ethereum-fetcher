package api

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	egateway "github.com/gsstoykov/go-ethereum-fetcher/ethereum"
	"github.com/gsstoykov/go-ethereum-fetcher/handlers"
	"github.com/gsstoykov/go-ethereum-fetcher/handlers/middleware"
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
	transactionHandler := handlers.NewTransactionHandler(
		repository.NewTransactionRepository(hm.db),
		egateway.NewEthereumGateway(hm.client),
	)
	// user routes
	hm.router.GET("/users", middleware.AuthenticateMiddleware(), userHandler.FetchUsers)
	hm.router.POST("/user", userHandler.CreateUser)
	hm.router.POST("/auth", userHandler.Authenticate)
	// transaction routes
	hm.router.GET("/transactions", transactionHandler.FetchTransactions)
	hm.router.POST("/transaction", transactionHandler.CreateTransaction)
	hm.router.GET("/eth/:hash", transactionHandler.FetchTransactionsList)
	return hm.router
}
