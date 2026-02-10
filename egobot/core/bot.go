package core

import (
	"egobot/egobot/core/methods"
	"egobot/egobot/models"
	"log"
	"time"
)

type Bot struct {
	Token         string
	requester     *methods.Requester
	handlers      *Handlers
	errorHandlers *ErrorHandlers
	*RegisterCommands
}

func NewBot(token string) *Bot {
	bot := &Bot{
		Token:         token,
		requester:     methods.NewRequester(token),
		handlers:      NewHandlers(),
		errorHandlers: NewErrorHandlers(),
	}
	bot.RegisterCommands = NewRegisterCommands(bot)
	return bot
}

// AddHandler adds a custom handler with a filter
func (b *Bot) AddHandler(filter FilterFunc, handler HandlerFunc, middlewares ...MiddlewareFunc) {
	b.handlers.AddHandler(filter, handler, middlewares...)
}

// RegisterGroup registers all handlers from a handler group
func (b *Bot) RegisterGroup(group *HandlerGroup) {
	for _, handler := range group.Handlers() {
		b.handlers.handlers = append(b.handlers.handlers, handler)
	}
}

// OnError registers an error handler with optional filter
// Pass nil as filter to handle all errors
func (b *Bot) OnError(filter ErrorFilter, handler ErrorHandlerFunc) {
	b.errorHandlers.AddHandler(filter, handler)
}

// OnTelegramError registers a handler for Telegram API errors
func (b *Bot) OnTelegramError(handler ErrorHandlerFunc) {
	b.errorHandlers.AddHandler(TelegramErrorFilter(), handler)
}

// OnRateLimitError registers a handler for rate limit errors (429)
func (b *Bot) OnRateLimitError(handler ErrorHandlerFunc) {
	b.errorHandlers.AddHandler(RateLimitErrorFilter(), handler)
}

// OnBadRequest registers a handler for bad request errors (400)
func (b *Bot) OnBadRequest(handler ErrorHandlerFunc) {
	b.errorHandlers.AddHandler(BadRequestErrorFilter(), handler)
}

// OnForbiddenError registers a handler for forbidden errors (403)
func (b *Bot) OnForbiddenError(handler ErrorHandlerFunc) {
	b.errorHandlers.AddHandler(ForbiddenErrorFilter(), handler)
}

// Message-specific error handlers

// OnMessageTextEmpty registers a handler for empty message text errors
func (b *Bot) OnMessageTextEmpty(handler ErrorHandlerFunc) {
	b.errorHandlers.AddHandler(MessageTextEmptyFilter(), handler)
}

// OnMessageTooLong registers a handler for message too long errors
func (b *Bot) OnMessageTooLong(handler ErrorHandlerFunc) {
	b.errorHandlers.AddHandler(MessageTooLongFilter(), handler)
}

// OnChatNotFound registers a handler for chat not found errors
func (b *Bot) OnChatNotFound(handler ErrorHandlerFunc) {
	b.errorHandlers.AddHandler(ChatNotFoundFilter(), handler)
}

// OnMessageNotFound registers a handler for message not found errors
func (b *Bot) OnMessageNotFound(handler ErrorHandlerFunc) {
	b.errorHandlers.AddHandler(MessageNotFoundFilter(), handler)
}

// OnMessageCantBeEdited registers a handler for message can't be edited errors
func (b *Bot) OnMessageCantBeEdited(handler ErrorHandlerFunc) {
	b.errorHandlers.AddHandler(MessageCantBeEditedFilter(), handler)
}

// OnMessageCantBeDeleted registers a handler for message can't be deleted errors
func (b *Bot) OnMessageCantBeDeleted(handler ErrorHandlerFunc) {
	b.errorHandlers.AddHandler(MessageCantBeDeletedFilter(), handler)
}

// OnBotBlocked registers a handler for bot was blocked by user errors
func (b *Bot) OnBotBlocked(handler ErrorHandlerFunc) {
	b.errorHandlers.AddHandler(BotBlockedFilter(), handler)
}

// OnBotKicked registers a handler for bot was kicked from chat errors
func (b *Bot) OnBotKicked(handler ErrorHandlerFunc) {
	b.errorHandlers.AddHandler(BotKickedFilter(), handler)
}

// OnInvalidFileID registers a handler for invalid file_id errors
func (b *Bot) OnInvalidFileID(handler ErrorHandlerFunc) {
	b.errorHandlers.AddHandler(InvalidFileIDFilter(), handler)
}

// OnButtonDataInvalid registers a handler for invalid button data errors
func (b *Bot) OnButtonDataInvalid(handler ErrorHandlerFunc) {
	b.errorHandlers.AddHandler(ButtonDataInvalidFilter(), handler)
}

// SetFallbackErrorHandler sets a fallback error handler for all unhandled errors
func (b *Bot) SetFallbackErrorHandler(handler ErrorHandlerFunc) {
	b.errorHandlers.SetFallbackHandler(handler)
}

// PollingOptions represents configuration options for polling
type PollingOptions struct {
	Timeout        int           // Timeout in seconds for long polling (default: 30)
	Limit          int           // Maximum number of updates to retrieve (default: 100)
	AllowedUpdates []string      // List of update types to receive (default: all)
	Async          bool          // Process updates asynchronously in goroutines (default: true)
	RetryDelay     int           // Delay in seconds before retrying after error (default: 3)
	OnStart        func()        // Callback when polling starts
	OnError        func(error)   // Callback when error occurs
}

// StartPolling starts polling for updates
// Pass nil to use default settings, or pass *PollingOptions to customize
func (b *Bot) StartPolling(options *PollingOptions) {
	// Use defaults if options is nil
	if options == nil {
		options = &PollingOptions{
			Timeout:    30,
			Limit:      100,
			Async:      true,
			RetryDelay: 3,
		}
	} else {
		// Fill in defaults for zero values
		if options.Timeout == 0 {
			options.Timeout = 30
		}
		if options.Limit == 0 {
			options.Limit = 100
		}
		if options.RetryDelay == 0 {
			options.RetryDelay = 3
		}
	}
	
	offset := int64(0)
	
	if options.OnStart != nil {
		options.OnStart()
	}
	
	log.Println("Bot started polling...")
	
	for {
		updates, err := b.GetUpdates(&models.GetUpdatesParams{
			Offset:         offset,
			Limit:          options.Limit,
			Timeout:        options.Timeout,
			AllowedUpdates: options.AllowedUpdates,
		})
		
		if err != nil {
			log.Printf("Error getting updates: %v", err)
			if options.OnError != nil {
				options.OnError(err)
			}
			time.Sleep(time.Duration(options.RetryDelay) * time.Second)
			continue
		}
		
		for _, update := range updates {
			offset = update.UpdateID + 1
			
			if options.Async {
				// Process update in a goroutine to handle multiple updates concurrently
				go b.handlers.Process(b, &update)
			} else {
				// Process update synchronously
				b.handlers.Process(b, &update)
			}
		}
	}
}

func (b *Bot) GetMe() (*models.User, error) {
	return b.requester.GetMe()
}

func (b *Bot) GetUpdates(params *models.GetUpdatesParams) ([]models.Update, error) {
	return b.requester.GetUpdates(params)
}

func (b *Bot) SendMessage(params *models.SendMessageParams) (*models.Message, error) {
	return b.requester.SendMessage(params)
}

func (b *Bot) SendPhoto(params *models.SendPhotoParams) (*models.Message, error) {
	return b.requester.SendPhoto(params)
}

func (b *Bot) SendDocument(params *models.SendDocumentParams) (*models.Message, error) {
	return b.requester.SendDocument(params)
}

func (b *Bot) SendVideo(params *models.SendVideoParams) (*models.Message, error) {
	return b.requester.SendVideo(params)
}

func (b *Bot) SendAudio(params *models.SendAudioParams) (*models.Message, error) {
	return b.requester.SendAudio(params)
}

func (b *Bot) EditMessageText(params *models.EditMessageTextParams) (*models.Message, error) {
	return b.requester.EditMessageText(params)
}

func (b *Bot) EditMessageCaption(params *models.EditMessageCaptionParams) (*models.Message, error) {
	return b.requester.EditMessageCaption(params)
}

func (b *Bot) EditMessageMedia(params *models.EditMessageMediaParams) (*models.Message, error) {
	return b.requester.EditMessageMedia(params)
}

func (b *Bot) EditMessageReplyMarkup(params *models.EditMessageReplyMarkupParams) (*models.Message, error) {
	return b.requester.EditMessageReplyMarkup(params)
}

func (b *Bot) EditMessageLiveLocation(params *models.EditMessageLiveLocationParams) (*models.Message, error) {
	return b.requester.EditMessageLiveLocation(params)
}

func (b *Bot) StopMessageLiveLocation(params *models.StopMessageLiveLocationParams) (*models.Message, error) {
	return b.requester.StopMessageLiveLocation(params)
}

func (b *Bot) EditMessageChecklist(params *models.EditMessageChecklistParams) (*models.Message, error) {
	return b.requester.EditMessageChecklist(params)
}

func (b *Bot) StopPoll(params *models.StopPollParams) (*models.Poll, error) {
	return b.requester.StopPoll(params)
}

func (b *Bot) DeleteMessage(params *models.DeleteMessageParams) (bool, error) {
	return b.requester.DeleteMessage(params)
}

func (b *Bot) AnswerCallbackQuery(callbackQueryID string, text string, showAlert bool) (bool, error) {
	return b.requester.AnswerCallbackQuery(callbackQueryID, text, showAlert)
}

func (b *Bot) SendChatAction(chatID interface{}, action string) (bool, error) {
	return b.requester.SendChatAction(chatID, action)
}

func (b *Bot) GetFile(fileID string) (*models.File, error) {
	return b.requester.GetFile(fileID)
}

func (b *Bot) BanChatMember(params *models.BanChatMemberParams) (bool, error) {
	return b.requester.BanChatMember(params)
}

func (b *Bot) UnbanChatMember(params *models.UnbanChatMemberParams) (bool, error) {
	return b.requester.UnbanChatMember(params)
}

func (b *Bot) RestrictChatMember(params *models.RestrictChatMemberParams) (bool, error) {
	return b.requester.RestrictChatMember(params)
}

func (b *Bot) PromoteChatMember(params *models.PromoteChatMemberParams) (bool, error) {
	return b.requester.PromoteChatMember(params)
}

func (b *Bot) SetChatAdministratorCustomTitle(params *models.SetChatAdministratorCustomTitleParams) (bool, error) {
	return b.requester.SetChatAdministratorCustomTitle(params)
}

func (b *Bot) GetChatMember(params *models.GetChatMemberParams) (*models.ChatMember, error) {
	return b.requester.GetChatMember(params)
}

func (b *Bot) PinChatMessage(params *models.PinChatMessageParams) (bool, error) {
	return b.requester.PinChatMessage(params)
}

func (b *Bot) UnpinChatMessage(params *models.UnpinChatMessageParams) (bool, error) {
	return b.requester.UnpinChatMessage(params)
}

func (b *Bot) UnpinAllChatMessages(chatID interface{}) (bool, error) {
	return b.requester.UnpinAllChatMessages(chatID)
}

func (b *Bot) LeaveChat(chatID interface{}) (bool, error) {
	return b.requester.LeaveChat(chatID)
}

func (b *Bot) GetChat(chatID interface{}) (*models.Chat, error) {
	return b.requester.GetChat(chatID)
}

func (b *Bot) GetChatAdministrators(chatID interface{}) ([]models.ChatMember, error) {
	return b.requester.GetChatAdministrators(chatID)
}

func (b *Bot) GetChatMemberCount(chatID interface{}) (int, error) {
	return b.requester.GetChatMemberCount(chatID)
}

func (b *Bot) ForwardMessage(params *models.ForwardMessageParams) (*models.Message, error) {
	return b.requester.ForwardMessage(params)
}

func (b *Bot) CopyMessage(params *models.CopyMessageParams) (*models.MessageID, error) {
	return b.requester.CopyMessage(params)
}

func (b *Bot) SendLocation(params *models.SendLocationParams) (*models.Message, error) {
	return b.requester.SendLocation(params)
}

func (b *Bot) SendContact(params *models.SendContactParams) (*models.Message, error) {
	return b.requester.SendContact(params)
}

func (b *Bot) SendPoll(params *models.SendPollParams) (*models.Message, error) {
	return b.requester.SendPoll(params)
}
func (b *Bot) SendAnimation(params models.SendAnimationParams) (*models.Message, error) {
	return b.requester.SendAnimation(params)
}

func (b *Bot) SendVoice(params models.SendVoiceParams) (*models.Message, error) {
	return b.requester.SendVoice(params)
}

func (b *Bot) SendVideoNote(params models.SendVideoNoteParams) (*models.Message, error) {
	return b.requester.SendVideoNote(params)
}

func (b *Bot) SendMediaGroup(params models.SendMediaGroupParams) ([]models.Message, error) {
	return b.requester.SendMediaGroup(params)
}

func (b *Bot) SendVenue(params models.SendVenueParams) (*models.Message, error) {
	return b.requester.SendVenue(params)
}

func (b *Bot) SendDice(params models.SendDiceParams) (*models.Message, error) {
	return b.requester.SendDice(params)
}

func (b *Bot) SendChecklist(params models.SendChecklistParams) (*models.Message, error) {
	return b.requester.SendChecklist(params)
}

func (b *Bot) SendPaidMedia(params models.SendPaidMediaParams) (*models.Message, error) {
	return b.requester.SendPaidMedia(params)
}

func (b *Bot) SendSticker(params models.SendStickerParams) (*models.Message, error) {
	return b.requester.SendSticker(params)
}

func (b *Bot) SendMessageDraft(params models.SendMessageDraftParams) (bool, error) {
	return b.requester.SendMessageDraft(params)
}

func (b *Bot) CopyMessages(params models.CopyMessagesParams) ([]models.MessageID, error) {
	return b.requester.CopyMessages(params)
}

func (b *Bot) ForwardMessages(params models.ForwardMessagesParams) ([]models.MessageID, error) {
	return b.requester.ForwardMessages(params)
}

func (b *Bot) DeleteMessages(params models.DeleteMessagesParams) (bool, error) {
	return b.requester.DeleteMessages(params)
}

// Chat Settings Methods

func (b *Bot) SetChatPhoto(params models.SetChatPhotoParams) (bool, error) {
	return b.requester.SetChatPhoto(params)
}

func (b *Bot) DeleteChatPhoto(params models.DeleteChatPhotoParams) (bool, error) {
	return b.requester.DeleteChatPhoto(params)
}

func (b *Bot) SetChatTitle(params models.SetChatTitleParams) (bool, error) {
	return b.requester.SetChatTitle(params)
}

func (b *Bot) SetChatDescription(params models.SetChatDescriptionParams) (bool, error) {
	return b.requester.SetChatDescription(params)
}

func (b *Bot) BanChatSenderChat(params models.BanChatSenderChatParams) (bool, error) {
	return b.requester.BanChatSenderChat(params)
}

func (b *Bot) UnbanChatSenderChat(params models.UnbanChatSenderChatParams) (bool, error) {
	return b.requester.UnbanChatSenderChat(params)
}

func (b *Bot) SetChatPermissions(params models.SetChatPermissionsParams) (bool, error) {
	return b.requester.SetChatPermissions(params)
}

func (b *Bot) ExportChatInviteLink(chatID interface{}) (string, error) {
	return b.requester.ExportChatInviteLink(chatID)
}

func (b *Bot) CreateChatInviteLink(params models.CreateChatInviteLinkParams) (*models.ChatInviteLink, error) {
	return b.requester.CreateChatInviteLink(params)
}

func (b *Bot) EditChatInviteLink(params models.EditChatInviteLinkParams) (*models.ChatInviteLink, error) {
	return b.requester.EditChatInviteLink(params)
}

func (b *Bot) RevokeChatInviteLink(params models.RevokeChatInviteLinkParams) (*models.ChatInviteLink, error) {
	return b.requester.RevokeChatInviteLink(params)
}

func (b *Bot) ApproveChatJoinRequest(params models.ApproveChatJoinRequestParams) (bool, error) {
	return b.requester.ApproveChatJoinRequest(params)
}

func (b *Bot) DeclineChatJoinRequest(params models.DeclineChatJoinRequestParams) (bool, error) {
	return b.requester.DeclineChatJoinRequest(params)
}

func (b *Bot) SetChatStickerSet(params models.SetChatStickerSetParams) (bool, error) {
	return b.requester.SetChatStickerSet(params)
}

func (b *Bot) DeleteChatStickerSet(params models.DeleteChatStickerSetParams) (bool, error) {
	return b.requester.DeleteChatStickerSet(params)
}

// Bot Configuration Methods

func (b *Bot) SetMyCommands(params models.SetMyCommandsParams) (bool, error) {
	return b.requester.SetMyCommands(params)
}

func (b *Bot) DeleteMyCommands(params models.DeleteMyCommandsParams) (bool, error) {
	return b.requester.DeleteMyCommands(params)
}

func (b *Bot) GetMyCommands(params models.GetMyCommandsParams) ([]models.BotCommand, error) {
	return b.requester.GetMyCommands(params)
}

func (b *Bot) SetMyName(params models.SetMyNameParams) (bool, error) {
	return b.requester.SetMyName(params)
}

func (b *Bot) GetMyName(params models.GetMyNameParams) (*models.BotName, error) {
	return b.requester.GetMyName(params)
}

func (b *Bot) SetMyDescription(params models.SetMyDescriptionParams) (bool, error) {
	return b.requester.SetMyDescription(params)
}

func (b *Bot) GetMyDescription(params models.GetMyDescriptionParams) (*models.BotDescription, error) {
	return b.requester.GetMyDescription(params)
}

func (b *Bot) SetMyShortDescription(params models.SetMyShortDescriptionParams) (bool, error) {
	return b.requester.SetMyShortDescription(params)
}

func (b *Bot) GetMyShortDescription(params models.GetMyShortDescriptionParams) (*models.BotShortDescription, error) {
	return b.requester.GetMyShortDescription(params)
}

func (b *Bot) SetChatMenuButton(params models.SetChatMenuButtonParams) (bool, error) {
	return b.requester.SetChatMenuButton(params)
}

func (b *Bot) GetChatMenuButton(params models.GetChatMenuButtonParams) (*models.MenuButton, error) {
	return b.requester.GetChatMenuButton(params)
}

func (b *Bot) GetUserProfilePhotos(params models.GetUserProfilePhotosParams) (*models.UserProfilePhotos, error) {
	return b.requester.GetUserProfilePhotos(params)
}

func (b *Bot) SetMessageReaction(params models.SetMessageReactionParams) (bool, error) {
	return b.requester.SetMessageReaction(params)
}

// Bot API 9.4: Set bot profile photo
func (b *Bot) SetMyProfilePhoto(photo *models.InputProfilePhoto) (bool, error) {
	return b.requester.SetMyProfilePhoto(photo)
}

// Bot API 9.4: Remove bot profile photo
func (b *Bot) RemoveMyProfilePhoto() (bool, error) {
	return b.requester.RemoveMyProfilePhoto()
}

// Bot API 9.4: Get user profile audios
func (b *Bot) GetUserProfileAudios(params *models.GetUserProfileAudiosParams) (*models.UserProfileAudios, error) {
	return b.requester.GetUserProfileAudios(params)
}

// Webhook Methods

func (b *Bot) SetWebhook(params models.SetWebhookParams) (bool, error) {
	return b.requester.SetWebhook(params)
}

func (b *Bot) DeleteWebhook(params models.DeleteWebhookParams) (bool, error) {
	return b.requester.DeleteWebhook(params)
}

func (b *Bot) GetWebhookInfo() (*models.WebhookInfo, error) {
	return b.requester.GetWebhookInfo()
}

// Sticker Methods

func (b *Bot) GetStickerSet(params models.GetStickerSetParams) (*models.StickerSet, error) {
	return b.requester.GetStickerSet(params)
}

func (b *Bot) GetCustomEmojiStickers(params models.GetCustomEmojiStickersParams) ([]models.Sticker, error) {
	return b.requester.GetCustomEmojiStickers(params)
}

func (b *Bot) UploadStickerFile(params models.UploadStickerFileParams) (*models.File, error) {
	return b.requester.UploadStickerFile(params)
}

func (b *Bot) CreateNewStickerSet(params models.CreateNewStickerSetParams) (bool, error) {
	return b.requester.CreateNewStickerSet(params)
}

func (b *Bot) AddStickerToSet(params models.AddStickerToSetParams) (bool, error) {
	return b.requester.AddStickerToSet(params)
}

func (b *Bot) SetStickerPositionInSet(params models.SetStickerPositionInSetParams) (bool, error) {
	return b.requester.SetStickerPositionInSet(params)
}

func (b *Bot) DeleteStickerFromSet(params models.DeleteStickerFromSetParams) (bool, error) {
	return b.requester.DeleteStickerFromSet(params)
}

func (b *Bot) SetStickerSetThumbnail(params models.SetStickerSetThumbnailParams) (bool, error) {
	return b.requester.SetStickerSetThumbnail(params)
}

// Inline Mode Methods

func (b *Bot) AnswerInlineQuery(params models.AnswerInlineQueryParams) (bool, error) {
	return b.requester.AnswerInlineQuery(params)
}

// Payment Methods

func (b *Bot) SendInvoice(params models.SendInvoiceParams) (*models.Message, error) {
	return b.requester.SendInvoice(params)
}

func (b *Bot) CreateInvoiceLink(params models.CreateInvoiceLinkParams) (string, error) {
	return b.requester.CreateInvoiceLink(params)
}

func (b *Bot) AnswerShippingQuery(params models.AnswerShippingQueryParams) (bool, error) {
	return b.requester.AnswerShippingQuery(params)
}

func (b *Bot) AnswerPreCheckoutQuery(params models.AnswerPreCheckoutQueryParams) (bool, error) {
	return b.requester.AnswerPreCheckoutQuery(params)
}

// Game Methods

func (b *Bot) SendGame(params models.SendGameParams) (*models.Message, error) {
	return b.requester.SendGame(params)
}

func (b *Bot) SetGameScore(params models.SetGameScoreParams) (*models.Message, error) {
	return b.requester.SetGameScore(params)
}

func (b *Bot) GetGameHighScores(params models.GetGameHighScoresParams) ([]models.GameHighScore, error) {
	return b.requester.GetGameHighScores(params)
}

// Forum Topic Methods

func (b *Bot) CreateForumTopic(params models.CreateForumTopicParams) (*models.ForumTopic, error) {
	return b.requester.CreateForumTopic(params)
}

func (b *Bot) EditForumTopic(params models.EditForumTopicParams) (bool, error) {
	return b.requester.EditForumTopic(params)
}

func (b *Bot) CloseForumTopic(params models.CloseForumTopicParams) (bool, error) {
	return b.requester.CloseForumTopic(params)
}

func (b *Bot) ReopenForumTopic(params models.ReopenForumTopicParams) (bool, error) {
	return b.requester.ReopenForumTopic(params)
}

func (b *Bot) DeleteForumTopic(params models.DeleteForumTopicParams) (bool, error) {
	return b.requester.DeleteForumTopic(params)
}

func (b *Bot) UnpinAllForumTopicMessages(params models.UnpinAllForumTopicMessagesParams) (bool, error) {
	return b.requester.UnpinAllForumTopicMessages(params)
}

func (b *Bot) EditGeneralForumTopic(params models.EditGeneralForumTopicParams) (bool, error) {
	return b.requester.EditGeneralForumTopic(params)
}

func (b *Bot) CloseGeneralForumTopic(params models.CloseGeneralForumTopicParams) (bool, error) {
	return b.requester.CloseGeneralForumTopic(params)
}

func (b *Bot) ReopenGeneralForumTopic(params models.ReopenGeneralForumTopicParams) (bool, error) {
	return b.requester.ReopenGeneralForumTopic(params)
}

func (b *Bot) HideGeneralForumTopic(params models.HideGeneralForumTopicParams) (bool, error) {
	return b.requester.HideGeneralForumTopic(params)
}

func (b *Bot) UnhideGeneralForumTopic(params models.UnhideGeneralForumTopicParams) (bool, error) {
	return b.requester.UnhideGeneralForumTopic(params)
}
