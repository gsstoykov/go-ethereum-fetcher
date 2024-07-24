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
		return fmt.Errorf("failed to parse ABI: %w", err)
	}

	// Define the event topic for PersonInfoUpdated
	personInfoTopic := crypto.Keccak256Hash([]byte("PersonInfoUpdated(uint256,string,uint256)"))
	contractAddress := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))
	if contractAddress == (common.Address{}) {
		return fmt.Errorf("CONTRACT_ADDRESS environment variable is not set or invalid")
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
		return fmt.Errorf("failed to subscribe to logs: %w", err)
	}
	defer sub.Unsubscribe()

	fmt.Println("Listening for events on SimplePersonInfoContract with address: ", contractAddress)
	// Process logs received on the channel
	for {
		select {
		case log := <-ch:
			var eventData struct {
				PersonIndex *big.Int
				NewName     string
				NewAge      *big.Int
			}
			if err := unpackPICLog(contractAbi, log, &eventData); err != nil {
				fmt.Printf("Failed to unpack log: %v\n", err)
				continue
			}

			fmt.Println("Received valid log from SimplePersonInfoContract with address: ", contractAddress)

			// Create or update person in the repository
			p := &model.Person{Name: eventData.NewName, Age: eventData.NewAge.Int64()}
			if _, err := prepo.Create(p); err != nil {
				fmt.Printf("Failed to create person: %v\n", err)
				continue
			}

		case <-ctx.Done():
			// Exit gracefully on context cancellation
			return nil
		}
	}
}

// unpackPICLog extracts event data from the log using the provided ABI.
// Returns an error if the log does not match the expected event or if parsing fails.
func unpackPICLog(contractAbi abi.ABI, log types.Log, eventData interface{}) error {
	for _, evInfo := range contractAbi.Events {
		if evInfo.ID.Hex() != log.Topics[0].Hex() {
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
		if err := abi.ParseTopics(eventData, indexed, log.Topics[1:]); err != nil {
			return fmt.Errorf("failed to parse topics: %w", err)
		}
		return contractAbi.UnpackIntoInterface(eventData, evInfo.Name, log.Data)
	}
	return fmt.Errorf("event not found in log")
}
