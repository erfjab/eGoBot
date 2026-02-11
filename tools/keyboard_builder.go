package tools

import "egobot/models"

// InlineKeyboardBuilder helps build inline keyboards easily
type InlineKeyboardBuilder struct {
	rows [][]models.InlineKeyboardButton
}

// NewInlineKeyboard creates a new inline keyboard builder
func NewInlineKeyboard() *InlineKeyboardBuilder {
	return &InlineKeyboardBuilder{
		rows: make([][]models.InlineKeyboardButton, 0),
	}
}

// Row adds a new row to the keyboard
func (kb *InlineKeyboardBuilder) Row(buttons ...models.InlineKeyboardButton) *InlineKeyboardBuilder {
	kb.rows = append(kb.rows, buttons)
	return kb
}

// Build returns the final InlineKeyboardMarkup
func (kb *InlineKeyboardBuilder) Build() *models.InlineKeyboardMarkup {
	return &models.InlineKeyboardMarkup{
		InlineKeyboard: kb.rows,
	}
}

// Button creates a simple text button with callback data
func Button(text, callbackData string) models.InlineKeyboardButton {
	return models.InlineKeyboardButton{
		Text:         text,
		CallbackData: callbackData,
	}
}

// URLButton creates a button that opens a URL
func URLButton(text, url string) models.InlineKeyboardButton {
	return models.InlineKeyboardButton{
		Text: text,
		URL:  url,
	}
}

// WebAppButton creates a button that opens a web app
func WebAppButton(text, url string) models.InlineKeyboardButton {
	return models.InlineKeyboardButton{
		Text: text,
		WebApp: &models.WebAppInfo{
			URL: url,
		},
	}
}

// SwitchInlineButton creates a button that switches to inline query
func SwitchInlineButton(text, query string) models.InlineKeyboardButton {
	return models.InlineKeyboardButton{
		Text:              text,
		SwitchInlineQuery: query,
	}
}

// SwitchInlineCurrentButton creates a button that switches to inline query in current chat
func SwitchInlineCurrentButton(text, query string) models.InlineKeyboardButton {
	return models.InlineKeyboardButton{
		Text:                         text,
		SwitchInlineQueryCurrentChat: query,
	}
}

// PayButton creates a payment button
func PayButton(text string) models.InlineKeyboardButton {
	return models.InlineKeyboardButton{
		Text: text,
		Pay:  true,
	}
}

// LoginButton creates a login URL button
func LoginButton(text, url string) models.InlineKeyboardButton {
	return models.InlineKeyboardButton{
		Text: text,
		LoginURL: &models.LoginURL{
			URL: url,
		},
	}
}

// ReplyKeyboardBuilder helps build reply keyboards easily
type ReplyKeyboardBuilder struct {
	rows                  [][]models.KeyboardButton
	isPersistent          bool
	resizeKeyboard        bool
	oneTimeKeyboard       bool
	inputFieldPlaceholder string
	selective             bool
}

// NewReplyKeyboard creates a new reply keyboard builder
func NewReplyKeyboard() *ReplyKeyboardBuilder {
	return &ReplyKeyboardBuilder{
		rows:           make([][]models.KeyboardButton, 0),
		resizeKeyboard: true, // Default to true for better UX
	}
}

// Row adds a new row to the keyboard
func (kb *ReplyKeyboardBuilder) Row(buttons ...models.KeyboardButton) *ReplyKeyboardBuilder {
	kb.rows = append(kb.rows, buttons)
	return kb
}

// Persistent sets whether the keyboard should be persistent
func (kb *ReplyKeyboardBuilder) Persistent(persistent bool) *ReplyKeyboardBuilder {
	kb.isPersistent = persistent
	return kb
}

// Resize sets whether the keyboard should be resized
func (kb *ReplyKeyboardBuilder) Resize(resize bool) *ReplyKeyboardBuilder {
	kb.resizeKeyboard = resize
	return kb
}

// OneTime sets whether the keyboard should be one-time
func (kb *ReplyKeyboardBuilder) OneTime(oneTime bool) *ReplyKeyboardBuilder {
	kb.oneTimeKeyboard = oneTime
	return kb
}

// Placeholder sets the input field placeholder
func (kb *ReplyKeyboardBuilder) Placeholder(placeholder string) *ReplyKeyboardBuilder {
	kb.inputFieldPlaceholder = placeholder
	return kb
}

// Selective sets whether the keyboard should be selective
func (kb *ReplyKeyboardBuilder) Selective(selective bool) *ReplyKeyboardBuilder {
	kb.selective = selective
	return kb
}

// Build returns the final ReplyKeyboardMarkup
func (kb *ReplyKeyboardBuilder) Build() *models.ReplyKeyboardMarkup {
	return &models.ReplyKeyboardMarkup{
		Keyboard:              kb.rows,
		IsPersistent:          kb.isPersistent,
		ResizeKeyboard:        kb.resizeKeyboard,
		OneTimeKeyboard:       kb.oneTimeKeyboard,
		InputFieldPlaceholder: kb.inputFieldPlaceholder,
		Selective:             kb.selective,
	}
}

// TextButton creates a simple text button
func TextButton(text string) models.KeyboardButton {
	return models.KeyboardButton{
		Text: text,
	}
}

// ContactButton creates a button that requests contact
func ContactButton(text string) models.KeyboardButton {
	return models.KeyboardButton{
		Text:           text,
		RequestContact: true,
	}
}

// LocationButton creates a button that requests location
func LocationButton(text string) models.KeyboardButton {
	return models.KeyboardButton{
		Text:            text,
		RequestLocation: true,
	}
}

// PollButton creates a button that requests a poll
func PollButton(text, pollType string) models.KeyboardButton {
	return models.KeyboardButton{
		Text: text,
		RequestPoll: &models.KeyboardButtonPollType{
			Type: pollType,
		},
	}
}

// WebAppReplyButton creates a button that opens a web app
func WebAppReplyButton(text, url string) models.KeyboardButton {
	return models.KeyboardButton{
		Text: text,
		WebApp: &models.WebAppInfo{
			URL: url,
		},
	}
}

// RequestUsersButton creates a button that requests users
func RequestUsersButton(text string, requestID int) models.KeyboardButton {
	return models.KeyboardButton{
		Text: text,
		RequestUsers: &models.KeyboardButtonRequestUsers{
			RequestID: requestID,
		},
	}
}

// RequestChatButton creates a button that requests a chat
func RequestChatButton(text string, requestID int, isChannel bool) models.KeyboardButton {
	return models.KeyboardButton{
		Text: text,
		RequestChat: &models.KeyboardButtonRequestChat{
			RequestID:     requestID,
			ChatIsChannel: isChannel,
		},
	}
}

// RemoveKeyboard creates a ReplyKeyboardRemove
func RemoveKeyboard() *models.ReplyKeyboardRemove {
	return &models.ReplyKeyboardRemove{
		RemoveKeyboard: true,
	}
}

// RemoveKeyboardSelective creates a selective ReplyKeyboardRemove
func RemoveKeyboardSelective() *models.ReplyKeyboardRemove {
	return &models.ReplyKeyboardRemove{
		RemoveKeyboard: true,
		Selective:      true,
	}
}
