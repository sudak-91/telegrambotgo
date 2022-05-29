package types

type TelegramAudio struct {
	FileID     string            `json:"file_id"`
	FileUniqId string            `json:"file_unique_id"`
	Duration   int32             `json:"duration"`
	Performer  string            `json:"performer,omitempty"`
	Title      string            `json:"title,omitempty"`
	FileName   string            `json:"file_name,omitempty"`
	MimeType   string            `json:"mime_type,omitempty"`
	FileSize   int32             `json:"file_size,omitempty"`
	Thumb      TelegramPhotoSize `json:"thumb,omitempty"`
}
