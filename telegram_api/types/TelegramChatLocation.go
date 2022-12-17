package types

type TelegramChatLocation struct {
	Location Location `json:"location"`
	Address  string   `json:"address"`
}
