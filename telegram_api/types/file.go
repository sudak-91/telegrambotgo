package types

type File struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	FileSize     int32  `json:"file_size,omitempty"`
	FilePath     string `json:"file_path,omitempty"`
}
