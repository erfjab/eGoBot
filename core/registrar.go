package core

// HandlerRegistrar is an interface for types that can register handlers
type HandlerRegistrar interface {
	AddHandler(filter FilterFunc, handler HandlerFunc, opts ...interface{})
}

// ErrorHandlerRegistrar is an interface for types that can register error handlers
type ErrorHandlerRegistrar interface {
	OnError(filter ErrorFilter, handler ErrorHandlerFunc)
	OnTelegramError(handler ErrorHandlerFunc)
	OnRateLimitError(handler ErrorHandlerFunc)
	OnBadRequest(handler ErrorHandlerFunc)
	OnForbiddenError(handler ErrorHandlerFunc)
	SetFallbackErrorHandler(handler ErrorHandlerFunc)
}

// RegisterCommands provides convenience methods for registering handlers
type RegisterCommands struct {
	registrar HandlerRegistrar
}

// NewRegisterCommands creates a new RegisterCommands instance
func NewRegisterCommands(registrar HandlerRegistrar) *RegisterCommands {
	return &RegisterCommands{registrar: registrar}
}

// OnCommand registers a handler for a specific command
func (r *RegisterCommands) OnCommand(command string, handler HandlerFunc, opts ...interface{}) {
	r.registrar.AddHandler(CommandFilter(command), handler, opts...)
}

// OnMessage registers a handler for all messages
func (r *RegisterCommands) OnMessage(handler HandlerFunc, opts ...interface{}) {
	r.registrar.AddHandler(MessageFilter(), handler, opts...)
}

// OnText registers a handler for text messages (non-command)
func (r *RegisterCommands) OnText(handler HandlerFunc, opts ...interface{}) {
	r.registrar.AddHandler(TextFilter(), handler, opts...)
}

// OnCallbackQuery registers a handler for all callback queries
func (r *RegisterCommands) OnCallbackQuery(handler HandlerFunc, opts ...interface{}) {
	r.registrar.AddHandler(CallbackQueryFilter(), handler, opts...)
}

// OnCallbackData registers a handler for callback queries with specific data
func (r *RegisterCommands) OnCallbackData(data string, handler HandlerFunc, opts ...interface{}) {
	r.registrar.AddHandler(CallbackDataFilter(data), handler, opts...)
}

// OnCallbackDataPrefix registers a handler for callback queries with data starting with prefix
func (r *RegisterCommands) OnCallbackDataPrefix(prefix string, handler HandlerFunc, opts ...interface{}) {
	r.registrar.AddHandler(CallbackDataPrefixFilter(prefix), handler, opts...)
}

// OnPhoto registers a handler for photo messages
func (r *RegisterCommands) OnPhoto(handler HandlerFunc, opts ...interface{}) {
	r.registrar.AddHandler(PhotoFilter(), handler, opts...)
}

// OnDocument registers a handler for document messages
func (r *RegisterCommands) OnDocument(handler HandlerFunc, opts ...interface{}) {
	r.registrar.AddHandler(DocumentFilter(), handler, opts...)
}

// OnVideo registers a handler for video messages
func (r *RegisterCommands) OnVideo(handler HandlerFunc, opts ...interface{}) {
	r.registrar.AddHandler(VideoFilter(), handler, opts...)
}

// OnAudio registers a handler for audio messages
func (r *RegisterCommands) OnAudio(handler HandlerFunc, opts ...interface{}) {
	r.registrar.AddHandler(AudioFilter(), handler, opts...)
}

// OnVoice registers a handler for voice messages
func (r *RegisterCommands) OnVoice(handler HandlerFunc, opts ...interface{}) {
	r.registrar.AddHandler(VoiceFilter(), handler, opts...)
}

// OnSticker registers a handler for sticker messages
func (r *RegisterCommands) OnSticker(handler HandlerFunc, opts ...interface{}) {
	r.registrar.AddHandler(StickerFilter(), handler, opts...)
}

// OnLocation registers a handler for location messages
func (r *RegisterCommands) OnLocation(handler HandlerFunc, opts ...interface{}) {
	r.registrar.AddHandler(LocationFilter(), handler, opts...)
}

// OnContact registers a handler for contact messages
func (r *RegisterCommands) OnContact(handler HandlerFunc, opts ...interface{}) {
	r.registrar.AddHandler(ContactFilter(), handler, opts...)
}

// OnEditedMessage registers a handler for edited messages
func (r *RegisterCommands) OnEditedMessage(handler HandlerFunc, opts ...interface{}) {
	r.registrar.AddHandler(EditedMessageFilter(), handler, opts...)
}

// OnInlineQuery registers a handler for inline queries
func (r *RegisterCommands) OnInlineQuery(handler HandlerFunc, opts ...interface{}) {
	r.registrar.AddHandler(InlineQueryFilter(), handler, opts...)
}

// OnChannelPost registers a handler for channel posts
func (r *RegisterCommands) OnChannelPost(handler HandlerFunc, opts ...interface{}) {
	r.registrar.AddHandler(ChannelPostFilter(), handler, opts...)
}
