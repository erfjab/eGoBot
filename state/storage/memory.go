package storage

import (
	"context"
	"sync"
)

// MemoryStorage implements BaseStorage using in-memory storage
type MemoryStorage struct {
	mu      sync.RWMutex
	storage map[string]*UserContext
}

// NewMemoryStorage creates a new memory storage instance
func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		storage: make(map[string]*UserContext),
	}
}

// GetContext retrieves the complete user context (state + data)
func (m *MemoryStorage) GetContext(ctx context.Context, key string) (*UserContext, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if uc, ok := m.storage[key]; ok {
		// Return a copy to prevent external modifications
		data := make(map[string]interface{})
		for k, v := range uc.Data {
			data[k] = v
		}
		return &UserContext{
			State: uc.State,
			Data:  data,
		}, nil
	}

	// Return empty context if not found
	return &UserContext{
		State: "",
		Data:  make(map[string]interface{}),
	}, nil
}

// GetState retrieves only the state for a user
func (m *MemoryStorage) GetState(ctx context.Context, key string) (string, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if uc, ok := m.storage[key]; ok {
		return uc.State, nil
	}
	return "", nil
}

// SetState sets the state for a user
func (m *MemoryStorage) SetState(ctx context.Context, key string, state string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if uc, ok := m.storage[key]; ok {
		uc.State = state
	} else {
		m.storage[key] = &UserContext{
			State: state,
			Data:  make(map[string]interface{}),
		}
	}
	return nil
}

// ClearState clears the state for a user
func (m *MemoryStorage) ClearState(ctx context.Context, key string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if uc, ok := m.storage[key]; ok {
		uc.State = ""
	}
	return nil
}

// GetData retrieves only the data for a user
func (m *MemoryStorage) GetData(ctx context.Context, key string) (map[string]interface{}, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if uc, ok := m.storage[key]; ok {
		// Return a copy to prevent external modifications
		data := make(map[string]interface{})
		for k, v := range uc.Data {
			data[k] = v
		}
		return data, nil
	}
	return make(map[string]interface{}), nil
}

// UpsertData updates or inserts data for a user
func (m *MemoryStorage) UpsertData(ctx context.Context, key string, data map[string]interface{}) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if uc, ok := m.storage[key]; ok {
		// Merge new data with existing data
		for k, v := range data {
			uc.Data[k] = v
		}
	} else {
		// Create new entry
		newData := make(map[string]interface{})
		for k, v := range data {
			newData[k] = v
		}
		m.storage[key] = &UserContext{
			State: "",
			Data:  newData,
		}
	}
	return nil
}

// ClearData clears all data for a user
func (m *MemoryStorage) ClearData(ctx context.Context, key string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if uc, ok := m.storage[key]; ok {
		uc.Data = make(map[string]interface{})
	}
	return nil
}

// UpsertContext updates or inserts the complete user context
func (m *MemoryStorage) UpsertContext(ctx context.Context, key string, state string, data map[string]interface{}) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if uc, ok := m.storage[key]; ok {
		// Update existing context
		if state != "" {
			uc.State = state
		}
		for k, v := range data {
			uc.Data[k] = v
		}
	} else {
		// Create new context
		newData := make(map[string]interface{})
		for k, v := range data {
			newData[k] = v
		}
		m.storage[key] = &UserContext{
			State: state,
			Data:  newData,
		}
	}
	return nil
}

// ClearAll removes all data and state for a user
func (m *MemoryStorage) ClearAll(ctx context.Context, key string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.storage, key)
	return nil
}

// Close closes the storage (no-op for memory storage)
func (m *MemoryStorage) Close() error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.storage = make(map[string]*UserContext)
	return nil
}

// Count returns the number of users in storage (useful for testing/debugging)
func (m *MemoryStorage) Count() int {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return len(m.storage)
}
