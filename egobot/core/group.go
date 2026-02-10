package core

// HandlerGroup represents a group of related handlers
type HandlerGroup struct {
	name        string
	handlers    []Handler
	filter      FilterFunc
	middlewares []MiddlewareFunc
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

// UseMiddleware adds middleware(s) to this group
// All handlers in this group will use these middlewares
func (g *HandlerGroup) UseMiddleware(middlewares ...MiddlewareFunc) *HandlerGroup {
	g.middlewares = append(g.middlewares, middlewares...)
	return g
}

// AddHandler adds a custom handler with a filter to the group
func (g *HandlerGroup) AddHandler(filter FilterFunc, handler HandlerFunc, middlewares ...MiddlewareFunc) {
	// Apply group filter if exists
	finalFilter := filter
	if g.filter != nil {
		finalFilter = AndFilter(g.filter, filter)
	}
	
	// Combine group middlewares with handler-specific middlewares
	allMiddlewares := append([]MiddlewareFunc{}, g.middlewares...)
	allMiddlewares = append(allMiddlewares, middlewares...)
	
	g.handlers = append(g.handlers, Handler{
		Filter:      finalFilter,
		Handler:     handler,
		Middlewares: allMiddlewares,
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
