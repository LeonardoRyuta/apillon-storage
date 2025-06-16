package storage

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/LeonardoRyuta/apillon-storage/requests"
)

// GetBucketContent retrieves the raw content of a storage bucket by its UUID.
// Returns the raw response as a string, or an error if the request fails.
func GetBucketContent(bucketUuid string) (string, error) {
	if bucketUuid == "" {
		return "", fmt.Errorf("bucket uuid is required")
	}

	path := "/storage/buckets/" + bucketUuid + "/content"

	res, err := requests.GetReq(path, nil)
	if err != nil {
		log.Printf("Failed to get bucket content %s: %v", bucketUuid, err)
		return "", err
	}

	log.Printf("Bucket content retrieved successfully for bucket %s: %s", bucketUuid, res)
	return res, nil
}

// ListFilesInBucket lists all files in a given bucket by its UUID.
// Returns a ListFilesResponse struct or an error if the request or unmarshalling fails.
func ListFilesInBucket(bucketUuid string) (ListFilesResponse, error) {
	if bucketUuid == "" {
		return ListFilesResponse{}, fmt.Errorf("bucket uuid is required")
	}

	path := "/storage/buckets/" + bucketUuid + "/files"
	res, err := requests.GetReq(path, nil)
	if err != nil {
		log.Printf("Failed to list files in bucket %s: %v", bucketUuid, err)
		return ListFilesResponse{}, err
	}

	var fileList ListFilesResponse
	if errUnmarshal := json.Unmarshal([]byte(res), &fileList); errUnmarshal != nil {
		log.Printf("Failed to unmarshal JSON response from list files in bucket %s: %v. Raw response: %s", bucketUuid, errUnmarshal, res)
		return ListFilesResponse{}, fmt.Errorf("failed to unmarshal list files response: %w. Raw response: %s", errUnmarshal, res)
	}

	log.Printf("Files in bucket %s: %s", bucketUuid, res)
	return fileList, nil
}

// GetFileDetails retrieves details for a specific file in a bucket using their UUIDs.
// Returns a FileDetails struct or an error if the request or unmarshalling fails.
func GetFileDetails(bucketUuid string, fileUuid string) (FileDetails, error) {
	if bucketUuid == "" || fileUuid == "" {
		return FileDetails{}, fmt.Errorf("bucket uuid and file uuid are required")
	}

	path := "/storage/buckets/" + bucketUuid + "/files/" + fileUuid
	res, err := requests.GetReq(path, nil)
	if err != nil {
		log.Printf("Failed to get file details for file %s in bucket %s: %v", fileUuid, bucketUuid, err)
		return FileDetails{}, err
	}

	var fileDetails FileDetails
	if errUnmarshal := json.Unmarshal([]byte(res), &fileDetails); errUnmarshal != nil {
		log.Printf("Failed to unmarshal JSON response from get file details for file %s in bucket %s: %v. Raw response: %s", fileUuid, bucketUuid, errUnmarshal, res)
		return FileDetails{}, fmt.Errorf("failed to unmarshal get file details response: %w. Raw response: %s", errUnmarshal, res)
	}

	log.Printf("File details for file %s in bucket %s: %s", fileUuid, bucketUuid, res)
	return fileDetails, nil
}

// DeleteFile deletes a specific file from a bucket using their UUIDs.
// Returns the raw response as a string, or an error if the request fails.
func DeleteFile(bucketUuid string, fileUuid string) (string, error) {
	if bucketUuid == "" || fileUuid == "" {
		return "", fmt.Errorf("bucket uuid and file uuid are required")
	}

	path := "/storage/buckets/" + bucketUuid + "/files/" + fileUuid

	res, err := requests.DeleteReq(path)
	if err != nil {
		log.Printf("Failed to delete file %s in bucket %s: %v", fileUuid, bucketUuid, err)
		return "", err
	}

	log.Printf("File %s deleted successfully from bucket %s: %s", fileUuid, bucketUuid, res)
	return res, nil
}

// DeleteDirectory deletes a directory from a bucket using their UUIDs.
// Returns a DeleteDirectoryResponse struct or an error if the request or unmarshalling fails.
// Handles known error codes for non-existent or already deleted directories.
func DeleteDirectory(bucketUuid string, directoryUuid string) (DeleteDirectoryResponse, error) {
	if bucketUuid == "" || directoryUuid == "" {
		return DeleteDirectoryResponse{}, fmt.Errorf("bucket uuid and directory uuid are required")
	}

	path := "/storage/buckets/" + bucketUuid + "/directories/" + directoryUuid

	res, err := requests.DeleteReq(path)
	if err != nil {
		log.Printf("Failed to delete directory %s in bucket %s: %v", directoryUuid, bucketUuid, err)
		return DeleteDirectoryResponse{}, err
	}

	var resp DeleteDirectoryResponse
	if errUnmarshal := json.Unmarshal([]byte(res), &resp); errUnmarshal != nil {
		log.Printf("Failed to unmarshal JSON response from delete directory %s in bucket %s: %v. Raw response: %s", directoryUuid, bucketUuid, errUnmarshal, res)
		return DeleteDirectoryResponse{}, fmt.Errorf("failed to unmarshal delete directory response: %w. Raw response: %s", errUnmarshal, res)
	}

	// Handle known error codes in the response if needed
	if resp.Status == 40406003 {
		return resp, fmt.Errorf("directory does not exist (error 40406003)")
	}
	if resp.Status == 40006007 {
		return resp, fmt.Errorf("directory is already marked for deletion (error 40006007)")
	}

	log.Printf("Directory %s deleted successfully from bucket %s: %+v", directoryUuid, bucketUuid, resp)
	return resp, nil
}

// GetOrGenerateIPFSLink retrieves or generates an IPFS link for a given CID.
// Returns the IPFS link as a string, or an error if the request or unmarshalling fails.
func GetOrGenerateIPFSLink(cid string) (string, error) {
	if cid == "" {
		log.Printf("CID is empty, cannot generate IPFS link")
		return "", fmt.Errorf("CID is empty, cannot generate IPFS link")
	}

	// Construct the correct API path. The previous implementation
	// mistakenly kept the ":cid" placeholder in the final URL which
	// resulted in requests like "/storage/link-on-ipfs/:cidQm...".
	// The API expects the CID directly appended without the colon.
	ipfsLink := "/storage/link-on-ipfs/" + cid
	res, err := requests.GetReq(ipfsLink, nil)
	if err != nil {
		log.Printf("Failed to get IPFS link for CID %s: %v", cid, err)
		return "", err
	}

	var ipfsLinkResponse IPFSLinkResponse
	if errUnmarshal := json.Unmarshal([]byte(res), &ipfsLinkResponse); errUnmarshal != nil {
		log.Printf("Failed to unmarshal JSON response from get IPFS link for CID %s: %v. Raw response: %s", cid, errUnmarshal, res)
		return "", fmt.Errorf("failed to unmarshal get IPFS link response: %w. Raw response: %s", errUnmarshal, res)
	}

	if ipfsLinkResponse.Data.Link == "" {
		log.Printf("No IPFS link found for CID %s", cid)
		return "", fmt.Errorf("no IPFS link found for CID %s", cid)
	}

	log.Printf("IPFS link for CID %s: %s", cid, ipfsLinkResponse.Data.Link)
	return ipfsLinkResponse.Data.Link, nil
}

// GetIPFSClusterInfo retrieves information about the IPFS cluster.
// Returns an IPFSClusterInfoResponse struct or an error if the request or unmarshalling fails.
func GetIPFSClusterInfo() (IPFSClusterInfoResponse, error) {
	path := "/storage/ipfs-cluster-info"

	res, err := requests.GetReq(path, nil)
	if err != nil {
		log.Printf("Failed to get IPFS cluster info: %v", err)
		return IPFSClusterInfoResponse{}, err
	}

	var infoResp IPFSClusterInfoResponse
	if errUnmarshal := json.Unmarshal([]byte(res), &infoResp); errUnmarshal != nil {
		log.Printf("Failed to unmarshal JSON response from get IPFS cluster info: %v. Raw response: %s", errUnmarshal, res)
		return IPFSClusterInfoResponse{}, fmt.Errorf("failed to unmarshal IPFS cluster info response: %w. Raw response: %s", errUnmarshal, res)
	}

	log.Printf("IPFS cluster info retrieved: %+v", infoResp)
	return infoResp, nil
}
