package types

type TelegramUpdate struct {
	UpdateId           int32                       `json:"update_id"`
	Message            *TelegramMessage            `json:"message,omitempty"`
	EditedMessage      *TelegramMessage            `json:"edited_message,omitempty"`
	ChannelPost        *TelegramMessage            `json:"channel_post,omitempty"`
	EditedChannelPost  *TelegramMessage            `json:"edited_channel_post,omitempty"`
	InlineQuery        *TelegramInlineQuery        `json:"inline_query,omitempty"`
	ChosenInlineResult *TelegramChosenInlineResult `json:"chosen_inline_result,omitempty"`
	CallbackQuery      *TelegramCallbackQuery      `json:"callback_query,omitempty"`
	ShippingQuery      *TelegramShippingQuery      `json:"shipping_query,omitempty"`
	PreCheckoutQuery   *TelegramPreCheckoutQuery   `json:"pre_checkout_query,omitempty"`
	Poll               *TelegramPoll               `json:"poll,omitempty"`
	PollAnswer         *TelegramPollAnwer          `json:"poll_answer,omitempty"`
	MyChatMember       *TelegramChatMemberUpdated  `json:"my_chat_member,omitempty"`
	ChatMember         *TelegramChatMemberUpdated  `json:"chat_member,omitempty"`
	ChatJoinRequest    *TelegramChatJoinRequest    `json:"chat_join_request,omitempty"`
}
