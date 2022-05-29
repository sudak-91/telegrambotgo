package types

type TelegramGame struct {
	Title        string                  `json:"title"`
	Description  string                  `json:"description"`
	Photo        []TelegramPhotoSize     `json:"photo"`
	Text         string                  `json:"text,omitempty"`
	TextEntities []TelegramMessageEntity `json:"text_entities,omitempty"`
	Animation    TelegramAnimation       `json:"animation,omitempty"`
}
