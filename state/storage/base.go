package storage

import (
	"context"
)

// UserContext contains user state and data
type UserContext struct {
	State string                 `json:"state"`
	Data  map[string]interface{} `json:"data"`
}

// BaseStorage defines the interface for state storage backends
type BaseStorage interface {
	// GetContext retrieves the complete user context (state + data)
	GetContext(ctx context.Context, key string) (*UserContext, error)

	// GetState retrieves only the state for a user
	GetState(ctx context.Context, key string) (string, error)

	// SetState sets the state for a user
	SetState(ctx context.Context, key string, state string) error

	// ClearState clears the state for a user
	ClearState(ctx context.Context, key string) error

	// GetData retrieves only the data for a user
	GetData(ctx context.Context, key string) (map[string]interface{}, error)

	// UpsertData updates or inserts data for a user
	UpsertData(ctx context.Context, key string, data map[string]interface{}) error

	// ClearData clears all data for a user
	ClearData(ctx context.Context, key string) error

	// UpsertContext updates or inserts the complete user context
	UpsertContext(ctx context.Context, key string, state string, data map[string]interface{}) error

	// ClearAll removes all data and state for a user
	ClearAll(ctx context.Context, key string) error

	// Close closes the storage connection (if applicable)
	Close() error
}
