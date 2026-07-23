package models

import "encoding/json"

// InputFile represents a file to send to Telegram.
// Exactly one of FileID, URL, or Data should be set:
//   - FileID: re-use an already-uploaded Telegram file by its file_id.
//   - URL:    let Telegram fetch the file from a public HTTP/HTTPS URL.
//   - Data:   upload raw bytes via multipart/form-data (Name is used as the filename).
//
// https://core.telegram.org/bots/api#inputfile
type InputFile struct {
	FileID string
	URL    string
	Data   []byte
	Name   string
}

// MarshalJSON serialises InputFile as a plain JSON string (file_id or URL).
// When Data is set the caller must use a multipart/form-data request instead;
// this method is only a fallback for the JSON path and returns an empty string
// in that case so that misconfiguration is obvious.
func (f InputFile) MarshalJSON() ([]byte, error) {
	if f.FileID != "" {
		return json.Marshal(f.FileID)
	}
	return json.Marshal(f.URL)
}

// https://core.telegram.org/bots/api#photosize
type PhotoSize struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	FileSize     int64  `json:"file_size,omitempty"`
}

// https://core.telegram.org/bots/api#document
type Document struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"`
	FileName     string     `json:"file_name,omitempty"`
	MimeType     string     `json:"mime_type,omitempty"`
	FileSize     int64      `json:"file_size,omitempty"`
}

// https://core.telegram.org/bots/api#video
type Video struct {
	FileID         string         `json:"file_id"`
	FileUniqueID   string         `json:"file_unique_id"`
	Width          int            `json:"width"`
	Height         int            `json:"height"`
	Duration       int            `json:"duration"`
	Thumbnail      *PhotoSize     `json:"thumbnail,omitempty"`
	Cover          []PhotoSize    `json:"cover,omitempty"`
	StartTimestamp int            `json:"start_timestamp,omitempty"`
	Qualities      []VideoQuality `json:"qualities,omitempty"`
	FileName       string         `json:"file_name,omitempty"`
	MimeType       string         `json:"mime_type,omitempty"`
	FileSize       int64          `json:"file_size,omitempty"`
}

// https://core.telegram.org/bots/api#audio
type Audio struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Duration     int        `json:"duration"`
	Performer    string     `json:"performer,omitempty"`
	Title        string     `json:"title,omitempty"`
	FileName     string     `json:"file_name,omitempty"`
	MimeType     string     `json:"mime_type,omitempty"`
	FileSize     int64      `json:"file_size,omitempty"`
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"`
}

// https://core.telegram.org/bots/api#voice
type Voice struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	Duration     int    `json:"duration"`
	MimeType     string `json:"mime_type,omitempty"`
	FileSize     int64  `json:"file_size,omitempty"`
}

// https://core.telegram.org/bots/api#videonote
type VideoNote struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Length       int        `json:"length"`
	Duration     int        `json:"duration"`
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"`
	FileSize     int64      `json:"file_size,omitempty"`
}

// https://core.telegram.org/bots/api#contact
type Contact struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name,omitempty"`
	UserID      int64  `json:"user_id,omitempty"`
	VCard       string `json:"vcard,omitempty"`
}

// https://core.telegram.org/bots/api#sticker
type Sticker struct {
	FileID           string        `json:"file_id"`
	FileUniqueID     string        `json:"file_unique_id"`
	Type             string        `json:"type"`
	Width            int           `json:"width"`
	Height           int           `json:"height"`
	IsAnimated       bool          `json:"is_animated"`
	IsVideo          bool          `json:"is_video"`
	Thumbnail        *PhotoSize    `json:"thumbnail,omitempty"`
	Emoji            string        `json:"emoji,omitempty"`
	SetName          string        `json:"set_name,omitempty"`
	PremiumAnimation *File         `json:"premium_animation,omitempty"`
	MaskPosition     *MaskPosition `json:"mask_position,omitempty"`
	CustomEmojiID    string        `json:"custom_emoji_id,omitempty"`
	FileSize         int64         `json:"file_size,omitempty"`
}

// https://core.telegram.org/bots/api#maskposition
type MaskPosition struct {
	Point  string  `json:"point"`
	XShift float64 `json:"x_shift"`
	YShift float64 `json:"y_shift"`
	Scale  float64 `json:"scale"`
}

// https://core.telegram.org/bots/api#file
type File struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	FileSize     int64  `json:"file_size,omitempty"`
	FilePath     string `json:"file_path,omitempty"`
}

// https://core.telegram.org/bots/api#animation
type Animation struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Width        int        `json:"width"`
	Height       int        `json:"height"`
	Duration     int        `json:"duration"`
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"`
	FileName     string     `json:"file_name,omitempty"`
	MimeType     string     `json:"mime_type,omitempty"`
	FileSize     int64      `json:"file_size,omitempty"`
}

// https://core.telegram.org/bots/api#location
type Location struct {
	Longitude            float64 `json:"longitude"`
	Latitude             float64 `json:"latitude"`
	HorizontalAccuracy   float64 `json:"horizontal_accuracy,omitempty"`
	LivePeriod           int     `json:"live_period,omitempty"`
	Heading              int     `json:"heading,omitempty"`
	ProximityAlertRadius int     `json:"proximity_alert_radius,omitempty"`
}

// https://core.telegram.org/bots/api#venue
type Venue struct {
	Location        Location `json:"location"`
	Title           string   `json:"title"`
	Address         string   `json:"address"`
	FoursquareID    string   `json:"foursquare_id,omitempty"`
	FoursquareType  string   `json:"foursquare_type,omitempty"`
	GooglePlaceID   string   `json:"google_place_id,omitempty"`
	GooglePlaceType string   `json:"google_place_type,omitempty"`
}

// https://core.telegram.org/bots/api#dice
type Dice struct {
	Emoji string `json:"emoji"`
	Value int    `json:"value"`
}

// https://core.telegram.org/bots/api#inputmedialink
type InputMediaLink struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

// https://core.telegram.org/bots/api#inputmedialivephoto
type InputMediaLivePhoto struct {
	Type                  string          `json:"type"`
	Media                 string          `json:"media"`
	Photo                 string          `json:"photo"`
	Caption               string          `json:"caption,omitempty"`
	ParseMode             string          `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool            `json:"show_caption_above_media,omitempty"`
	HasSpoiler            bool            `json:"has_spoiler,omitempty"`
}

// https://core.telegram.org/bots/api#inputmedialocation
type InputMediaLocation struct {
	Type               string  `json:"type"`
	Latitude           float64 `json:"latitude"`
	Longitude          float64 `json:"longitude"`
	HorizontalAccuracy float64 `json:"horizontal_accuracy,omitempty"`
}

// https://core.telegram.org/bots/api#inputmediasticker
type InputMediaSticker struct {
	Type  string `json:"type"`
	Media string `json:"media"`
	Emoji string `json:"emoji,omitempty"`
}

// https://core.telegram.org/bots/api#inputmediavenue
type InputMediaVenue struct {
	Type            string  `json:"type"`
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
	Title           string  `json:"title"`
	Address         string  `json:"address"`
	FoursquareID    string  `json:"foursquare_id,omitempty"`
	FoursquareType  string  `json:"foursquare_type,omitempty"`
	GooglePlaceID   string  `json:"google_place_id,omitempty"`
	GooglePlaceType string  `json:"google_place_type,omitempty"`
}

// https://core.telegram.org/bots/api#inputmediavoicenote
type InputMediaVoiceNote struct {
	Type            string          `json:"type"`
	Media           string          `json:"media"`
	Caption         string          `json:"caption,omitempty"`
	ParseMode       string          `json:"parse_mode,omitempty"`
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	Duration        int             `json:"duration,omitempty"`
}

// https://core.telegram.org/bots/api#inputpaidmedialivephoto
type InputPaidMediaLivePhoto struct {
	Type  string `json:"type"`
	Media string `json:"media"`
	Photo string `json:"photo"`
}

// https://core.telegram.org/bots/api#paidmedialivephoto
type PaidMediaLivePhoto struct {
	Type      string    `json:"type"`
	LivePhoto LivePhoto `json:"live_photo"`
}

// https://core.telegram.org/bots/api#paidmediaphoto
type PaidMediaPhoto struct {
	Type  string      `json:"type"`
	Photo []PhotoSize `json:"photo"`
}

// https://core.telegram.org/bots/api#paidmediapreview
type PaidMediaPreview struct {
	Type     string `json:"type"`
	Width    int    `json:"width,omitempty"`
	Height   int    `json:"height,omitempty"`
	Duration int    `json:"duration,omitempty"`
}

// https://core.telegram.org/bots/api#paidmediavideo
type PaidMediaVideo struct {
	Type  string `json:"type"`
	Video Video  `json:"video"`
}

// https://core.telegram.org/bots/api#inputstorycontentphoto
type InputStoryContentPhoto struct {
	Type  string `json:"type"`
	Photo string `json:"photo"`
}

// https://core.telegram.org/bots/api#inputstorycontentvideo
type InputStoryContentVideo struct {
	Type                string  `json:"type"`
	Video               string  `json:"video"`
	Duration            float64 `json:"duration,omitempty"`
	CoverFrameTimestamp float64 `json:"cover_frame_timestamp,omitempty"`
	IsAnimation         bool    `json:"is_animation,omitempty"`
}

// https://core.telegram.org/bots/api#inputprofilephotoanimated
type InputProfilePhotoAnimated struct {
	Type               string  `json:"type"`
	Animation          string  `json:"animation"`
	MainFrameTimestamp float64 `json:"main_frame_timestamp,omitempty"`
}
