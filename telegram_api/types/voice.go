package types

type Voice struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	Duration     int32  `json:"duration"`
	MimeType     string `json:"mime_type,omitempty"`
	FileSize     int32  `json:"file_size,omitempty"`
}
