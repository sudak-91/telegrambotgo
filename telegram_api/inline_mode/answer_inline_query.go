package inline

type AnswerInlineQuery struct {
	InlineQueryId string                   `json:"inline_query_id"`
	Results       []any                    `json:"results"`
	CacheTime     int32                    `json:"cache_time,omitempty"`
	IsPersonal    bool                     `json:"is_personal,omitempty"`
	NextOffset    string                   `json:"next_offset,omitempty"`
	Button        *InlineQueryResultButton `json:"button,omitempty"`
}
