package types

type ResponceParameters struct {
	MigrateToChatId int32 `json:"migrate_to_chat_id,omitempty"`
	RetryAfter      int32 `json:"retry_after,omitempty"`
}
