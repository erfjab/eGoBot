package models

// https://core.telegram.org/bots/api#labeledprice
type LabeledPrice struct {
	Label  string `json:"label"`
	Amount int    `json:"amount"`
}

// https://core.telegram.org/bots/api#invoice
type Invoice struct {
	Title          string `json:"title"`
	Description    string `json:"description"`
	StartParameter string `json:"start_parameter"`
	Currency       string `json:"currency"`
	TotalAmount    int    `json:"total_amount"`
}

// https://core.telegram.org/bots/api#shippingaddress
type ShippingAddress struct {
	CountryCode string `json:"country_code"`
	State       string `json:"state"`
	City        string `json:"city"`
	StreetLine1 string `json:"street_line1"`
	StreetLine2 string `json:"street_line2"`
	PostCode    string `json:"post_code"`
}

// https://core.telegram.org/bots/api#orderinfo
type OrderInfo struct {
	Name            string           `json:"name,omitempty"`
	PhoneNumber     string           `json:"phone_number,omitempty"`
	Email           string           `json:"email,omitempty"`
	ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"`
}

// https://core.telegram.org/bots/api#shippingoption
type ShippingOption struct {
	ID     string         `json:"id"`
	Title  string         `json:"title"`
	Prices []LabeledPrice `json:"prices"`
}

// https://core.telegram.org/bots/api#successfulpayment
type SuccessfulPayment struct {
	Currency                string     `json:"currency"`
	TotalAmount             int        `json:"total_amount"`
	InvoicePayload          string     `json:"invoice_payload"`
	ShippingOptionID        string     `json:"shipping_option_id,omitempty"`
	OrderInfo               *OrderInfo `json:"order_info,omitempty"`
	TelegramPaymentChargeID string     `json:"telegram_payment_charge_id"`
	ProviderPaymentChargeID string     `json:"provider_payment_charge_id"`
}

// https://core.telegram.org/bots/api#shippingquery
type ShippingQuery struct {
	ID              string          `json:"id"`
	From            User            `json:"from"`
	InvoicePayload  string          `json:"invoice_payload"`
	ShippingAddress ShippingAddress `json:"shipping_address"`
}

// https://core.telegram.org/bots/api#precheckoutquery
type PreCheckoutQuery struct {
	ID               string     `json:"id"`
	From             User       `json:"from"`
	Currency         string     `json:"currency"`
	TotalAmount      int        `json:"total_amount"`
	InvoicePayload   string     `json:"invoice_payload"`
	ShippingOptionID string     `json:"shipping_option_id,omitempty"`
	OrderInfo        *OrderInfo `json:"order_info,omitempty"`
}

// https://core.telegram.org/bots/api#sendinvoice
type SendInvoiceParams struct {
	ChatID                    interface{}   `json:"chat_id"`
	MessageThreadID           int           `json:"message_thread_id,omitempty"`
	Title                     string        `json:"title"`
	Description               string        `json:"description"`
	Payload                   string        `json:"payload"`
	ProviderToken             string        `json:"provider_token,omitempty"`
	Currency                  string        `json:"currency"`
	Prices                    []LabeledPrice `json:"prices"`
	MaxTipAmount              int           `json:"max_tip_amount,omitempty"`
	SuggestedTipAmounts       []int         `json:"suggested_tip_amounts,omitempty"`
	StartParameter            string        `json:"start_parameter,omitempty"`
	ProviderData              string        `json:"provider_data,omitempty"`
	PhotoURL                  string        `json:"photo_url,omitempty"`
	PhotoSize                 int           `json:"photo_size,omitempty"`
	PhotoWidth                int           `json:"photo_width,omitempty"`
	PhotoHeight               int           `json:"photo_height,omitempty"`
	NeedName                  bool          `json:"need_name,omitempty"`
	NeedPhoneNumber           bool          `json:"need_phone_number,omitempty"`
	NeedEmail                 bool          `json:"need_email,omitempty"`
	NeedShippingAddress       bool          `json:"need_shipping_address,omitempty"`
	SendPhoneNumberToProvider bool          `json:"send_phone_number_to_provider,omitempty"`
	SendEmailToProvider       bool          `json:"send_email_to_provider,omitempty"`
	IsFlexible                bool          `json:"is_flexible,omitempty"`
	DisableNotification       bool          `json:"disable_notification,omitempty"`
	ProtectContent            bool          `json:"protect_content,omitempty"`
	ReplyParameters           *ReplyParameters `json:"reply_parameters,omitempty"`
	ReplyMarkup               *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#createinvoicelink
type CreateInvoiceLinkParams struct {
	Title                     string         `json:"title"`
	Description               string         `json:"description"`
	Payload                   string         `json:"payload"`
	ProviderToken             string         `json:"provider_token,omitempty"`
	Currency                  string         `json:"currency"`
	Prices                    []LabeledPrice `json:"prices"`
	MaxTipAmount              int            `json:"max_tip_amount,omitempty"`
	SuggestedTipAmounts       []int          `json:"suggested_tip_amounts,omitempty"`
	ProviderData              string         `json:"provider_data,omitempty"`
	PhotoURL                  string         `json:"photo_url,omitempty"`
	PhotoSize                 int            `json:"photo_size,omitempty"`
	PhotoWidth                int            `json:"photo_width,omitempty"`
	PhotoHeight               int            `json:"photo_height,omitempty"`
	NeedName                  bool           `json:"need_name,omitempty"`
	NeedPhoneNumber           bool           `json:"need_phone_number,omitempty"`
	NeedEmail                 bool           `json:"need_email,omitempty"`
	NeedShippingAddress       bool           `json:"need_shipping_address,omitempty"`
	SendPhoneNumberToProvider bool           `json:"send_phone_number_to_provider,omitempty"`
	SendEmailToProvider       bool           `json:"send_email_to_provider,omitempty"`
	IsFlexible                bool           `json:"is_flexible,omitempty"`
}

// https://core.telegram.org/bots/api#answershippingquery
type AnswerShippingQueryParams struct {
	ShippingQueryID string           `json:"shipping_query_id"`
	OK              bool             `json:"ok"`
	ShippingOptions []ShippingOption `json:"shipping_options,omitempty"`
	ErrorMessage    string           `json:"error_message,omitempty"`
}

// https://core.telegram.org/bots/api#answerprecheckoutquery
type AnswerPreCheckoutQueryParams struct {
	PreCheckoutQueryID string `json:"pre_checkout_query_id"`
	OK                 bool   `json:"ok"`
	ErrorMessage       string `json:"error_message,omitempty"`
}

// https://core.telegram.org/bots/api#refundedpayment
type RefundedPayment struct {
	Currency                string `json:"currency"`
	TotalAmount             int    `json:"total_amount"`
	InvoicePayload          string `json:"invoice_payload"`
	TelegramPaymentChargeID string `json:"telegram_payment_charge_id"`
	ProviderPaymentChargeID string `json:"provider_payment_charge_id,omitempty"`
}
