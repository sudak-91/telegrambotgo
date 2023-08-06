package inline

import types "github.com/sudak-91/telegrambotgo/telegram_api/types"

type InlineQuery struct {
	ID       string          `json:"id"`
	From     *types.User     `json:"from"`
	Query    string          `json:"query"`
	Offset   string          `json:"offset"`
	ChatType string          `json:"chat_type,omitempty"`
	Location *types.Location `json:"location"`
}
