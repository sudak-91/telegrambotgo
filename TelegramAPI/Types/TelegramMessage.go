package types

type TelegramMessage struct {
	MessageID                    int32                                 `json:"message_id,omitempty"`
	From                         *TelegramUser                         `json:"from,omitempty"`
	SenderChat                   *TelegramChat                         `json:"sender_chat,omitempty"`
	Date                         int32                                 `json:"date"`
	Chat                         *TelegramChat                         `json:"chat,omitempty"`
	ForwardFrom                  *TelegramUser                         `json:"forward_from,omitempty"`
	ForwardFromChat              *TelegramChat                         `json:"forward_from_chat,omitempty"`
	ForwardFromMessageId         int32                                 `json:"forward_from_message_id,omitempty"`
	ForwardSignature             string                                `json:"forward_signature,omitempty"`
	ForwardSenderName            string                                `json:"forward_sender_name,omitempty"`
	ForwardDate                  int32                                 `json:"forward_date,omitempty"`
	IsAutomaticForward           bool                                  `json:"is_automatic_forward,omitempty"`
	ReplyToMessage               *TelegramMessage                      `json:"reply_to_message,omitempty"`
	ViaBot                       *TelegramUser                         `json:"via_bot,omitempty"`
	EditDate                     int32                                 `json:"edit_date,omitempty"`
	HasProtectedContent          bool                                  `json:"has_protected_content,omitempty"`
	MediaGroupId                 string                                `json:"media_group_id,omitempty"`
	AuthorSignature              string                                `json:"author_signature,omitempty"`
	Text                         string                                `json:"text,omitempty"`
	Entities                     []TelegramMessageEntity               `json:"entities,omitempty"`
	Animation                    *TelegramAnimation                    `json:"animation,omitempty"`
	Audio                        *TelegramAudio                        `json:"audio,omitempty"`
	Document                     *TelegramDocument                     `json:"document,omitempty"`
	Photo                        []TelegramPhotoSize                   `json:"photo,omitempty"`
	Sticker                      *TelegramSticker                      `json:"sticker,omitempty"`
	Video                        *TelegramVideo                        `json:"video,omitempty"`
	VideoNote                    *TelegramVideoNote                    `json:"video_note,omitempty"`
	Voice                        *TelegramVoice                        `json:"voice,omitempty"`
	Caption                      string                                `json:"caption,omitempty"`
	CaptionEntities              []TelegramMessageEntity               `json:"caption_entities,omitempty"`
	Contact                      *TelegramContact                      `json:"contact,omitempty"`
	Dice                         *TelegramDice                         `json:"dice,omitempty"`
	Game                         *TelegramGame                         `json:"game,omitempty"`
	Poll                         *TelegramPoll                         `json:"poll,omitempty"`
	Venue                        *TelegramVenue                        `json:"venue,omitempty"`
	Location                     *TelegramLocation                     `json:"location,omitempty"`
	NewChatMember                *[]TelegramUser                       `json:"new_chat_members,omitempty"`
	LeftChatMember               *TelegramUser                         `json:"left_chat_member,omitempty"`
	NewChatTitle                 string                                `json:"new_chat_title,omitempty"`
	NewChatPhoto                 []TelegramPhotoSize                   `json:"new_chat_photo,omitempty"`
	DeletChatPhoto               bool                                  `json:"delete_chat_photo,omitempty"`
	GroupChatCreated             bool                                  `json:"group_chat_created,omitempty"`
	SuperGroupChatCreated        bool                                  `json:"supergroup_chat_created,omitempty"`
	ChannelChatCreated           bool                                  `json:"channel_chat_created,omitempty"`
	MsgAutoDeletTimerChanged     *TelegramMsgAutoDelTimerChanged       `json:"message_auto_delete_timer_changed,omitempty"`
	MigrateToChatId              int32                                 `json:"migrate_to_chat_id,omitempty"`
	MigrateFromChatId            int32                                 `json:"migrate_from_chat_id,omitempty"`
	PinnedMessage                *TelegramMessage                      `json:"pinned_message,omitempty"`
	Invoice                      *TelegramInvoice                      `json:"invoice,omitempty"`
	SuccessfulPayment            *TelegramSuccessfulPayment            `json:"successful_payment,omitempty"`
	ConnectedWebsite             string                                `json:"connected_website,omitempty"`
	PassportData                 *TelegramPassportData                 `json:"passport_data,omitempty"`
	ProximityAlertTrigered       *TelegramProximityAlertTrigered       `json:"proximity_alert_triggered,omitempty"`
	VoiceChatSheduled            *TelegramVoiceChatSheduled            `json:"voice_chat_scheduled,omitempty"`
	VoiceChatStarted             *TelegramVoiceChatStarted             `json:"voice_chat_started,omitempty"`
	VoiceChatEnded               *TelegramVoiceChatEnded               `json:"voice_chat_ended,omitempty"`
	VoiceChatParticipantsInvited *TelegramVoiceChatParticipantsInvited `json:"voice_chat_participants_invited,omitempty"`
	WebAppData                   *TelegramWebAppData                   `json:"web_app_data"`
	ReplyMarkup                  *TelegramInlineKeyboardMarkup         `json:"reply_markup,omitempty"`
}
