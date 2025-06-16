package storage

import (
	"io"
	"net/http"
	"strings"
	"testing"

	gock "gopkg.in/h2non/gock.v1"
)

func TestStartUploadFilesToBucket_DefaultContentType(t *testing.T) {
	defer gock.Off()

	bucketUUID := "test-bucket-uuid"
	path := "/storage/buckets/" + bucketUUID + "/upload"

	var body string

	// Mock the POST request and capture the request body
	gock.New("https://api.apillon.io").
		Post(path).
		AddMatcher(func(req *http.Request, ereq *gock.Request) (bool, error) {
			b, _ := io.ReadAll(req.Body)
			body = string(b)
			return true, nil
		}).
		Reply(200).
		JSON(map[string]any{
			"data": map[string]any{
				"sessionUuid": "session",
				"files":       []map[string]any{{"url": "http://example.com"}},
			},
		})

	files := []FileMetadata{{FileName: "test.txt"}}

	_, err := StartUploadFilesToBucket(bucketUUID, files)
	if err != nil {
		t.Fatalf("StartUploadFilesToBucket returned error: %v", err)
	}

	if !gock.IsDone() {
		t.Fatalf("expected HTTP request was not made")
	}

	if !strings.Contains(body, "\"contentType\":\"text/plain\"") {
		t.Errorf("request body does not contain default content type: %s", body)
	}
}
