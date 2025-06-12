package requests

import (
	"io"
	"net/http"
	"os"
)

var apiKey string

func SetAPIKey(key string) {
	apiKey = key
}

func getAPIKey() string {
	if apiKey != "" {
		return apiKey
	}
	return os.Getenv("APILLON_API_KEY")
}

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

	client := &http.Client{}
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

func PostReq(path string, body io.Reader) (string, error) {
	url := "https://api.apillon.io" + path

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic "+getAPIKey())

	client := &http.Client{}
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

func DeleteReq(path string) (string, error) {
	url := "https://api.apillon.io" + path

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Basic "+getAPIKey())

	client := &http.Client{}
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
