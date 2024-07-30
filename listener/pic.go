package listener

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
	"github.com/gsstoykov/go-ethereum-fetcher/contract/repository"
)

const PIUpdatedABI = `[
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

// SubPIC listens for PersonInfoUpdated events on the Ethereum blockchain and updates the person repository.
// Returns an error if it encounters issues during the subscription or processing.
func SubPIC(ctx context.Context, client *ethclient.Client, prepo repository.IPersonRepository) error {
	// Parse the ABI definition of the contract
	contractAbi, err := abi.JSON(strings.NewReader(PIUpdatedABI))
	if err != nil {
		log.Printf("Failed to parse ABI: %v", err)
		return err
	}

	// Define the event topic for PersonInfoUpdated
	personInfoTopic := crypto.Keccak256Hash([]byte("PersonInfoUpdated(uint256,string,uint256)"))
	contractAddress := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))
	if contractAddress == (common.Address{}) {
		err := fmt.Errorf("CONTRACT_ADDRESS environment variable is not set or invalid")
		log.Print(err)
		return err
	}

	// Create a filter query to listen for the specific event on the given contract address
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
		Topics:    [][]common.Hash{{personInfoTopic}},
	}

	// Create a channel to receive log updates
	ch := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(ctx, query, ch)
	if err != nil {
		log.Printf("Failed to subscribe to logs: %v", err)
		return err
	}
	defer sub.Unsubscribe()

	log.Printf("Listening for events on SimplePersonInfoContract with address: %s", contractAddress.Hex())
	// Process logs received on the channel
	for {
		select {
		case clog := <-ch:
			var eventData struct {
				PersonIndex *big.Int
				NewName     string
				NewAge      *big.Int
			}
			if err := unpackPICLog(contractAbi, clog, &eventData); err != nil {
				log.Printf("Failed to unpack log: %v", err)
				continue
			}

			log.Printf("Received valid log from SimplePersonInfoContract with address: %s", contractAddress.Hex())

			// Create or update person in the repository
			p := &model.Person{Name: eventData.NewName, Age: eventData.NewAge.Int64()}
			if _, err := prepo.Create(p); err != nil {
				log.Printf("Failed to create person: %v", err)
				continue
			}

		case <-ctx.Done():
			// Exit gracefully on context cancellation
			log.Println("Context done, exiting gracefully")
			return nil
		}
	}
}

// unpackPICLog extracts event data from the log using the provided ABI.
// Returns an error if the log does not match the expected event or if parsing fails.
func unpackPICLog(contractAbi abi.ABI, clog types.Log, eventData interface{}) error {
	for _, evInfo := range contractAbi.Events {
		if evInfo.ID.Hex() != clog.Topics[0].Hex() {
			continue
		}
		// Filter indexed arguments
		indexed := make([]abi.Argument, 0, len(evInfo.Inputs))
		for _, input := range evInfo.Inputs {
			if input.Indexed {
				indexed = append(indexed, input)
			}
		}
		// Parse topics and unpack data
		if err := abi.ParseTopics(eventData, indexed, clog.Topics[1:]); err != nil {
			log.Printf("Failed to parse topics: %v", err)
			return err
		}
		if err := contractAbi.UnpackIntoInterface(eventData, evInfo.Name, clog.Data); err != nil {
			log.Printf("Failed to unpack data: %v", err)
			return err
		}
		return nil
	}
	log.Printf("Event not found in log: %v", clog)
	return fmt.Errorf("event not found in log")
}
