package types

type TelegramGame struct {
	Title        string            `json:"title"`
	Description  string            `json:"description"`
	Photo        []PhotoSize       `json:"photo"`
	Text         string            `json:"text,omitempty"`
	TextEntities []MessageEntity   `json:"text_entities,omitempty"`
	Animation    TelegramAnimation `json:"animation,omitempty"`
}
