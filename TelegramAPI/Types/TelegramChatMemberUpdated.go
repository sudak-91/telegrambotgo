package types

type TelegramChatMemberUpdated struct {
	Chat          *TelegramChat           `json:"chat"`
	From          *TelegramUser           `json:"from"`
	Date          int32                   `json:"date"`
	OldChatMember *TelegramChatMember     `json:"old_chat_member"`
	NewChatMember *TelegramChatMember     `json:"new_chat_member"`
	InviteLink    *TelegramChatInviteLink `json:"invite_link,omitempty"`
}
