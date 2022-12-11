package inline

import types "github.com/sudak-91/telegrambotgo/TelegramAPI/Types"

type InlineQuery struct {
	ID       string                  `json:"id"`
	From     *types.TelegramUser     `json:"from"`
	Query    string                  `json:"query"`
	Offset   string                  `json:"offset"`
	ChatType string                  `json:"chat_type,omitempty"`
	Location *types.TelegramLocation `json:"location"`
}
