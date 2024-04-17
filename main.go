package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gsstoykov/go-ethereum-fetcher/handlers"
)

func main() {
	router := gin.Default()
	router.GET("users", handlers.GetUsers)
	router.Run(":8080")
	fmt.Println("Hello, Go!")
}
