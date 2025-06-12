package storage

type FileMetadata struct {
	FileName    string `json:"fileName" validate:"required"`
	ContentType string `json:"contentType"`
}

type WholeFile struct {
	Content  string       `json:"content"`
	Metadata FileMetadata `json:"metadata"`
}

type FileItem struct {
	Path        *string `json:"path"` // Use pointer for nullable fields like path
	FileName    string  `json:"fileName"`
	ContentType string  `json:"contentType"`
	URL         string  `json:"url"`
	FileUUID    string  `json:"fileUuid"`
}

// ProcessData represents the data object in the ProcessAPIResponse.
type ProcessData struct {
	SessionUUID string     `json:"sessionUuid"`
	Files       []FileItem `json:"files"`
}

// Generic API response wrapper
type APIResponse[T any] struct {
	ID     string `json:"id"`
	Status int    `json:"status"`
	Data   T      `json:"data"`
}

type ListBucketsResponse = APIResponse[BucketListData]
type ListFilesResponse = APIResponse[FileListData]
type ProcessAPIResponse = APIResponse[ProcessData]
type FileDetails = APIResponse[FileInfo]
type IPFSLinkResponse = APIResponse[struct {
	Link string `json:"link"`
}]

// Generic list structure
type ListData[T any] struct {
	Items []T `json:"items"`
	Total int `json:"total"`
}

// Replace specific list data types
type BucketListData = ListData[BucketItem]
type FileListData = ListData[FileInfo]

// Common timestamp fields
type Timestamps struct {
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
}

type FileInfo struct {
	Timestamps
	FileUUID      string  `json:"fileUuid"`
	CID           string  `json:"CID"`
	Name          string  `json:"name"`
	ContentType   string  `json:"contentType"`
	Path          *string `json:"path"`
	Size          int64   `json:"size"`
	FileStatus    int     `json:"fileStatus"`
	Link          string  `json:"link"`
	DirectoryUUID *string `json:"directoryUuid,omitempty"`
}

type BucketItem struct {
	Timestamps
	BucketUUID  string `json:"bucketUuid"`
	BucketType  int    `json:"bucketType"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Size        int64  `json:"size"`
}
