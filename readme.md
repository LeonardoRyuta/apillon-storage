# Unofficial Apillon Storage Go SDK

A Go SDK for interacting with the [Apillon Storage API](https://apillon.io/). This SDK allows you to manage storage buckets, upload and manage files, and retrieve IPFS links programmatically. This is a very underdeveloped version so feel free to contribute!

**Originally developed by Leonardo Iara.**

---

## Features

- **Bucket Management:** Create, list, and retrieve storage buckets.
- **File Upload:** Upload single or multiple files to a bucket.
- **File Management:** List, retrieve details, and delete files.
- **Directory Management:** Delete directories from a bucket.
- **IPFS Integration:** Retrieve or generate IPFS links for files.
- **IPFS Cluster Info:** Retrieve IPFS cluster information.
- **Session Management:** Manage upload sessions for batch file uploads.

---

## Installation

```sh
go get github.com/LeonardoRyuta/apillon-storage
```

---

## Authentication

The SDK requires an Apillon API key for authentication.  
You can provide the API key in two ways:

### 1. Environment Variable (Recommended)

Set the environment variable `APILLON_API_KEY` before running your application:

**Windows (Command Prompt):**
```sh
set APILLON_API_KEY=your_api_key_here
```

**Linux/macOS:**
```sh
export APILLON_API_KEY=your_api_key_here
```

### 2. Programmatically

You can set the API key at runtime in your Go code:

```go
import "github.com/LeonardoRyuta/apillon-storage/requests"

func main() {
    requests.SetAPIKey("your_api_key_here")
    // ... your code ...
}
```

---

## Usage

### Import the SDK

```go
import (
    "github.com/LeonardoRyuta/apillon-storage/storage"
)
```

---

### Create a Bucket

```go
err := storage.CreateBucket("my-bucket", "A description for my bucket")
if err != nil {
    // handle error
}
```

---

### List Buckets

```go
buckets, err := storage.GetBucket("my-bucket") // or storage.GetBucket("") for all the buckets
if err != nil {
    // handle error
}
for _, bucket := range buckets.Data.Items {
    fmt.Println(bucket.Name, bucket.BucketUUID)
}
```

---

### Upload Files

```go
files := []storage.WholeFile{
    {
        Metadata: storage.FileMetadata{
            FileName:    "example.txt",
            ContentType: "text/plain",
        },
        Content: "Hello, Apillon!",
    },
}

bucketUUID := "your-bucket-uuid"
result, err := storage.UploadFileProcess(bucketUUID, files)
if err != nil {
    // handle error
}
fmt.Println("Upload result:", result)
```

---

### List Files in a Bucket

```go
fileList, err := storage.ListFilesInBucket(bucketUUID)
if err != nil {
    // handle error
}
for _, file := range fileList.Data.Items {
    fmt.Println(file.Name, file.FileUUID)
}
```

---

### Get File Details

```go
fileDetails, err := storage.GetFileDetails(bucketUUID, fileUUID)
if err != nil {
    // handle error
}
fmt.Printf("File details: %+v\n", fileDetails.Data)
```

---

### Delete a File

```go
_, err := storage.DeleteFile(bucketUUID, fileUUID)
if err != nil {
    // handle error
}
```

---

### Delete a Directory

```go
resp, err := storage.DeleteDirectory(bucketUUID, directoryUUID)
if err != nil {
    // handle error
}
fmt.Printf("Delete directory response: %+v\n", resp)
```

---

### Get or Generate IPFS Link

```go
ipfsLink, err := storage.GetOrGenerateIPFSLink(cid)
if err != nil {
    // handle error
}
fmt.Println("IPFS Link:", ipfsLink)
```

---

### Get IPFS Cluster Info

```go
info, err := storage.GetIPFSClusterInfo()
if err != nil {
    // handle error
}
fmt.Printf("IPFS Cluster Info: %+v\n", info.Data)
```

---

### Get Bucket Content

```go
content, err := storage.GetBucketContent(bucketUUID)
if err != nil {
    // handle error
}
fmt.Println("Bucket Content:", content)
```

---

### Advanced: Manual Upload Session Control

#### Start an Upload Session

```go
files := []storage.FileMetadata{
    {FileName: "file1.txt", ContentType: "text/plain"},
    {FileName: "file2.json", ContentType: "application/json"},
}
resp, err := storage.StartUploadFilesToBucket(bucketUUID, files)
if err != nil {
    // handle error
}
fmt.Println("Start upload session response:", resp)
```

#### Upload File Content to Signed URL

```go
result, err := storage.UploadFiles(signedURL, fileContent)
if err != nil {
    // handle error
}
fmt.Println("Upload result:", result)
```

#### End an Upload Session

```go
resp, err := storage.EndSession(bucketUUID, sessionID)
if err != nil {
    // handle error
}
fmt.Println("End session response:", resp)
```

---

## Running Tests

The SDK includes comprehensive unit tests.  
To run all tests:

```sh
go test ./...
```

---

## Contributing

Pull requests are welcome! For major changes, please open an issue first to discuss what you would like to change.

---

## License

MIT License

---

## Notes

- Ensure your API key is kept secure and **never** committed to version control.
- For more details on the Apillon API, see the [official documentation](https://apillon.io/).
