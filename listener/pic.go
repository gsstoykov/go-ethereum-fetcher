package listener

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gsstoykov/go-ethereum-fetcher/contract/model"
	"github.com/gsstoykov/go-ethereum-fetcher/contract/repository"
)

const personInfoAbi = `[
    {
        "anonymous": false,
        "inputs": [
            {
                "indexed": true,
                "internalType": "uint256",
                "name": "personIndex",
                "type": "uint256"
            },
            {
                "indexed": false,
                "internalType": "string",
                "name": "newName",
                "type": "string"
            },
            {
                "indexed": false,
                "internalType": "uint256",
                "name": "newAge",
                "type": "uint256"
            }
        ],
        "name": "PersonInfoUpdated",
        "type": "event"
    }
]`

func SubPIC(ctx context.Context, client *ethclient.Client, prepo repository.IPersonRepository) {
	contractAbi, err := abi.JSON(strings.NewReader(personInfoAbi))
	if err != nil {
		panic(err)
	}
	personInfoTopic := crypto.Keccak256Hash([]byte("PersonInfoUpdated(uint256,string,uint256)"))
	contractAddress := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))

	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
		Topics:    [][]common.Hash{{personInfoTopic}},
	}

	ch := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(ctx, query, ch)
	if err != nil {
		panic(err)
	}
	defer sub.Unsubscribe()

	for {
		select {
		case log := <-ch:
			var eventData struct {
				PersonIndex *big.Int
				NewName     string
				NewAge      *big.Int
			}
			if err := unpackPICLog(contractAbi, log, &eventData); err != nil {
				continue
			}
			p := &model.Person{Name: eventData.NewName, Age: eventData.NewAge.Int64()}
			if _, err := prepo.Create(p); err != nil {
				continue
			}
		case <-ctx.Done():
			return
		}
	}
}

func unpackPICLog(contractAbi abi.ABI, log types.Log, eventData interface{}) error {
	for _, evInfo := range contractAbi.Events {
		if evInfo.ID.Hex() != log.Topics[0].Hex() {
			continue
		}
		indexed := make([]abi.Argument, 0, len(evInfo.Inputs))
		for _, input := range evInfo.Inputs {
			if input.Indexed {
				indexed = append(indexed, input)
			}
		}
		if err := abi.ParseTopics(eventData, indexed, log.Topics[1:]); err != nil {
			return err
		}
		return contractAbi.UnpackIntoInterface(eventData, evInfo.Name, log.Data)
	}
	return fmt.Errorf("event not found")
}
