package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gsstoykov/go-ethereum-fetcher/handlers"
	"github.com/gsstoykov/go-ethereum-fetcher/repository"
	"gorm.io/gorm"
)

type HandleRouter struct {
	gin_router *gin.Engine
}

func NewRouter(db *gorm.DB) *HandleRouter {
	router := &HandleRouter{
		gin_router: gin.Default(),
	}
	userHandler := handlers.NewUserHandler(repository.NewUserRepository(db))
	transactionHandler := handlers.NewTransactionHandler(repository.NewTransactionRepository(db))
	// user routes
	router.gin_router.GET("users", userHandler.FetchUsers)
	router.gin_router.POST("user", userHandler.CreateUser)
	// transaction routes
	router.gin_router.GET("transactions", transactionHandler.FetchTransactions)
	router.gin_router.POST("transaction", transactionHandler.CreateTransaction)
	return router
}
