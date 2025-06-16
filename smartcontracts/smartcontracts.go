package smartcontracts

import (
	"fmt"
	"strings"

	"github.com/LeonardoRyuta/apillon-storage/requests"
)

// ListContracts lists available contracts.
func ListContracts() (string, error) {
	return requests.GetReq("/contracts", nil)
}

// GetContract retrieves a contract by UUID.
func GetContract(uuid string) (string, error) {
	if uuid == "" {
		return "", fmt.Errorf("uuid is required")
	}

	path := "/contracts/" + uuid
	return requests.GetReq(path, nil)
}

// GetContractABI retrieves contract ABI.
func GetContractABI(uuid string) (string, error) {
	if uuid == "" {
		return "", fmt.Errorf("uuid is required")
	}

	path := "/contracts/" + uuid + "/abi"
	return requests.GetReq(path, nil)
}

// DeployContract deploys a contract.
func DeployContract(uuid, body string) (string, error) {
	if uuid == "" {
		return "", fmt.Errorf("uuid is required")
	}
	if body == "" {
		return "", fmt.Errorf("request body is required")
	}

	path := "/contracts/" + uuid + "/deploy"
	return requests.PostReq(path, strings.NewReader(body))
}

// GetDeployedContract retrieves deployed contract details.
func GetDeployedContract(uuid string) (string, error) {
	if uuid == "" {
		return "", fmt.Errorf("uuid is required")
	}

	path := "/contracts/deployed/" + uuid
	return requests.GetReq(path, nil)
}

// ListDeployedContracts lists deployed contracts.
func ListDeployedContracts() (string, error) {
	return requests.GetReq("/contracts/deployed", nil)
}

// CallDeployedContract executes a call on a deployed contract.
func CallDeployedContract(uuid string, body string) (string, error) {
	if uuid == "" {
		return "", fmt.Errorf("uuid is required")
	}
	if body == "" {
		return "", fmt.Errorf("request body is required")
	}

	path := "/contracts/deployed/" + uuid + "/call"
	return requests.PostReq(path, strings.NewReader(body))
}

// GetDeployedABI retrieves ABI of a deployed contract.
func GetDeployedABI(uuid string) (string, error) {
	if uuid == "" {
		return "", fmt.Errorf("uuid is required")
	}

	path := "/contracts/deployed/" + uuid + "/abi"
	return requests.GetReq(path, nil)
}

// DeleteDeployedContract deletes a deployed contract.
func DeleteDeployedContract(uuid string) (string, error) {
	if uuid == "" {
		return "", fmt.Errorf("uuid is required")
	}

	path := "/contracts/deployed/" + uuid
	return requests.DeleteReq(path)
}

// ListTransactions lists transactions for deployed contract.
func ListTransactions(uuid string) (string, error) {
	if uuid == "" {
		return "", fmt.Errorf("uuid is required")
	}

	path := "/contracts/deployed/" + uuid + "/transactions"
	return requests.GetReq(path, nil)
}
