package models

// https://core.telegram.org/bots/api#inlinequery
type InlineQuery struct {
	ID       string    `json:"id"`
	From     User      `json:"from"`
	Query    string    `json:"query"`
	Offset   string    `json:"offset"`
	ChatType string    `json:"chat_type,omitempty"`
	Location *Location `json:"location,omitempty"`
}

// https://core.telegram.org/bots/api#choseninlineresult
type ChosenInlineResult struct {
	ResultID        string    `json:"result_id"`
	From            User      `json:"from"`
	Location        *Location `json:"location,omitempty"`
	InlineMessageID string    `json:"inline_message_id,omitempty"`
	Query           string    `json:"query"`
}

// https://core.telegram.org/bots/api#callbackquery
type CallbackQuery struct {
	ID              string   `json:"id"`
	From            User     `json:"from"`
	Message         *Message `json:"message,omitempty"`
	InlineMessageID string   `json:"inline_message_id,omitempty"`
	ChatInstance    string   `json:"chat_instance"`
	Data            string   `json:"data,omitempty"`
	GameShortName   string   `json:"game_short_name,omitempty"`
}

// https://core.telegram.org/bots/api#answerinlinequery
type AnswerInlineQueryParams struct {
	InlineQueryID string        `json:"inline_query_id"`
	Results       []interface{} `json:"results"`
	CacheTime     int           `json:"cache_time,omitempty"`
	IsPersonal    bool          `json:"is_personal,omitempty"`
	NextOffset    string        `json:"next_offset,omitempty"`
	Button        interface{}   `json:"button,omitempty"`
}

// https://core.telegram.org/bots/api#inlinequeryresultarticle
type InlineQueryResultArticle struct {
	Type                string          `json:"type"`
	ID                  string          `json:"id"`
	Title               string          `json:"title"`
	InputMessageContent interface{}     `json:"input_message_content"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	URL                 string          `json:"url,omitempty"`
	HideURL             bool            `json:"hide_url,omitempty"`
	Description         string          `json:"description,omitempty"`
	ThumbnailURL        string          `json:"thumbnail_url,omitempty"`
	ThumbnailWidth      int             `json:"thumbnail_width,omitempty"`
	ThumbnailHeight     int             `json:"thumbnail_height,omitempty"`
}

// https://core.telegram.org/bots/api#inlinequeryresultphoto
type InlineQueryResultPhoto struct {
	Type                  string                `json:"type"`
	ID                    string                `json:"id"`
	PhotoURL              string                `json:"photo_url"`
	ThumbnailURL          string                `json:"thumbnail_url"`
	PhotoWidth            int                   `json:"photo_width,omitempty"`
	PhotoHeight           int                   `json:"photo_height,omitempty"`
	Title                 string                `json:"title,omitempty"`
	Description           string                `json:"description,omitempty"`
	Caption               string                `json:"caption,omitempty"`
	ParseMode             string                `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity       `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool                  `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent   interface{}           `json:"input_message_content,omitempty"`
}

// https://core.telegram.org/bots/api#inlinequeryresultcachedphoto
type InlineQueryResultCachedPhoto struct {
	Type                  string                `json:"type"`
	ID                    string                `json:"id"`
	PhotoFileID           string                `json:"photo_file_id"`
	Title                 string                `json:"title,omitempty"`
	Description           string                `json:"description,omitempty"`
	Caption               string                `json:"caption,omitempty"`
	ParseMode             string                `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity       `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool                  `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent   interface{}           `json:"input_message_content,omitempty"`
}

// https://core.telegram.org/bots/api#inputtextmessagecontent
type InputTextMessageContent struct {
	MessageText           string          `json:"message_text"`
	ParseMode             string          `json:"parse_mode,omitempty"`
	Entities              []MessageEntity `json:"entities,omitempty"`
	LinkPreviewOptions    interface{}     `json:"link_preview_options,omitempty"`
}
