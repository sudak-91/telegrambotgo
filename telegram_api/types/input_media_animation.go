package types

type InputMediaAnimation struct {
	Media           string           `json:"media"`
	Thumb           string           `json:"thumb,omitempty"` //InpuFile
	Caption         string           `json:"caption,omitempty"`
	ParseMode       string           `json:"parse_mode,omitempty"`
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`
	Width           int32            `json:"width,omitempty"`
	Height          int32            `json:"height,omitempty"`
	Duration        int32            `json:"duration,omitempty"`
}
