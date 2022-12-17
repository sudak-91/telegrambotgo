package types

type InputMediaAudio struct {
	Media           string           `json:"media"`
	Thumb           string           `json:"thumb,omitempty"` //InputFile
	Caption         string           `json:"caption,omitempty"`
	ParseMode       string           `json:"parse_mode,omitempty"`
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`
	Duration        int32            `json:"duration,omitempty"`
	Performer       string           `json:"performer,omitempty"`
	Title           string           `json:"title,omitempty"`
}
