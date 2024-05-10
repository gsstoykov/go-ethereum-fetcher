package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gsstoykov/go-ethereum-fetcher/handlers"
)

type Router struct {
	gin_router *gin.Engine
}

func NewRouter() *Router {
	router := &Router{
		gin_router: gin.Default(),
	}
	router.gin_router.GET("users", handlers.GetUsers)
	return router
}
