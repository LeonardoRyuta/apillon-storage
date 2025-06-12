// Package storage provides functions to handle file uploads to Apillon storage buckets.
// It manages the upload session lifecycle, including starting uploads, uploading files via signed URLs, and ending sessions.
package storage

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/LeonardoRyuta/apillon-storage/requests"
)

// StartUploadFilesToBucket initiates an upload session for a set of files in a given bucket.
// It sends file metadata to the Apillon API and returns the raw API response or an error.
func StartUploadFilesToBucket(bucketUuid string, files []FileMetadata) (string, error) {
	var filesJsonArrayElements []string
	for _, file := range files {
		contentType := file.ContentType
		if contentType == "" {
			contentType = "text/plain"
		}
		// Prepare JSON element for each file's metadata
		element := `{"fileName":"` + file.FileName + `", "contentType":"` + file.ContentType + `"}` // Note: contentType variable is not used here
		filesJsonArrayElements = append(filesJsonArrayElements, element)
	}

	filesJsonArray := strings.Join(filesJsonArrayElements, ",")
	bodyString := `{"files":[` + filesJsonArray + `]}`

	path := "/storage/buckets/" + bucketUuid + "/upload"

	res, err := requests.PostReq(path, strings.NewReader(bodyString))
	if err != nil {
		log.Printf("Failed to upload files to bucket %s via /upload endpoint: %v", bucketUuid, err)
		return "", err
	}

	log.Printf("Files uploaded successfully to bucket %s via /upload endpoint: %s", bucketUuid, res)
	return res, nil
}

// UploadFiles uploads a file's raw content to a signed URL using HTTP PUT.
// Returns a success message or an error if the upload fails.
func UploadFiles(signedURL string, rawFile string) (string, error) {
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPut, signedURL, strings.NewReader(rawFile))
	if err != nil {
		log.Printf("Failed to create request for signed URL %s: %v", signedURL, err)
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Failed to upload file to signed URL %s: %v", signedURL, err)
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		log.Printf("Failed to upload file to signed URL %s, status code: %d, response: %s", signedURL, resp.StatusCode, string(bodyBytes))
		return "", fmt.Errorf("upload failed with status code %d: %s", resp.StatusCode, string(bodyBytes))
	}

	log.Printf("Successfully uploaded file %d to signed URL: %s", len(rawFile), signedURL)
	log.Printf("%d files uploaded successfully to bucket via signed URL: %s", len(rawFile), signedURL)
	return "upload-success", nil
}

// EndSession finalizes an upload session for a given bucket and session ID.
// Returns the API response or an error.
func EndSession(bucketUuid string, sessionId string) (string, error) {
	path := "/storage/buckets/" + bucketUuid + "/upload/" + sessionId + "/end"

	res, err := requests.PostReq(path, nil)
	if err != nil {
		log.Printf("Failed to end session for bucket %s: %v", bucketUuid, err)
		return "", err
	}

	log.Printf("Session ended successfully for bucket %s: %s", bucketUuid, res)
	return res, nil
}

// UploadFileProcess orchestrates the full upload process for multiple files:
// 1. Starts an upload session and retrieves signed URLs.
// 2. Uploads each file to its corresponding signed URL.
// 3. Ends the upload session.
// Returns the final API response or an error.
func UploadFileProcess(bucketUuid string, files []WholeFile) (string, error) {
	// Extract only the metadata for the upload session initiation
	onlyMetadata := make([]FileMetadata, len(files))
	for i, file := range files {
		if file.Content == "" || file.Metadata.FileName == "" {
			log.Printf("File content or metadata is empty for file %s in bucket %s", file.Metadata.FileName, bucketUuid)
			return "", fmt.Errorf("file content or metadata is empty for file %s in bucket %s", file.Metadata.FileName, bucketUuid)
		}
		onlyMetadata[i] = file.Metadata
	}

	// Step 1: Start upload session and get signed URLs
	res, err := StartUploadFilesToBucket(bucketUuid, onlyMetadata)
	if err != nil {
		log.Printf("Failed to start upload files for bucket %s: %v", bucketUuid, err)
		return "", fmt.Errorf("failed to start upload files for bucket %s: %w", bucketUuid, err)
	}

	var apiResp ProcessAPIResponse
	if errUnmarshal := json.Unmarshal([]byte(res), &apiResp); errUnmarshal != nil {
		log.Printf("Failed to unmarshal JSON response from process upload for bucket %s: %v. Raw response: %s", bucketUuid, errUnmarshal, res)
		return "", fmt.Errorf("failed to unmarshal process upload response: %w. Raw response: %s", errUnmarshal, res)
	}

	// Extract signed URLs from API response
	var urls []string
	if apiResp.Data.Files != nil {
		for _, fileItem := range apiResp.Data.Files {
			if fileItem.URL != "" {
				urls = append(urls, fileItem.URL)
			}
		}
	}
	if len(urls) == 0 {
		log.Printf("No URLs found in process upload response for bucket %s", bucketUuid)
		return "", fmt.Errorf("no URLs found in process upload response for bucket %s", bucketUuid)
	}

	log.Printf("Extracted URLs from process upload response for bucket %s: %v", bucketUuid, urls)

	time.Sleep(2 * time.Second) // Wait for the URLs to be ready

	// Step 2: Upload each file to its signed URL
	for i, file := range files {
		if i >= len(urls) {
			log.Printf("Not enough URLs provided for the number of files. Expected %d URLs, got %d", len(files), len(urls))
			return "", fmt.Errorf("not enough URLs provided for the number of files. Expected %d URLs, got %d", len(files), len(urls))
		}
		signedURL := urls[i]
		rawFile := file.Content
		if rawFile == "" {
			log.Printf("File content is empty for file %s in bucket %s", file.Metadata.FileName, bucketUuid)
			return "", fmt.Errorf("file content is empty for file %s in bucket %s", file.Metadata.FileName, bucketUuid)
		}

		uploadRes, err := UploadFiles(signedURL, rawFile)
		if err != nil {
			log.Printf("Failed to upload file %s to signed URL %s for bucket %s: %v", file.Metadata.FileName, signedURL, bucketUuid, err)
			return "", fmt.Errorf("failed to upload file %s to signed URL %s for bucket %s: %w", file.Metadata.FileName, signedURL, bucketUuid, err)
		}
		log.Printf("File %s uploaded successfully to signed URL %s for bucket %s: %s", file.Metadata.FileName, signedURL, bucketUuid, uploadRes)
	}

	// Step 3: End the upload session
	res, err = EndSession(bucketUuid, apiResp.Data.SessionUUID)
	if err != nil {
		log.Printf("Failed to end session for bucket %s: %v", bucketUuid, err)
		return "", fmt.Errorf("failed to end session for bucket %s: %w", bucketUuid, err)
	}

	log.Printf("Files processed successfully for bucket %s: %s", bucketUuid, res)
	return res, nil
}
