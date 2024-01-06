package keyboardmaker

import types "github.com/sudak-91/telegrambotgo/telegram_api/types"

// InlineCommandKeyboard создает Inline клавиатуру для сообщения
type InlineCommandKeyboard struct {
	types.InlineKeyboardMarkup
}

func (k *InlineCommandKeyboard) MakeGrid(raw int, column int) {
	k.InlineKeyboard = make([][]types.InlineKeyboardButton, raw)
	for i := range k.InlineKeyboard {
		k.InlineKeyboard[i] = make([]types.InlineKeyboardButton, column)
	}

}

// AddButton добавляется кнопку с CommandCallback
func (k *InlineCommandKeyboard) AddButton(text string, commandcallback string, raw int, column int) {
	var Button types.InlineKeyboardButton
	Button.Text = text
	Button.CallbackData = commandcallback
	k.InlineKeyboard[raw][column] = Button
}

func (k *InlineCommandKeyboard) AddSwitchInlineButton(text string, url string, raw int, column int) {
	var Button types.InlineKeyboardButton
	Button.Text = text
	Button.SwitchInlineQuery = url
	Button.Url = url
	k.InlineKeyboard[raw][column] = Button

}
func (k *InlineCommandKeyboard) AddLoginUrlButton(text string, button *types.LoginUrl, raw int, column int) {
	var Button types.InlineKeyboardButton
	Button.Text = text
	Button.LoginUrl = button
	k.InlineKeyboard[raw][column] = Button
}

func (k *InlineCommandKeyboard) addCallbackQuery(btn *types.InlineKeyboardButton) {
	btn.CallbackData = "/start"
}
func (k InlineCommandKeyboard) ClearButton(raw int, column int) {
	return
}
func (k *InlineCommandKeyboard) GetKeyboard() types.InlineKeyboardMarkup {
	return k.InlineKeyboardMarkup
}
