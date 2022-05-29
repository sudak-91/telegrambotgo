package types

type TelegramChatLocation struct {
	Location TelegramLocation `json:"location"`
	Address  string           `json:"address"`
}
