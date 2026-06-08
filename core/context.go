package core

import (
	"context"
	"sync"
	"time"
)

// Context provides a way to store and retrieve data during request processing.
// It implements the standard context.Context interface so it can be passed
// directly to database drivers, HTTP clients, and any function that accepts
// context.Context — just like fiber.Ctx does.
//
// It wraps a parent context.Context (defaults to context.Background()) and
// adds a thread-safe data store for passing values between middlewares and
// handlers.
type Context struct {
	context.Context
	data map[string]interface{}
	mu   sync.RWMutex
}

// NewContext creates a new Context instance with context.Background() as parent.
func NewContext() *Context {
	return NewContextWithParent(context.Background())
}

// NewContextWithParent creates a new Context with the given parent context.
// Use this when you have an existing context.Context (e.g. from an HTTP request,
// fiber.Ctx, or a database transaction) that you want to propagate cancellation
// and deadlines through.
func NewContextWithParent(parent context.Context) *Context {
	return &Context{
		Context: parent,
		data:    make(map[string]interface{}),
	}
}

// Set stores a value in the context's local data store.
// This is separate from the parent context's Value() chain and is thread-safe.
func (c *Context) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}

// Get retrieves a value from the context's local data store.
// Returns nil if the key doesn't exist.
func (c *Context) Get(key string) interface{} {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.data[key]
}

// GetString retrieves a string value from the context.
// Returns empty string if the key doesn't exist or value is not a string.
func (c *Context) GetString(key string) string {
	val := c.Get(key)
	if str, ok := val.(string); ok {
		return str
	}
	return ""
}

// GetInt retrieves an int value from the context.
// Returns 0 if the key doesn't exist or value is not an int.
func (c *Context) GetInt(key string) int {
	val := c.Get(key)
	if i, ok := val.(int); ok {
		return i
	}
	return 0
}

// GetInt64 retrieves an int64 value from the context.
// Returns 0 if the key doesn't exist or value is not an int64.
func (c *Context) GetInt64(key string) int64 {
	val := c.Get(key)
	if i, ok := val.(int64); ok {
		return i
	}
	return 0
}

// GetBool retrieves a bool value from the context.
// Returns false if the key doesn't exist or value is not a bool.
func (c *Context) GetBool(key string) bool {
	val := c.Get(key)
	if b, ok := val.(bool); ok {
		return b
	}
	return false
}

// Has checks if a key exists in the context's local data store.
func (c *Context) Has(key string) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	_, exists := c.data[key]
	return exists
}

// Delete removes a key from the context's local data store.
func (c *Context) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.data, key)
}

// Clear removes all data from the context's local data store.
func (c *Context) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data = make(map[string]interface{})
}

// Keys returns all keys in the context's local data store.
func (c *Context) Keys() []string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	keys := make([]string, 0, len(c.data))
	for k := range c.data {
		keys = append(keys, k)
	}
	return keys
}

// Deadline returns the deadline of the parent context.
// Implements context.Context.
func (c *Context) Deadline() (deadline time.Time, ok bool) {
	return c.Context.Deadline()
}

// Done returns the Done channel of the parent context.
// Implements context.Context.
func (c *Context) Done() <-chan struct{} {
	return c.Context.Done()
}

// Err returns the error of the parent context.
// Implements context.Context.
func (c *Context) Err() error {
	return c.Context.Err()
}

// Value looks up a key in the context chain. It first checks the local data
// store (for string keys), then falls back to the parent context.
// Implements context.Context.
func (c *Context) Value(key any) any {
	// First try the local data store (only for string keys)
	if k, ok := key.(string); ok {
		c.mu.RLock()
		if val, exists := c.data[k]; exists {
			c.mu.RUnlock()
			return val
		}
		c.mu.RUnlock()
	}
	// Fall back to parent context
	return c.Context.Value(key)
}

// SetParentContext replaces the parent context. This is useful when you need
// to inject a different context (e.g., with a timeout or cancellation) after
// the Context has been created.
func (c *Context) SetParentContext(parent context.Context) {
	c.Context = parent
}

// GetStateData is a convenience method that returns the user state name and
// data map from the context, as injected by the framework during update processing.
func (c *Context) GetStateData() (state string, data map[string]interface{}) {
	state = c.GetString("state")
	if d, ok := c.Get("data").(map[string]interface{}); ok {
		data = d
	}
	return
}

// GetStateObj returns the state object from the context if present.
func (c *Context) GetStateObj() interface{} {
	return c.Get("state_obj")
}

// GetStateName returns the user's current state name.
// Returns empty string if no state exists.
func (c *Context) GetStateName() string {
	return c.GetString("state")
}

func (c *Context) setCallbackData(value interface{}) {
	c.Set(callbackDataContextKey, value)
}

// GetCallbackData returns parsed callback payload stored by OnCallbackStruct.
func (c *Context) GetCallbackData() interface{} {
	return c.Get(callbackDataContextKey)
}

// LoadCallbackData loads parsed callback payload from context into out.
// out must be a non-nil pointer to struct.
func (c *Context) LoadCallbackData(out interface{}) bool {
	if c == nil {
		return false
	}
	return loadCallbackData(out, c.GetCallbackData())
}

// MatchCallbackData checks whether parsed callback payload in context matches
// non-zero exported fields of expect.
//
// expect must be a struct (or pointer to struct) of the same type as payload.
// Zero-value fields in expect are ignored.
func (c *Context) MatchCallbackData(expect interface{}) bool {
	if c == nil {
		return false
	}
	return callbackStructPatternMatches(c.GetCallbackData(), expect)
}
