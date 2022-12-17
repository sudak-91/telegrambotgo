package types

type Animation struct {
	FileID     string     `json:"file_id"`
	FileUniqID string     `json:"file_unique_id"`
	Width      int32      `json:"width"`
	Height     int32      `json:"height"`
	Duration   int32      `json:"duration"`
	Thumb      *PhotoSize `json:"thumb,omitempty"`
	FileName   string     `json:"file_name,omitempty"`
	MimeType   string     `json:"mime_type,omitempty"`
	FileSize   int32      `json:"file_size,omitempty"`
}
