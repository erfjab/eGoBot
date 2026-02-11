package core

import (
	"github.com/erfjab/egobot/models"
)

// MiddlewareFunc represents a middleware function
// It receives the bot, update, and a next function
// Call next() to continue to the next middleware or handler
// Return without calling next() to stop the chain
type MiddlewareFunc func(*Bot, *models.Update, NextFunc)

// NextFunc is the function to call the next middleware or handler in the chain
type NextFunc func()

// MiddlewareChain manages a chain of middlewares
type MiddlewareChain struct {
	middlewares []MiddlewareFunc
	handler     HandlerFunc
}

// NewMiddlewareChain creates a new middleware chain
func NewMiddlewareChain(handler HandlerFunc, middlewares ...MiddlewareFunc) *MiddlewareChain {
	return &MiddlewareChain{
		middlewares: middlewares,
		handler:     handler,
	}
}

// Execute executes the middleware chain
func (mc *MiddlewareChain) Execute(bot *Bot, update *models.Update) error {
	if len(mc.middlewares) == 0 {
		// No middlewares, just execute the handler
		return mc.handler(bot, update)
	}

	var currentIndex int
	var err error

	// Define the next function
	var next NextFunc
	next = func() {
		if currentIndex < len(mc.middlewares) {
			// Execute current middleware
			middleware := mc.middlewares[currentIndex]
			currentIndex++
			middleware(bot, update, next)
		} else {
			// All middlewares passed, execute the handler
			err = mc.handler(bot, update)
		}
	}

	// Start the chain
	next()
	return err
}
