package types

type InputMediaPhoto struct {
	Media           string           `json:"media"`
	Caption         string           `json:"caption,omitempty"`
	ParseMode       string           `json:"parse_mode,omitempty"`
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`
}
