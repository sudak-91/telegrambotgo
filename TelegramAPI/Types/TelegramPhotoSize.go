package types

//@TODO
type TelegramPhotoSize struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	Width        int32  `json:"width"`
	Height       int32  `json:"height"`
	FileSize     int32  `json:"file_size"`
}
