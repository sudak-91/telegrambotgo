package types

type ChatMemberRestricted struct {
	User                  *User `json:"user"`
	IsMember              bool  `json:"is_member"`
	CanChangeInfo         bool  `json:"can_change_info"`
	CanInviteUsers        bool  `json:"can_invite_users"`
	CanPinMessages        bool  `json:"can_pin_messages"`
	CanManageTopics       bool  `json:"can_manage_topics"`
	CanSendMessages       bool  `json:"can_send_messages"`
	CanSendMediaMessages  bool  `json:"can_send_media_messages"`
	CanSendPolls          bool  `json:"can_send_polls"`
	CanSendOtherMessages  bool  `json:"can_send_other_messages"`
	CanAddWebPagePreviews bool  `json:"can_add_web_page_previews"`
	UntilDate             int32 `json"until_date"`
}
