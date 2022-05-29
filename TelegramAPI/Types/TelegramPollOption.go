package types

type TelegramPollOption struct {
	Text       string `json:"text"`
	VoterCount int32  `json:"voter_count"`
}
