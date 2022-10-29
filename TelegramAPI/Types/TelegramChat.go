package types

type TelegramChat struct {
	ID                                 int64                   `json:"id"`
	Type                               string                  `json:"type"`
	Title                              string                  `json:"title,omitempty"`
	UserName                           string                  `json:"username,omitempty"`
	FirstName                          string                  `json:"first_name,omitempty"`
	LastName                           string                  `json:"last_name,omitempty"`
	Photo                              TelegramChatPhoto       `json:"photo,omitempty"`
	Bio                                string                  `json:"bio,omitempty"`
	HasPrivateForwards                 bool                    `json:"has_private_forwards,omitempty"`
	HasRestrictedVoiceAndVideoMessages bool                    `json:"has_restricted_voice_and_video_messages"`
	JoinToSendMessages                 bool                    `json:"join_to_send_messages,omitempty"`
	JoinByRequest                      bool                    `json:"join_by_request,omitempty"`
	Description                        string                  `json:"description,omitempty"`
	InviteLink                         string                  `json:"invite_link,omitempty"`
	PinnedMessage                      *TelegramMessage        `json:"pinned_message,omitempty"`
	Permissions                        TelegramChatPermissions `json:"permissions,omitempty"`
	SlowModeDelay                      int32                   `json:"slow_mode_delay,omitempty"`
	MessageAutoDeleteTime              int32                   `json:"message_auto_delete_time,omitempty"`
	HasProtectedContent                bool                    `json:"has_protected_content,omitempty"`
	StickerSetName                     string                  `json:"sticker_set_name,omitempty"`
	CanSetStickerSet                   bool                    `json:"can_set_sticker_set,omitempty"`
	LinkedChatID                       int32                   `json:"linked_chat_id,omitempty"`
	Location                           TelegramChatLocation    `json:"location,omitempty"`
}
