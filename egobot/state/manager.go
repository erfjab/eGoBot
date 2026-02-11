package state

import (
	"context"
	"fmt"
	"strconv"

	"egobot/egobot/state/storage"
)

// Manager handles state management for users
type Manager struct {
	storage storage.BaseStorage
}

// NewManager creates a new state manager with the given storage backend
func NewManager(store storage.BaseStorage) *Manager {
	if store == nil {
		store = storage.NewMemoryStorage()
	}
	return &Manager{
		storage: store,
	}
}

// GetStorage returns the underlying storage
func (m *Manager) GetStorage() storage.BaseStorage {
	return m.storage
}

// ForUser returns a user-specific context manager
func (m *Manager) ForUser(userID interface{}) *UserStateManager {
	key := m.getUserKey(userID)
	return &UserStateManager{
		key:     key,
		storage: m.storage,
	}
}

// getUserKey converts various user ID types to string key
func (m *Manager) getUserKey(userID interface{}) string {
	switch v := userID.(type) {
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	case int64:
		return strconv.FormatInt(v, 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	default:
		return fmt.Sprintf("%v", v)
	}
}

// UserStateManager manages state for a specific user
type UserStateManager struct {
	key     string
	storage storage.BaseStorage
}

// GetContext retrieves the complete user context
func (u *UserStateManager) GetContext(ctx context.Context) (*storage.UserContext, error) {
	return u.storage.GetContext(ctx, u.key)
}

// GetState retrieves the current state
func (u *UserStateManager) GetState(ctx context.Context) (*State, error) {
	stateName, err := u.storage.GetState(ctx, u.key)
	if err != nil {
		return nil, err
	}
	if stateName == "" {
		return nil, nil
	}
	return NewState(stateName), nil
}

// SetState sets the user's state
func (u *UserStateManager) SetState(ctx context.Context, state *State) error {
	if state == nil {
		return u.storage.ClearState(ctx, u.key)
	}
	return u.storage.SetState(ctx, u.key, state.Name)
}

// ClearState clears the user's state
func (u *UserStateManager) ClearState(ctx context.Context) error {
	return u.storage.ClearState(ctx, u.key)
}

// GetData retrieves the user's data
func (u *UserStateManager) GetData(ctx context.Context) (map[string]interface{}, error) {
	return u.storage.GetData(ctx, u.key)
}

// SetData sets specific data fields for the user
func (u *UserStateManager) SetData(ctx context.Context, data map[string]interface{}) error {
	return u.storage.UpsertData(ctx, u.key, data)
}

// GetDataValue retrieves a specific data value
func (u *UserStateManager) GetDataValue(ctx context.Context, key string) (interface{}, error) {
	data, err := u.storage.GetData(ctx, u.key)
	if err != nil {
		return nil, err
	}
	return data[key], nil
}

// SetDataValue sets a specific data value
func (u *UserStateManager) SetDataValue(ctx context.Context, key string, value interface{}) error {
	return u.storage.UpsertData(ctx, u.key, map[string]interface{}{key: value})
}

// ClearData clears all user data
func (u *UserStateManager) ClearData(ctx context.Context) error {
	return u.storage.ClearData(ctx, u.key)
}

// UpdateContext updates the complete user context
func (u *UserStateManager) UpdateContext(ctx context.Context, state *State, data map[string]interface{}) error {
	stateName := ""
	if state != nil {
		stateName = state.Name
	}
	return u.storage.UpsertContext(ctx, u.key, stateName, data)
}

// ClearAll clears all state and data for the user
func (u *UserStateManager) ClearAll(ctx context.Context) error {
	return u.storage.ClearAll(ctx, u.key)
}

// Key returns the user's storage key
func (u *UserStateManager) Key() string {
	return u.key
}
