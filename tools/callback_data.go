package tools

import (
	"github.com/erfjab/egobot/core"
	"github.com/erfjab/egobot/models"
)

// CallbackData is a marker type for struct-based callback payloads.
//
// Example:
//
//	type AdminCB struct {
//		tools.CallbackData `prefix:"admin"`
//		Section   string
//		Action    string
//		SubAction string
//		UserID    int
//	}
//
// If prefix tag is omitted, prefix is derived from struct name.
// Optional tag: separator:"|".
type CallbackData struct{}

// PackCallback packs a struct payload into Telegram callback data.
func PackCallback(v interface{}) (string, error) {
	data, err := core.NewCallbackDataFromStruct(v)
	if err != nil {
		return "", err
	}
	return data.PackStruct(v)
}

// ParseCallback parses callback data string into provided struct pointer.
func ParseCallback(raw string, out interface{}) bool {
	data, err := core.NewCallbackDataFromStruct(out)
	if err != nil {
		return false
	}
	return data.ParseToStruct(raw, out)
}

// CallbackButton builds an inline button and auto-packs callback payload.
func CallbackButton(text string, payload interface{}) (models.InlineKeyboardButton, error) {
	packed, err := PackCallback(payload)
	if err != nil {
		return models.InlineKeyboardButton{}, err
	}
	return Button(text, packed), nil
}

// MustCallbackButton is like CallbackButton but panics on invalid payload.
// Useful for concise keyboard definitions.
func MustCallbackButton(text string, payload interface{}) models.InlineKeyboardButton {
	btn, err := CallbackButton(text, payload)
	if err != nil {
		panic(err)
	}
	return btn
}
