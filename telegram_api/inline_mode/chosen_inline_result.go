package inline

import types "github.com/sudak-91/telegrambotgo/telegram_api/types"

type ChosenInlineResult struct {
	ResultID        string          `json:"result_id"`
	From            *types.User     `json:"from"`
	Location        *types.Location `json:"location,omitempty"`
	InlineMessageID string          `json:"inline_message_id,omitempty"`
	Query           string          `json:"query"`
}
