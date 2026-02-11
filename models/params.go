package models

// https://core.telegram.org/bots/api#sendmessage
type SendMessageParams struct {
	ChatID                   interface{} `json:"chat_id"`
	MessageThreadID          int64       `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID    int64       `json:"direct_messages_topic_id,omitempty"`
	Text                     string      `json:"text"`
	ParseMode                string      `json:"parse_mode,omitempty"`
	Entities                 []MessageEntity `json:"entities,omitempty"`
	DisableWebPagePreview    bool        `json:"disable_web_page_preview,omitempty"`
	DisableNotification      bool        `json:"disable_notification,omitempty"`
	ProtectContent           bool        `json:"protect_content,omitempty"`
	ReplyToMessageID         int64       `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool        `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              interface{} `json:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#sendphoto
type SendPhotoParams struct {
	ChatID                   interface{} `json:"chat_id"`
	MessageThreadID          int64       `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID    int64       `json:"direct_messages_topic_id,omitempty"`
	Photo                    interface{} `json:"photo"`
	Caption                  string      `json:"caption,omitempty"`
	ParseMode                string      `json:"parse_mode,omitempty"`
	CaptionEntities          []MessageEntity `json:"caption_entities,omitempty"`
	HasSpoiler               bool        `json:"has_spoiler,omitempty"`
	DisableNotification      bool        `json:"disable_notification,omitempty"`
	ProtectContent           bool        `json:"protect_content,omitempty"`
	ReplyToMessageID         int64       `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool        `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              interface{} `json:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#senddocument
type SendDocumentParams struct {
	ChatID                      interface{} `json:"chat_id"`
	MessageThreadID             int64       `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID       int64       `json:"direct_messages_topic_id,omitempty"`
	Document                    interface{} `json:"document"`
	Thumbnail                   interface{} `json:"thumbnail,omitempty"`
	Caption                     string      `json:"caption,omitempty"`
	ParseMode                   string      `json:"parse_mode,omitempty"`
	CaptionEntities             []MessageEntity `json:"caption_entities,omitempty"`
	DisableContentTypeDetection bool        `json:"disable_content_type_detection,omitempty"`
	DisableNotification         bool        `json:"disable_notification,omitempty"`
	ProtectContent              bool        `json:"protect_content,omitempty"`
	ReplyToMessageID            int64       `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply    bool        `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup                 interface{} `json:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#sendvideo
type SendVideoParams struct {
	ChatID                   interface{} `json:"chat_id"`
	MessageThreadID          int64       `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID    int64       `json:"direct_messages_topic_id,omitempty"`
	Video                    interface{} `json:"video"`
	Duration                 int         `json:"duration,omitempty"`
	Width                    int         `json:"width,omitempty"`
	Height                   int         `json:"height,omitempty"`
	Thumbnail                interface{} `json:"thumbnail,omitempty"`
	Caption                  string      `json:"caption,omitempty"`
	ParseMode                string      `json:"parse_mode,omitempty"`
	CaptionEntities          []MessageEntity `json:"caption_entities,omitempty"`
	HasSpoiler               bool        `json:"has_spoiler,omitempty"`
	SupportsStreaming        bool        `json:"supports_streaming,omitempty"`
	DisableNotification      bool        `json:"disable_notification,omitempty"`
	ProtectContent           bool        `json:"protect_content,omitempty"`
	ReplyToMessageID         int64       `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool        `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              interface{} `json:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#sendaudio
type SendAudioParams struct {
	ChatID                   interface{} `json:"chat_id"`
	MessageThreadID          int64       `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID    int64       `json:"direct_messages_topic_id,omitempty"`
	Audio                    interface{} `json:"audio"`
	Caption                  string      `json:"caption,omitempty"`
	ParseMode                string      `json:"parse_mode,omitempty"`
	CaptionEntities          []MessageEntity `json:"caption_entities,omitempty"`
	Duration                 int         `json:"duration,omitempty"`
	Performer                string      `json:"performer,omitempty"`
	Title                    string      `json:"title,omitempty"`
	Thumbnail                interface{} `json:"thumbnail,omitempty"`
	DisableNotification      bool        `json:"disable_notification,omitempty"`
	ProtectContent           bool        `json:"protect_content,omitempty"`
	ReplyToMessageID         int64       `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool        `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              interface{} `json:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#sendanimation
type SendAnimationParams struct {
	BusinessConnectionID     string                `json:"business_connection_id,omitempty"`
	ChatID                   interface{}           `json:"chat_id"`
	MessageThreadID          int                   `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID    int                   `json:"direct_messages_topic_id,omitempty"`
	Animation                interface{}           `json:"animation"`
	Duration                 int                   `json:"duration,omitempty"`
	Width                    int                   `json:"width,omitempty"`
	Height                   int                   `json:"height,omitempty"`
	Thumbnail                interface{}           `json:"thumbnail,omitempty"`
	Caption                  string                `json:"caption,omitempty"`
	ParseMode                string                `json:"parse_mode,omitempty"`
	CaptionEntities          []MessageEntity       `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia    bool                  `json:"show_caption_above_media,omitempty"`
	HasSpoiler               bool                  `json:"has_spoiler,omitempty"`
	DisableNotification      bool                  `json:"disable_notification,omitempty"`
	ProtectContent           bool                  `json:"protect_content,omitempty"`
	AllowPaidBroadcast       bool                  `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID          string                `json:"message_effect_id,omitempty"`
	SuggestedPostParameters  *SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters          *ReplyParameters      `json:"reply_parameters,omitempty"`
	ReplyMarkup              interface{}           `json:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#sendvoice
type SendVoiceParams struct {
	BusinessConnectionID    string                `json:"business_connection_id,omitempty"`
	ChatID                  interface{}           `json:"chat_id"`
	MessageThreadID         int                   `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID   int                   `json:"direct_messages_topic_id,omitempty"`
	Voice                   interface{}           `json:"voice"`
	Caption                 string                `json:"caption,omitempty"`
	ParseMode               string                `json:"parse_mode,omitempty"`
	CaptionEntities         []MessageEntity       `json:"caption_entities,omitempty"`
	Duration                int                   `json:"duration,omitempty"`
	DisableNotification     bool                  `json:"disable_notification,omitempty"`
	ProtectContent          bool                  `json:"protect_content,omitempty"`
	AllowPaidBroadcast      bool                  `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID         string                `json:"message_effect_id,omitempty"`
	SuggestedPostParameters *SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters         *ReplyParameters      `json:"reply_parameters,omitempty"`
	ReplyMarkup             interface{}           `json:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#sendvideonote
type SendVideoNoteParams struct {
	BusinessConnectionID    string                `json:"business_connection_id,omitempty"`
	ChatID                  interface{}           `json:"chat_id"`
	MessageThreadID         int                   `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID   int                   `json:"direct_messages_topic_id,omitempty"`
	VideoNote               interface{}           `json:"video_note"`
	Duration                int                   `json:"duration,omitempty"`
	Length                  int                   `json:"length,omitempty"`
	Thumbnail               interface{}           `json:"thumbnail,omitempty"`
	DisableNotification     bool                  `json:"disable_notification,omitempty"`
	ProtectContent          bool                  `json:"protect_content,omitempty"`
	AllowPaidBroadcast      bool                  `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID         string                `json:"message_effect_id,omitempty"`
	SuggestedPostParameters *SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters         *ReplyParameters      `json:"reply_parameters,omitempty"`
	ReplyMarkup             interface{}           `json:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#sendvenue
type SendVenueParams struct {
	BusinessConnectionID    string                `json:"business_connection_id,omitempty"`
	ChatID                  interface{}           `json:"chat_id"`
	MessageThreadID         int                   `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID   int                   `json:"direct_messages_topic_id,omitempty"`
	Latitude                float64               `json:"latitude"`
	Longitude               float64               `json:"longitude"`
	Title                   string                `json:"title"`
	Address                 string                `json:"address"`
	FoursquareID            string                `json:"foursquare_id,omitempty"`
	FoursquareType          string                `json:"foursquare_type,omitempty"`
	GooglePlaceID           string                `json:"google_place_id,omitempty"`
	GooglePlaceType         string                `json:"google_place_type,omitempty"`
	DisableNotification     bool                  `json:"disable_notification,omitempty"`
	ProtectContent          bool                  `json:"protect_content,omitempty"`
	AllowPaidBroadcast      bool                  `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID         string                `json:"message_effect_id,omitempty"`
	SuggestedPostParameters *SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters         *ReplyParameters      `json:"reply_parameters,omitempty"`
	ReplyMarkup             interface{}           `json:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#senddice
type SendDiceParams struct {
	BusinessConnectionID    string                `json:"business_connection_id,omitempty"`
	ChatID                  interface{}           `json:"chat_id"`
	MessageThreadID         int                   `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID   int                   `json:"direct_messages_topic_id,omitempty"`
	Emoji                   string                `json:"emoji,omitempty"`
	DisableNotification     bool                  `json:"disable_notification,omitempty"`
	ProtectContent          bool                  `json:"protect_content,omitempty"`
	AllowPaidBroadcast      bool                  `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID         string                `json:"message_effect_id,omitempty"`
	SuggestedPostParameters *SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters         *ReplyParameters      `json:"reply_parameters,omitempty"`
	ReplyMarkup             interface{}           `json:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#sendlocation
type SendLocationParams struct {
	ChatID                   interface{} `json:"chat_id"`
	MessageThreadID          int64       `json:"message_thread_id,omitempty"`
	Latitude                 float64     `json:"latitude"`
	Longitude                float64     `json:"longitude"`
	HorizontalAccuracy       float64     `json:"horizontal_accuracy,omitempty"`
	LivePeriod               int         `json:"live_period,omitempty"`
	Heading                  int         `json:"heading,omitempty"`
	ProximityAlertRadius     int         `json:"proximity_alert_radius,omitempty"`
	DisableNotification      bool        `json:"disable_notification,omitempty"`
	ProtectContent           bool        `json:"protect_content,omitempty"`
	ReplyToMessageID         int64       `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool        `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              interface{} `json:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#sendcontact
type SendContactParams struct {
	ChatID                   interface{} `json:"chat_id"`
	MessageThreadID          int64       `json:"message_thread_id,omitempty"`
	PhoneNumber              string      `json:"phone_number"`
	FirstName                string      `json:"first_name"`
	LastName                 string      `json:"last_name,omitempty"`
	VCard                    string      `json:"vcard,omitempty"`
	DisableNotification      bool        `json:"disable_notification,omitempty"`
	ProtectContent           bool        `json:"protect_content,omitempty"`
	ReplyToMessageID         int64       `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool        `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              interface{} `json:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#sendpoll
type SendPollParams struct {
	ChatID                   interface{} `json:"chat_id"`
	MessageThreadID          int64       `json:"message_thread_id,omitempty"`
	Question                 string      `json:"question"`
	Options                  []string    `json:"options"`
	IsAnonymous              bool        `json:"is_anonymous,omitempty"`
	Type                     string      `json:"type,omitempty"`
	AllowsMultipleAnswers    bool        `json:"allows_multiple_answers,omitempty"`
	CorrectOptionID          int         `json:"correct_option_id,omitempty"`
	Explanation              string      `json:"explanation,omitempty"`
	ExplanationParseMode     string      `json:"explanation_parse_mode,omitempty"`
	ExplanationEntities      []MessageEntity `json:"explanation_entities,omitempty"`
	OpenPeriod               int         `json:"open_period,omitempty"`
	CloseDate                int64       `json:"close_date,omitempty"`
	IsClosed                 bool        `json:"is_closed,omitempty"`
	DisableNotification      bool        `json:"disable_notification,omitempty"`
	ProtectContent           bool        `json:"protect_content,omitempty"`
	ReplyToMessageID         int64       `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool        `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              interface{} `json:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#sendchecklist
type SendChecklistParams struct {
	BusinessConnectionID string             `json:"business_connection_id"`
	ChatID               int64              `json:"chat_id"`
	Checklist            InputChecklist     `json:"checklist"`
	DisableNotification  bool               `json:"disable_notification,omitempty"`
	ProtectContent       bool               `json:"protect_content,omitempty"`
	MessageEffectID      string             `json:"message_effect_id,omitempty"`
	ReplyParameters      *ReplyParameters   `json:"reply_parameters,omitempty"`
	ReplyMarkup          *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#forwardmessage
type ForwardMessageParams struct {
	ChatID              interface{} `json:"chat_id"`
	MessageThreadID     int64       `json:"message_thread_id,omitempty"`
	FromChatID          interface{} `json:"from_chat_id"`
	DisableNotification bool        `json:"disable_notification,omitempty"`
	ProtectContent      bool        `json:"protect_content,omitempty"`
	MessageID           int64       `json:"message_id"`
}

// https://core.telegram.org/bots/api#copymessage
type CopyMessageParams struct {
	ChatID              interface{} `json:"chat_id"`
	MessageThreadID     int64       `json:"message_thread_id,omitempty"`
	FromChatID          interface{} `json:"from_chat_id"`
	MessageID           int64       `json:"message_id"`
	Caption             string      `json:"caption,omitempty"`
	ParseMode           string      `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity `json:"caption_entities,omitempty"`
	DisableNotification bool        `json:"disable_notification,omitempty"`
	ProtectContent      bool        `json:"protect_content,omitempty"`
	ReplyToMessageID    int64       `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool   `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup         interface{} `json:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#messageid
type MessageID struct {
	MessageID int64 `json:"message_id"`
}

// https://core.telegram.org/bots/api#editmessagetext
type EditMessageTextParams struct {
	ChatID                interface{} `json:"chat_id,omitempty"`
	MessageID             int64       `json:"message_id,omitempty"`
	InlineMessageID       string      `json:"inline_message_id,omitempty"`
	Text                  string      `json:"text"`
	ParseMode             string      `json:"parse_mode,omitempty"`
	Entities              []MessageEntity `json:"entities,omitempty"`
	DisableWebPagePreview bool        `json:"disable_web_page_preview,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#deletemessage
type DeleteMessageParams struct {
	ChatID    interface{} `json:"chat_id"`
	MessageID int64       `json:"message_id"`
}

// https://core.telegram.org/bots/api#editmessagecaption
type EditMessageCaptionParams struct {
	BusinessConnectionID  string               `json:"business_connection_id,omitempty"`
	ChatID                interface{}          `json:"chat_id,omitempty"`
	MessageID             int                  `json:"message_id,omitempty"`
	InlineMessageID       string               `json:"inline_message_id,omitempty"`
	Caption               string               `json:"caption,omitempty"`
	ParseMode             string               `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity      `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool                 `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#editmessagemedia
type EditMessageMediaParams struct {
	BusinessConnectionID string                `json:"business_connection_id,omitempty"`
	ChatID               interface{}           `json:"chat_id,omitempty"`
	MessageID            int                   `json:"message_id,omitempty"`
	InlineMessageID      string                `json:"inline_message_id,omitempty"`
	Media                interface{}           `json:"media"`
	ReplyMarkup          *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#editmessagereplymarkup
type EditMessageReplyMarkupParams struct {
	BusinessConnectionID string                `json:"business_connection_id,omitempty"`
	ChatID               interface{}           `json:"chat_id,omitempty"`
	MessageID            int                   `json:"message_id,omitempty"`
	InlineMessageID      string                `json:"inline_message_id,omitempty"`
	ReplyMarkup          *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#editmessagelivelocation
type EditMessageLiveLocationParams struct {
	BusinessConnectionID    string                `json:"business_connection_id,omitempty"`
	ChatID                  interface{}           `json:"chat_id,omitempty"`
	MessageID               int                   `json:"message_id,omitempty"`
	InlineMessageID         string                `json:"inline_message_id,omitempty"`
	Latitude                float64               `json:"latitude"`
	Longitude               float64               `json:"longitude"`
	LivePeriod              int                   `json:"live_period,omitempty"`
	HorizontalAccuracy      float64               `json:"horizontal_accuracy,omitempty"`
	Heading                 int                   `json:"heading,omitempty"`
	ProximityAlertRadius    int                   `json:"proximity_alert_radius,omitempty"`
	ReplyMarkup             *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#stopmessagelivelocation
type StopMessageLiveLocationParams struct {
	BusinessConnectionID string                `json:"business_connection_id,omitempty"`
	ChatID               interface{}           `json:"chat_id,omitempty"`
	MessageID            int                   `json:"message_id,omitempty"`
	InlineMessageID      string                `json:"inline_message_id,omitempty"`
	ReplyMarkup          *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#editmessagechecklist
type EditMessageChecklistParams struct {
	BusinessConnectionID string                `json:"business_connection_id"`
	ChatID               int64                 `json:"chat_id"`
	MessageID            int                   `json:"message_id"`
	Checklist            InputChecklist        `json:"checklist"`
	ReplyMarkup          *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#stoppoll
type StopPollParams struct {
	BusinessConnectionID string                `json:"business_connection_id,omitempty"`
	ChatID               interface{}           `json:"chat_id"`
	MessageID            int                   `json:"message_id"`
	ReplyMarkup          *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#sendmediagroup
type SendMediaGroupParams struct {
	BusinessConnectionID    string                `json:"business_connection_id,omitempty"`
	ChatID                  interface{}           `json:"chat_id"`
	MessageThreadID         int                   `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID   int                   `json:"direct_messages_topic_id,omitempty"`
	Media                   []interface{}         `json:"media"`
	DisableNotification     bool                  `json:"disable_notification,omitempty"`
	ProtectContent          bool                  `json:"protect_content,omitempty"`
	AllowPaidBroadcast      bool                  `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID         string                `json:"message_effect_id,omitempty"`
	ReplyParameters         *ReplyParameters      `json:"reply_parameters,omitempty"`
}

// https://core.telegram.org/bots/api#sendpaidmedia
type SendPaidMediaParams struct {
	BusinessConnectionID    string                `json:"business_connection_id,omitempty"`
	ChatID                  interface{}           `json:"chat_id"`
	MessageThreadID         int                   `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID   int                   `json:"direct_messages_topic_id,omitempty"`
	StarCount               int                   `json:"star_count"`
	Media                   []interface{}         `json:"media"`
	Payload                 string                `json:"payload,omitempty"`
	Caption                 string                `json:"caption,omitempty"`
	ParseMode               string                `json:"parse_mode,omitempty"`
	CaptionEntities         []MessageEntity       `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia   bool                  `json:"show_caption_above_media,omitempty"`
	DisableNotification     bool                  `json:"disable_notification,omitempty"`
	ProtectContent          bool                  `json:"protect_content,omitempty"`
	AllowPaidBroadcast      bool                  `json:"allow_paid_broadcast,omitempty"`
	SuggestedPostParameters *SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters         *ReplyParameters      `json:"reply_parameters,omitempty"`
	ReplyMarkup             interface{}           `json:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#sendsticker
type SendStickerParams struct {
	BusinessConnectionID    string                `json:"business_connection_id,omitempty"`
	ChatID                  interface{}           `json:"chat_id"`
	MessageThreadID         int                   `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID   int                   `json:"direct_messages_topic_id,omitempty"`
	Sticker                 interface{}           `json:"sticker"`
	Emoji                   string                `json:"emoji,omitempty"`
	DisableNotification     bool                  `json:"disable_notification,omitempty"`
	ProtectContent          bool                  `json:"protect_content,omitempty"`
	AllowPaidBroadcast      bool                  `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID         string                `json:"message_effect_id,omitempty"`
	SuggestedPostParameters *SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters         *ReplyParameters      `json:"reply_parameters,omitempty"`
	ReplyMarkup             interface{}           `json:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#sendmessagedraft
type SendMessageDraftParams struct {
	ChatID          int64           `json:"chat_id"`
	MessageThreadID int             `json:"message_thread_id,omitempty"`
	DraftID         int             `json:"draft_id"`
	Text            string          `json:"text"`
	ParseMode       string          `json:"parse_mode,omitempty"`
	Entities        []MessageEntity `json:"entities,omitempty"`
}

// https://core.telegram.org/bots/api#copymessages
type CopyMessagesParams struct {
	ChatID                  interface{} `json:"chat_id"`
	MessageThreadID         int         `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID   int         `json:"direct_messages_topic_id,omitempty"`
	FromChatID              interface{} `json:"from_chat_id"`
	MessageIDs              []int       `json:"message_ids"`
	DisableNotification     bool        `json:"disable_notification,omitempty"`
	ProtectContent          bool        `json:"protect_content,omitempty"`
	RemoveCaption           bool        `json:"remove_caption,omitempty"`
}

// https://core.telegram.org/bots/api#forwardmessages
type ForwardMessagesParams struct {
	ChatID                interface{} `json:"chat_id"`
	MessageThreadID       int         `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID int         `json:"direct_messages_topic_id,omitempty"`
	FromChatID            interface{} `json:"from_chat_id"`
	MessageIDs            []int       `json:"message_ids"`
	DisableNotification   bool        `json:"disable_notification,omitempty"`
	ProtectContent        bool        `json:"protect_content,omitempty"`
}

// https://core.telegram.org/bots/api#deletemessages
type DeleteMessagesParams struct {
	ChatID     interface{} `json:"chat_id"`
	MessageIDs []int       `json:"message_ids"`
}

// https://core.telegram.org/bots/api#inputmedia
type InputMedia interface{}

// https://core.telegram.org/bots/api#inputmediaphoto
type InputMediaPhoto struct {
	Type                  string          `json:"type"`
	Media                 string          `json:"media"`
	Caption               string          `json:"caption,omitempty"`
	ParseMode             string          `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool            `json:"show_caption_above_media,omitempty"`
	HasSpoiler            bool            `json:"has_spoiler,omitempty"`
}

// https://core.telegram.org/bots/api#inputmediavideo
type InputMediaVideo struct {
	Type                  string          `json:"type"`
	Media                 string          `json:"media"`
	Thumbnail             interface{}     `json:"thumbnail,omitempty"`
	Caption               string          `json:"caption,omitempty"`
	ParseMode             string          `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool            `json:"show_caption_above_media,omitempty"`
	Width                 int             `json:"width,omitempty"`
	Height                int             `json:"height,omitempty"`
	Duration              int             `json:"duration,omitempty"`
	SupportsStreaming     bool            `json:"supports_streaming,omitempty"`
	HasSpoiler            bool            `json:"has_spoiler,omitempty"`
}

// https://core.telegram.org/bots/api#inputmediaanimation
type InputMediaAnimation struct {
	Type                  string          `json:"type"`
	Media                 string          `json:"media"`
	Thumbnail             interface{}     `json:"thumbnail,omitempty"`
	Caption               string          `json:"caption,omitempty"`
	ParseMode             string          `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool            `json:"show_caption_above_media,omitempty"`
	Width                 int             `json:"width,omitempty"`
	Height                int             `json:"height,omitempty"`
	Duration              int             `json:"duration,omitempty"`
	HasSpoiler            bool            `json:"has_spoiler,omitempty"`
}

// https://core.telegram.org/bots/api#inputmediaaudio
type InputMediaAudio struct {
	Type            string          `json:"type"`
	Media           string          `json:"media"`
	Thumbnail       interface{}     `json:"thumbnail,omitempty"`
	Caption         string          `json:"caption,omitempty"`
	ParseMode       string          `json:"parse_mode,omitempty"`
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	Duration        int             `json:"duration,omitempty"`
	Performer       string          `json:"performer,omitempty"`
	Title           string          `json:"title,omitempty"`
}

// https://core.telegram.org/bots/api#inputmediadocument
type InputMediaDocument struct {
	Type                        string          `json:"type"`
	Media                       string          `json:"media"`
	Thumbnail                   interface{}     `json:"thumbnail,omitempty"`
	Caption                     string          `json:"caption,omitempty"`
	ParseMode                   string          `json:"parse_mode,omitempty"`
	CaptionEntities             []MessageEntity `json:"caption_entities,omitempty"`
	DisableContentTypeDetection bool            `json:"disable_content_type_detection,omitempty"`
}

// https://core.telegram.org/bots/api#paidmediainfo
type PaidMediaInfo struct {
	StarCount int           `json:"star_count"`
	PaidMedia []interface{} `json:"paid_media"`
}

// https://core.telegram.org/bots/api#inputpaidmediaphoto
type InputPaidMediaPhoto struct {
	Type  string `json:"type"`
	Media string `json:"media"`
}

// https://core.telegram.org/bots/api#inputpaidmediavideo
type InputPaidMediaVideo struct {
	Type      string      `json:"type"`
	Media     string      `json:"media"`
	Thumbnail interface{} `json:"thumbnail,omitempty"`
	Width     int         `json:"width,omitempty"`
	Height    int         `json:"height,omitempty"`
	Duration  int         `json:"duration,omitempty"`
}

// https://core.telegram.org/bots/api#replyparameters
type ReplyParameters struct {
	MessageID                int    `json:"message_id"`
	ChatID                   interface{} `json:"chat_id,omitempty"`
	AllowSendingWithoutReply bool   `json:"allow_sending_without_reply,omitempty"`
	Quote                    string `json:"quote,omitempty"`
	QuoteParseMode           string `json:"quote_parse_mode,omitempty"`
	QuoteEntities            []MessageEntity `json:"quote_entities,omitempty"`
	QuotePosition            int    `json:"quote_position,omitempty"`
}

// https://core.telegram.org/bots/api#suggestedpostparameters
type SuggestedPostParameters struct {
	Price    *SuggestedPostPrice `json:"price,omitempty"`
	SendDate int64               `json:"send_date,omitempty"`
}

// https://core.telegram.org/bots/api#suggestedpostprice
type SuggestedPostPrice struct {
	StarCount int `json:"star_count"`
}

// https://core.telegram.org/bots/api#suggestedpostinfo
type SuggestedPostInfo struct {
	State    string              `json:"state"`
	Price    *SuggestedPostPrice `json:"price,omitempty"`
	SendDate int64               `json:"send_date,omitempty"`
}

// https://core.telegram.org/bots/api#polloption
type PollOption struct {
	Text       string `json:"text"`
	VoterCount int    `json:"voter_count"`
}

// https://core.telegram.org/bots/api#inputpolloption
type InputPollOption struct {
	Text          string          `json:"text"`
	TextParseMode string          `json:"text_parse_mode,omitempty"`
	TextEntities  []MessageEntity `json:"text_entities,omitempty"`
}

// https://core.telegram.org/bots/api#pollanswer
type PollAnswer struct {
	PollID    string `json:"poll_id"`
	VoterChat *Chat  `json:"voter_chat,omitempty"`
	User      *User  `json:"user,omitempty"`
	OptionIDs []int  `json:"option_ids"`
}

// https://core.telegram.org/bots/api#poll
type Poll struct {
	ID                    string           `json:"id"`
	Question              string           `json:"question"`
	QuestionEntities      []MessageEntity  `json:"question_entities,omitempty"`
	Options               []PollOption     `json:"options"`
	TotalVoterCount       int              `json:"total_voter_count"`
	IsClosed              bool             `json:"is_closed"`
	IsAnonymous           bool             `json:"is_anonymous"`
	Type                  string           `json:"type"`
	AllowsMultipleAnswers bool             `json:"allows_multiple_answers"`
	CorrectOptionID       *int             `json:"correct_option_id,omitempty"`
	Explanation           string           `json:"explanation,omitempty"`
	ExplanationEntities   []MessageEntity  `json:"explanation_entities,omitempty"`
	OpenPeriod            int              `json:"open_period,omitempty"`
	CloseDate             int64            `json:"close_date,omitempty"`
}

// https://core.telegram.org/bots/api#checklisttask
type ChecklistTask struct {
	Text               string          `json:"text"`
	TextEntities       []MessageEntity `json:"text_entities,omitempty"`
	IsDone             bool            `json:"is_done"`
	CompletedByUser    *User           `json:"completed_by_user,omitempty"`
	CompletedByChat    *Chat           `json:"completed_by_chat,omitempty"`
}

// https://core.telegram.org/bots/api#checklist
type Checklist struct {
	Title                    string          `json:"title"`
	TitleEntities            []MessageEntity `json:"title_entities,omitempty"`
	Tasks                    []ChecklistTask `json:"tasks"`
	OthersCanAddTasks        bool            `json:"others_can_add_tasks,omitempty"`
	OthersCanMarkTasksAsDone bool            `json:"others_can_mark_tasks_as_done,omitempty"`
}

// https://core.telegram.org/bots/api#inputchecklisttask
type InputChecklistTask struct {
	Text          string          `json:"text"`
	TextParseMode string          `json:"text_parse_mode,omitempty"`
	TextEntities  []MessageEntity `json:"text_entities,omitempty"`
}

// https://core.telegram.org/bots/api#inputchecklist
type InputChecklist struct {
	Title                    string               `json:"title"`
	TitleParseMode           string               `json:"title_parse_mode,omitempty"`
	TitleEntities            []MessageEntity      `json:"title_entities,omitempty"`
	Tasks                    []InputChecklistTask `json:"tasks"`
	OthersCanAddTasks        bool                 `json:"others_can_add_tasks,omitempty"`
	OthersCanMarkTasksAsDone bool                 `json:"others_can_mark_tasks_as_done,omitempty"`
}

// https://core.telegram.org/bots/api#game
type Game struct {
	Title        string          `json:"title"`
	Description  string          `json:"description"`
	Photo        []PhotoSize     `json:"photo"`
	Text         string          `json:"text,omitempty"`
	TextEntities []MessageEntity `json:"text_entities,omitempty"`
	Animation    *Animation      `json:"animation,omitempty"`
}

// https://core.telegram.org/bots/api#gamehighscore
type GameHighScore struct {
	Position int  `json:"position"`
	User     User `json:"user"`
	Score    int  `json:"score"`
}

// https://core.telegram.org/bots/api#sendgame
type SendGameParams struct {
	BusinessConnectionID      string                `json:"business_connection_id,omitempty"`
	ChatID                    int64                 `json:"chat_id"`
	MessageThreadID           int                   `json:"message_thread_id,omitempty"`
	GameShortName             string                `json:"game_short_name"`
	DisableNotification       bool                  `json:"disable_notification,omitempty"`
	ProtectContent            bool                  `json:"protect_content,omitempty"`
	MessageEffectID           string                `json:"message_effect_id,omitempty"`
	ReplyParameters           *ReplyParameters      `json:"reply_parameters,omitempty"`
	ReplyMarkup               *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#setgamescore
type SetGameScoreParams struct {
	UserID             int64  `json:"user_id"`
	Score              int    `json:"score"`
	Force              bool   `json:"force,omitempty"`
	DisableEditMessage bool   `json:"disable_edit_message,omitempty"`
	ChatID             int64  `json:"chat_id,omitempty"`
	MessageID          int    `json:"message_id,omitempty"`
	InlineMessageID    string `json:"inline_message_id,omitempty"`
}

// https://core.telegram.org/bots/api#getgamehighscores
type GetGameHighScoresParams struct {
	UserID          int64  `json:"user_id"`
	ChatID          int64  `json:"chat_id,omitempty"`
	MessageID       int    `json:"message_id,omitempty"`
	InlineMessageID string `json:"inline_message_id,omitempty"`
}
