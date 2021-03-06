package methods

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type AnswerCallBackQuery struct {
	CallbackQueryId string `json:"callback_query_id"`
	Text            string `json:"text,omitempty"`
	ShowAlert       bool   `json:"show_alert,omitempty"`
	Url             string `json:"url,omitempty"`
	CasheTime       int    `json:"cache_time,omitempty"`
}

func AnswerCallBackQueryMethod(Key string, CallbackQueryId string, Text string, Url string) ([]byte, error) {
	var Answer AnswerCallBackQuery
	Answer.CallbackQueryId = CallbackQueryId
	Answer.Text = Text
	Answer.Url = Url
	URL := fmt.Sprintf("https://api.telegram.org/bot%s/answerCallbackQuery", Key)
	postBody, err := json.Marshal(Answer)
	if err != nil {
		return nil, err
	}
	Body := bytes.NewBuffer(postBody)

	resp, err := http.Post(URL, "application/json", Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respbody, err := ioutil.ReadAll(resp.Body)
	return respbody, nil
}
