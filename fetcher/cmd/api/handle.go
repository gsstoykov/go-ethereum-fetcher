package api

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	chandler "github.com/gsstoykov/go-ethereum-fetcher/contract/handler"
	crepo "github.com/gsstoykov/go-ethereum-fetcher/contract/repository"
	egateway "github.com/gsstoykov/go-ethereum-fetcher/fetcher/ethereum"
	"github.com/gsstoykov/go-ethereum-fetcher/fetcher/handlers"
	"github.com/gsstoykov/go-ethereum-fetcher/fetcher/handlers/middleware"
	"github.com/gsstoykov/go-ethereum-fetcher/fetcher/repository"
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
		repository.NewUserRepository(hm.db),
		egateway.NewEthereumGateway(hm.client),
	)
	personHandler := chandler.NewPersonHandler(
		crepo.NewPersonRepository(hm.db),
		hm.client,
	)
	// user routes
	hm.router.GET("/users", userHandler.FetchUsers)
	hm.router.POST("/user", userHandler.CreateUser)
	hm.router.POST("/auth", userHandler.Authenticate)
	hm.router.GET("/my", middleware.AuthenticateMiddleware(), userHandler.FetchUserTransactions)
	// transaction routes
	hm.router.GET("/transactions", transactionHandler.FetchTransactions)
	hm.router.GET("/eth", middleware.AuthenticateMiddleware(), transactionHandler.FetchTransactionsList)
	// person routes
	hm.router.POST("/savePerson", personHandler.SavePerson)
	hm.router.GET("/listPeople", personHandler.ListPeople)
	return hm.router
}
