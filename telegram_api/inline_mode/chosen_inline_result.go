package inline

import types "github.com/sudak-91/telegrambotgo/TelegramAPI/Types"

type ChosenInlineResult struct {
	ResultID        string                  `json:"result_id"`
	From            *types.TelegramUser     `json:"from"`
	Location        *types.TelegramLocation `json:"location,omitempty"`
	InlineMessageID string                  `json:"inline_message_id,omitempty"`
	Query           string                  `json:"query"`
}
