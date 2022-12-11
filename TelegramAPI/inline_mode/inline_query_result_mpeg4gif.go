package inline

import types "github.com/sudak-91/telegrambotgo/TelegramAPI/Types"

type InlineQueryResultMpeg4Gif struct {
	Type                string                              `json:"type"`
	ID                  string                              `json:"id"`
	Mpeg4Url            string                              `json:"mpeg4_url"`
	Mpeg4Width          int32                               `json:"mpeg4_width,omitempty"`
	Mpeg4Height         int32                               `json:"mpeg4_height,omitempty"`
	ThumbUrl            string                              `json:"thumb_url"`
	ThumbMimeType       string                              `json:"thumb_mime_url,omitempty"`
	Title               string                              `json:"title,omitempty"`
	Caption             string                              `json:"caption,omitempty"`
	ParseMode           string                              `json:"pars_mode,omitempty"`
	CaptionEntities     []*types.TelegramMessageEntity      `json:"caption_entities,omitempty"`
	ReplyMarkup         *types.TelegramInlineKeyboardMarkup `json:"reply_markup, omitempty"`
	InputMessageContent *InputMessageContent                `json:"input_message_content"`
}
