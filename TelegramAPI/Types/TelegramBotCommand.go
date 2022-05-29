package types

type TelegramBotCommand struct {
	Command     string `json:"command"`
	Description string `json:"description"`
}
