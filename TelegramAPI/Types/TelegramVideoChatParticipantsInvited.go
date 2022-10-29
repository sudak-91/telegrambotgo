package types

type TelegramVideoChatParticipantsInvited struct {
	Users []*TelegramUser `json:"users"`
}
