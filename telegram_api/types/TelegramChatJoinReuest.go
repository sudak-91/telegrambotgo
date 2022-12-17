package types

type TelegramChatJoinRequest struct {
	Chat       *TelegramChat  `json:"chat"`
	From       *TelegramUser  `json:"from"`
	Data       int32          `json:"date"`
	Bio        string         `json:"bio,omitempty"`
	InviteLink ChatInviteLink `json:"invite_link,omitempty"`
}
