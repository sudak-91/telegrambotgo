package types

type ForumTopicCreated struct {
	Name              string `json:"name"`
	IconColor         int32  `json:"icon_color"`
	IconCustomEmojiId string `json:"icon_custom_emoji_id,omitempty"`
}
