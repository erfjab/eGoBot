package core

import (
	"egobot/models"
	"fmt"
)

// ErrorHandlerFunc represents a function that handles errors
type ErrorHandlerFunc func(*Bot, *models.Update, error) error

// TelegramError represents an error from the Telegram Bot API
type TelegramError struct {
	ErrorCode   int
	Description string
	Parameters  *ResponseParameters
	Update      *models.Update
}

// ResponseParameters contains information about why a request was unsuccessful
type ResponseParameters struct {
	MigrateToChatID int64 // The group has been migrated to a supergroup with the specified identifier
	RetryAfter      int   // In case of exceeding flood control, the number of seconds left to wait
}

// Error implements the error interface for TelegramError
func (e *TelegramError) Error() string {
	if e.Parameters != nil {
		if e.Parameters.MigrateToChatID != 0 {
			return fmt.Sprintf("Telegram API error [%d]: %s (migrate to chat ID: %d)", 
				e.ErrorCode, e.Description, e.Parameters.MigrateToChatID)
		}
		if e.Parameters.RetryAfter != 0 {
			return fmt.Sprintf("Telegram API error [%d]: %s (retry after %d seconds)", 
				e.ErrorCode, e.Description, e.Parameters.RetryAfter)
		}
	}
	return fmt.Sprintf("Telegram API error [%d]: %s", e.ErrorCode, e.Description)
}

// IsTelegramError checks if an error is a TelegramError
func IsTelegramError(err error) bool {
	_, ok := err.(*TelegramError)
	return ok
}

// NewTelegramError creates a new TelegramError
func NewTelegramError(code int, description string, update *models.Update) *TelegramError {
	return &TelegramError{
		ErrorCode:   code,
		Description: description,
		Update:      update,
	}
}

// NewTelegramErrorWithParams creates a new TelegramError with ResponseParameters
func NewTelegramErrorWithParams(code int, description string, params *ResponseParameters, update *models.Update) *TelegramError {
	return &TelegramError{
		ErrorCode:   code,
		Description: description,
		Parameters:  params,
		Update:      update,
	}
}

// Common Telegram Error Codes
const (
	ErrorCodeBadRequest          = 400 // Bad Request
	ErrorCodeUnauthorized        = 401 // Unauthorized
	ErrorCodeForbidden           = 403 // Forbidden
	ErrorCodeNotFound            = 404 // Not Found
	ErrorCodeConflict            = 409 // Conflict
	ErrorCodeTooManyRequests     = 429 // Too Many Requests
	ErrorCodeInternalServerError = 500 // Internal Server Error
)

// IsRateLimitError checks if the error is a rate limit error (429)
func (e *TelegramError) IsRateLimitError() bool {
	return e.ErrorCode == ErrorCodeTooManyRequests
}

// IsBadRequest checks if the error is a bad request (400)
func (e *TelegramError) IsBadRequest() bool {
	return e.ErrorCode == ErrorCodeBadRequest
}

// IsUnauthorized checks if the error is an unauthorized error (401)
func (e *TelegramError) IsUnauthorized() bool {
	return e.ErrorCode == ErrorCodeUnauthorized
}

// IsForbidden checks if the error is a forbidden error (403)
func (e *TelegramError) IsForbidden() bool {
	return e.ErrorCode == ErrorCodeForbidden
}

// IsNotFound checks if the error is a not found error (404)
func (e *TelegramError) IsNotFound() bool {
	return e.ErrorCode == ErrorCodeNotFound
}

// IsConflict checks if the error is a conflict error (409)
func (e *TelegramError) IsConflict() bool {
	return e.ErrorCode == ErrorCodeConflict
}

// IsServerError checks if the error is a server error (5xx)
func (e *TelegramError) IsServerError() bool {
	return e.ErrorCode >= 500 && e.ErrorCode < 600
}

// Message-specific error checks
// IsMessageTextEmpty checks if error is about empty message text
func (e *TelegramError) IsMessageTextEmpty() bool {
	return contains(e.Description, "message text is empty")
}

// IsMessageTooLong checks if error is about message being too long
func (e *TelegramError) IsMessageTooLong() bool {
	return contains(e.Description, "message is too long")
}

// IsChatNotFound checks if error is about chat not being found
func (e *TelegramError) IsChatNotFound() bool {
	return contains(e.Description, "chat not found")
}

// IsMessageNotFound checks if error is about message not being found
func (e *TelegramError) IsMessageNotFound() bool {
	return contains(e.Description, "message to delete not found") ||
		contains(e.Description, "message to edit not found") ||
		contains(e.Description, "message not found")
}

// IsMessageCantBeEdited checks if error is about message that can't be edited
func (e *TelegramError) IsMessageCantBeEdited() bool {
	return contains(e.Description, "message can't be edited") ||
		contains(e.Description, "message to be edited was not found")
}

// IsMessageCantBeDeleted checks if error is about message that can't be deleted
func (e *TelegramError) IsMessageCantBeDeleted() bool {
	return contains(e.Description, "message can't be deleted") ||
		contains(e.Description, "message to delete not found")
}

// IsBotWasBlocked checks if the bot was blocked by user
func (e *TelegramError) IsBotWasBlocked() bool {
	return contains(e.Description, "bot was blocked by the user") ||
		contains(e.Description, "user is deactivated") ||
		(e.ErrorCode == ErrorCodeForbidden && contains(e.Description, "blocked"))
}

// IsBotKicked checks if the bot was kicked from chat
func (e *TelegramError) IsBotKicked() bool {
	return contains(e.Description, "bot was kicked") ||
		contains(e.Description, "bot is not a member")
}

// IsInvalidFileID checks if the file_id is invalid
func (e *TelegramError) IsInvalidFileID() bool {
	return contains(e.Description, "wrong file identifier") ||
		contains(e.Description, "file_id")
}

// IsButtonDataInvalid checks if callback data is invalid
func (e *TelegramError) IsButtonDataInvalid() bool {
	return contains(e.Description, "BUTTON_DATA_INVALID") ||
		contains(e.Description, "data is too long")
}

// Helper function to check if string contains substring (case-insensitive)
func contains(s, substr string) bool {
	if len(substr) == 0 {
		return true
	}
	if len(s) < len(substr) {
		return false
	}
	
	// Simple case-insensitive search
	for i := 0; i <= len(s)-len(substr); i++ {
		match := true
		for j := 0; j < len(substr); j++ {
			c1 := s[i+j]
			c2 := substr[j]
			// Convert to lowercase for comparison
			if c1 >= 'A' && c1 <= 'Z' {
				c1 += 32
			}
			if c2 >= 'A' && c2 <= 'Z' {
				c2 += 32
			}
			if c1 != c2 {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}
	return false
}

// ErrorFilter filters errors based on a condition
type ErrorFilter func(error) bool

// ErrorHandler represents an error handler with optional filter
type ErrorHandler struct {
	Filter  ErrorFilter      // Optional filter to match specific errors
	Handler ErrorHandlerFunc // Handler function
}

// ErrorHandlers holds all registered error handlers
type ErrorHandlers struct {
	handlers []ErrorHandler
	fallback ErrorHandlerFunc // Default fallback handler
}

// NewErrorHandlers creates a new ErrorHandlers instance
func NewErrorHandlers() *ErrorHandlers {
	return &ErrorHandlers{
		handlers: make([]ErrorHandler, 0),
	}
}

// AddHandler adds a new error handler
func (eh *ErrorHandlers) AddHandler(filter ErrorFilter, handler ErrorHandlerFunc) {
	eh.handlers = append(eh.handlers, ErrorHandler{
		Filter:  filter,
		Handler: handler,
	})
}

// SetFallbackHandler sets a fallback error handler that runs when no other handler matches
func (eh *ErrorHandlers) SetFallbackHandler(handler ErrorHandlerFunc) {
	eh.fallback = handler
}

// Process processes an error through all registered handlers
func (eh *ErrorHandlers) Process(bot *Bot, update *models.Update, err error) error {
	// Try to match specific error handlers
	for _, handler := range eh.handlers {
		if handler.Filter == nil || handler.Filter(err) {
			if handlerErr := handler.Handler(bot, update, err); handlerErr != nil {
				return handlerErr
			}
			return nil // Stop after first matching handler
		}
	}
	
	// If no handler matched and we have a fallback, use it
	if eh.fallback != nil {
		return eh.fallback(bot, update, err)
	}
	
	return err // Return original error if no handler matched
}

// Error Filter Builders

// TelegramErrorFilter creates a filter for Telegram API errors
func TelegramErrorFilter() ErrorFilter {
	return func(err error) bool {
		return IsTelegramError(err)
	}
}

// ErrorCodeFilter creates a filter for specific error codes
func ErrorCodeFilter(code int) ErrorFilter {
	return func(err error) bool {
		if teleErr, ok := err.(*TelegramError); ok {
			return teleErr.ErrorCode == code
		}
		return false
	}
}

// RateLimitErrorFilter creates a filter for rate limit errors (429)
func RateLimitErrorFilter() ErrorFilter {
	return ErrorCodeFilter(ErrorCodeTooManyRequests)
}

// BadRequestErrorFilter creates a filter for bad request errors (400)
func BadRequestErrorFilter() ErrorFilter {
	return ErrorCodeFilter(ErrorCodeBadRequest)
}

// ForbiddenErrorFilter creates a filter for forbidden errors (403)
func ForbiddenErrorFilter() ErrorFilter {
	return ErrorCodeFilter(ErrorCodeForbidden)
}

// UnauthorizedErrorFilter creates a filter for unauthorized errors (401)
func UnauthorizedErrorFilter() ErrorFilter {
	return ErrorCodeFilter(ErrorCodeUnauthorized)
}

// ConflictErrorFilter creates a filter for conflict errors (409)
func ConflictErrorFilter() ErrorFilter {
	return ErrorCodeFilter(ErrorCodeConflict)
}

// ServerErrorFilter creates a filter for server errors (5xx)
func ServerErrorFilter() ErrorFilter {
	return func(err error) bool {
		if teleErr, ok := err.(*TelegramError); ok {
			return teleErr.IsServerError()
		}
		return false
	}
}

// AllErrorsFilter creates a filter that matches all errors
func AllErrorsFilter() ErrorFilter {
	return func(err error) bool {
		return true
	}
}

// Message-specific error filters

// MessageTextEmptyFilter creates a filter for empty message text errors
func MessageTextEmptyFilter() ErrorFilter {
	return func(err error) bool {
		if teleErr, ok := err.(*TelegramError); ok {
			return teleErr.IsMessageTextEmpty()
		}
		return false
	}
}

// MessageTooLongFilter creates a filter for message too long errors
func MessageTooLongFilter() ErrorFilter {
	return func(err error) bool {
		if teleErr, ok := err.(*TelegramError); ok {
			return teleErr.IsMessageTooLong()
		}
		return false
	}
}

// ChatNotFoundFilter creates a filter for chat not found errors
func ChatNotFoundFilter() ErrorFilter {
	return func(err error) bool {
		if teleErr, ok := err.(*TelegramError); ok {
			return teleErr.IsChatNotFound()
		}
		return false
	}
}

// MessageNotFoundFilter creates a filter for message not found errors
func MessageNotFoundFilter() ErrorFilter {
	return func(err error) bool {
		if teleErr, ok := err.(*TelegramError); ok {
			return teleErr.IsMessageNotFound()
		}
		return false
	}
}

// MessageCantBeEditedFilter creates a filter for message can't be edited errors
func MessageCantBeEditedFilter() ErrorFilter {
	return func(err error) bool {
		if teleErr, ok := err.(*TelegramError); ok {
			return teleErr.IsMessageCantBeEdited()
		}
		return false
	}
}

// MessageCantBeDeletedFilter creates a filter for message can't be deleted errors
func MessageCantBeDeletedFilter() ErrorFilter {
	return func(err error) bool {
		if teleErr, ok := err.(*TelegramError); ok {
			return teleErr.IsMessageCantBeDeleted()
		}
		return false
	}
}

// BotBlockedFilter creates a filter for bot was blocked by user errors
func BotBlockedFilter() ErrorFilter {
	return func(err error) bool {
		if teleErr, ok := err.(*TelegramError); ok {
			return teleErr.IsBotWasBlocked()
		}
		return false
	}
}

// BotKickedFilter creates a filter for bot was kicked from chat errors
func BotKickedFilter() ErrorFilter {
	return func(err error) bool {
		if teleErr, ok := err.(*TelegramError); ok {
			return teleErr.IsBotKicked()
		}
		return false
	}
}

// InvalidFileIDFilter creates a filter for invalid file_id errors
func InvalidFileIDFilter() ErrorFilter {
	return func(err error) bool {
		if teleErr, ok := err.(*TelegramError); ok {
			return teleErr.IsInvalidFileID()
		}
		return false
	}
}

// ButtonDataInvalidFilter creates a filter for invalid button data errors
func ButtonDataInvalidFilter() ErrorFilter {
	return func(err error) bool {
		if teleErr, ok := err.(*TelegramError); ok {
			return teleErr.IsButtonDataInvalid()
		}
		return false
	}
}
