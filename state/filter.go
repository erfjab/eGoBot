package state

import (
	"context"
)

// Filter represents a state filter that can be used in handlers
type Filter struct {
	states        []*State
	ignoreState   bool
	matchAny      bool // if true, matches any of the states; if false, requires exact match
	allowNoState  bool // if true, allows execution when no state is set
}

// InState creates a state filter for specific states
// Usage: bot.AddHandler(filter, handler, state.InState(MyState))
func InState(states ...*State) *Filter {
	return &Filter{
		states:   states,
		matchAny: len(states) > 1,
	}
}

// IgnoreState creates a filter that ignores state (always passes)
// Usage: bot.AddHandler(filter, handler, state.IgnoreState())
func IgnoreState() *Filter {
	return &Filter{
		ignoreState: true,
	}
}

// NoState creates a filter that only matches when no state is set
// Usage: bot.AddHandler(filter, handler, state.NoState())
func NoState() *Filter {
	return &Filter{
		allowNoState: true,
	}
}

// Deprecated: Use InState instead
func NewStateFilter(states ...*State) *Filter {
	return InState(states...)
}

// Deprecated: Use IgnoreState instead
func NewIgnoreStateFilter() *Filter {
	return IgnoreState()
}

// Deprecated: Use NoState instead
func NewNoStateFilter() *Filter {
	return NoState()
}

// Check checks if the given state matches the filter
func (f *Filter) Check(currentState *State) bool {
	// If we ignore state, always return true
	if f.ignoreState {
		return true
	}

	// If current state is nil
	if currentState == nil || currentState.Name == "" {
		return f.allowNoState || len(f.states) == 0
	}

	// If allowNoState is true but state is not nil, reject
	if f.allowNoState && len(f.states) == 0 {
		return false
	}

	// If no specific states are required, allow any state
	if len(f.states) == 0 {
		return true
	}

	// Check if current state matches any of the filter states
	for _, state := range f.states {
		if state.Equals(currentState) {
			return true
		}
	}

	return false
}

// MatchesAny returns true if the filter matches any of the given states
func (f *Filter) MatchesAny() bool {
	return f.matchAny
}

// GetStates returns the states this filter matches
func (f *Filter) GetStates() []*State {
	return f.states
}

// IsIgnoreState returns true if this filter ignores state
func (f *Filter) IsIgnoreState() bool {
	return f.ignoreState
}

// AllowNoState returns true if this filter allows no state
func (f *Filter) AllowNoState() bool {
	return f.allowNoState
}

// Middleware creates a middleware function for the state filter
func (f *Filter) Middleware(stateManager *Manager) func(next func(ctx context.Context) error) func(ctx context.Context) error {
	return func(next func(ctx context.Context) error) func(ctx context.Context) error {
		return func(ctx context.Context) error {
			// Extract user ID from context (should be set by previous middleware)
			userID := ctx.Value("user_id")
			if userID == nil {
				// If no user ID, skip state check
				return next(ctx)
			}

			// Get user's current state
			userManager := stateManager.ForUser(userID)
			currentState, err := userManager.GetState(ctx)
			if err != nil {
				return err
			}

			// Check if state matches filter
			if !f.Check(currentState) {
				// State doesn't match, skip this handler
				return nil
			}

			// State matches, continue to handler
			return next(ctx)
		}
	}
}

// Or combines this filter with another using OR logic
func (f *Filter) Or(other *Filter) *Filter {
	combined := &Filter{
		states:   append(f.states, other.states...),
		matchAny: true,
	}
	if f.ignoreState || other.ignoreState {
		combined.ignoreState = true
	}
	if f.allowNoState || other.allowNoState {
		combined.allowNoState = true
	}
	return combined
}

// And combines this filter with another using AND logic (both must match)
func (f *Filter) And(other *Filter) *Filter {
	return &Filter{
		states:       f.states, // For AND, we keep only the first filter's states
		ignoreState:  f.ignoreState && other.ignoreState,
		matchAny:     false,
		allowNoState: f.allowNoState && other.allowNoState,
	}
}
