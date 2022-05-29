package types

type TelegramDocument struct {
	FileID       string            `json:"file_id"`
	FileUniqueID string            `json:"file_unique_id"`
	Thumb        TelegramPhotoSize `json:"thumb,omitempty"`
	FileName     string            `json:"file_name,omitempty"`
	MimeType     string            `json:"mime_type,omitempty"`
	Filesize     int32             `json:"file_size,omitempty"`
}
