package models

// https://core.telegram.org/bots/api#update
type Update struct {
	UpdateID                int64                        `json:"update_id"`
	Message                 *Message                     `json:"message,omitempty"`
	EditedMessage           *Message                     `json:"edited_message,omitempty"`
	ChannelPost             *Message                     `json:"channel_post,omitempty"`
	EditedChannelPost       *Message                     `json:"edited_channel_post,omitempty"`
	BusinessConnection      *BusinessConnection          `json:"business_connection,omitempty"`
	BusinessMessage         *Message                     `json:"business_message,omitempty"`
	EditedBusinessMessage   *Message                     `json:"edited_business_message,omitempty"`
	DeletedBusinessMessages *BusinessMessagesDeleted     `json:"deleted_business_messages,omitempty"`
	GuestMessage            *Message                     `json:"guest_message,omitempty"`
	MessageReaction         *MessageReactionUpdated      `json:"message_reaction,omitempty"`
	MessageReactionCount    *MessageReactionCountUpdated `json:"message_reaction_count,omitempty"`
	InlineQuery             *InlineQuery                 `json:"inline_query,omitempty"`
	ChosenInlineResult      *ChosenInlineResult          `json:"chosen_inline_result,omitempty"`
	CallbackQuery           *CallbackQuery               `json:"callback_query,omitempty"`
	ShippingQuery           *ShippingQuery               `json:"shipping_query,omitempty"`
	PreCheckoutQuery        *PreCheckoutQuery            `json:"pre_checkout_query,omitempty"`
	PurchasedPaidMedia      *PaidMediaPurchased          `json:"purchased_paid_media,omitempty"`
	Poll                    *Poll                        `json:"poll,omitempty"`
	PollAnswer              *PollAnswer                  `json:"poll_answer,omitempty"`
	MyChatMember            *ChatMemberUpdated           `json:"my_chat_member,omitempty"`
	ChatMember              *ChatMemberUpdated           `json:"chat_member,omitempty"`
	ChatJoinRequest         *ChatJoinRequest             `json:"chat_join_request,omitempty"`
	ChatBoost               *ChatBoostUpdated            `json:"chat_boost,omitempty"`
	RemovedChatBoost        *ChatBoostRemoved            `json:"removed_chat_boost,omitempty"`
	ManagedBot              *ManagedBotUpdated           `json:"managed_bot,omitempty"`
	Subscription            *BotSubscriptionUpdated      `json:"subscription,omitempty"`
}

// https://core.telegram.org/bots/api#getupdates
type GetUpdatesParams struct {
	Offset         int64    `json:"offset,omitempty"`
	Limit          int      `json:"limit,omitempty"`
	Timeout        int      `json:"timeout,omitempty"`
	AllowedUpdates []string `json:"allowed_updates,omitempty"`
}

// https://core.telegram.org/bots/api#webhookinfo
type WebhookInfo struct {
	URL                          string   `json:"url"`
	HasCustomCertificate         bool     `json:"has_custom_certificate"`
	PendingUpdateCount           int      `json:"pending_update_count"`
	IPAddress                    string   `json:"ip_address,omitempty"`
	LastErrorDate                int64    `json:"last_error_date,omitempty"`
	LastErrorMessage             string   `json:"last_error_message,omitempty"`
	LastSynchronizationErrorDate int64    `json:"last_synchronization_error_date,omitempty"`
	MaxConnections               int      `json:"max_connections,omitempty"`
	AllowedUpdates               []string `json:"allowed_updates,omitempty"`
}

// https://core.telegram.org/bots/api#setwebhook
type SetWebhookParams struct {
	URL                string      `json:"url"`
	Certificate        interface{} `json:"certificate,omitempty"`
	IPAddress          string      `json:"ip_address,omitempty"`
	MaxConnections     int         `json:"max_connections,omitempty"`
	AllowedUpdates     []string    `json:"allowed_updates,omitempty"`
	DropPendingUpdates bool        `json:"drop_pending_updates,omitempty"`
	SecretToken        string      `json:"secret_token,omitempty"`
}

// https://core.telegram.org/bots/api#deletewebhook
type DeleteWebhookParams struct {
	DropPendingUpdates bool `json:"drop_pending_updates,omitempty"`
}

// https://core.telegram.org/bots/api#inputprofilephoto
type InputProfilePhoto interface{}

// https://core.telegram.org/bots/api#inputprofilephotostatic
type InputProfilePhotoStatic struct {
	Type  string      `json:"type"`  // "static"
	Photo interface{} `json:"photo"` // File to upload
}

// https://core.telegram.org/bots/api#userprofileaudios
type UserProfileAudios struct {
	TotalCount int     `json:"total_count"`
	Audios     []Audio `json:"audios"`
}

// https://core.telegram.org/bots/api#getuserprofileaudios
type GetUserProfileAudiosParams struct {
	UserID int64 `json:"user_id"`
	Offset int   `json:"offset,omitempty"`
	Limit  int   `json:"limit,omitempty"` // 1-100, default 100
}

// https://core.telegram.org/bots/api#botcommand
type BotCommand struct {
	Command     string `json:"command"`
	Description string `json:"description"`
}

// https://core.telegram.org/bots/api#setmycommands
type SetMyCommandsParams struct {
	Commands     []BotCommand `json:"commands"`
	Scope        interface{}  `json:"scope,omitempty"`
	LanguageCode string       `json:"language_code,omitempty"`
}

// https://core.telegram.org/bots/api#deletemycommands
type DeleteMyCommandsParams struct {
	Scope        interface{} `json:"scope,omitempty"`
	LanguageCode string      `json:"language_code,omitempty"`
}

// https://core.telegram.org/bots/api#getmycommands
type GetMyCommandsParams struct {
	Scope        interface{} `json:"scope,omitempty"`
	LanguageCode string      `json:"language_code,omitempty"`
}

// https://core.telegram.org/bots/api#setmyname
type SetMyNameParams struct {
	Name         string `json:"name,omitempty"`
	LanguageCode string `json:"language_code,omitempty"`
}

// https://core.telegram.org/bots/api#getmyname
type GetMyNameParams struct {
	LanguageCode string `json:"language_code,omitempty"`
}

// https://core.telegram.org/bots/api#botname
type BotName struct {
	Name string `json:"name"`
}

// https://core.telegram.org/bots/api#setmydescription
type SetMyDescriptionParams struct {
	Description  string `json:"description,omitempty"`
	LanguageCode string `json:"language_code,omitempty"`
}

// https://core.telegram.org/bots/api#getmydescription
type GetMyDescriptionParams struct {
	LanguageCode string `json:"language_code,omitempty"`
}

// https://core.telegram.org/bots/api#botdescription
type BotDescription struct {
	Description string `json:"description"`
}

// https://core.telegram.org/bots/api#setmyshortdescription
type SetMyShortDescriptionParams struct {
	ShortDescription string `json:"short_description,omitempty"`
	LanguageCode     string `json:"language_code,omitempty"`
}

// https://core.telegram.org/bots/api#getmyshortdescription
type GetMyShortDescriptionParams struct {
	LanguageCode string `json:"language_code,omitempty"`
}

// https://core.telegram.org/bots/api#botshortdescription
type BotShortDescription struct {
	ShortDescription string `json:"short_description"`
}

// https://core.telegram.org/bots/api#setchatmenubutton
type SetChatMenuButtonParams struct {
	ChatID     int64       `json:"chat_id,omitempty"`
	MenuButton interface{} `json:"menu_button,omitempty"`
}

// https://core.telegram.org/bots/api#getchatmenubutton
type GetChatMenuButtonParams struct {
	ChatID int64 `json:"chat_id,omitempty"`
}

// https://core.telegram.org/bots/api#menubutton
type MenuButton struct {
	Type   string      `json:"type"`
	Text   string      `json:"text,omitempty"`
	WebApp *WebAppInfo `json:"web_app,omitempty"`
}

// https://core.telegram.org/bots/api#userprofilephotos
type UserProfilePhotos struct {
	TotalCount int           `json:"total_count"`
	Photos     [][]PhotoSize `json:"photos"`
}

// https://core.telegram.org/bots/api#getuserprofilephotos
type GetUserProfilePhotosParams struct {
	UserID int64 `json:"user_id"`
	Offset int   `json:"offset,omitempty"`
	Limit  int   `json:"limit,omitempty"`
}

// https://core.telegram.org/bots/api#setmessagereaction
type SetMessageReactionParams struct {
	ChatID    interface{}   `json:"chat_id"`
	MessageID int           `json:"message_id"`
	Reaction  []interface{} `json:"reaction,omitempty"`
	IsBig     bool          `json:"is_big,omitempty"`
}

// https://core.telegram.org/bots/api#webappdata
type WebAppData struct {
	Data       string `json:"data"`
	ButtonText string `json:"button_text"`
}

// https://core.telegram.org/bots/api#proximityalerttriggered
type ProximityAlertTriggered struct {
	Traveler User `json:"traveler"`
	Watcher  User `json:"watcher"`
	Distance int  `json:"distance"`
}

// https://core.telegram.org/bots/api#writeaccessallowed
type WriteAccessAllowed struct {
	FromRequest        bool   `json:"from_request,omitempty"`
	WebAppName         string `json:"web_app_name,omitempty"`
	FromAttachmentMenu bool   `json:"from_attachment_menu,omitempty"`
}

// https://core.telegram.org/bots/api#videochatscheduled
type VideoChatScheduled struct {
	StartDate int64 `json:"start_date"`
}

// https://core.telegram.org/bots/api#videochatstarted
type VideoChatStarted struct{}

// https://core.telegram.org/bots/api#videochatended
type VideoChatEnded struct {
	Duration int `json:"duration"`
}

// https://core.telegram.org/bots/api#videochatparticipantsinvited
type VideoChatParticipantsInvited struct {
	Users []User `json:"users"`
}

// https://core.telegram.org/bots/api#forumtopiccreated
type ForumTopicCreated struct {
	Name              string `json:"name"`
	IconColor         int    `json:"icon_color"`
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"`
	IsNameImplicit    bool   `json:"is_name_implicit,omitempty"`
}

// https://core.telegram.org/bots/api#forumtopicclosed
type ForumTopicClosed struct{}

// https://core.telegram.org/bots/api#forumtopicreopened
type ForumTopicReopened struct{}

// https://core.telegram.org/bots/api#forumtopicedited
type ForumTopicEdited struct {
	Name              string `json:"name,omitempty"`
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"`
}

// https://core.telegram.org/bots/api#generalforumtopichidden
type GeneralForumTopicHidden struct{}

// https://core.telegram.org/bots/api#generalforumtopicunhidden
type GeneralForumTopicUnhidden struct{}

// https://core.telegram.org/bots/api#shareduser
type SharedUser struct {
	UserID    int64       `json:"user_id"`
	FirstName string      `json:"first_name,omitempty"`
	LastName  string      `json:"last_name,omitempty"`
	Username  string      `json:"username,omitempty"`
	Photo     []PhotoSize `json:"photo,omitempty"`
}

// https://core.telegram.org/bots/api#usersshared
type UsersShared struct {
	RequestID int          `json:"request_id"`
	Users     []SharedUser `json:"users"`
}

// https://core.telegram.org/bots/api#chatshared
type ChatShared struct {
	RequestID int         `json:"request_id"`
	ChatID    int64       `json:"chat_id"`
	Title     string      `json:"title,omitempty"`
	Username  string      `json:"username,omitempty"`
	Photo     []PhotoSize `json:"photo,omitempty"`
}

// https://core.telegram.org/bots/api#giveaway
type Giveaway struct {
	Chats                         []Chat   `json:"chats"`
	WinnersSelectionDate          int64    `json:"winners_selection_date"`
	WinnerCount                   int      `json:"winner_count"`
	OnlyNewMembers                bool     `json:"only_new_members,omitempty"`
	HasPublicWinners              bool     `json:"has_public_winners,omitempty"`
	PrizeDescription              string   `json:"prize_description,omitempty"`
	CountryCodes                  []string `json:"country_codes,omitempty"`
	PrizeStarCount                int      `json:"prize_star_count,omitempty"`
	PremiumSubscriptionMonthCount int      `json:"premium_subscription_month_count,omitempty"`
}

// https://core.telegram.org/bots/api#giveawaywinners
type GiveawayWinners struct {
	Chat                          Chat   `json:"chat"`
	GiveawayMessageID             int    `json:"giveaway_message_id"`
	WinnersSelectionDate          int64  `json:"winners_selection_date"`
	WinnerCount                   int    `json:"winner_count"`
	Winners                       []User `json:"winners"`
	AdditionalChatCount           int    `json:"additional_chat_count,omitempty"`
	PrizeStarCount                int    `json:"prize_star_count,omitempty"`
	PremiumSubscriptionMonthCount int    `json:"premium_subscription_month_count,omitempty"`
	UnclaimedPrizeCount           int    `json:"unclaimed_prize_count,omitempty"`
	OnlyNewMembers                bool   `json:"only_new_members,omitempty"`
	WasRefunded                   bool   `json:"was_refunded,omitempty"`
	PrizeDescription              string `json:"prize_description,omitempty"`
}

// https://core.telegram.org/bots/api#giveawaycompleted
type GiveawayCompleted struct {
	WinnerCount         int      `json:"winner_count"`
	UnclaimedPrizeCount int      `json:"unclaimed_prize_count,omitempty"`
	GiveawayMessage     *Message `json:"giveaway_message,omitempty"`
	IsStarGiveaway      bool     `json:"is_star_giveaway,omitempty"`
}

// https://core.telegram.org/bots/api#giveawaycreated
type GiveawayCreated struct {
	PrizeStarCount int `json:"prize_star_count,omitempty"`
}

// https://core.telegram.org/bots/api#chatboostadded
type ChatBoostAdded struct {
	BoostCount int `json:"boost_count"`
}

// https://core.telegram.org/bots/api#chatbackground
type ChatBackground struct {
	Type BackgroundType `json:"type"`
}

// https://core.telegram.org/bots/api#backgroundtype
type BackgroundType interface{}

// https://core.telegram.org/bots/api#checklisttasksdone
type ChecklistTasksDone struct {
	ChecklistMessage       *Message `json:"checklist_message,omitempty"`
	MarkedAsDoneTaskIDs    []int    `json:"marked_as_done_task_ids,omitempty"`
	MarkedAsNotDoneTaskIDs []int    `json:"marked_as_not_done_task_ids,omitempty"`
}

// https://core.telegram.org/bots/api#checklisttasksadded
type ChecklistTasksAdded struct {
	ChecklistMessage *Message        `json:"checklist_message,omitempty"`
	Tasks            []ChecklistTask `json:"tasks"`
}

// https://core.telegram.org/bots/api#communitychatadded
type CommunityChatAdded struct {
	Community Community `json:"community"`
}

// https://core.telegram.org/bots/api#communitychatremoved
type CommunityChatRemoved struct{}

// https://core.telegram.org/bots/api#paidmessagepricechanged
type PaidMessagePriceChanged struct {
	PaidMessageStarCount int `json:"paid_message_star_count"`
}

// https://core.telegram.org/bots/api#directmessagepricechanged
type DirectMessagePriceChanged struct {
	AreDirectMessagesEnabled bool `json:"are_direct_messages_enabled"`
	DirectMessageStarCount   int  `json:"direct_message_star_count,omitempty"`
}

// https://core.telegram.org/bots/api#suggestedpostapproved
type SuggestedPostApproved struct {
	SuggestedPostMessage *Message            `json:"suggested_post_message,omitempty"`
	Price                *SuggestedPostPrice `json:"price,omitempty"`
	SendDate             int64               `json:"send_date"`
}

// https://core.telegram.org/bots/api#suggestedpostapprovalfailed
type SuggestedPostApprovalFailed struct {
	SuggestedPostMessage *Message            `json:"suggested_post_message,omitempty"`
	Price                *SuggestedPostPrice `json:"price,omitempty"`
}

// https://core.telegram.org/bots/api#suggestedpostdeclined
type SuggestedPostDeclined struct {
	SuggestedPostMessage *Message `json:"suggested_post_message,omitempty"`
	Comment              string   `json:"comment,omitempty"`
}

// https://core.telegram.org/bots/api#suggestedpostpaid
type SuggestedPostPaid struct {
	SuggestedPostMessage *Message    `json:"suggested_post_message,omitempty"`
	Currency             string      `json:"currency"`
	Amount               int         `json:"amount,omitempty"`
	StarAmount           *StarAmount `json:"star_amount,omitempty"`
}

// https://core.telegram.org/bots/api#suggestedpostrefunded
type SuggestedPostRefunded struct {
	SuggestedPostMessage *Message `json:"suggested_post_message,omitempty"`
	Reason               string   `json:"reason"`
}

// https://core.telegram.org/bots/api#polloptionadded
type PollOptionAdded struct {
	PollMessage        *Message        `json:"poll_message,omitempty"`
	OptionPersistentID string          `json:"option_persistent_id"`
	OptionText         string          `json:"option_text"`
	OptionTextEntities []MessageEntity `json:"option_text_entities,omitempty"`
}

// https://core.telegram.org/bots/api#polloptiondeleted
type PollOptionDeleted struct {
	PollMessage        *Message        `json:"poll_message,omitempty"`
	OptionPersistentID string          `json:"option_persistent_id"`
	OptionText         string          `json:"option_text"`
	OptionTextEntities []MessageEntity `json:"option_text_entities,omitempty"`
}

// https://core.telegram.org/bots/api#managedbotcreated
type ManagedBotCreated struct {
	Bot User `json:"bot"`
}

// https://core.telegram.org/bots/api#inaccessiblemessage
type InaccessibleMessage struct {
	Chat      Chat  `json:"chat"`
	MessageID int   `json:"message_id"`
	Date      int64 `json:"date"`
}

// https://core.telegram.org/bots/api#keyboardbuttonrequestmanagedbot
type KeyboardButtonRequestManagedBot struct {
	RequestID         int    `json:"request_id"`
	SuggestedName     string `json:"suggested_name,omitempty"`
	SuggestedUsername string `json:"suggested_username,omitempty"`
}

// https://core.telegram.org/bots/api#links
type Link struct {
	URL string `json:"url"`
}

// https://core.telegram.org/bots/api#richmessage
type RichMessage struct {
	Blocks []RichBlock `json:"blocks"`
	IsRTL  bool        `json:"is_rtl,omitempty"`
}

// https://core.telegram.org/bots/api#inputrichmessage
type InputRichMessage struct {
	Blocks              []InputRichBlock        `json:"blocks,omitempty"`
	HTML                string                  `json:"html,omitempty"`
	Markdown            string                  `json:"markdown,omitempty"`
	Media               []InputRichMessageMedia `json:"media,omitempty"`
	IsRTL               bool                    `json:"is_rtl,omitempty"`
	SkipEntityDetection bool                    `json:"skip_entity_detection,omitempty"`
}

// https://core.telegram.org/bots/api#inputrichmessagemedia
type InputRichMessageMedia struct {
	ID    string      `json:"id"`
	Media interface{} `json:"media"`
}

// Placeholder types for RichMessage formatting
type RichBlock interface{}
type InputRichBlock interface{}

// https://core.telegram.org/bots/api#giftbackground
type GiftBackground struct {
	CenterColor int `json:"center_color"`
	EdgeColor   int `json:"edge_color"`
	TextColor   int `json:"text_color"`
}

// https://core.telegram.org/bots/api#gift
type Gift struct {
	ID                     string          `json:"id"`
	Sticker                Sticker         `json:"sticker"`
	StarCount              int             `json:"star_count"`
	UpgradeStarCount       int             `json:"upgrade_star_count,omitempty"`
	IsPremium              bool            `json:"is_premium,omitempty"`
	HasColors              bool            `json:"has_colors,omitempty"`
	TotalCount             int             `json:"total_count,omitempty"`
	RemainingCount         int             `json:"remaining_count,omitempty"`
	PersonalTotalCount     int             `json:"personal_total_count,omitempty"`
	PersonalRemainingCount int             `json:"personal_remaining_count,omitempty"`
	Background             *GiftBackground `json:"background,omitempty"`
	UniqueGiftVariantCount int             `json:"unique_gift_variant_count,omitempty"`
	PublisherChat          *Chat           `json:"publisher_chat,omitempty"`
}

// https://core.telegram.org/bots/api#uniquegiftmodel
type UniqueGiftModel struct {
	Name           string  `json:"name"`
	Sticker        Sticker `json:"sticker"`
	RarityPerMille int     `json:"rarity_per_mille"`
	Rarity         string  `json:"rarity,omitempty"`
}

// https://core.telegram.org/bots/api#uniquegiftsymbol
type UniqueGiftSymbol struct {
	Name           string  `json:"name"`
	Sticker        Sticker `json:"sticker"`
	RarityPerMille int     `json:"rarity_per_mille"`
}

// https://core.telegram.org/bots/api#uniquegiftbackdropcolors
type UniqueGiftBackdropColors struct {
	CenterColor int `json:"center_color"`
	EdgeColor   int `json:"edge_color"`
	SymbolColor int `json:"symbol_color"`
	TextColor   int `json:"text_color"`
}

// https://core.telegram.org/bots/api#uniquegiftbackdrop
type UniqueGiftBackdrop struct {
	Name           string                   `json:"name"`
	Colors         UniqueGiftBackdropColors `json:"colors"`
	RarityPerMille int                      `json:"rarity_per_mille"`
}

// https://core.telegram.org/bots/api#uniquegiftcolors
type UniqueGiftColors struct {
	ModelCustomEmojiID    string `json:"model_custom_emoji_id"`
	SymbolCustomEmojiID   string `json:"symbol_custom_emoji_id"`
	LightThemeMainColor   int    `json:"light_theme_main_color"`
	LightThemeOtherColors []int  `json:"light_theme_other_colors,omitempty"`
	DarkThemeMainColor    int    `json:"dark_theme_main_color"`
	DarkThemeOtherColors  []int  `json:"dark_theme_other_colors,omitempty"`
}

// https://core.telegram.org/bots/api#uniquegift
type UniqueGift struct {
	GiftID           string             `json:"gift_id"`
	BaseName         string             `json:"base_name"`
	Name             string             `json:"name"`
	Number           int                `json:"number"`
	Model            UniqueGiftModel    `json:"model"`
	Symbol           UniqueGiftSymbol   `json:"symbol"`
	Backdrop         UniqueGiftBackdrop `json:"backdrop"`
	IsPremium        bool               `json:"is_premium,omitempty"`
	IsBurned         bool               `json:"is_burned,omitempty"`
	IsFromBlockchain bool               `json:"is_from_blockchain,omitempty"`
	Colors           *UniqueGiftColors  `json:"colors,omitempty"`
	PublisherChat    *Chat              `json:"publisher_chat,omitempty"`
}

// https://core.telegram.org/bots/api#giftinfo
type GiftInfo struct {
	Gift                    Gift            `json:"gift"`
	OwnedGiftID             string          `json:"owned_gift_id,omitempty"`
	ConvertStarCount        int             `json:"convert_star_count,omitempty"`
	PrepaidUpgradeStarCount int             `json:"prepaid_upgrade_star_count,omitempty"`
	IsUpgradeSeparate       bool            `json:"is_upgrade_separate,omitempty"`
	CanBeUpgraded           bool            `json:"can_be_upgraded,omitempty"`
	Text                    string          `json:"text,omitempty"`
	Entities                []MessageEntity `json:"entities,omitempty"`
	IsPrivate               bool            `json:"is_private,omitempty"`
	UniqueGiftNumber        int             `json:"unique_gift_number,omitempty"`
}

// https://core.telegram.org/bots/api#uniquegiftinfo
type UniqueGiftInfo struct {
	Gift               UniqueGift `json:"gift"`
	Origin             string     `json:"origin"`
	LastResaleCurrency string     `json:"last_resale_currency,omitempty"`
	LastResaleAmount   int        `json:"last_resale_amount,omitempty"`
	OwnedGiftID        string     `json:"owned_gift_id,omitempty"`
	TransferStarCount  int        `json:"transfer_star_count,omitempty"`
	NextTransferDate   int64      `json:"next_transfer_date,omitempty"`
}

// https://core.telegram.org/bots/api#passportdata
type PassportData struct {
	Data        []EncryptedPassportElement `json:"data"`
	Credentials EncryptedCredentials       `json:"credentials"`
}

// https://core.telegram.org/bots/api#encryptedpassportelement
type EncryptedPassportElement struct {
	Type        string         `json:"type"`
	Data        string         `json:"data,omitempty"`
	PhoneNumber string         `json:"phone_number,omitempty"`
	Email       string         `json:"email,omitempty"`
	Files       []PassportFile `json:"files,omitempty"`
	FrontSide   *PassportFile  `json:"front_side,omitempty"`
	ReverseSide *PassportFile  `json:"reverse_side,omitempty"`
	Selfie      *PassportFile  `json:"selfie,omitempty"`
	Translation []PassportFile `json:"translation,omitempty"`
	Hash        string         `json:"hash"`
}

// https://core.telegram.org/bots/api#passportfile
type PassportFile struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	FileSize     int64  `json:"file_size"`
	FileDate     int64  `json:"file_date"`
}

// https://core.telegram.org/bots/api#encryptedcredentials
type EncryptedCredentials struct {
	Data   string `json:"data"`
	Hash   string `json:"hash"`
	Secret string `json:"secret"`
}

// https://core.telegram.org/bots/api#ownedgifts
type OwnedGifts struct {
	TotalCount int         `json:"total_count"`
	Gifts      []OwnedGift `json:"gifts"`
	NextOffset string      `json:"next_offset,omitempty"`
}

// https://core.telegram.org/bots/api#ownedgift
type OwnedGift interface{}

// https://core.telegram.org/bots/api#gifts
type Gifts struct {
	Gifts []Gift `json:"gifts"`
}

// https://core.telegram.org/bots/api#startransactions
type StarTransactions struct {
	Transactions []StarTransaction `json:"transactions"`
}

// https://core.telegram.org/bots/api#startransaction
type StarTransaction struct {
	ID             string      `json:"id"`
	Amount         int         `json:"amount"`
	NanostarAmount int         `json:"nanostar_amount,omitempty"`
	Date           int64       `json:"date"`
	Source         interface{} `json:"source,omitempty"`
	Receiver       interface{} `json:"receiver,omitempty"`
}

// https://core.telegram.org/bots/api#sentguestmessage
type SentGuestMessage struct {
	InlineMessageID string `json:"inline_message_id,omitempty"`
}

// https://core.telegram.org/bots/api#sentwebappmessage
type SentWebAppMessage struct {
	InlineMessageID string `json:"inline_message_id,omitempty"`
}

// https://core.telegram.org/bots/api#preparedinlinemessage
type PreparedInlineMessage struct {
	ID             string `json:"id"`
	ExpirationDate int64  `json:"expiration_date"`
}

// https://core.telegram.org/bots/api#preparedkeyboardbutton
type PreparedKeyboardButton struct {
	ID string `json:"id"`
}

// https://core.telegram.org/bots/api#botaccesssettings
type BotAccessSettings struct {
	IsAccessRestricted bool   `json:"is_access_restricted"`
	AddedUsers         []User `json:"added_users,omitempty"`
}
