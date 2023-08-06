package types

type Update struct {
	UpdateId           int32                      `json:"update_id"`
	Message            *Message                   `json:"message,omitempty"`
	EditedMessage      *Message                   `json:"edited_message,omitempty"`
	ChannelPost        *Message                   `json:"channel_post,omitempty"`
	EditedChannelPost  *Message                   `json:"edited_channel_post,omitempty"`
	InlineQuery        *TelegramInlineQuery       `json:"inline_query,omitempty"`
	ChosenInlineResult *ChosenInlineResult        `json:"chosen_inline_result,omitempty"`
	CallbackQuery      *CallbackQuery             `json:"callback_query,omitempty"`
	ShippingQuery      *ShippingQuery             `json:"shipping_query,omitempty"`
	PreCheckoutQuery   *PreCheckoutQuery          `json:"pre_checkout_query,omitempty"`
	Poll               *TelegramPoll              `json:"poll,omitempty"`
	PollAnswer         *PollAnwer                 `json:"poll_answer,omitempty"`
	MyChatMember       *TelegramChatMemberUpdated `json:"my_chat_member,omitempty"`
	ChatMember         *TelegramChatMemberUpdated `json:"chat_member,omitempty"`
	ChatJoinRequest    *TelegramChatJoinRequest   `json:"chat_join_request,omitempty"`
}
