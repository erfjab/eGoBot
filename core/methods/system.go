package methods

import (
	"egobot/models"
	"fmt"
)

// https://core.telegram.org/bots/api#getupdates
func (r *Requester) GetUpdates(params *models.GetUpdatesParams) ([]models.Update, error) {
	respBody, err := r.Request("getUpdates", params)
	if err != nil {
		return nil, err
	}

	var updates []models.Update
	if err := r.ParseResponse(respBody, &updates); err != nil {
		return nil, err
	}

	return updates, nil
}

// https://core.telegram.org/bots/api#setwebhook
func (r *Requester) SetWebhook(params models.SetWebhookParams) (bool, error) {
	if params.URL == "" {
		return false, fmt.Errorf("url cannot be empty")
	}

	respBody, err := r.Request("setWebhook", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#deletewebhook
func (r *Requester) DeleteWebhook(params models.DeleteWebhookParams) (bool, error) {
	respBody, err := r.Request("deleteWebhook", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#getwebhookinfo
func (r *Requester) GetWebhookInfo() (*models.WebhookInfo, error) {
	respBody, err := r.Request("getWebhookInfo", nil)
	if err != nil {
		return nil, err
	}

	var result models.WebhookInfo
	if err := r.ParseResponse(respBody, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// https://core.telegram.org/bots/api#getfile
func (r *Requester) GetFile(fileID string) (*models.File, error) {
	if fileID == "" {
		return nil, fmt.Errorf("file_id cannot be empty")
	}

	params := map[string]interface{}{
		"file_id": fileID,
	}

	respBody, err := r.Request("getFile", params)
	if err != nil {
		return nil, err
	}

	var file models.File
	if err := r.ParseResponse(respBody, &file); err != nil {
		return nil, err
	}

	return &file, nil
}

// https://core.telegram.org/bots/api#getuserprofilephotos
func (r *Requester) GetUserProfilePhotos(params models.GetUserProfilePhotosParams) (*models.UserProfilePhotos, error) {
	respBody, err := r.Request("getUserProfilePhotos", params)
	if err != nil {
		return nil, err
	}

	var result models.UserProfilePhotos
	if err := r.ParseResponse(respBody, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// https://core.telegram.org/bots/api#sendchataction
func (r *Requester) SendChatAction(chatID interface{}, action string) (bool, error) {
	if action == "" {
		return false, fmt.Errorf("action cannot be empty")
	}

	params := map[string]interface{}{
		"chat_id": chatID,
		"action":  action,
	}

	respBody, err := r.Request("sendChatAction", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}

	return result, nil
}

// https://core.telegram.org/bots/api#setmessagereaction
func (r *Requester) SetMessageReaction(params models.SetMessageReactionParams) (bool, error) {
	respBody, err := r.Request("setMessageReaction", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}
	return result, nil
}
