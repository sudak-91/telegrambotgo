package types

type TelegramChatMember struct {
	Status              string        `json:"status"`
	User                *TelegramUser `json:"user"`
	CanBeEdited         bool          `json:"can_be_edited,omitempty"`
	IsAnonymous         bool          `json:"is_anonymous,omitempty"`
	CanManageChat       bool          `json:"can_manage_chat,omitempty"`
	CanDeleteMessages   bool          `json:"can_delete_messages,omitempty"`
	CanManageVideoChats bool          `json:"can_manage_video_chats,omitempty"`
	CanRestrictMembers  bool          `json:"can_restrict_members,omitempty"`
	CanPromoteMembers   bool          `json:"can_promote_members,omitempty"`
	CanChangeInfo       bool          `json:"can_change_info,omitempty"`
	CanInviteUsers      bool          `json:"can_invite_users,omitempty"`
	CanPostMessages     bool          `json:"can_post_messages,omitempty"`
	CanEditMessages     bool          `json:"can_edit_messages,omitempty"`
	CanPinMessages      bool          `json:"can_pin_messages,omitempty"`
	CustomTitle         string        `json:"custom_title,omitempty"`
}
