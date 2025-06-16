package hosting

import (
	"fmt"
	"strings"

	"github.com/LeonardoRyuta/apillon-storage/requests"
)

// ListWebsites retrieves all websites.
func ListWebsites() (string, error) {
	return requests.GetReq("/hosting/websites", nil)
}

// CreateWebsite creates a new website with the provided JSON body.
func CreateWebsite(body string) (string, error) {
	if body == "" {
		return "", fmt.Errorf("request body is required")
	}

	return requests.PostReq("/hosting/websites", strings.NewReader(body))
}

// GetWebsite returns details for a specific website.
func GetWebsite(uuid string) (string, error) {
	if uuid == "" {
		return "", fmt.Errorf("uuid is required")
	}

	path := "/hosting/websites/" + uuid
	return requests.GetReq(path, nil)
}

// StartUpload initiates an upload session for a website.
func StartUpload(uuid string, body string) (string, error) {
	if uuid == "" {
		return "", fmt.Errorf("uuid is required")
	}
	if body == "" {
		return "", fmt.Errorf("request body is required")
	}

	path := "/hosting/websites/" + uuid + "/upload"
	return requests.PostReq(path, strings.NewReader(body))
}

// EndUpload ends an upload session for a website.
func EndUpload(uuid, session string) (string, error) {
	if uuid == "" || session == "" {
		return "", fmt.Errorf("uuid and session are required")
	}

	path := "/hosting/websites/" + uuid + "/upload/" + session + "/end"
	return requests.PostReq(path, nil)
}

// DeployWebsite triggers a deployment of the website.
func DeployWebsite(uuid string, body string) (string, error) {
	if uuid == "" {
		return "", fmt.Errorf("uuid is required")
	}
	if body == "" {
		return "", fmt.Errorf("request body is required")
	}

	path := "/hosting/websites/" + uuid + "/deploy"
	return requests.PostReq(path, strings.NewReader(body))
}

// ListDeployments lists deployments for a website.
func ListDeployments(uuid string) (string, error) {
	if uuid == "" {
		return "", fmt.Errorf("uuid is required")
	}

	path := "/hosting/websites/" + uuid + "/deployments"
	return requests.GetReq(path, nil)
}

// GetDeployment returns details of a website deployment.
func GetDeployment(uuid, deployment string) (string, error) {
	if uuid == "" || deployment == "" {
		return "", fmt.Errorf("uuid and deployment are required")
	}

	path := "/hosting/websites/" + uuid + "/deployments/" + deployment
	return requests.GetReq(path, nil)
}

// CreateShortURL creates a new short URL.
func CreateShortURL(body string) (string, error) {
	if body == "" {
		return "", fmt.Errorf("request body is required")
	}

	return requests.PostReq("/hosting/short-url", strings.NewReader(body))
}
