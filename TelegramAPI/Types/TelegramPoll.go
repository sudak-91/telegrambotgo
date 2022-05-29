package types

type TelegramPoll struct {
	ID                    string                  `json:"id"`
	Question              string                  `json:"question"`
	Options               []TelegramPollOption    `json:"options"`
	TotalVoterCount       int32                   `json:"total_voter_count"`
	IsClosed              bool                    `json:"is_closed"`
	IsAnonymous           bool                    `json:"is_anonymous"`
	Type                  string                  `json:"type"`
	AllowsMultipleAnswers bool                    `json:"allows_multiple_answers"`
	CorrectOptionID       int32                   `json:"correct_option_id,omitempty"`
	Explanation           string                  `json:"explanation,omitempty"`
	ExplanationEntities   []TelegramMessageEntity `json:"explanation_entitites,omitempty"`
	OpenPeriod            int32                   `json:"open_period,omitempty"`
	CloseDate             int32                   `json:"close_data,omitempty"`
}
