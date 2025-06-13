package storage

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/LeonardoRyuta/apillon-storage/requests"
)

// CreateBucket creates a new storage bucket with the specified name and optional description.
// Sends a POST request to the storage API to create the bucket.
// Returns an error if the request fails or the API returns an error.
func CreateBucket(name string, description string) error {
	body := `{"name":"` + name + `"`
	if description != "" {
		body += `, "description":"` + description + `"`
	}
	body += `}`

	res, err := requests.PostReq("/storage/buckets", strings.NewReader(body))
	if err != nil {
		log.Printf("Failed to create bucket: %v", err)
		return err
	}

	log.Printf("Bucket created successfully: %s", res)
	return nil
}

// GetBucket retrieves information about storage buckets, optionally filtered by name.
// Sends a GET request to the storage API with the provided name as a query parameter.
// Returns a ListBucketsResponse containing the bucket(s) information, or an error if the request or unmarshalling fails.
func GetBucket(name string) (ListBucketsResponse, error) {

	params := map[string]string{}

	if name != "" {
		params["name"] = name
	}

	res, err := requests.GetReq("/storage/buckets/", params)
	if err != nil {
		log.Printf("Failed to get bucket: %v", err)
		return ListBucketsResponse{}, err
	}

	var bucketList ListBucketsResponse
	if errUnmarshal := json.Unmarshal([]byte(res), &bucketList); errUnmarshal != nil {
		log.Printf("Failed to unmarshal JSON response: %v. Raw response: %s", errUnmarshal, res)
		return ListBucketsResponse{}, fmt.Errorf("failed to unmarshal JSON response: %w. Raw response: %s", errUnmarshal, res)
	}

	log.Printf("Bucket details: %s", res)
	return bucketList, nil
}
