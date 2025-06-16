package computing

import (
	"fmt"
	"strings"

	"github.com/LeonardoRyuta/apillon-storage/requests"
)

// CreateContract creates a new computing contract.
func CreateContract(body string) (string, error) {
	if body == "" {
		return "", fmt.Errorf("request body is required")
	}

	return requests.PostReq("/computing/contracts", strings.NewReader(body))
}

// ListContracts lists computing contracts.
func ListContracts() (string, error) {
	return requests.GetReq("/computing/contracts", nil)
}

// GetContract returns details of a contract.
func GetContract(uuid string) (string, error) {
	if uuid == "" {
		return "", fmt.Errorf("uuid is required")
	}

	path := "/computing/contracts/" + uuid
	return requests.GetReq(path, nil)
}

// ListTransactions lists contract transactions.
func ListTransactions(uuid string) (string, error) {
	if uuid == "" {
		return "", fmt.Errorf("uuid is required")
	}

	path := "/computing/contracts/" + uuid + "/transactions"
	return requests.GetReq(path, nil)
}

// TransferOwnership transfers contract ownership.
func TransferOwnership(uuid string, body string) (string, error) {
	if uuid == "" {
		return "", fmt.Errorf("uuid is required")
	}
	if body == "" {
		return "", fmt.Errorf("request body is required")
	}

	path := "/computing/contracts/" + uuid + "/transfer-ownership"
	return requests.PostReq(path, strings.NewReader(body))
}

// Encrypt encrypts contract data.
func Encrypt(uuid string, body string) (string, error) {
	if uuid == "" {
		return "", fmt.Errorf("uuid is required")
	}
	if body == "" {
		return "", fmt.Errorf("request body is required")
	}

	path := "/computing/contracts/" + uuid + "/encrypt"
	return requests.PostReq(path, strings.NewReader(body))
}

// AssignCIDToNFT assigns a CID to an NFT.
func AssignCIDToNFT(uuid string, body string) (string, error) {
	if uuid == "" {
		return "", fmt.Errorf("uuid is required")
	}
	if body == "" {
		return "", fmt.Errorf("request body is required")
	}

	path := "/computing/contracts/" + uuid + "/assign-cid-to-nft"
	return requests.PostReq(path, strings.NewReader(body))
}
