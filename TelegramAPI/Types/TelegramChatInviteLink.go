package types

type TelegramChatInviteLink struct {
	InviteLink              string       `json:"invite_link"`
	Creator                 TelegramUser `json:"creator"`
	CreatesJoinRequest      bool         `json:"reates_join_request"`
	IsPrimary               bool         `json:"is_primary"`
	IsRevoked               bool         `json:"is_revoked"`
	Name                    string       `json:"name,omitempty"`
	ExpireData              int32        `json:"expire_data,omitempty"`
	MemberLimit             int32        `json:"member_limit,omitempty"`
	PendingJoinRequestCount int32        `json:"pending_join_request_count,omitempty"`
}
