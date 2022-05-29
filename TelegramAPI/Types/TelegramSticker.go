package types

type TelegramSticker struct {
	FileID       string               `json:"file_id"`
	FileUniqueID string               `json:"file_unique_id"`
	Width        int32                `json:"width"`
	Height       int32                `json:"height"`
	IsAnimated   bool                 `json:"is_animated"`
	IsVideo      bool                 `json:"is_video"`
	Thumb        TelegramPhotoSize    `json:"thumb,omitempty"`
	Emoji        string               `json:"emoji,omitempty"`
	SetName      string               `json:"set_name,omitempty"`
	MaskPosition TelegramMaskPosition `json:"mask_position,omitempty"`
	FileSize     int32                `json:"file_size,omitempty"`
}
