package models

// https://core.telegram.org/bots/api#user
type User struct {
	ID                         int64  `json:"id"`
	IsBot                      bool   `json:"is_bot"`
	FirstName                  string `json:"first_name"`
	LastName                   string `json:"last_name,omitempty"`
	Username                   string `json:"username,omitempty"`
	LanguageCode               string `json:"language_code,omitempty"`
	CanJoinGroups              bool   `json:"can_join_groups,omitempty"`
	CanReadAllGroupMessages    bool   `json:"can_read_all_group_messages,omitempty"`
	SupportsInlineQueries      bool   `json:"supports_inline_queries,omitempty"`
	HasTopicsEnabled           bool   `json:"has_topics_enabled,omitempty"`
	AllowsUsersToCreateTopics  bool   `json:"allows_users_to_create_topics,omitempty"`
}

// https://core.telegram.org/bots/api#chatphoto
type ChatPhoto struct {
	SmallFileID       string `json:"small_file_id"`
	SmallFileUniqueID string `json:"small_file_unique_id"`
	BigFileID         string `json:"big_file_id"`
	BigFileUniqueID   string `json:"big_file_unique_id"`
}

// https://core.telegram.org/bots/api#message
type Message struct {
	MessageID                     int64                  `json:"message_id"`
	MessageThreadID               int64                  `json:"message_thread_id,omitempty"`
	From                          *User                  `json:"from,omitempty"`
	SenderChat                    *Chat                  `json:"sender_chat,omitempty"`
	Date                          int64                  `json:"date"`
	Chat                          Chat                   `json:"chat"`
	ForwardFrom                   *User                  `json:"forward_from,omitempty"`
	ForwardFromChat               *Chat                  `json:"forward_from_chat,omitempty"`
	ForwardFromMessageID          int64                  `json:"forward_from_message_id,omitempty"`
	ForwardSignature              string                 `json:"forward_signature,omitempty"`
	ForwardSenderName             string                 `json:"forward_sender_name,omitempty"`
	ForwardDate                   int64                  `json:"forward_date,omitempty"`
	IsTopicMessage                bool                   `json:"is_topic_message,omitempty"`
	IsAutomaticForward            bool                   `json:"is_automatic_forward,omitempty"`
	ReplyToMessage                *Message               `json:"reply_to_message,omitempty"`
	ViaBot                        *User                  `json:"via_bot,omitempty"`
	EditDate                      int64                  `json:"edit_date,omitempty"`
	HasProtectedContent           bool                   `json:"has_protected_content,omitempty"`
	MediaGroupID                  string                 `json:"media_group_id,omitempty"`
	AuthorSignature               string                 `json:"author_signature,omitempty"`
	Text                          string                 `json:"text,omitempty"`
	Entities                      []MessageEntity        `json:"entities,omitempty"`
	Animation                     *Animation             `json:"animation,omitempty"`
	Audio                         *Audio                 `json:"audio,omitempty"`
	Document                      *Document              `json:"document,omitempty"`
	Photo                         []PhotoSize            `json:"photo,omitempty"`
	Sticker                       *Sticker               `json:"sticker,omitempty"`
	Video                         *Video                 `json:"video,omitempty"`
	VideoNote                     *VideoNote             `json:"video_note,omitempty"`
	Voice                         *Voice                 `json:"voice,omitempty"`
	Caption                       string                 `json:"caption,omitempty"`
	CaptionEntities               []MessageEntity        `json:"caption_entities,omitempty"`
	Contact                       *Contact               `json:"contact,omitempty"`
	Location                      *Location              `json:"location,omitempty"`
	NewChatMembers                []User                         `json:"new_chat_members,omitempty"`
	LeftChatMember                *User                          `json:"left_chat_member,omitempty"`
	NewChatTitle                  string                         `json:"new_chat_title,omitempty"`
	NewChatPhoto                  []PhotoSize                    `json:"new_chat_photo,omitempty"`
	DeleteChatPhoto               bool                           `json:"delete_chat_photo,omitempty"`
	GroupChatCreated              bool                           `json:"group_chat_created,omitempty"`
	SupergroupChatCreated         bool                           `json:"supergroup_chat_created,omitempty"`
	ChannelChatCreated            bool                           `json:"channel_chat_created,omitempty"`
	MessageAutoDeleteTimerChanged *MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed,omitempty"`
	MigrateToChatID               int64                          `json:"migrate_to_chat_id,omitempty"`
	MigrateFromChatID             int64                          `json:"migrate_from_chat_id,omitempty"`
	PinnedMessage                 *Message                       `json:"pinned_message,omitempty"`
	ReplyMarkup                   *InlineKeyboardMarkup          `json:"reply_markup,omitempty"`
	
	ReplyToChecklistTaskID int64                 `json:"reply_to_checklist_task_id,omitempty"`
	DirectMessagesTopic    *DirectMessagesTopic  `json:"direct_messages_topic,omitempty"`
	
	GiftUpgradeSent *GiftInfo `json:"gift_upgrade_sent,omitempty"`
	
	ChatOwnerLeft    *ChatOwnerLeft    `json:"chat_owner_left,omitempty"`
	ChatOwnerChanged *ChatOwnerChanged `json:"chat_owner_changed,omitempty"`
}

// https://core.telegram.org/bots/api#messageautodeletetimerchanged
type MessageAutoDeleteTimerChanged struct {
	MessageAutoDeleteTime int `json:"message_auto_delete_time"`
}

// https://core.telegram.org/bots/api#chatownerleft
type ChatOwnerLeft struct {
	NewOwner *User `json:"new_owner,omitempty"`
}

// https://core.telegram.org/bots/api#chatownerchanged
type ChatOwnerChanged struct {
	NewOwner User `json:"new_owner"`
}

// https://core.telegram.org/bots/api#giftinfo
type GiftInfo struct {
}

// https://core.telegram.org/bots/api#messageentity
type MessageEntity struct {
	Type          string `json:"type"`
	Offset        int    `json:"offset"`
	Length        int    `json:"length"`
	URL           string `json:"url,omitempty"`
	User          *User  `json:"user,omitempty"`
	Language      string `json:"language,omitempty"`
	CustomEmojiID string `json:"custom_emoji_id,omitempty"`
}

// https://core.telegram.org/bots/api#directmessagestopic
type DirectMessagesTopic struct {
	ID int `json:"id"`
}

// https://core.telegram.org/bots/api#making-requests
type Response struct {
	Ok          bool        `json:"ok"`
	Result      interface{} `json:"result,omitempty"`
	ErrorCode   int         `json:"error_code,omitempty"`
	Description string      `json:"description,omitempty"`
}
