package methods

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"net/http"

	tgtypes "github.com/sudak-91/telegrambotgo/TelegramAPI/Types"
)

//TODO: Переделать метод
type EditMessageText struct {
	ChatID    string `json:"chat_id,omitempty"`
	MessageID int    `json:"message_id,omitempty"`
	//InlineMessageID       int                                   `json:"inline_message_id,omitempty"`
	Text string `json:"text"`
	//ParseMode             string                                `json:"parse_mode,omitempty"`
	//Entities              []*tgtypes.TelegramMessageEntity      `json:"entites,omitempty"`
	//DisableWebPagePreview bool                                  `json:"disable_web_page_preview,omitempty"`
	//ReplyMarkup           *tgtypes.TelegramInlineKeyboardMarkup `json:"reply_markup,omitempty"`
}
type EditMessageTextResult struct {
	Ok     bool `json:"ok"`
	Result *tgtypes.TelegramMessage
}

func EditMessageTextMethod(ChatID string, MessageID int, Key string) {
	var editMessageText EditMessageText
	editMessageText.ChatID = ChatID
	editMessageText.MessageID = MessageID
	editMessageText.Text = "NEW TEXT"
	URL := fmt.Sprintf("https://api.telegram.org/bot%s/editMessageText", Key)
	postBody, _ := json.Marshal(editMessageText)

	Body := bytes.NewBuffer(postBody)

	resp, _ := http.Post(URL, "application/json", Body)

	defer resp.Body.Close()
	var Result GetChatMemberResult
	respbody, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal(respbody, &Result)

}
