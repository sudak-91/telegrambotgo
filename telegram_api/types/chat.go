package types

type Chat struct {
	ID                                 int64            `json:"id"`
	Type                               string           `json:"type"`
	Title                              string           `json:"title,omitempty"`
	UserName                           string           `json:"username,omitempty"`
	FirstName                          string           `json:"first_name,omitempty"`
	LastName                           string           `json:"last_name,omitempty"`
	IsForum                            bool             `json:"is_forum,omitempty"`
	ActiveUsernames                    []string         `json:"active_usernames,omitempty"`
	EmojiStatusCustomEmojiId           string           `json:"emoji_status_custom_emoji_id,omitempty"`
	Photo                              *ChatPhoto       `json:"photo,omitempty"`
	Bio                                string           `json:"bio,omitempty"`
	HasPrivateForwards                 bool             `json:"has_private_forwards,omitempty"`
	HasRestrictedVoiceAndVideoMessages bool             `json:"has_restricted_voice_and_video_messages"`
	JoinToSendMessages                 bool             `json:"join_to_send_messages,omitempty"`
	JoinByRequest                      bool             `json:"join_by_request,omitempty"`
	Description                        string           `json:"description,omitempty"`
	InviteLink                         string           `json:"invite_link,omitempty"`
	PinnedMessage                      *Message         `json:"pinned_message,omitempty"`
	Permissions                        *ChatPermissions `json:"permissions,omitempty"`
	SlowModeDelay                      int32            `json:"slow_mode_delay,omitempty"`
	MessageAutoDeleteTime              int32            `json:"message_auto_delete_time,omitempty"`
	HasProtectedContent                bool             `json:"has_protected_content,omitempty"`
	StickerSetName                     string           `json:"sticker_set_name,omitempty"`
	CanSetStickerSet                   bool             `json:"can_set_sticker_set,omitempty"`
	LinkedChatID                       int32            `json:"linked_chat_id,omitempty"`
	Location                           *ChatLocation    `json:"location,omitempty"`
}
