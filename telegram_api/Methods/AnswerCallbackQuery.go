package methods

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type AnswerCallBackQuery struct {
	CallbackQueryId string `json:"callback_query_id"`
	Text            string `json:"text,omitempty"`
	ShowAlert       bool   `json:"show_alert,omitempty"`
	Url             string `json:"url,omitempty"`
	CasheTime       int    `json:"cache_time,omitempty"`
}

func AnswerCallBackQueryMethod(Key string, answerCallBackQuery AnswerCallBackQuery) error {
	URL := fmt.Sprintf("https://api.telegram.org/bot%s/answerCallbackQuery", Key)
	postBody, err := json.Marshal(answerCallBackQuery)
	if err != nil {
		return nil
	}
	Body := bytes.NewBuffer(postBody)
	_, err = http.Post(URL, "application/json", Body)
	if err != nil {
		return err
	}
	return nil
	// defer resp.Body.Close()
	// respbody, err := ioutil.ReadAll(resp.Body)
	// return respbody, nil
}
