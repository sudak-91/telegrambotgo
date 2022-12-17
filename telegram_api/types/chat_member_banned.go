package types

type ChatMemberBanned struct {
	User      *User `json:"user"`
	UntilDate int32 `json:"until_date"`
}
