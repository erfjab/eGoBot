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
	IsForum          bool     `json:"is_forum,omitempty"`
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
	ChatID                        interface{}     `json:"chat_id"`
	UserID                        int64           `json:"user_id"`
	Permissions                   ChatPermissions `json:"permissions"`
	UseIndependentChatPermissions bool            `json:"use_independent_chat_permissions,omitempty"`
	UntilDate                     int64           `json:"until_date,omitempty"`
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
	IsNameImplicit    bool   `json:"is_name_implicit,omitempty"`
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
	ChatID                        interface{}     `json:"chat_id"`
	Permissions                   ChatPermissions `json:"permissions"`
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
	SubscriptionPeriod      int    `json:"subscription_period,omitempty"`
	SubscriptionPrice       int    `json:"subscription_price,omitempty"`
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

// https://core.telegram.org/bots/api#chatfullinfo
type ChatFullInfo struct {
	ID                                 int64                 `json:"id"`
	Type                               ChatType              `json:"type"`
	Title                              string                `json:"title,omitempty"`
	Username                           string                `json:"username,omitempty"`
	FirstName                          string                `json:"first_name,omitempty"`
	LastName                           string                `json:"last_name,omitempty"`
	IsForum                            bool                  `json:"is_forum,omitempty"`
	IsDirectMessages                   bool                  `json:"is_direct_messages,omitempty"`
	AccentColorID                      int                   `json:"accent_color_id"`
	MaxReactionCount                   int                   `json:"max_reaction_count"`
	Photo                              *ChatPhoto            `json:"photo,omitempty"`
	ActiveUsernames                    []string              `json:"active_usernames,omitempty"`
	Birthdate                          *Birthdate            `json:"birthdate,omitempty"`
	BusinessIntro                      *BusinessIntro        `json:"business_intro,omitempty"`
	BusinessLocation                   *BusinessLocation     `json:"business_location,omitempty"`
	BusinessOpeningHours               *BusinessOpeningHours `json:"business_opening_hours,omitempty"`
	PersonalChat                       *Chat                 `json:"personal_chat,omitempty"`
	ParentChat                         *Chat                 `json:"parent_chat,omitempty"`
	AvailableReactions                 []ReactionType        `json:"available_reactions,omitempty"`
	BackgroundCustomEmojiID            string                `json:"background_custom_emoji_id,omitempty"`
	ProfileAccentColorID               int                   `json:"profile_accent_color_id,omitempty"`
	ProfileBackgroundCustomEmojiID     string                `json:"profile_background_custom_emoji_id,omitempty"`
	EmojiStatusCustomEmojiID           string                `json:"emoji_status_custom_emoji_id,omitempty"`
	EmojiStatusExpirationDate          int64                 `json:"emoji_status_expiration_date,omitempty"`
	Bio                                string                `json:"bio,omitempty"`
	HasPrivateForwards                 bool                  `json:"has_private_forwards,omitempty"`
	HasRestrictedVoiceAndVideoMessages bool                  `json:"has_restricted_voice_and_video_messages,omitempty"`
	JoinToSendMessages                 bool                  `json:"join_to_send_messages,omitempty"`
	JoinByRequest                      bool                  `json:"join_by_request,omitempty"`
	Description                        string                `json:"description,omitempty"`
	InviteLink                         string                `json:"invite_link,omitempty"`
	PinnedMessage                      *Message              `json:"pinned_message,omitempty"`
	Permissions                        *ChatPermissions      `json:"permissions,omitempty"`
	AcceptedGiftTypes                  *AcceptedGiftTypes    `json:"accepted_gift_types"`
	CanSendPaidMedia                   bool                  `json:"can_send_paid_media,omitempty"`
	SlowModeDelay                      int                   `json:"slow_mode_delay,omitempty"`
	UnrestrictBoostCount               int                   `json:"unrestrict_boost_count,omitempty"`
	MessageAutoDeleteTime              int                   `json:"message_auto_delete_time,omitempty"`
	HasAggressiveAntiSpamEnabled       bool                  `json:"has_aggressive_anti_spam_enabled,omitempty"`
	HasHiddenMembers                   bool                  `json:"has_hidden_members,omitempty"`
	HasProtectedContent                bool                  `json:"has_protected_content,omitempty"`
	HasVisibleHistory                  bool                  `json:"has_visible_history,omitempty"`
	StickerSetName                     string                `json:"sticker_set_name,omitempty"`
	CanSetStickerSet                   bool                  `json:"can_set_sticker_set,omitempty"`
	CustomEmojiStickerSetName          string                `json:"custom_emoji_sticker_set_name,omitempty"`
	LinkedChatID                       int64                 `json:"linked_chat_id,omitempty"`
	Location                           *ChatLocation         `json:"location,omitempty"`
	Rating                             *UserRating           `json:"rating,omitempty"`
	FirstProfileAudio                  *Audio                `json:"first_profile_audio,omitempty"`
	UniqueGiftColors                   *UniqueGiftColors     `json:"unique_gift_colors,omitempty"`
	PaidMessageStarCount               int                   `json:"paid_message_star_count,omitempty"`
	GuardBot                           *User                 `json:"guard_bot,omitempty"`
	Community                          *Community            `json:"community,omitempty"`
}

// https://core.telegram.org/bots/api#chatlocation
type ChatLocation struct {
	Location Location `json:"location"`
	Address  string   `json:"address"`
}

// https://core.telegram.org/bots/api#birthdate
type Birthdate struct {
	Day   int `json:"day"`
	Month int `json:"month"`
	Year  int `json:"year,omitempty"`
}

// https://core.telegram.org/bots/api#businessintro
type BusinessIntro struct {
	Title   string   `json:"title,omitempty"`
	Message string   `json:"message,omitempty"`
	Sticker *Sticker `json:"sticker,omitempty"`
}

// https://core.telegram.org/bots/api#businesslocation
type BusinessLocation struct {
	Address  string    `json:"address"`
	Location *Location `json:"location,omitempty"`
}

// https://core.telegram.org/bots/api#businessopeninghoursinterval
type BusinessOpeningHoursInterval struct {
	OpeningMinute int `json:"opening_minute"`
	ClosingMinute int `json:"closing_minute"`
}

// https://core.telegram.org/bots/api#businessopeninghours
type BusinessOpeningHours struct {
	TimeZoneName string                         `json:"time_zone_name"`
	OpeningHours []BusinessOpeningHoursInterval `json:"opening_hours"`
}

// https://core.telegram.org/bots/api#userrating
type UserRating struct {
	Level              int  `json:"level"`
	Rating             int  `json:"rating"`
	CurrentLevelRating int  `json:"current_level_rating"`
	NextLevelRating    *int `json:"next_level_rating,omitempty"`
}

// https://core.telegram.org/bots/api#businessconnection
type BusinessConnection struct {
	ID         string             `json:"id"`
	User       User               `json:"user"`
	UserChatID int64              `json:"user_chat_id"`
	Date       int64              `json:"date"`
	Rights     *BusinessBotRights `json:"rights,omitempty"`
	IsEnabled  bool               `json:"is_enabled"`
}

// https://core.telegram.org/bots/api#businessbotrights
type BusinessBotRights struct {
	CanReply                   bool `json:"can_reply,omitempty"`
	CanReadMessages            bool `json:"can_read_messages,omitempty"`
	CanDeleteSentMessages      bool `json:"can_delete_sent_messages,omitempty"`
	CanDeleteAllMessages       bool `json:"can_delete_all_messages,omitempty"`
	CanEditName                bool `json:"can_edit_name,omitempty"`
	CanEditBio                 bool `json:"can_edit_bio,omitempty"`
	CanEditProfilePhoto        bool `json:"can_edit_profile_photo,omitempty"`
	CanEditUsername            bool `json:"can_edit_username,omitempty"`
	CanChangeGiftSettings      bool `json:"can_change_gift_settings,omitempty"`
	CanViewGiftsAndStars       bool `json:"can_view_gifts_and_stars,omitempty"`
	CanConvertGiftsToStars     bool `json:"can_convert_gifts_to_stars,omitempty"`
	CanTransferAndUpgradeGifts bool `json:"can_transfer_and_upgrade_gifts,omitempty"`
	CanTransferStars           bool `json:"can_transfer_stars,omitempty"`
	CanManageStories           bool `json:"can_manage_stories,omitempty"`
}

// https://core.telegram.org/bots/api#businessmessagesdeleted
type BusinessMessagesDeleted struct {
	BusinessConnectionID string `json:"business_connection_id"`
	Chat                 Chat   `json:"chat"`
	MessageIDs           []int  `json:"message_ids"`
}

// https://core.telegram.org/bots/api#community
type Community struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// https://core.telegram.org/bots/api#chatmemberupdated
type ChatMemberUpdated struct {
	Chat                    Chat            `json:"chat"`
	From                    User            `json:"from"`
	Date                    int64           `json:"date"`
	OldChatMember           ChatMember      `json:"old_chat_member"`
	NewChatMember           ChatMember      `json:"new_chat_member"`
	InviteLink              *ChatInviteLink `json:"invite_link,omitempty"`
	ViaJoinRequest          bool            `json:"via_join_request,omitempty"`
	ViaChatFolderInviteLink bool            `json:"via_chat_folder_invite_link,omitempty"`
}

// https://core.telegram.org/bots/api#chatjoinrequest
type ChatJoinRequest struct {
	Chat       Chat            `json:"chat"`
	From       User            `json:"from"`
	UserChatID int64           `json:"user_chat_id"`
	Date       int64           `json:"date"`
	Bio        string          `json:"bio,omitempty"`
	InviteLink *ChatInviteLink `json:"invite_link,omitempty"`
	QueryID    string          `json:"query_id,omitempty"`
}

// https://core.telegram.org/bots/api#chatboost
type ChatBoost struct {
	BoostID        string          `json:"boost_id"`
	AddDate        int64           `json:"add_date"`
	ExpirationDate int64           `json:"expiration_date"`
	Source         ChatBoostSource `json:"source"`
}

// https://core.telegram.org/bots/api#chatboostsource
type ChatBoostSource interface{}

// https://core.telegram.org/bots/api#chatboostupdated
type ChatBoostUpdated struct {
	Chat  Chat      `json:"chat"`
	Boost ChatBoost `json:"boost"`
}

// https://core.telegram.org/bots/api#chatboostremoved
type ChatBoostRemoved struct {
	Chat       Chat            `json:"chat"`
	BoostID    string          `json:"boost_id"`
	RemoveDate int64           `json:"remove_date"`
	Source     ChatBoostSource `json:"source"`
}

// https://core.telegram.org/bots/api#userchatboosts
type UserChatBoosts struct {
	Boosts []ChatBoost `json:"boosts"`
}

// https://core.telegram.org/bots/api#messagereactionupdated
type MessageReactionUpdated struct {
	Chat        Chat           `json:"chat"`
	MessageID   int            `json:"message_id"`
	User        *User          `json:"user,omitempty"`
	ActorChat   *Chat          `json:"actor_chat,omitempty"`
	Date        int64          `json:"date"`
	OldReaction []ReactionType `json:"old_reaction"`
	NewReaction []ReactionType `json:"new_reaction"`
}

// https://core.telegram.org/bots/api#messagereactioncountupdated
type MessageReactionCountUpdated struct {
	Chat      Chat            `json:"chat"`
	MessageID int             `json:"message_id"`
	Date      int64           `json:"date"`
	Reactions []ReactionCount `json:"reactions"`
}

// https://core.telegram.org/bots/api#reactioncount
type ReactionCount struct {
	Type       ReactionType `json:"type"`
	TotalCount int          `json:"total_count"`
}

// https://core.telegram.org/bots/api#reactiontype
type ReactionType struct {
	Type          string `json:"type"`
	Emoji         string `json:"emoji,omitempty"`
	CustomEmojiID string `json:"custom_emoji_id,omitempty"`
}

// https://core.telegram.org/bots/api#managedbotupdated
type ManagedBotUpdated struct {
	User User `json:"user"`
	Bot  User `json:"bot"`
}

// https://core.telegram.org/bots/api#botsubscriptionupdated
type BotSubscriptionUpdated struct {
	User           User   `json:"user"`
	InvoicePayload string `json:"invoice_payload"`
	State          string `json:"state"`
}

// https://core.telegram.org/bots/api#paidmediapurchased
type PaidMediaPurchased struct {
	From             User   `json:"from"`
	PaidMediaPayload string `json:"paid_media_payload"`
}

// https://core.telegram.org/bots/api#acceptedgifttypes
type AcceptedGiftTypes struct {
	UnlimitedGifts      bool `json:"unlimited_gifts"`
	LimitedGifts        bool `json:"limited_gifts"`
	UniqueGifts         bool `json:"unique_gifts"`
	PremiumSubscription bool `json:"premium_subscription"`
	GiftsFromChannels   bool `json:"gifts_from_channels"`
}

// https://core.telegram.org/bots/api#staramount
type StarAmount struct {
	Amount         int  `json:"amount"`
	NanostarAmount *int `json:"nanostar_amount,omitempty"`
}

// https://core.telegram.org/bots/api#approvedchatjoinrequest
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

// https://core.telegram.org/bots/api#getchatadministrators
type GetChatAdministratorsParams struct {
	ChatID     interface{} `json:"chat_id"`
	ReturnBots bool        `json:"return_bots,omitempty"`
}
