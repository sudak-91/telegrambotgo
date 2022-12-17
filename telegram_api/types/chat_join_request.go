package types

type ChatJoinRequest struct {
	Chat       *Chat           `json:"chat"`
	From       *User           `json:"from"`
	Data       int32           `json:"date"`
	Bio        string          `json:"bio,omitempty"`
	InviteLink *ChatInviteLink `json:"invite_link,omitempty"`
}
