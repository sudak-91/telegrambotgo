package types

type TelegramVoiceChatParticipantsInvited struct {
	Users []TelegramUser `json:"users,omitempty"`
}
