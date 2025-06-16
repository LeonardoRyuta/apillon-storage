package nfts

import (
	"fmt"
	"strings"

	"github.com/LeonardoRyuta/apillon-storage/requests"
)

// ListCollections returns all NFT collections.
func ListCollections() (string, error) {
	return requests.GetReq("/nfts/collections", nil)
}

// GetCollection returns details about a collection.
func GetCollection(uuid string) (string, error) {
	if uuid == "" {
		return "", fmt.Errorf("uuid is required")
	}

	path := "/nfts/collections/" + uuid
	return requests.GetReq(path, nil)
}

// ListTransactions lists collection transactions.
func ListTransactions(uuid string) (string, error) {
	if uuid == "" {
		return "", fmt.Errorf("uuid is required")
	}

	path := "/nfts/collections/" + uuid + "/transactions"
	return requests.GetReq(path, nil)
}

// CreateSubstrateCollection creates a Substrate collection.
func CreateSubstrateCollection(body string) (string, error) {
	if body == "" {
		return "", fmt.Errorf("request body is required")
	}

	return requests.PostReq("/nfts/collections/substrate", strings.NewReader(body))
}

// CreateEvmCollection creates an EVM collection.
func CreateEvmCollection(body string) (string, error) {
	if body == "" {
		return "", fmt.Errorf("request body is required")
	}

	return requests.PostReq("/nfts/collections/evm", strings.NewReader(body))
}

// CreateUniqueCollection creates a Unique collection.
func CreateUniqueCollection(body string) (string, error) {
	if body == "" {
		return "", fmt.Errorf("request body is required")
	}

	return requests.PostReq("/nfts/collections/unique", strings.NewReader(body))
}
