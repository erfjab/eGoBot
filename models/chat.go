package models

// https://core.telegram.org/bots/api#chattype
type ChatType string

const (
	ChatTypePrivate    ChatType = "private"
	ChatTypeGroup      ChatType = "group"
	ChatTypeSupergroup ChatType = "supergroup"
	ChatTypeChannel    ChatType = "channel"
)

// https://core.telegram.org/bots/api#chat
type Chat struct {
	ID               int64    `json:"id"`
	Type             ChatType `json:"type"`
	Title            string   `json:"title,omitempty"`
	Username         string   `json:"username,omitempty"`
	FirstName        string   `json:"first_name,omitempty"`
	LastName         string   `json:"last_name,omitempty"`
	IsDirectMessages bool     `json:"is_direct_messages,omitempty"`
}

// https://core.telegram.org/bots/api#chatmember
type ChatMember struct {
	Status                string `json:"status"`
	User                  User   `json:"user"`
	IsAnonymous           bool   `json:"is_anonymous,omitempty"`
	CustomTitle           string `json:"custom_title,omitempty"`
	CanBeEdited           bool   `json:"can_be_edited,omitempty"`
	CanManageChat         bool   `json:"can_manage_chat,omitempty"`
	CanDeleteMessages     bool   `json:"can_delete_messages,omitempty"`
	CanManageVideoChats   bool   `json:"can_manage_video_chats,omitempty"`
	CanRestrictMembers    bool   `json:"can_restrict_members,omitempty"`
	CanPromoteMembers     bool   `json:"can_promote_members,omitempty"`
	CanChangeInfo         bool   `json:"can_change_info,omitempty"`
	CanInviteUsers        bool   `json:"can_invite_users,omitempty"`
	CanPostMessages       bool   `json:"can_post_messages,omitempty"`
	CanEditMessages       bool   `json:"can_edit_messages,omitempty"`
	CanPinMessages        bool   `json:"can_pin_messages,omitempty"`
	CanManageTopics       bool   `json:"can_manage_topics,omitempty"`
	IsMember              bool   `json:"is_member,omitempty"`
	CanSendMessages       bool   `json:"can_send_messages,omitempty"`
	CanSendAudios         bool   `json:"can_send_audios,omitempty"`
	CanSendDocuments      bool   `json:"can_send_documents,omitempty"`
	CanSendPhotos         bool   `json:"can_send_photos,omitempty"`
	CanSendVideos         bool   `json:"can_send_videos,omitempty"`
	CanSendVideoNotes     bool   `json:"can_send_video_notes,omitempty"`
	CanSendVoiceNotes     bool   `json:"can_send_voice_notes,omitempty"`
	CanSendPolls          bool   `json:"can_send_polls,omitempty"`
	CanSendOtherMessages  bool   `json:"can_send_other_messages,omitempty"`
	CanAddWebPagePreviews bool   `json:"can_add_web_page_previews,omitempty"`
	UntilDate             int64  `json:"until_date,omitempty"`
}

// https://core.telegram.org/bots/api#chatpermissions
type ChatPermissions struct {
	CanSendMessages       bool `json:"can_send_messages,omitempty"`
	CanSendAudios         bool `json:"can_send_audios,omitempty"`
	CanSendDocuments      bool `json:"can_send_documents,omitempty"`
	CanSendPhotos         bool `json:"can_send_photos,omitempty"`
	CanSendVideos         bool `json:"can_send_videos,omitempty"`
	CanSendVideoNotes     bool `json:"can_send_video_notes,omitempty"`
	CanSendVoiceNotes     bool `json:"can_send_voice_notes,omitempty"`
	CanSendPolls          bool `json:"can_send_polls,omitempty"`
	CanSendOtherMessages  bool `json:"can_send_other_messages,omitempty"`
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews,omitempty"`
	CanChangeInfo         bool `json:"can_change_info,omitempty"`
	CanInviteUsers        bool `json:"can_invite_users,omitempty"`
	CanPinMessages        bool `json:"can_pin_messages,omitempty"`
	CanManageTopics       bool `json:"can_manage_topics,omitempty"`
}

// https://core.telegram.org/bots/api#banchatmember
type BanChatMemberParams struct {
	ChatID         interface{} `json:"chat_id"`
	UserID         int64       `json:"user_id"`
	UntilDate      int64       `json:"until_date,omitempty"`
	RevokeMessages bool        `json:"revoke_messages,omitempty"`
}

// https://core.telegram.org/bots/api#unbanchatmember
type UnbanChatMemberParams struct {
	ChatID       interface{} `json:"chat_id"`
	UserID       int64       `json:"user_id"`
	OnlyIfBanned bool        `json:"only_if_banned,omitempty"`
}

// https://core.telegram.org/bots/api#restrictchatmember
type RestrictChatMemberParams struct {
	ChatID                        interface{}           `json:"chat_id"`
	UserID                        int64                 `json:"user_id"`
	Permissions                   ChatPermissions       `json:"permissions"`
	UseIndependentChatPermissions bool                  `json:"use_independent_chat_permissions,omitempty"`
	UntilDate                     int64                 `json:"until_date,omitempty"`
}

// https://core.telegram.org/bots/api#promotechatmember
type PromoteChatMemberParams struct {
	ChatID              interface{} `json:"chat_id"`
	UserID              int64       `json:"user_id"`
	IsAnonymous         bool        `json:"is_anonymous,omitempty"`
	CanManageChat       bool        `json:"can_manage_chat,omitempty"`
	CanDeleteMessages   bool        `json:"can_delete_messages,omitempty"`
	CanManageVideoChats bool        `json:"can_manage_video_chats,omitempty"`
	CanRestrictMembers  bool        `json:"can_restrict_members,omitempty"`
	CanPromoteMembers   bool        `json:"can_promote_members,omitempty"`
	CanChangeInfo       bool        `json:"can_change_info,omitempty"`
	CanInviteUsers      bool        `json:"can_invite_users,omitempty"`
	CanPostMessages     bool        `json:"can_post_messages,omitempty"`
	CanEditMessages     bool        `json:"can_edit_messages,omitempty"`
	CanPinMessages      bool        `json:"can_pin_messages,omitempty"`
	CanManageTopics     bool        `json:"can_manage_topics,omitempty"`
}

// https://core.telegram.org/bots/api#setchatadministratorcustomtitle
type SetChatAdministratorCustomTitleParams struct {
	ChatID      interface{} `json:"chat_id"`
	UserID      int64       `json:"user_id"`
	CustomTitle string      `json:"custom_title"`
}

// https://core.telegram.org/bots/api#getchatmember
type GetChatMemberParams struct {
	ChatID interface{} `json:"chat_id"`
	UserID int64       `json:"user_id"`
}

// https://core.telegram.org/bots/api#pinchatmessage
type PinChatMessageParams struct {
	ChatID              interface{} `json:"chat_id"`
	MessageID           int64       `json:"message_id"`
	DisableNotification bool        `json:"disable_notification,omitempty"`
}

// https://core.telegram.org/bots/api#unpinchatmessage
type UnpinChatMessageParams struct {
	ChatID    interface{} `json:"chat_id"`
	MessageID int64       `json:"message_id,omitempty"`
}

// https://core.telegram.org/bots/api#forumtopic
type ForumTopic struct {
	MessageThreadID   int    `json:"message_thread_id"`
	Name              string `json:"name"`
	IconColor         int    `json:"icon_color"`
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"`
}

// https://core.telegram.org/bots/api#createforumtopic
type CreateForumTopicParams struct {
	ChatID            interface{} `json:"chat_id"`
	Name              string      `json:"name"`
	IconColor         int         `json:"icon_color,omitempty"`
	IconCustomEmojiID string      `json:"icon_custom_emoji_id,omitempty"`
}

// https://core.telegram.org/bots/api#editforumtopic
type EditForumTopicParams struct {
	ChatID            interface{} `json:"chat_id"`
	MessageThreadID   int         `json:"message_thread_id"`
	Name              string      `json:"name,omitempty"`
	IconCustomEmojiID string      `json:"icon_custom_emoji_id,omitempty"`
}

// https://core.telegram.org/bots/api#closeforumtopic
type CloseForumTopicParams struct {
	ChatID          interface{} `json:"chat_id"`
	MessageThreadID int         `json:"message_thread_id"`
}

// https://core.telegram.org/bots/api#reopenforumtopic
type ReopenForumTopicParams struct {
	ChatID          interface{} `json:"chat_id"`
	MessageThreadID int         `json:"message_thread_id"`
}

// https://core.telegram.org/bots/api#deleteforumtopic
type DeleteForumTopicParams struct {
	ChatID          interface{} `json:"chat_id"`
	MessageThreadID int         `json:"message_thread_id"`
}

// https://core.telegram.org/bots/api#unpinallforumtopicmessages
type UnpinAllForumTopicMessagesParams struct {
	ChatID          interface{} `json:"chat_id"`
	MessageThreadID int         `json:"message_thread_id"`
}

// https://core.telegram.org/bots/api#editgeneralforumtopic
type EditGeneralForumTopicParams struct {
	ChatID interface{} `json:"chat_id"`
	Name   string      `json:"name"`
}

// https://core.telegram.org/bots/api#closegeneralforumtopic
type CloseGeneralForumTopicParams struct {
	ChatID interface{} `json:"chat_id"`
}

// https://core.telegram.org/bots/api#reopengeneralforumtopic
type ReopenGeneralForumTopicParams struct {
	ChatID interface{} `json:"chat_id"`
}

// https://core.telegram.org/bots/api#hidegeneralforumtopic
type HideGeneralForumTopicParams struct {
	ChatID interface{} `json:"chat_id"`
}

// https://core.telegram.org/bots/api#unhidegeneralforumtopic
type UnhideGeneralForumTopicParams struct {
	ChatID interface{} `json:"chat_id"`
}

// https://core.telegram.org/bots/api#setchatphoto
type SetChatPhotoParams struct {
	ChatID interface{} `json:"chat_id"`
	Photo  interface{} `json:"photo"`
}

// https://core.telegram.org/bots/api#deletechatphoto
type DeleteChatPhotoParams struct {
	ChatID interface{} `json:"chat_id"`
}

// https://core.telegram.org/bots/api#setchattitle
type SetChatTitleParams struct {
	ChatID interface{} `json:"chat_id"`
	Title  string      `json:"title"`
}

// https://core.telegram.org/bots/api#setchatdescription
type SetChatDescriptionParams struct {
	ChatID      interface{} `json:"chat_id"`
	Description string      `json:"description,omitempty"`
}

// https://core.telegram.org/bots/api#banchatsenderchat
type BanChatSenderChatParams struct {
	ChatID       interface{} `json:"chat_id"`
	SenderChatID int64       `json:"sender_chat_id"`
}

// https://core.telegram.org/bots/api#unbanchatsenderchat
type UnbanChatSenderChatParams struct {
	ChatID       interface{} `json:"chat_id"`
	SenderChatID int64       `json:"sender_chat_id"`
}

// https://core.telegram.org/bots/api#setchatpermissions
type SetChatPermissionsParams struct {
	ChatID                       interface{}     `json:"chat_id"`
	Permissions                  ChatPermissions `json:"permissions"`
	UseIndependentChatPermissions bool            `json:"use_independent_chat_permissions,omitempty"`
}

// https://core.telegram.org/bots/api#exportchatinvitelink
type ExportChatInviteLinkParams struct {
	ChatID interface{} `json:"chat_id"`
}

// https://core.telegram.org/bots/api#chatinvitelink
type ChatInviteLink struct {
	InviteLink              string `json:"invite_link"`
	Creator                 User   `json:"creator"`
	CreatesJoinRequest      bool   `json:"creates_join_request"`
	IsPrimary               bool   `json:"is_primary"`
	IsRevoked               bool   `json:"is_revoked"`
	Name                    string `json:"name,omitempty"`
	ExpireDate              int64  `json:"expire_date,omitempty"`
	MemberLimit             int    `json:"member_limit,omitempty"`
	PendingJoinRequestCount int    `json:"pending_join_request_count,omitempty"`
}

// https://core.telegram.org/bots/api#createchatinvitelink
type CreateChatInviteLinkParams struct {
	ChatID             interface{} `json:"chat_id"`
	Name               string      `json:"name,omitempty"`
	ExpireDate         int64       `json:"expire_date,omitempty"`
	MemberLimit        int         `json:"member_limit,omitempty"`
	CreatesJoinRequest bool        `json:"creates_join_request,omitempty"`
}

// https://core.telegram.org/bots/api#editchatinvitelink
type EditChatInviteLinkParams struct {
	ChatID             interface{} `json:"chat_id"`
	InviteLink         string      `json:"invite_link"`
	Name               string      `json:"name,omitempty"`
	ExpireDate         int64       `json:"expire_date,omitempty"`
	MemberLimit        int         `json:"member_limit,omitempty"`
	CreatesJoinRequest bool        `json:"creates_join_request,omitempty"`
}

// https://core.telegram.org/bots/api#revokechatinvitelink
type RevokeChatInviteLinkParams struct {
	ChatID     interface{} `json:"chat_id"`
	InviteLink string      `json:"invite_link"`
}

// https://core.telegram.org/bots/api#approvechatjoinrequest
type ApproveChatJoinRequestParams struct {
	ChatID interface{} `json:"chat_id"`
	UserID int64       `json:"user_id"`
}

// https://core.telegram.org/bots/api#declinechatjoinrequest
type DeclineChatJoinRequestParams struct {
	ChatID interface{} `json:"chat_id"`
	UserID int64       `json:"user_id"`
}

// https://core.telegram.org/bots/api#setchatstickerset
type SetChatStickerSetParams struct {
	ChatID         interface{} `json:"chat_id"`
	StickerSetName string      `json:"sticker_set_name"`
}

// https://core.telegram.org/bots/api#deletechatstickerset
type DeleteChatStickerSetParams struct {
	ChatID interface{} `json:"chat_id"`
}

// https://core.telegram.org/bots/api#chatjoinrequest
type ChatJoinRequest struct {
	Chat       Chat       `json:"chat"`
	From       User       `json:"from"`
	UserChatID int64      `json:"user_chat_id"`
	Date       int64      `json:"date"`
	Bio        string     `json:"bio,omitempty"`
	InviteLink *ChatInviteLink `json:"invite_link,omitempty"`
}
