package storage

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/HealthTrust/MVP-TEE-Backend/apillon/requests"
)

func GetBucketContent(bucketUuid string) (string, error) {
	path := "/storage/buckets/" + bucketUuid + "/content"

	res, err := requests.GetReq(path, nil)
	if err != nil {
		log.Printf("Failed to get bucket content %s: %v", bucketUuid, err)
		return "", err
	}

	log.Printf("Bucket content retrieved successfully for bucket %s: %s", bucketUuid, res)
	return res, nil
}

func ListFilesInBucket(bucketUuid string) (ListFilesResponse, error) {
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

func GetFileDetails(bucketUuid string, fileUuid string) (FileDetails, error) {
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

func DeleteFile(bucketUuid string, fileUuid string) (string, error) {
	path := "/storage/buckets/" + bucketUuid + "/files/" + fileUuid

	res, err := requests.DeleteReq(path)
	if err != nil {
		log.Printf("Failed to delete file %s in bucket %s: %v", fileUuid, bucketUuid, err)
		return "", err
	}

	log.Printf("File %s deleted successfully from bucket %s: %s", fileUuid, bucketUuid, res)
	return res, nil
}

func GetOrGenerateIPFSLink(cid string) (string, error) {
	if cid == "" {
		log.Printf("CID is empty, cannot generate IPFS link")
		return "", fmt.Errorf("CID is empty, cannot generate IPFS link")
	}

	ipfsLink := "/storage/link-on-ipfs/:cid" + cid
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
