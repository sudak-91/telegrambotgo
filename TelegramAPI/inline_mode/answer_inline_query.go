package inline

type AnswerInlineQuery struct {
	InlineQueryId     string               `json:"inline_query_id"`
	Results           []*InlineQueryResult `json:"results"`
	CacheTime         int32                `json:"cache_time,omitepmty"`
	IsPersonal        bool                 `json:"is_personal,omitempty"`
	NextOffset        string               `json:"next_offset,omitempty"`
	SwitchPmText      string               `json:"switch_pm_text,omitempty"`
	SwitchPmParameter string               `json:"switch_pm_patameter,omitepmty"`
}
