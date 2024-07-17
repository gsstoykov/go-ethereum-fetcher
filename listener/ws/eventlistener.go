package ws

import (
	"context"
	"fmt"
	"log"
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
	crepo "github.com/gsstoykov/go-ethereum-fetcher/contract/repository"
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

type EventListener struct {
}

func (el *EventListener) Subscirbe(ctx context.Context, client *ethclient.Client, prepo crepo.IPersonRepository) {
	contractAbi, err := abi.JSON(strings.NewReader(personInfoAbi))
	if err != nil {
		panic(err)
	}
	personInfoTopic := crypto.Keccak256Hash([]byte("PersonInfoUpdated(uint256,string,uint256)"))
	log.Println("Subscribed to PersonInfoUpdated events")

	contractAddressStr := os.Getenv("CONTRACT_ADDRESS")
	contractAddress := common.HexToAddress(contractAddressStr)

	query := ethereum.FilterQuery{
		FromBlock: nil, // Start from latest block
		ToBlock:   nil, // Monitor new blocks
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
				NewName     string
				NewAge      *big.Int
				PersonIndex *big.Int
			}

			for _, evInfo := range contractAbi.Events {
				if evInfo.ID.Hex() != log.Topics[0].Hex() {
					continue
				}

				indexed := make([]abi.Argument, 0)
				for _, input := range evInfo.Inputs {
					if input.Indexed {
						indexed = append(indexed, input)
					}
				}

				if err := abi.ParseTopics(&eventData, indexed, log.Topics[1:]); err != nil {
					continue
				}

				if err := contractAbi.UnpackIntoInterface(&eventData, evInfo.Name, log.Data); err != nil {
					continue
				}
				break
			}

			fmt.Println(eventData)

			p := &model.Person{
				Name: eventData.NewName,
				Age:  eventData.NewAge.Int64(),
			}

			pp, _ := prepo.Create(p)

			fmt.Println("Created: ", pp)

			if err != nil {
				continue
			}
		case <-ctx.Done():
			sub.Unsubscribe()
			return
		}
	}
}
