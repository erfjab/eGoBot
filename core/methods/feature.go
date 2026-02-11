package methods

import (
	"egobot/models"
	"fmt"
)

// https://core.telegram.org/bots/api#answerinlinequery
func (r *Requester) AnswerInlineQuery(params models.AnswerInlineQueryParams) (bool, error) {
	if params.InlineQueryID == "" {
		return false, fmt.Errorf("inline_query_id cannot be empty")
	}
	if len(params.Results) == 0 {
		return false, fmt.Errorf("results cannot be empty")
	}

	respBody, err := r.Request("answerInlineQuery", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#answercallbackquery
func (r *Requester) AnswerCallbackQuery(callbackQueryID string, text string, showAlert bool) (bool, error) {
	if callbackQueryID == "" {
		return false, fmt.Errorf("callback_query_id cannot be empty")
	}

	params := map[string]interface{}{
		"callback_query_id": callbackQueryID,
		"text":              text,
		"show_alert":        showAlert,
	}

	respBody, err := r.Request("answerCallbackQuery", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}

	return result, nil
}

// https://core.telegram.org/bots/api#getstickerset
func (r *Requester) GetStickerSet(params models.GetStickerSetParams) (*models.StickerSet, error) {
	if params.Name == "" {
		return nil, fmt.Errorf("name cannot be empty")
	}

	respBody, err := r.Request("getStickerSet", params)
	if err != nil {
		return nil, err
	}

	var result models.StickerSet
	if err := r.ParseResponse(respBody, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// https://core.telegram.org/bots/api#getcustomemojistickers
func (r *Requester) GetCustomEmojiStickers(params models.GetCustomEmojiStickersParams) ([]models.Sticker, error) {
	if len(params.CustomEmojiIDs) == 0 {
		return nil, fmt.Errorf("custom_emoji_ids cannot be empty")
	}

	respBody, err := r.Request("getCustomEmojiStickers", params)
	if err != nil {
		return nil, err
	}

	var result []models.Sticker
	if err := r.ParseResponse(respBody, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#uploadstickerfile
func (r *Requester) UploadStickerFile(params models.UploadStickerFileParams) (*models.File, error) {
	if params.Sticker == nil {
		return nil, fmt.Errorf("sticker cannot be nil")
	}

	respBody, err := r.Request("uploadStickerFile", params)
	if err != nil {
		return nil, err
	}

	var result models.File
	if err := r.ParseResponse(respBody, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// https://core.telegram.org/bots/api#createnewstickerset
func (r *Requester) CreateNewStickerSet(params models.CreateNewStickerSetParams) (bool, error) {
	if params.Name == "" || params.Title == "" {
		return false, fmt.Errorf("name and title cannot be empty")
	}
	if len(params.Stickers) == 0 {
		return false, fmt.Errorf("stickers cannot be empty")
	}

	respBody, err := r.Request("createNewStickerSet", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#addstickertoset
func (r *Requester) AddStickerToSet(params models.AddStickerToSetParams) (bool, error) {
	if params.Name == "" {
		return false, fmt.Errorf("name cannot be empty")
	}
	if params.Sticker.Sticker == nil {
		return false, fmt.Errorf("sticker cannot be nil")
	}

	respBody, err := r.Request("addStickerToSet", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#setstickerpositioninset
func (r *Requester) SetStickerPositionInSet(params models.SetStickerPositionInSetParams) (bool, error) {
	if params.Sticker == "" {
		return false, fmt.Errorf("sticker cannot be empty")
	}

	respBody, err := r.Request("setStickerPositionInSet", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#deletestickerfromset
func (r *Requester) DeleteStickerFromSet(params models.DeleteStickerFromSetParams) (bool, error) {
	if params.Sticker == "" {
		return false, fmt.Errorf("sticker cannot be empty")
	}

	respBody, err := r.Request("deleteStickerFromSet", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#setstickersetthumbnail
func (r *Requester) SetStickerSetThumbnail(params models.SetStickerSetThumbnailParams) (bool, error) {
	if params.Name == "" {
		return false, fmt.Errorf("name cannot be empty")
	}

	respBody, err := r.Request("setStickerSetThumbnail", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#sendinvoice
func (r *Requester) SendInvoice(params models.SendInvoiceParams) (*models.Message, error) {
	if params.Title == "" || params.Description == "" || params.Payload == "" || params.Currency == "" {
		return nil, fmt.Errorf("title, description, payload and currency are required")
	}
	if len(params.Prices) == 0 {
		return nil, fmt.Errorf("prices cannot be empty")
	}

	respBody, err := r.Request("sendInvoice", params)
	if err != nil {
		return nil, err
	}

	var result models.Message
	if err := r.ParseResponse(respBody, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// https://core.telegram.org/bots/api#createinvoicelink
func (r *Requester) CreateInvoiceLink(params models.CreateInvoiceLinkParams) (string, error) {
	if params.Title == "" || params.Description == "" || params.Payload == "" || params.Currency == "" {
		return "", fmt.Errorf("title, description, payload and currency are required")
	}
	if len(params.Prices) == 0 {
		return "", fmt.Errorf("prices cannot be empty")
	}

	respBody, err := r.Request("createInvoiceLink", params)
	if err != nil {
		return "", err
	}

	var result string
	if err := r.ParseResponse(respBody, &result); err != nil {
		return "", err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#answershippingquery
func (r *Requester) AnswerShippingQuery(params models.AnswerShippingQueryParams) (bool, error) {
	if params.ShippingQueryID == "" {
		return false, fmt.Errorf("shipping_query_id cannot be empty")
	}

	respBody, err := r.Request("answerShippingQuery", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#answerprecheckoutquery
func (r *Requester) AnswerPreCheckoutQuery(params models.AnswerPreCheckoutQueryParams) (bool, error) {
	if params.PreCheckoutQueryID == "" {
		return false, fmt.Errorf("pre_checkout_query_id cannot be empty")
	}

	respBody, err := r.Request("answerPreCheckoutQuery", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#sendgame
func (r *Requester) SendGame(params models.SendGameParams) (*models.Message, error) {
	if params.GameShortName == "" {
		return nil, fmt.Errorf("game_short_name cannot be empty")
	}

	respBody, err := r.Request("sendGame", params)
	if err != nil {
		return nil, err
	}

	var result models.Message
	if err := r.ParseResponse(respBody, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// https://core.telegram.org/bots/api#setgamescore
func (r *Requester) SetGameScore(params models.SetGameScoreParams) (*models.Message, error) {
	respBody, err := r.Request("setGameScore", params)
	if err != nil {
		return nil, err
	}

	var result models.Message
	if err := r.ParseResponse(respBody, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// https://core.telegram.org/bots/api#getgamehighscores
func (r *Requester) GetGameHighScores(params models.GetGameHighScoresParams) ([]models.GameHighScore, error) {
	respBody, err := r.Request("getGameHighScores", params)
	if err != nil {
		return nil, err
	}

	var result []models.GameHighScore
	if err := r.ParseResponse(respBody, &result); err != nil {
		return nil, err
	}
	return result, nil
}
