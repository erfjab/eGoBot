package state

import (
	"fmt"
	"reflect"
	"strings"
)

// State represents a single state in the state machine
type State struct {
	Name string
}

// NewState creates a new state with the given name
func NewState(name string) *State {
	return &State{Name: name}
}

// String returns the state name
func (s *State) String() string {
	return s.Name
}

// Equals checks if two states are equal
func (s *State) Equals(other interface{}) bool {
	switch v := other.(type) {
	case *State:
		return s.Name == v.Name
	case string:
		return s.Name == v
	default:
		return false
	}
}

// StateGroup is a container for related states
type StateGroup struct {
	name   string
	states map[string]*State
}

// NewStateGroup creates a new state group with the given name
func NewStateGroup(name string) *StateGroup {
	return &StateGroup{
		name:   name,
		states: make(map[string]*State),
	}
}

// Add adds a state to the group
func (sg *StateGroup) Add(stateName string) *State {
	fullName := fmt.Sprintf("%s.%s", sg.name, stateName)
	state := NewState(fullName)
	sg.states[stateName] = state
	return state
}

// Get retrieves a state from the group by name
func (sg *StateGroup) Get(stateName string) (*State, bool) {
	state, ok := sg.states[stateName]
	return state, ok
}

// GetStates returns all states in the group
func (sg *StateGroup) GetStates() map[string]*State {
	return sg.states
}

// Name returns the group name
func (sg *StateGroup) Name() string {
	return sg.name
}

// StateGroupFromStruct creates a StateGroup from a struct
// Each exported field becomes a state
func StateGroupFromStruct(v interface{}) *StateGroup {
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		panic("StateGroupFromStruct requires a struct type")
	}

	groupName := t.Name()
	sg := NewStateGroup(groupName)

	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.IsExported() && !strings.HasPrefix(field.Name, "_") {
			state := sg.Add(field.Name)
			// Set the field value to the state pointer
			fieldVal := val.Field(i)
			if fieldVal.CanSet() && fieldVal.Type() == reflect.TypeOf(&State{}) {
				fieldVal.Set(reflect.ValueOf(state))
			}
		}
	}

	return sg
}
