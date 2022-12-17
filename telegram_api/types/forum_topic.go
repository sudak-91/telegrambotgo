package types

type ForumTopic struct {
	MessageThreadId   int32  `json:"message_thread_id"`
	Name              string `json:"name"`
	IconColor         int32  `json:"icon_color"`
	IconCustomEmojiId string `json:"icon_custom_emoji_id,omitempty"`
}
