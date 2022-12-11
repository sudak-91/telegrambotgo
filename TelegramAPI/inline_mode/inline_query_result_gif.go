package inline

import types "github.com/sudak-91/telegrambotgo/TelegramAPI/Types"

type InlineQueryResultGif struct {
	Type                string                              `json:"type"`
	ID                  string                              `json:"id"`
	GifUrl              string                              `json:"gif_url"`
	GifWidth            int32                               `json:"gif_widht,omitempty"`
	GifHeight           int32                               `json:"gif_height,omitempty"`
	GifDuration         int32                               `json:"gif_duration,omitempty"`
	ThumbUrl            string                              `json:"thumb_url,omitempty"`
	ThumbMimeType       string                              `json:"thumb_mime_type,omitempty"`
	Title               string                              `json:"title,omitempty"`
	Caption             string                              `json:"caption,omitempty"`
	ParseMode           string                              `json:"parse_mode,omitempty"`
	CaptionEntities     []*types.TelegramMessageEntity      `json:"caption_entites,omitempty"`
	ReplyMarkup         *types.TelegramInlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent                `json:"input_message_content,omitempty"`
}
