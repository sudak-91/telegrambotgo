package types

type PollAnwer struct {
	PollID    string  `json:"poll_id"`
	User      *User   `json:"user"`
	OptionIds []int32 `json:"option_ids"`
}
