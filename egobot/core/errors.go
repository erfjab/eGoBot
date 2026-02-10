package core

import (
	"egobot/egobot/models"
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
