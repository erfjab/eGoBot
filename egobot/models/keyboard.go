package models

// https://core.telegram.org/bots/api#inlinekeyboardmarkup
type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

// https://core.telegram.org/bots/api#inlinekeyboardbutton
type InlineKeyboardButton struct {
	Text                         string        `json:"text"`
	URL                          string        `json:"url,omitempty"`
	CallbackData                 string        `json:"callback_data,omitempty"`
	WebApp                       *WebAppInfo   `json:"web_app,omitempty"`
	LoginURL                     *LoginURL     `json:"login_url,omitempty"`
	SwitchInlineQuery            string        `json:"switch_inline_query,omitempty"`
	SwitchInlineQueryCurrentChat string        `json:"switch_inline_query_current_chat,omitempty"`
	CallbackGame                 *CallbackGame `json:"callback_game,omitempty"`
	Pay                          bool          `json:"pay,omitempty"`
	IconCustomEmojiID            string        `json:"icon_custom_emoji_id,omitempty"`
	Style                        string        `json:"style,omitempty"`
}

// https://core.telegram.org/bots/api#replykeyboardmarkup
type ReplyKeyboardMarkup struct {
	Keyboard              [][]KeyboardButton `json:"keyboard"`
	IsPersistent          bool               `json:"is_persistent,omitempty"`
	ResizeKeyboard        bool               `json:"resize_keyboard,omitempty"`
	OneTimeKeyboard       bool               `json:"one_time_keyboard,omitempty"`
	InputFieldPlaceholder string             `json:"input_field_placeholder,omitempty"`
	Selective             bool               `json:"selective,omitempty"`
}

// https://core.telegram.org/bots/api#keyboardbutton
type KeyboardButton struct {
	Text              string                      `json:"text"`
	RequestUsers      *KeyboardButtonRequestUsers `json:"request_users,omitempty"`
	RequestChat       *KeyboardButtonRequestChat  `json:"request_chat,omitempty"`
	RequestContact    bool                        `json:"request_contact,omitempty"`
	RequestLocation   bool                        `json:"request_location,omitempty"`
	RequestPoll       *KeyboardButtonPollType     `json:"request_poll,omitempty"`
	WebApp            *WebAppInfo                 `json:"web_app,omitempty"`
	IconCustomEmojiID string                      `json:"icon_custom_emoji_id,omitempty"`
	Style             string                      `json:"style,omitempty"`
}

// https://core.telegram.org/bots/api#keyboardbuttonrequestusers
type KeyboardButtonRequestUsers struct {
	RequestID     int  `json:"request_id"`
	UserIsBot     bool `json:"user_is_bot,omitempty"`
	UserIsPremium bool `json:"user_is_premium,omitempty"`
	MaxQuantity   int  `json:"max_quantity,omitempty"`
}

// https://core.telegram.org/bots/api#keyboardbuttonrequestchat
type KeyboardButtonRequestChat struct {
	RequestID               int                     `json:"request_id"`
	ChatIsChannel           bool                    `json:"chat_is_channel"`
	ChatIsForum             bool                    `json:"chat_is_forum,omitempty"`
	ChatHasUsername         bool                    `json:"chat_has_username,omitempty"`
	ChatIsCreated           bool                    `json:"chat_is_created,omitempty"`
	UserAdministratorRights *ChatAdministratorRights `json:"user_administrator_rights,omitempty"`
	BotAdministratorRights  *ChatAdministratorRights `json:"bot_administrator_rights,omitempty"`
	BotIsMember             bool                    `json:"bot_is_member,omitempty"`
}

// https://core.telegram.org/bots/api#chatadministratorrights
type ChatAdministratorRights struct {
	IsAnonymous         bool `json:"is_anonymous"`
	CanManageChat       bool `json:"can_manage_chat"`
	CanDeleteMessages   bool `json:"can_delete_messages"`
	CanManageVideoChats bool `json:"can_manage_video_chats"`
	CanRestrictMembers  bool `json:"can_restrict_members"`
	CanPromoteMembers   bool `json:"can_promote_members"`
	CanChangeInfo       bool `json:"can_change_info"`
	CanInviteUsers      bool `json:"can_invite_users"`
	CanPostMessages     bool `json:"can_post_messages,omitempty"`
	CanEditMessages     bool `json:"can_edit_messages,omitempty"`
	CanPinMessages      bool `json:"can_pin_messages,omitempty"`
	CanManageTopics     bool `json:"can_manage_topics,omitempty"`
}

// https://core.telegram.org/bots/api#keyboardbuttonpolltype
type KeyboardButtonPollType struct {
	Type string `json:"type,omitempty"`
}

// https://core.telegram.org/bots/api#replykeyboardremove
type ReplyKeyboardRemove struct {
	RemoveKeyboard bool `json:"remove_keyboard"`
	Selective      bool `json:"selective,omitempty"`
}

// https://core.telegram.org/bots/api#forcereply
type ForceReply struct {
	ForceReply            bool   `json:"force_reply"`
	InputFieldPlaceholder string `json:"input_field_placeholder,omitempty"`
	Selective             bool   `json:"selective,omitempty"`
}

// https://core.telegram.org/bots/api#webappinfo
type WebAppInfo struct {
	URL string `json:"url"`
}

// https://core.telegram.org/bots/api#loginurl
type LoginURL struct {
	URL                string `json:"url"`
	ForwardText        string `json:"forward_text,omitempty"`
	BotUsername        string `json:"bot_username,omitempty"`
	RequestWriteAccess bool   `json:"request_write_access,omitempty"`
}

// https://core.telegram.org/bots/api#callbackgame
type CallbackGame struct{}
