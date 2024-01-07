package types

type MessageEntity struct {
	Type          string `json:"type"`
	Offset        int32  `json:"offset"`
	Length        int32  `json:"length"`
	Url           string `json:"url,omitempty"`
	User          *User  `json:"user,omitempty"`
	Language      string `json:"language,omitempty"`
	CustomEmojiID string `json:"custom_emoji_id,omitempty"`
}
