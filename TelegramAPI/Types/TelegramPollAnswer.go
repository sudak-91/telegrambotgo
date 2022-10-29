package types

type TelegramPollAnwer struct {
	PollID    string        `json:"poll_id"`
	User      *TelegramUser `json:"user"`
	OptionIds []int32       `json:"option_ids"`
}
