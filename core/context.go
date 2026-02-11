package core

import (
	"sync"
)

// Context provides a way to store and retrieve data during request processing
// It allows middlewares to pass data to handlers
type Context struct {
	data map[string]interface{}
	mu   sync.RWMutex
}

// NewContext creates a new Context instance
func NewContext() *Context {
	return &Context{
		data: make(map[string]interface{}),
	}
}

// Set stores a value in the context
func (c *Context) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}

// Get retrieves a value from the context
// Returns nil if the key doesn't exist
func (c *Context) Get(key string) interface{} {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.data[key]
}

// GetString retrieves a string value from the context
// Returns empty string if the key doesn't exist or value is not a string
func (c *Context) GetString(key string) string {
	val := c.Get(key)
	if str, ok := val.(string); ok {
		return str
	}
	return ""
}

// GetInt retrieves an int value from the context
// Returns 0 if the key doesn't exist or value is not an int
func (c *Context) GetInt(key string) int {
	val := c.Get(key)
	if i, ok := val.(int); ok {
		return i
	}
	return 0
}

// GetInt64 retrieves an int64 value from the context
// Returns 0 if the key doesn't exist or value is not an int64
func (c *Context) GetInt64(key string) int64 {
	val := c.Get(key)
	if i, ok := val.(int64); ok {
		return i
	}
	return 0
}

// GetBool retrieves a bool value from the context
// Returns false if the key doesn't exist or value is not a bool
func (c *Context) GetBool(key string) bool {
	val := c.Get(key)
	if b, ok := val.(bool); ok {
		return b
	}
	return false
}

// Has checks if a key exists in the context
func (c *Context) Has(key string) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	_, exists := c.data[key]
	return exists
}

// Delete removes a key from the context
func (c *Context) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.data, key)
}

// Clear removes all data from the context
func (c *Context) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data = make(map[string]interface{})
}

// Keys returns all keys in the context
func (c *Context) Keys() []string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	keys := make([]string, 0, len(c.data))
	for k := range c.data {
		keys = append(keys, k)
	}
	return keys
}
