package models

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type ChatMemberUpdated struct {
	Chat          Chat           `json:"chat"`
	From          *User          `json:"from"`
	Date          int            `json:"date"`
	OldChatMember ChatMember     `json:"old_chat_member"`
	NewChatMember ChatMember     `json:"new_chat_member"`
	InviteLink    ChatInviteLink `json:"invite_link"`
}

type ChatMemberType int

const (
	ChatMemberTypeOwner ChatMemberType = iota
	ChatMemberTypeAdministrator
	ChatMemberTypeMember
	ChatMemberTypeRestricted
	ChatMemberTypeLeft
	ChatMemberTypeBanned
)

type ChatMember struct {
	Type ChatMemberType

	Owner         *ChatMemberOwner
	Administrator *ChatMemberAdministrator
	Member        *ChatMemberMember
	Restricted    *ChatMemberRestricted
	Left          *ChatMemberLeft
	Banned        *ChatMemberBanned
}

func (c *ChatMember) UnmarshalJSON(data []byte) error {
	if bytes.Contains(data, []byte(`"status":"creator"`)) {
		c.Type = ChatMemberTypeOwner
		c.Owner = &ChatMemberOwner{}
		return json.Unmarshal(data, c.Owner)
	}
	if bytes.Contains(data, []byte(`"status":"administrator"`)) {
		c.Type = ChatMemberTypeAdministrator
		c.Administrator = &ChatMemberAdministrator{}
		return json.Unmarshal(data, c.Administrator)
	}
	if bytes.Contains(data, []byte(`"status":"member"`)) {
		c.Type = ChatMemberTypeMember
		c.Member = &ChatMemberMember{}
		return json.Unmarshal(data, c.Member)
	}
	if bytes.Contains(data, []byte(`"status":"restricted"`)) {
		c.Type = ChatMemberTypeRestricted
		c.Restricted = &ChatMemberRestricted{}
		return json.Unmarshal(data, c.Restricted)
	}
	if bytes.Contains(data, []byte(`"status":"left"`)) {
		c.Type = ChatMemberTypeLeft
		c.Left = &ChatMemberLeft{}
		return json.Unmarshal(data, c.Left)
	}
	if bytes.Contains(data, []byte(`"status":"kicked"`)) {
		c.Type = ChatMemberTypeBanned
		c.Banned = &ChatMemberBanned{}
		return json.Unmarshal(data, c.Banned)
	}

	return fmt.Errorf("unsupported ChatMember type")
}

type ChatMemberOwner struct {
	Status      string `json:"status"` // The member's status in the chat, always “creator”
	User        *User  `json:"user"`
	IsAnonymous bool   `json:"is_anonymous"`
	CustomTitle string `json:"custom_title,omitempty"`
}

func (ChatMemberOwner) Type() ChatMemberType {
	return ChatMemberTypeOwner
}

type ChatMemberAdministrator struct {
	Status              string `json:"status"` // The member's status in the chat, always “administrator”
	User                *User  `json:"user"`
	CanBeEdited         bool   `json:"can_be_edited"`
	IsAnonymous         bool   `json:"is_anonymous"`
	CanManageChat       bool   `json:"can_manage_chat"`
	CanDeleteMessages   bool   `json:"can_delete_messages"`
	CanManageVideoChats bool   `json:"can_manage_video_chats"`
	CanRestrictMembers  bool   `json:"can_restrict_members"`
	CanPromoteMembers   bool   `json:"can_promote_members"`
	CanChangeInfo       bool   `json:"can_change_info"`
	CanInviteUsers      bool   `json:"can_invite_users"`
	CanPostMessages     bool   `json:"can_post_messages,omitempty"`
	CanEditMessages     bool   `json:"can_edit_messages,omitempty"`
	CanPinMessages      bool   `json:"can_pin_messages,omitempty"`
	CustomTitle         string `json:"custom_title,omitempty"`
}

type ChatMemberMember struct {
	Status string `json:"status"` // The member's status in the chat, always “member”
	User   *User  `json:"user"`
}

type ChatMemberRestricted struct {
	Status                string `json:"status"` // The member's status in the chat, always “restricted”
	User                  *User  `json:"user"`
	IsMember              bool   `json:"is_member"`
	CanChangeInfo         bool   `json:"can_change_info"`
	CanInviteUsers        bool   `json:"can_invite_users"`
	CanPinMessages        bool   `json:"can_pin_messages"`
	CanSendMessages       bool   `json:"can_send_messages"`
	CanSendMediaMessages  bool   `json:"can_send_media_messages"`
	CanSendPolls          bool   `json:"can_send_polls"`
	CanSendOtherMessages  bool   `json:"can_send_other_messages"`
	CanAddWebPagePreviews bool   `json:"can_add_web_page_previews"`
	UntilDate             int    `json:"until_date"`
}

type ChatMemberLeft struct {
	Status string `json:"status"` // The member's status in the chat, always “left”
	User   *User  `json:"user"`
}

type ChatMemberBanned struct {
	Status    string `json:"status"` // The member's status in the chat, always “kicked”
	User      *User  `json:"user"`
	UntilDate int    `json:"until_date"`
}