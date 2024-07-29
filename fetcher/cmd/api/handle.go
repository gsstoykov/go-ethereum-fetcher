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

// HandleManager manages the setup and initialization of routes and handlers.
type HandleManager struct {
	router *gin.Engine
	client *ethclient.Client
	db     *gorm.DB
	ethurl string
}

// NewHandleManager creates a new HandleManager with the given database, Ethereum client, and Ethereum URL.
func NewHandleManager(db *gorm.DB, client *ethclient.Client, ethurl string) *HandleManager {
	return &HandleManager{
		router: gin.Default(),
		client: client,
		db:     db,
		ethurl: ethurl,
	}
}

// InitRouter initializes all routes and their respective handlers.
func (hm *HandleManager) InitRouter() *gin.Engine {
	hm.setupUserRoutes()
	hm.setupTransactionRoutes()
	hm.setupPersonRoutes()
	return hm.router
}

// setupUserRoutes sets up routes related to user operations.
func (hm *HandleManager) setupUserRoutes() {
	userHandler := handlers.NewUserHandler(repository.NewUserRepository(hm.db))

	hm.router.GET("/users", userHandler.FetchUsers)
	hm.router.POST("/user", userHandler.CreateUser)
	hm.router.POST("/auth", userHandler.Authenticate)
	hm.router.GET("/my", middleware.AuthenticateMiddleware(), userHandler.FetchUserTransactions)
}

// setupTransactionRoutes sets up routes related to transaction operations.
func (hm *HandleManager) setupTransactionRoutes() {
	transactionHandler := handlers.NewTransactionHandler(
		repository.NewTransactionRepository(hm.db),
		repository.NewUserRepository(hm.db),
		egateway.NewEthereumGateway(hm.ethurl),
	)

	hm.router.GET("/transactions", transactionHandler.FetchTransactions)
	hm.router.GET("/eth", middleware.AuthenticateMiddleware(), transactionHandler.FetchTransactionsList)
	hm.router.GET("/eth/:rlphex", middleware.RLPDecodeMiddleware(), middleware.AuthenticateMiddleware(), transactionHandler.FetchTransactionsList)
}

// setupPersonRoutes sets up routes related to person operations.
func (hm *HandleManager) setupPersonRoutes() {
	personHandler := chandler.NewPersonHandler(
		crepo.NewPersonRepository(hm.db),
		hm.client,
	)

	hm.router.POST("/savePerson", personHandler.SavePerson)
	hm.router.GET("/listPeople", personHandler.ListPeople)
}
