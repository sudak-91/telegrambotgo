package methods

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"net/http"

	tgtypes "github.com/sudak-91/telegrambotgo/telegram_api/types"
)

type SetMyCommand struct {
	Commands []tgtypes.BotCommand `json:"commands"`
}

func (s *SetMyCommand) AddCommand(command, description string) {
	var Command tgtypes.BotCommand
	Command.Command = command
	Command.Description = description
	s.Commands = append(s.Commands, Command)

}

func SetMyCommandMethos(Key string, Commands SetMyCommand) ([]byte, error) {
	URL := fmt.Sprintf("https://api.telegram.org/bot%s/setMyCommands", Key)
	postBody, err := json.Marshal(Commands)
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
