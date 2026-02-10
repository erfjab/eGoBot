package models

// https://core.telegram.org/bots/api#stickerset
type StickerSet struct {
	Name        string     `json:"name"`
	Title       string     `json:"title"`
	StickerType string     `json:"sticker_type"`
	Stickers    []Sticker  `json:"stickers"`
	Thumbnail   *PhotoSize `json:"thumbnail,omitempty"`
}

// https://core.telegram.org/bots/api#inputsticker
type InputSticker struct {
	Sticker      interface{}   `json:"sticker"`
	Format       string        `json:"format"`
	EmojiList    []string      `json:"emoji_list"`
	MaskPosition *MaskPosition `json:"mask_position,omitempty"`
	Keywords     []string      `json:"keywords,omitempty"`
}

// https://core.telegram.org/bots/api#getstickerset
type GetStickerSetParams struct {
	Name string `json:"name"`
}

// https://core.telegram.org/bots/api#getcustomemojistickers
type GetCustomEmojiStickersParams struct {
	CustomEmojiIDs []string `json:"custom_emoji_ids"`
}

// https://core.telegram.org/bots/api#uploadstickerfile
type UploadStickerFileParams struct {
	UserID        int64       `json:"user_id"`
	Sticker       interface{} `json:"sticker"`
	StickerFormat string      `json:"sticker_format"`
}

// https://core.telegram.org/bots/api#createnewstickerset
type CreateNewStickerSetParams struct {
	UserID          int64          `json:"user_id"`
	Name            string         `json:"name"`
	Title           string         `json:"title"`
	Stickers        []InputSticker `json:"stickers"`
	StickerType     string         `json:"sticker_type,omitempty"`
	NeedsRepainting bool           `json:"needs_repainting,omitempty"`
}

// https://core.telegram.org/bots/api#addstickertoset
type AddStickerToSetParams struct {
	UserID  int64        `json:"user_id"`
	Name    string       `json:"name"`
	Sticker InputSticker `json:"sticker"`
}

// https://core.telegram.org/bots/api#setstickerpositioninset
type SetStickerPositionInSetParams struct {
	Sticker  string `json:"sticker"`
	Position int    `json:"position"`
}

// https://core.telegram.org/bots/api#deletestickerfromset
type DeleteStickerFromSetParams struct {
	Sticker string `json:"sticker"`
}

// https://core.telegram.org/bots/api#setstickersetthumbnail
type SetStickerSetThumbnailParams struct {
	Name      string      `json:"name"`
	UserID    int64       `json:"user_id"`
	Thumbnail interface{} `json:"thumbnail,omitempty"`
	Format    string      `json:"format"`
}
