package command

import (
	"fmt"

	"strings"

	tgtypes "github.com/sudak-91/telegrambotgo/telegram_api/types"
)

type CommandHandler interface {
	Handl(interface{}) ([]byte, error)
}

type ICommandService interface {
	AddNewCommand(string, CommandHandler)
	Execute(string, interface{}) ([]byte, error)
	DefaultAnswer(*tgtypes.TelegramMessage) ([]byte, error)
}

type telegramCommandService struct {
	commandlist map[string]CommandHandler
}

func NewTelegramCommandService() *telegramCommandService {
	m := make(map[string]CommandHandler)
	return &telegramCommandService{
		commandlist: m,
	}
}
func (t *telegramCommandService) AddNewCommand(command string, handler CommandHandler) {
	t.commandlist[command] = handler
}

func (t *telegramCommandService) Execute(Command string, data interface{}) ([]byte, error) {
	SplitCommand := strings.Split(Command, " ")
	command := SplitCommand[0]
	val, ok := t.commandlist[command]
	if !ok {
		defval, ok := t.commandlist["default"]
		if !ok {
			return nil, fmt.Errorf("Execute command has no key: %v", command)
		}
		result, err := defval.Handl(data)
		if err != nil {
			return nil, fmt.Errorf("Excute command has error in Handl fucntion: %v", err)
		}
		return result, nil
	}

	result, err := val.Handl(data)
	if err != nil {
		return nil, fmt.Errorf("Excute command has error in Handl fucntion: %v", err)
	}
	return result, err
}

func (t *telegramCommandService) DefaultAnswer(Message *tgtypes.TelegramMessage) ([]byte, error) {
	command := "default"
	val, ok := t.commandlist[command]
	if !ok {
		return nil, fmt.Errorf("Execute command has no key: %v", command)
	}
	result, err := val.Handl(Message)
	if err != nil {
		return nil, fmt.Errorf("Excute command has error in Handl fucntion: %v", err)
	}
	return result, err
}
