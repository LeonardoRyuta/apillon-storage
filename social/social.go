package social

import (
	"fmt"
	"strings"

	"github.com/LeonardoRyuta/apillon-storage/requests"
)

// ListChannels lists social channels.
func ListChannels() (string, error) {
	return requests.GetReq("/social/channels", nil)
}

// GetChannel retrieves a channel by UUID.
func GetChannel(uuid string) (string, error) {
	if uuid == "" {
		return "", fmt.Errorf("uuid is required")
	}

	path := "/social/channels/" + uuid
	return requests.GetReq(path, nil)
}

// CreateChannel creates a new channel.
func CreateChannel(body string) (string, error) {
	if body == "" {
		return "", fmt.Errorf("request body is required")
	}

	return requests.PostReq("/social/channels", strings.NewReader(body))
}

// ListHubs lists social hubs.
func ListHubs() (string, error) {
	return requests.GetReq("/social/hubs", nil)
}

// GetHub gets details of a hub.
func GetHub(uuid string) (string, error) {
	if uuid == "" {
		return "", fmt.Errorf("uuid is required")
	}

	path := "/social/hubs/" + uuid
	return requests.GetReq(path, nil)
}

// CreateHub creates a new hub.
func CreateHub(body string) (string, error) {
	if body == "" {
		return "", fmt.Errorf("request body is required")
	}

	return requests.PostReq("/social/hubs", strings.NewReader(body))
}
