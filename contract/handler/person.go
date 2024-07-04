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

type PersonHandler struct {
	pr     repository.IPersonRepository
	client *ethclient.Client
}

func NewPersonHandler(pr repository.IPersonRepository, client *ethclient.Client) *PersonHandler {
	return &PersonHandler{
		pr:     pr,
		client: client,
	}
}

func (ph PersonHandler) SavePerson(ctx *gin.Context) {
	var person model.Person
	if err := ctx.BindJSON(&person); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	contractAddressStr := os.Getenv("CONTRACT_ADDRESS")

	contractAddress := common.HexToAddress(contractAddressStr)
	instance, err := contract.NewSimplePersonInfoContract(contractAddress, ph.client)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to instantiate smart contract instance"})
		return
	}

	t, err := contract.BuildTransactor(ph.client)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not build transactor"})
		return
	}

	tx, err := instance.SetPersonInfo(t, person.Name, big.NewInt(person.Age))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// Return transaction hash and status
	ctx.JSON(http.StatusOK, gin.H{"txHash": tx.Hash().Hex(), "status": "pending"})
}

func (ph PersonHandler) ListPeople(ctx *gin.Context) {
	var ps []model.Person
	ps, err := ph.pr.FindAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"people": ps})
}
