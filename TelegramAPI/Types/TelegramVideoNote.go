package types

type TelegramVideoNote struct {
	FileID       string            `json:"file_id"`
	FileUniqueID string            `json:"file_unique_id"`
	Length       int32             `json:"length"`
	Duration     int32             `json:"duration"`
	Thumb        TelegramPhotoSize `json:"thumb,omitempty"`
	FileSize     int32             `json:"file_size,omitempty"`
}
