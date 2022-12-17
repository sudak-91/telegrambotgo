package types

type InputMediaDocument struct {
	Media                       string           `json:"media"`
	Thumb                       string           `json:"thumb,omitempty"` //TODO: Inputfile
	Caption                     string           `json:"caption,omitempty"`
	ParseMode                   string           `json:"parse_mode,omitempty"`
	CaptionEntities             []*MessageEntity `json:"caption_entities,omitempty"`
	DisableContentTypeDetection bool             `json:"disable_content_type_detection,omitempty"`
}
