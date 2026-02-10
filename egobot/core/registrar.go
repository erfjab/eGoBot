package core

// HandlerRegistrar is an interface for types that can register handlers
type HandlerRegistrar interface {
	AddHandler(filter FilterFunc, handler HandlerFunc)
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
func (r *RegisterCommands) OnCommand(command string, handler HandlerFunc) {
	r.registrar.AddHandler(CommandFilter(command), handler)
}

// OnMessage registers a handler for all messages
func (r *RegisterCommands) OnMessage(handler HandlerFunc) {
	r.registrar.AddHandler(MessageFilter(), handler)
}

// OnText registers a handler for text messages (non-command)
func (r *RegisterCommands) OnText(handler HandlerFunc) {
	r.registrar.AddHandler(TextFilter(), handler)
}

// OnCallbackQuery registers a handler for all callback queries
func (r *RegisterCommands) OnCallbackQuery(handler HandlerFunc) {
	r.registrar.AddHandler(CallbackQueryFilter(), handler)
}

// OnCallbackData registers a handler for callback queries with specific data
func (r *RegisterCommands) OnCallbackData(data string, handler HandlerFunc) {
	r.registrar.AddHandler(CallbackDataFilter(data), handler)
}

// OnCallbackDataPrefix registers a handler for callback queries with data starting with prefix
func (r *RegisterCommands) OnCallbackDataPrefix(prefix string, handler HandlerFunc) {
	r.registrar.AddHandler(CallbackDataPrefixFilter(prefix), handler)
}

// OnPhoto registers a handler for photo messages
func (r *RegisterCommands) OnPhoto(handler HandlerFunc) {
	r.registrar.AddHandler(PhotoFilter(), handler)
}

// OnDocument registers a handler for document messages
func (r *RegisterCommands) OnDocument(handler HandlerFunc) {
	r.registrar.AddHandler(DocumentFilter(), handler)
}

// OnVideo registers a handler for video messages
func (r *RegisterCommands) OnVideo(handler HandlerFunc) {
	r.registrar.AddHandler(VideoFilter(), handler)
}

// OnAudio registers a handler for audio messages
func (r *RegisterCommands) OnAudio(handler HandlerFunc) {
	r.registrar.AddHandler(AudioFilter(), handler)
}

// OnVoice registers a handler for voice messages
func (r *RegisterCommands) OnVoice(handler HandlerFunc) {
	r.registrar.AddHandler(VoiceFilter(), handler)
}

// OnSticker registers a handler for sticker messages
func (r *RegisterCommands) OnSticker(handler HandlerFunc) {
	r.registrar.AddHandler(StickerFilter(), handler)
}

// OnLocation registers a handler for location messages
func (r *RegisterCommands) OnLocation(handler HandlerFunc) {
	r.registrar.AddHandler(LocationFilter(), handler)
}

// OnContact registers a handler for contact messages
func (r *RegisterCommands) OnContact(handler HandlerFunc) {
	r.registrar.AddHandler(ContactFilter(), handler)
}

// OnEditedMessage registers a handler for edited messages
func (r *RegisterCommands) OnEditedMessage(handler HandlerFunc) {
	r.registrar.AddHandler(EditedMessageFilter(), handler)
}

// OnInlineQuery registers a handler for inline queries
func (r *RegisterCommands) OnInlineQuery(handler HandlerFunc) {
	r.registrar.AddHandler(InlineQueryFilter(), handler)
}

// OnChannelPost registers a handler for channel posts
func (r *RegisterCommands) OnChannelPost(handler HandlerFunc) {
	r.registrar.AddHandler(ChannelPostFilter(), handler)
}
