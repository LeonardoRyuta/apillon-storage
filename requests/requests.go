// Package requests provides helper functions for making authenticated HTTP requests
// to the Apillon API. It supports GET, POST, and DELETE methods, and manages API key authentication.
package requests

import (
	"io"
	"net/http"
	"os"
)

var apiKey string

// SetAPIKey sets the API key to be used for authentication in all requests.
//
// If not set, the package will attempt to read the API key from the APILLON_API_KEY environment variable.
func SetAPIKey(key string) {
	apiKey = key
}

// getAPIKey retrieves the API key for authentication.
// It returns the key set by SetAPIKey, or falls back to the APILLON_API_KEY environment variable.
func getAPIKey() string {
	if apiKey != "" {
		return apiKey
	}
	return os.Getenv("APILLON_API_KEY")
}

// GetReq sends an authenticated HTTP GET request to the Apillon API.
//
// Parameters:
//   - path: The API endpoint path (e.g., "/storage/buckets").
//   - params: Optional query parameters as a map[string]string.
//
// Returns:
//   - string: The response body as a string.
//   - error: An error if the request fails or the response cannot be read.
func GetReq(path string, params map[string]string) (string, error) {
	url := "https://api.apillon.io" + path

	if len(params) > 0 {
		url += "?"
		for key, value := range params {
			url += key + "=" + value + "&"
		}
		url = url[:len(url)-1] // Remove the trailing '&'
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Basic "+getAPIKey())

	client := &http.Client{
		Timeout: 30 * 1e9, // 30 seconds
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(responseBody), nil
}

// PostReq sends an authenticated HTTP POST request to the Apillon API.
//
// Parameters:
//   - path: The API endpoint path (e.g., "/storage/buckets").
//   - body: The request body as an io.Reader (should be JSON).
//
// Returns:
//   - string: The response body as a string.
//   - error: An error if the request fails or the response cannot be read.
func PostReq(path string, body io.Reader) (string, error) {
	url := "https://api.apillon.io" + path

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic "+getAPIKey())

	client := &http.Client{
		Timeout: 60 * 1e9, // 1 minute
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(responseBody), nil
}

// DeleteReq sends an authenticated HTTP DELETE request to the Apillon API.
//
// Parameters:
//   - path: The API endpoint path (e.g., "/storage/buckets/{uuid}").
//
// Returns:
//   - string: The response body as a string.
//   - error: An error if the request fails or the response cannot be read.
func DeleteReq(path string) (string, error) {
	url := "https://api.apillon.io" + path

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Basic "+getAPIKey())

	client := &http.Client{
		Timeout: 30 * 1e9, // 30 seconds
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(responseBody), nil
}
