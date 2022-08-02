package keyboardmaker

import types "github.com/sudak-91/telegrambotgo/TelegramAPI/Types"

//InlineCommandKeyboard создает Inline клавиатуру для сообщения
type InlineCommandKeyboard struct {
	types.TelegramInlineKeyboardMarkup
}

func (k *InlineCommandKeyboard) MakeGrid(raw int, column int) {
	k.InlineKeyboard = make([][]types.TelegramInlineKeyboardButton, raw)
	for i := range k.InlineKeyboard {
		k.InlineKeyboard[i] = make([]types.TelegramInlineKeyboardButton, column)
	}

}

//AddButton добавляется кнопку с CommandCallback
func (k *InlineCommandKeyboard) AddButton(text string, commandcallback string, raw int, column int) {
	var Button types.TelegramInlineKeyboardButton
	Button.Text = text
	Button.CallbackData = commandcallback
	k.InlineKeyboard[raw][column] = Button
}

func (k *InlineCommandKeyboard) AddSwitchInlineButton(text string, url string, raw int, column int) {
	var Button types.TelegramInlineKeyboardButton
	Button.Text = text
	Button.SwitchInlineQuery = url
	Button.Url = url
	k.InlineKeyboard[raw][column] = Button

}
func (k *InlineCommandKeyboard) AddLoginUrlButton(text string, button types.TelegramLoginUrl, raw int, column int) {
	var Button types.TelegramInlineKeyboardButton
	Button.Text = text
	Button.LoginUrl = button
	k.InlineKeyboard[raw][column] = Button
}

func (k *InlineCommandKeyboard) addCallbackQuery(btn *types.TelegramInlineKeyboardButton) {
	btn.CallbackData = "/start"
}
func (k InlineCommandKeyboard) ClearButton(raw int, column int) {
	return
}
func (k *InlineCommandKeyboard) GetKeyboard() types.TelegramInlineKeyboardMarkup {
	return k.TelegramInlineKeyboardMarkup
}
