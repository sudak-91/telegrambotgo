package types

//@TODO Generic for ReplyMarkUp

type SendMessage struct {
	Method                   string                `json:"method"`
	ChatID                   int64                 `json:"chat_id"`
	Text                     string                `json:"text"`
	ParseMode                string                `json:"parse_mode,omitempty"`
	Entities                 []MessageEntity       `json:"entities,omitempty"`
	DisableWebPagePreview    bool                  `json:"disable_web_page_preview,omitempty"`
	DisableNotification      bool                  `json:"disable_notification,omitempty"`
	ProtectContent           bool                  `json:"protect_contetn,omitempty"`
	ReplyToMessageID         int32                 `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                  `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              *InlineKeyboardMarkup `json:"reply_markup,omitempty"` //InlineKeybordMarkup, ReplyKeyboardMarkup,ReplyKeyboardRemove,ForceReplay
}

type SendMessageInlineKeyboardMarkup struct {
	Method                   string                `json:"method"`
	ChatID                   int64                 `json:"chat_id"`
	Text                     string                `json:"text"`
	ParseMode                string                `json:"parse_mode,omitempty"`
	Entities                 []MessageEntity       `json:"entities,omitempty"`
	DisableWebPagePreview    bool                  `json:"disable_web_page_preview,omitempty"`
	DisableNotification      bool                  `json:"disable_notification,omitempty"`
	ProtectContent           bool                  `json:"protect_contetn,omitempty"`
	ReplyToMessageID         int32                 `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                  `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              *InlineKeyboardMarkup `json:"reply_markup,omitempty"` //InlineKeybordMarkup, ReplyKeyboardMarkup,ReplyKeyboardRemove,ForceReplay
}
