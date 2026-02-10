package core

// HandlerGroup represents a group of related handlers
type HandlerGroup struct {
	name     string
	handlers []Handler
	filter   FilterFunc
	*RegisterCommands
}

// NewHandlerGroup creates a new handler group
func NewHandlerGroup(name string) *HandlerGroup {
	group := &HandlerGroup{
		name:     name,
		handlers: make([]Handler, 0),
	}
	group.RegisterCommands = NewRegisterCommands(group)
	return group
}

// WithFilter sets a filter for all handlers in this group
func (g *HandlerGroup) WithFilter(filter FilterFunc) *HandlerGroup {
	g.filter = filter
	return g
}

// AddHandler adds a custom handler with a filter to the group
func (g *HandlerGroup) AddHandler(filter FilterFunc, handler HandlerFunc) {
	// Apply group filter if exists
	finalFilter := filter
	if g.filter != nil {
		finalFilter = AndFilter(g.filter, filter)
	}
	
	g.handlers = append(g.handlers, Handler{
		Filter:  finalFilter,
		Handler: handler,
	})
}

// Handlers returns all handlers in this group
func (g *HandlerGroup) Handlers() []Handler {
	return g.handlers
}

// Name returns the name of this group
func (g *HandlerGroup) Name() string {
	return g.name
}
