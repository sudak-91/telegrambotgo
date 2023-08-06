package inline

import types "github.com/sudak-91/telegrambotgo/telegram_api/types"

type InlineQueryResultCachedMoeg4Gif struct {
	Type                string                     `json:"type"`
	ID                  string                     `json:"id"`
	Mpeg4FileID         string                     `json:"mpeg4_file_id"`
	Title               string                     `json:"titile,omitempty"`
	Caption             string                     `json:"caption,omitempty"`
	ParseMode           string                     `json:"parse_mode,omitempty"`
	CaptionEntities     []*types.MessageEntity     `json:"caption_entities,omitempty"`
	ReplyMarkup         types.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent       `json:"input_message_content,omitempty"`
}
