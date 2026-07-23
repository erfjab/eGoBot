package core

import (
	"context"
	"log"
	"strings"

	"github.com/erfjab/egobot/models"
	"github.com/erfjab/egobot/state"
	"github.com/erfjab/egobot/state/storage"
)

// HandlerFunc represents a function that handles an update
type HandlerFunc func(*Bot, *models.Update, *Context) error

// Handler represents a handler with its filter and handler function
type Handler struct {
	Filter      FilterFunc
	Handler     HandlerFunc
	Middlewares []MiddlewareFunc
	StateFilter *state.Filter // Optional state filter
}

// FilterFunc represents a function that filters updates
type FilterFunc func(*models.Update) bool

// Handlers holds all registered handlers
type Handlers struct {
	handlers []Handler
}

// NewHandlers creates a new Handlers instance
func NewHandlers() *Handlers {
	return &Handlers{
		handlers: make([]Handler, 0),
	}
}

// AddHandler adds a new handler
func (h *Handlers) AddHandler(filter FilterFunc, handler HandlerFunc, middlewares ...MiddlewareFunc) {
	h.handlers = append(h.handlers, Handler{
		Filter:      filter,
		Handler:     handler,
		Middlewares: middlewares,
		StateFilter: nil, // No state filter by default
	})
}

// AddHandlerWithState adds a new handler with state filter
func (h *Handlers) AddHandlerWithState(filter FilterFunc, stateFilter *state.Filter, handler HandlerFunc, middlewares ...MiddlewareFunc) {
	h.handlers = append(h.handlers, Handler{
		Filter:      filter,
		Handler:     handler,
		Middlewares: middlewares,
		StateFilter: stateFilter,
	})
}

// Process processes an update through all handlers
func (h *Handlers) Process(bot *Bot, update *models.Update) {
	// Get user ID for state checking
	var userID interface{}
	if update.Message != nil && update.Message.From != nil {
		userID = update.Message.From.ID
	} else if update.CallbackQuery != nil {
		userID = update.CallbackQuery.From.ID
	} else if update.EditedMessage != nil && update.EditedMessage.From != nil {
		userID = update.EditedMessage.From.ID
	} else if update.InlineQuery != nil {
		userID = update.InlineQuery.From.ID
	} else if update.BusinessMessage != nil && update.BusinessMessage.From != nil {
		userID = update.BusinessMessage.From.ID
	} else if update.EditedBusinessMessage != nil && update.EditedBusinessMessage.From != nil {
		userID = update.EditedBusinessMessage.From.ID
	} else if update.GuestMessage != nil && update.GuestMessage.From != nil {
		userID = update.GuestMessage.From.ID
	} else if update.ChannelPost != nil && update.ChannelPost.From != nil {
		userID = update.ChannelPost.From.ID
	} else if update.EditedChannelPost != nil && update.EditedChannelPost.From != nil {
		userID = update.EditedChannelPost.From.ID
	} else if update.ShippingQuery != nil {
		userID = update.ShippingQuery.From.ID
	} else if update.PreCheckoutQuery != nil {
		userID = update.PreCheckoutQuery.From.ID
	} else if update.PurchasedPaidMedia != nil {
		userID = update.PurchasedPaidMedia.From.ID
	} else if update.MyChatMember != nil {
		userID = update.MyChatMember.From.ID
	} else if update.ChatMember != nil {
		userID = update.ChatMember.From.ID
	} else if update.ChatJoinRequest != nil {
		userID = update.ChatJoinRequest.From.ID
	} else if update.BusinessConnection != nil {
		userID = update.BusinessConnection.User.ID
	} else if update.PollAnswer != nil && update.PollAnswer.User != nil {
		userID = update.PollAnswer.User.ID
	} else if update.Subscription != nil {
		userID = update.Subscription.User.ID
	}

	for _, handler := range h.handlers {
		if handler.Filter(update) {
			// Check state filter if present and load user context
			var userContext *storage.UserContext
			if handler.StateFilter != nil && userID != nil {
				ctx := context.Background()
				userManager := bot.StateManager.ForUser(userID)
				var err error
				userContext, err = userManager.GetContext(ctx)
				if err != nil {
					log.Printf("Error getting user context: %v", err)
					continue
				}

				// Check if state matches
				var currentState *state.State
				if userContext != nil && userContext.State != "" {
					currentState = state.NewState(userContext.State)
				}
				if !handler.StateFilter.Check(currentState) {
					// State doesn't match, skip this handler
					continue
				}
			}

			var err error
			// Create context and inject user state/data if available
			handlerCtx := NewContext()
			if userContext != nil {
				handlerCtx.Set("state", userContext.State)
				handlerCtx.Set("data", userContext.Data)
				if userContext.State != "" {
					handlerCtx.Set("state_obj", state.NewState(userContext.State))
				}
			}

			// Execute with middleware chain
			if len(handler.Middlewares) > 0 {
				chain := NewMiddlewareChainWithContext(handler.Handler, handlerCtx, handler.Middlewares...)
				err = chain.Execute(bot, update)
			} else {
				// No middlewares, execute handler directly
				err = handler.Handler(bot, update, handlerCtx)
			}

			// If there was an error, pass it to error handlers
			if err != nil {
				log.Printf("Error handling update: %v", err)
				if handlerErr := bot.errorHandlers.Process(bot, update, err); handlerErr != nil {
					log.Printf("Error handler failed: %v", handlerErr)
				}
			}
			return // Stop after first matching handler
		}
	}
}

// Filter builders for common use cases

// MessageFilter filters messages only
func MessageFilter() FilterFunc {
	return func(update *models.Update) bool {
		return update.Message != nil
	}
}

// CommandFilter filters commands (messages starting with /)
func CommandFilter(command string) FilterFunc {
	return func(update *models.Update) bool {
		if update.Message == nil || update.Message.Text == "" {
			return false
		}
		text := update.Message.Text
		// Handle both /command and /command@botname
		if strings.HasPrefix(text, "/"+command+" ") || text == "/"+command {
			return true
		}
		if strings.HasPrefix(text, "/"+command+"@") {
			return true
		}
		return false
	}
}

// TextFilter filters text messages (non-command)
func TextFilter() FilterFunc {
	return func(update *models.Update) bool {
		if update.Message == nil || update.Message.Text == "" {
			return false
		}
		return !strings.HasPrefix(update.Message.Text, "/")
	}
}

// TextContainsFilter filters text messages containing specific text
func TextContainsFilter(substring string) FilterFunc {
	return func(update *models.Update) bool {
		if update.Message == nil || update.Message.Text == "" {
			return false
		}
		return strings.Contains(strings.ToLower(update.Message.Text), strings.ToLower(substring))
	}
}

// CallbackQueryFilter filters callback queries
func CallbackQueryFilter() FilterFunc {
	return func(update *models.Update) bool {
		return update.CallbackQuery != nil
	}
}

// CallbackDataFilter filters callback queries with specific data
func CallbackDataFilter(data string) FilterFunc {
	return func(update *models.Update) bool {
		if update.CallbackQuery == nil {
			return false
		}
		return update.CallbackQuery.Data == data
	}
}

// CallbackDataPrefixFilter filters callback queries with data starting with prefix
func CallbackDataPrefixFilter(prefix string) FilterFunc {
	return func(update *models.Update) bool {
		if update.CallbackQuery == nil {
			return false
		}
		return strings.HasPrefix(update.CallbackQuery.Data, prefix)
	}
}

// PhotoFilter filters photo messages
func PhotoFilter() FilterFunc {
	return func(update *models.Update) bool {
		return update.Message != nil && len(update.Message.Photo) > 0
	}
}

// DocumentFilter filters document messages
func DocumentFilter() FilterFunc {
	return func(update *models.Update) bool {
		return update.Message != nil && update.Message.Document != nil
	}
}

// VideoFilter filters video messages
func VideoFilter() FilterFunc {
	return func(update *models.Update) bool {
		return update.Message != nil && update.Message.Video != nil
	}
}

// AudioFilter filters audio messages
func AudioFilter() FilterFunc {
	return func(update *models.Update) bool {
		return update.Message != nil && update.Message.Audio != nil
	}
}

// VoiceFilter filters voice messages
func VoiceFilter() FilterFunc {
	return func(update *models.Update) bool {
		return update.Message != nil && update.Message.Voice != nil
	}
}

// StickerFilter filters sticker messages
func StickerFilter() FilterFunc {
	return func(update *models.Update) bool {
		return update.Message != nil && update.Message.Sticker != nil
	}
}

// LocationFilter filters location messages
func LocationFilter() FilterFunc {
	return func(update *models.Update) bool {
		return update.Message != nil && update.Message.Location != nil
	}
}

// ContactFilter filters contact messages
func ContactFilter() FilterFunc {
	return func(update *models.Update) bool {
		return update.Message != nil && update.Message.Contact != nil
	}
}

// InlineQueryFilter filters inline queries
func InlineQueryFilter() FilterFunc {
	return func(update *models.Update) bool {
		return update.InlineQuery != nil
	}
}

// ChannelPostFilter filters channel posts
func ChannelPostFilter() FilterFunc {
	return func(update *models.Update) bool {
		return update.ChannelPost != nil
	}
}

// EditedChannelPostFilter filters edited channel posts
func EditedChannelPostFilter() FilterFunc {
	return func(update *models.Update) bool {
		return update.EditedChannelPost != nil
	}
}

// EditedMessageFilter filters edited messages
func EditedMessageFilter() FilterFunc {
	return func(update *models.Update) bool {
		return update.EditedMessage != nil
	}
}

// ChatTypeFilter filters messages by chat type (private, group, supergroup, channel)
func ChatTypeFilter(chatType models.ChatType) FilterFunc {
	return func(update *models.Update) bool {
		var chat *models.Chat
		if update.Message != nil {
			chat = &update.Message.Chat
		} else if update.EditedMessage != nil {
			chat = &update.EditedMessage.Chat
		} else if update.ChannelPost != nil {
			chat = &update.ChannelPost.Chat
		} else if update.EditedChannelPost != nil {
			chat = &update.EditedChannelPost.Chat
		} else if update.CallbackQuery != nil && update.CallbackQuery.Message != nil {
			chat = &update.CallbackQuery.Message.Chat
		}

		if chat == nil {
			return false
		}
		return chat.Type == chatType
	}
}

// PrivateChatFilter filters private chat messages
func PrivateChatFilter() FilterFunc {
	return ChatTypeFilter(models.ChatTypePrivate)
}

// GroupChatFilter filters group chat messages
func GroupChatFilter() FilterFunc {
	return ChatTypeFilter(models.ChatTypeGroup)
}

// SupergroupChatFilter filters supergroup chat messages
func SupergroupChatFilter() FilterFunc {
	return ChatTypeFilter(models.ChatTypeSupergroup)
}

// ChannelChatFilter filters channel chat messages
func ChannelChatFilter() FilterFunc {
	return ChatTypeFilter(models.ChatTypeChannel)
}

// AndFilter combines multiple filters with AND logic
func AndFilter(filters ...FilterFunc) FilterFunc {
	return func(update *models.Update) bool {
		for _, filter := range filters {
			if !filter(update) {
				return false
			}
		}
		return true
	}
}

// OrFilter combines multiple filters with OR logic
func OrFilter(filters ...FilterFunc) FilterFunc {
	return func(update *models.Update) bool {
		for _, filter := range filters {
			if filter(update) {
				return true
			}
		}
		return false
	}
}

// NotFilter inverts a filter
func NotFilter(filter FilterFunc) FilterFunc {
	return func(update *models.Update) bool {
		return !filter(update)
	}
}

// BusinessConnectionFilter filters business connection updates
func BusinessConnectionFilter() FilterFunc {
	return func(update *models.Update) bool {
		return update.BusinessConnection != nil
	}
}

// BusinessMessageFilter filters business messages
func BusinessMessageFilter() FilterFunc {
	return func(update *models.Update) bool {
		return update.BusinessMessage != nil
	}
}

// EditedBusinessMessageFilter filters edited business messages
func EditedBusinessMessageFilter() FilterFunc {
	return func(update *models.Update) bool {
		return update.EditedBusinessMessage != nil
	}
}

// DeletedBusinessMessagesFilter filters deleted business messages
func DeletedBusinessMessagesFilter() FilterFunc {
	return func(update *models.Update) bool {
		return update.DeletedBusinessMessages != nil
	}
}

// GuestMessageFilter filters guest messages
func GuestMessageFilter() FilterFunc {
	return func(update *models.Update) bool {
		return update.GuestMessage != nil
	}
}

// MessageReactionFilter filters message reaction updates
func MessageReactionFilter() FilterFunc {
	return func(update *models.Update) bool {
		return update.MessageReaction != nil
	}
}

// MessageReactionCountFilter filters message reaction count updates
func MessageReactionCountFilter() FilterFunc {
	return func(update *models.Update) bool {
		return update.MessageReactionCount != nil
	}
}

// ShippingQueryFilter filters shipping query updates
func ShippingQueryFilter() FilterFunc {
	return func(update *models.Update) bool {
		return update.ShippingQuery != nil
	}
}

// PreCheckoutQueryFilter filters pre-checkout query updates
func PreCheckoutQueryFilter() FilterFunc {
	return func(update *models.Update) bool {
		return update.PreCheckoutQuery != nil
	}
}

// PurchasedPaidMediaFilter filters purchased paid media updates
func PurchasedPaidMediaFilter() FilterFunc {
	return func(update *models.Update) bool {
		return update.PurchasedPaidMedia != nil
	}
}

// PollFilter filters poll updates
func PollFilter() FilterFunc {
	return func(update *models.Update) bool {
		return update.Poll != nil
	}
}

// PollAnswerFilter filters poll answer updates
func PollAnswerFilter() FilterFunc {
	return func(update *models.Update) bool {
		return update.PollAnswer != nil
	}
}

// MyChatMemberFilter filters my chat member updates
func MyChatMemberFilter() FilterFunc {
	return func(update *models.Update) bool {
		return update.MyChatMember != nil
	}
}

// ChatMemberFilter filters chat member updates
func ChatMemberFilter() FilterFunc {
	return func(update *models.Update) bool {
		return update.ChatMember != nil
	}
}

// ChatJoinRequestFilter filters chat join request updates
func ChatJoinRequestFilter() FilterFunc {
	return func(update *models.Update) bool {
		return update.ChatJoinRequest != nil
	}
}

// ChatBoostFilter filters chat boost updates
func ChatBoostFilter() FilterFunc {
	return func(update *models.Update) bool {
		return update.ChatBoost != nil
	}
}

// RemovedChatBoostFilter filters removed chat boost updates
func RemovedChatBoostFilter() FilterFunc {
	return func(update *models.Update) bool {
		return update.RemovedChatBoost != nil
	}
}

// ManagedBotFilter filters managed bot updates
func ManagedBotFilter() FilterFunc {
	return func(update *models.Update) bool {
		return update.ManagedBot != nil
	}
}

// SubscriptionFilter filters subscription updates
func SubscriptionFilter() FilterFunc {
	return func(update *models.Update) bool {
		return update.Subscription != nil
	}
}

// AnyFilter matches any update
func AnyFilter() FilterFunc {
	return func(update *models.Update) bool {
		return update != nil
	}
}
