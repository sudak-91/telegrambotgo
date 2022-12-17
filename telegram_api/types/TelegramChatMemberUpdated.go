package types

type TelegramChatMemberUpdated struct {
	Chat          *TelegramChat   `json:"chat"`
	From          *TelegramUser   `json:"from"`
	Date          int32           `json:"date"`
	OldChatMember *ChatMember     `json:"old_chat_member"`
	NewChatMember *ChatMember     `json:"new_chat_member"`
	InviteLink    *ChatInviteLink `json:"invite_link,omitempty"`
}
