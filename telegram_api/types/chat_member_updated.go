package types

type ChatMemberUpdated struct {
	Chat          *Chat           `json:"chat"`
	From          *User           `json:"from"`
	Date          int32           `json:"date"`
	OldChatMember *ChatMember     `json:"old_chat_member"`
	NewChatMember *ChatMember     `json:"new_chat_member"`
	InviteLink    *ChatInviteLink `json:"invite_link"`
}
