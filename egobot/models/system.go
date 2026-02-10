package models

// https://core.telegram.org/bots/api#update
type Update struct {
	UpdateID           int64               `json:"update_id"`
	Message            *Message            `json:"message,omitempty"`
	EditedMessage      *Message            `json:"edited_message,omitempty"`
	ChannelPost        *Message            `json:"channel_post,omitempty"`
	EditedChannelPost  *Message            `json:"edited_channel_post,omitempty"`
	CallbackQuery      *CallbackQuery      `json:"callback_query,omitempty"`
	InlineQuery        *InlineQuery        `json:"inline_query,omitempty"`
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result,omitempty"`
}

// https://core.telegram.org/bots/api#getupdates
type GetUpdatesParams struct {
	Offset         int64    `json:"offset,omitempty"`
	Limit          int      `json:"limit,omitempty"`
	Timeout        int      `json:"timeout,omitempty"`
	AllowedUpdates []string `json:"allowed_updates,omitempty"`
}

// https://core.telegram.org/bots/api#webhookinfo
type WebhookInfo struct {
	URL                          string   `json:"url"`
	HasCustomCertificate         bool     `json:"has_custom_certificate"`
	PendingUpdateCount           int      `json:"pending_update_count"`
	IPAddress                    string   `json:"ip_address,omitempty"`
	LastErrorDate                int64    `json:"last_error_date,omitempty"`
	LastErrorMessage             string   `json:"last_error_message,omitempty"`
	LastSynchronizationErrorDate int64    `json:"last_synchronization_error_date,omitempty"`
	MaxConnections               int      `json:"max_connections,omitempty"`
	AllowedUpdates               []string `json:"allowed_updates,omitempty"`
}

// https://core.telegram.org/bots/api#setwebhook
type SetWebhookParams struct {
	URL                string      `json:"url"`
	Certificate        interface{} `json:"certificate,omitempty"`
	IPAddress          string      `json:"ip_address,omitempty"`
	MaxConnections     int         `json:"max_connections,omitempty"`
	AllowedUpdates     []string    `json:"allowed_updates,omitempty"`
	DropPendingUpdates bool        `json:"drop_pending_updates,omitempty"`
	SecretToken        string      `json:"secret_token,omitempty"`
}

// https://core.telegram.org/bots/api#deletewebhook
type DeleteWebhookParams struct {
	DropPendingUpdates bool `json:"drop_pending_updates,omitempty"`
}

// https://core.telegram.org/bots/api#inputprofilephoto
type InputProfilePhoto interface{}

// https://core.telegram.org/bots/api#inputprofilephotostatic
type InputProfilePhotoStatic struct {
	Type  string      `json:"type"`  // "static"
	Photo interface{} `json:"photo"` // File to upload
}

// https://core.telegram.org/bots/api#inputprofilephotoanimated
type InputProfilePhotoAnimated struct {
	Type      string      `json:"type"`       // "animated"
	Video     interface{} `json:"video"`      // Video file to upload
	Thumbnail interface{} `json:"thumbnail,omitempty"`
}

// https://core.telegram.org/bots/api#userprofileaudios
type UserProfileAudios struct {
	TotalCount int     `json:"total_count"`
	Audios     []Audio `json:"audios"`
}

// https://core.telegram.org/bots/api#getuserprofileaudios
type GetUserProfileAudiosParams struct {
	UserID int64 `json:"user_id"`
	Offset int   `json:"offset,omitempty"`
	Limit  int   `json:"limit,omitempty"` // 1-100, default 100
}

// https://core.telegram.org/bots/api#botcommand
type BotCommand struct {
	Command     string `json:"command"`
	Description string `json:"description"`
}

// https://core.telegram.org/bots/api#setmycommands
type SetMyCommandsParams struct {
	Commands     []BotCommand `json:"commands"`
	Scope        interface{}  `json:"scope,omitempty"`
	LanguageCode string       `json:"language_code,omitempty"`
}

// https://core.telegram.org/bots/api#deletemycommands
type DeleteMyCommandsParams struct {
	Scope        interface{} `json:"scope,omitempty"`
	LanguageCode string      `json:"language_code,omitempty"`
}

// https://core.telegram.org/bots/api#getmycommands
type GetMyCommandsParams struct {
	Scope        interface{} `json:"scope,omitempty"`
	LanguageCode string      `json:"language_code,omitempty"`
}

// https://core.telegram.org/bots/api#setmyname
type SetMyNameParams struct {
	Name         string `json:"name,omitempty"`
	LanguageCode string `json:"language_code,omitempty"`
}

// https://core.telegram.org/bots/api#getmyname
type GetMyNameParams struct {
	LanguageCode string `json:"language_code,omitempty"`
}

// https://core.telegram.org/bots/api#botname
type BotName struct {
	Name string `json:"name"`
}

// https://core.telegram.org/bots/api#setmydescription
type SetMyDescriptionParams struct {
	Description  string `json:"description,omitempty"`
	LanguageCode string `json:"language_code,omitempty"`
}

// https://core.telegram.org/bots/api#getmydescription
type GetMyDescriptionParams struct {
	LanguageCode string `json:"language_code,omitempty"`
}

// https://core.telegram.org/bots/api#botdescription
type BotDescription struct {
	Description string `json:"description"`
}

// https://core.telegram.org/bots/api#setmyshortdescription
type SetMyShortDescriptionParams struct {
	ShortDescription string `json:"short_description,omitempty"`
	LanguageCode     string `json:"language_code,omitempty"`
}

// https://core.telegram.org/bots/api#getmyshortdescription
type GetMyShortDescriptionParams struct {
	LanguageCode string `json:"language_code,omitempty"`
}

// https://core.telegram.org/bots/api#botshortdescription
type BotShortDescription struct {
	ShortDescription string `json:"short_description"`
}

// https://core.telegram.org/bots/api#setchatmenubutton
type SetChatMenuButtonParams struct {
	ChatID     int64       `json:"chat_id,omitempty"`
	MenuButton interface{} `json:"menu_button,omitempty"`
}

// https://core.telegram.org/bots/api#getchatmenubutton
type GetChatMenuButtonParams struct {
	ChatID int64 `json:"chat_id,omitempty"`
}

// https://core.telegram.org/bots/api#menubutton
type MenuButton struct {
	Type   string      `json:"type"`
	Text   string      `json:"text,omitempty"`
	WebApp *WebAppInfo `json:"web_app,omitempty"`
}

// https://core.telegram.org/bots/api#userprofilephotos
type UserProfilePhotos struct {
	TotalCount int           `json:"total_count"`
	Photos     [][]PhotoSize `json:"photos"`
}

// https://core.telegram.org/bots/api#getuserprofilephotos
type GetUserProfilePhotosParams struct {
	UserID int64 `json:"user_id"`
	Offset int   `json:"offset,omitempty"`
	Limit  int   `json:"limit,omitempty"`
}

// https://core.telegram.org/bots/api#setmessagereaction
type SetMessageReactionParams struct {
	ChatID    interface{}   `json:"chat_id"`
	MessageID int           `json:"message_id"`
	Reaction  []interface{} `json:"reaction,omitempty"`
	IsBig     bool          `json:"is_big,omitempty"`
}

// https://core.telegram.org/bots/api#reactiontype
type ReactionType struct {
	Type          string `json:"type"`
	Emoji         string `json:"emoji,omitempty"`
	CustomEmojiID string `json:"custom_emoji_id,omitempty"`
}
