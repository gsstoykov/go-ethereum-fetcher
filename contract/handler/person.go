package handler

import (
	"math/big"
	"net/http"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/gsstoykov/go-ethereum-fetcher/contract"
	"github.com/gsstoykov/go-ethereum-fetcher/contract/model"
	"github.com/gsstoykov/go-ethereum-fetcher/contract/repository"
)

// PersonHandler handles HTTP requests related to Person entities.
type PersonHandler struct {
	pr     repository.IPersonRepository
	client *ethclient.Client
}

// NewPersonHandler creates a new instance of PersonHandler.
func NewPersonHandler(pr repository.IPersonRepository, client *ethclient.Client) *PersonHandler {
	return &PersonHandler{
		pr:     pr,
		client: client,
	}
}

// SavePerson saves a new Person entity and sets its info on the Ethereum smart contract.
func (ph *PersonHandler) SavePerson(ctx *gin.Context) {
	var person model.Person

	// Bind JSON input to the person struct.
	if err := ctx.BindJSON(&person); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	contractAddressStr := os.Getenv("CONTRACT_ADDRESS")
	contractAddress := common.HexToAddress(contractAddressStr)

	// Instantiate the smart contract.
	instance, err := contract.NewSimplePersonInfoContract(contractAddress, ph.client)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to instantiate smart contract instance"})
		return
	}

	// Build the transactor.
	transactor, err := contract.BuildTransactor(ph.client)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not build transactor"})
		return
	}

	// Set person info on the smart contract.
	tx, err := instance.SetPersonInfo(transactor, person.Name, big.NewInt(person.Age))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set person info on smart contract"})
		return
	}

	// Return the transaction hash and status.
	ctx.JSON(http.StatusOK, gin.H{"txHash": tx.Hash().Hex(), "status": "pending"})
}

// ListPeople lists all saved Person entities.
func (ph *PersonHandler) ListPeople(ctx *gin.Context) {
	people, err := ph.pr.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch people"})
		return
	}

	// Return the list of people.
	ctx.JSON(http.StatusOK, gin.H{"people": people})
}
