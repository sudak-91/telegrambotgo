package update

import (
	"encoding/json"
	"fmt"

	types "github.com/sudak-91/telegrambotgo/TelegramAPI/Types"
)

//Updater - интерфейс для обработки входящих данных от телеграмм бота(webhook)
//Метод Update является роутером для данных обнолвений
type Updater interface {
	Update([]byte) ([]byte, error)
}

//TelegramUpdater - интерфейс предназначенный для реализации методов обработки входящих данных
type TelegramUpdater interface {
	CallbackQueryService(types.TelegramCallbackQuery) ([]byte, error)
	ChannelPostService(types.TelegramMessage) ([]byte, error)
	ChatJoinRequsetService(types.TelegramChatJoinRequest) ([]byte, error)
	ChatMemberService(types.TelegramChatMemberUpdated) ([]byte, error)
	ChosenInlineResultService(types.TelegramChosenInlineResult) ([]byte, error)
	EditedChannelPostService(types.TelegramMessage) ([]byte, error)
	EditedMessageService(types.TelegramMessage) ([]byte, error)
	InlineQueryService(types.TelegramInlineQuery) ([]byte, error)
	MessageService(types.TelegramMessage) ([]byte, error)
	MyChatMemberService(types.TelegramChatMemberUpdated) ([]byte, error)
	PollService(types.TelegramPoll) ([]byte, error)
	PollAnswerService(types.TelegramPollAnwer) ([]byte, error)
	PreCheckoutPollService(types.TelegramPreCheckoutQuery) ([]byte, error)
	ShippingService(types.TelegramShippingQuery) ([]byte, error)
	ChatUserUpdateService(types.TelegramUpdate) ([]byte, error)
	Default(types.TelegramUpdate) ([]byte, error)
}

type TelegramService struct {
	TelegramUpdater
}

//NewTelegramService - конструктор возвращающий объект со стандартной реализацией интерфейса Updater
//Param Upd - пользовательская реализация интерфейса TelegramUpdater
func NewTelegramService(Upd TelegramUpdater) *TelegramService {
	return &TelegramService{
		TelegramUpdater: Upd,
	}
}

func (ts *TelegramService) Update(data []byte) ([]byte, error) {
	var Update types.TelegramUpdate

	err := json.Unmarshal(data, &Update)
	if err != nil {
		return nil, fmt.Errorf("GetRequesData from Telegram Service return: %v", err.Error())
	}

	switch {
	case Update.CallbackQuery != nil:
		return ts.CallbackQueryService(*Update.CallbackQuery)
	case Update.ChannelPost != nil:
		return ts.ChannelPostService(*Update.ChannelPost)
	case Update.ChatJoinRequest != nil:
		return ts.ChatJoinRequsetService(*Update.ChatJoinRequest)
	case Update.ChatMember != nil:
		return ts.ChatMemberService(*Update.ChatMember)
	case Update.ChosenInlineResult != nil:
		return ts.ChosenInlineResultService(*Update.ChosenInlineResult)
	case Update.EditedChannelPost != nil:
		return ts.EditedChannelPostService(*Update.Message)
	case Update.EditedMessage != nil:
		return ts.EditedMessageService(*Update.EditedMessage)
	case Update.InlineQuery != nil:
		return ts.InlineQueryService(*Update.InlineQuery)
	case Update.Message != nil:
		return ts.MessageService(*Update.Message)
	case Update.MyChatMember != nil:
		return ts.MyChatMemberService(*Update.MyChatMember)
	case Update.Poll != nil:
		return ts.PollService(*Update.Poll)
	case Update.PollAnswer != nil:
		return ts.PollAnswerService(*Update.PollAnswer)
	case Update.PreCheckoutQuery != nil:
		return ts.PreCheckoutPollService(*Update.PreCheckoutQuery)
	case Update.ShippingQuery != nil:
		return ts.ShippingService(*Update.ShippingQuery)
	default:
		return ts.Default(Update)
	}

}
