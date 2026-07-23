package methods

import (
	"errors"
	"fmt"

	"github.com/erfjab/egobot/models"
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

// https://core.telegram.org/bots/api#approvesuggestedpost
func (r *Requester) ApproveSuggestedPost(params models.ApproveSuggestedPostParams) (bool, error) {
	respBody, err := r.Request("approveSuggestedPost", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#declinesuggestedpost
func (r *Requester) DeclineSuggestedPost(params models.DeclineSuggestedPostParams) (bool, error) {
	respBody, err := r.Request("declineSuggestedPost", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#answerguestquery
func (r *Requester) AnswerGuestQuery(params models.AnswerGuestQueryParams) (*models.SentGuestMessage, error) {
	respBody, err := r.Request("answerGuestQuery", params)
	if err != nil {
		return nil, err
	}

	var result models.SentGuestMessage
	if err := r.ParseResponse(respBody, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// https://core.telegram.org/bots/api#answerchatjoinrequestquery
func (r *Requester) AnswerChatJoinRequestQuery(params models.AnswerChatJoinRequestQueryParams) (bool, error) {
	respBody, err := r.Request("answerChatJoinRequestQuery", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#sendchatjoinrequestwebapp
func (r *Requester) SendChatJoinRequestWebApp(params models.SendChatJoinRequestWebAppParams) (bool, error) {
	respBody, err := r.Request("sendChatJoinRequestWebApp", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#readbusinessmessage
func (r *Requester) ReadBusinessMessage(params models.ReadBusinessMessageParams) (bool, error) {
	respBody, err := r.Request("readBusinessMessage", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#deletebusinessmessages
func (r *Requester) DeleteBusinessMessages(params models.DeleteBusinessMessagesParams) (bool, error) {
	if len(params.MessageIDs) == 0 {
		return false, errors.New("message_ids is required")
	}

	respBody, err := r.Request("deleteBusinessMessages", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#setbusinessaccountname
func (r *Requester) SetBusinessAccountName(params models.SetBusinessAccountNameParams) (bool, error) {
	respBody, err := r.Request("setBusinessAccountName", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#setbusinessaccountusername
func (r *Requester) SetBusinessAccountUsername(params models.SetBusinessAccountUsernameParams) (bool, error) {
	respBody, err := r.Request("setBusinessAccountUsername", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#setbusinessaccountbio
func (r *Requester) SetBusinessAccountBio(params models.SetBusinessAccountBioParams) (bool, error) {
	respBody, err := r.Request("setBusinessAccountBio", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#setbusinessaccountprofilephoto
func (r *Requester) SetBusinessAccountProfilePhoto(params models.SetBusinessAccountProfilePhotoParams) (bool, error) {
	respBody, err := r.Request("setBusinessAccountProfilePhoto", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#removebusinessaccountprofilephoto
func (r *Requester) RemoveBusinessAccountProfilePhoto(params models.RemoveBusinessAccountProfilePhotoParams) (bool, error) {
	respBody, err := r.Request("removeBusinessAccountProfilePhoto", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#setbusinessaccountgiftsettings
func (r *Requester) SetBusinessAccountGiftSettings(params models.SetBusinessAccountGiftSettingsParams) (bool, error) {
	respBody, err := r.Request("setBusinessAccountGiftSettings", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#getbusinessaccountstarbalance
func (r *Requester) GetBusinessAccountStarBalance(businessConnectionID string) (*models.StarAmount, error) {
	params := map[string]interface{}{
		"business_connection_id": businessConnectionID,
	}

	respBody, err := r.Request("getBusinessAccountStarBalance", params)
	if err != nil {
		return nil, err
	}

	var result models.StarAmount
	if err := r.ParseResponse(respBody, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// https://core.telegram.org/bots/api#transferbusinessaccountstars
func (r *Requester) TransferBusinessAccountStars(params models.TransferBusinessAccountStarsParams) (bool, error) {
	respBody, err := r.Request("transferBusinessAccountStars", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#convertgifttostars
func (r *Requester) ConvertGiftToStars(params models.ConvertGiftToStarsParams) (bool, error) {
	respBody, err := r.Request("convertGiftToStars", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#upgradegift
func (r *Requester) UpgradeGift(params models.UpgradeGiftParams) (bool, error) {
	respBody, err := r.Request("upgradeGift", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#transfergift
func (r *Requester) TransferGift(params models.TransferGiftParams) (bool, error) {
	respBody, err := r.Request("transferGift", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#poststory
func (r *Requester) PostStory(params models.PostStoryParams) (*models.Story, error) {
	respBody, err := r.Request("postStory", params)
	if err != nil {
		return nil, err
	}

	var result models.Story
	if err := r.ParseResponse(respBody, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// https://core.telegram.org/bots/api#repoststory
func (r *Requester) RepostStory(params models.RepostStoryParams) (*models.Story, error) {
	respBody, err := r.Request("repostStory", params)
	if err != nil {
		return nil, err
	}

	var result models.Story
	if err := r.ParseResponse(respBody, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// https://core.telegram.org/bots/api#editstory
func (r *Requester) EditStory(params models.EditStoryParams) (*models.Story, error) {
	respBody, err := r.Request("editStory", params)
	if err != nil {
		return nil, err
	}

	var result models.Story
	if err := r.ParseResponse(respBody, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// https://core.telegram.org/bots/api#deletestory
func (r *Requester) DeleteStory(params models.DeleteStoryParams) (bool, error) {
	respBody, err := r.Request("deleteStory", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#savepreparedinlinemessage
func (r *Requester) SavePreparedInlineMessage(params models.SavePreparedInlineMessageParams) (*models.PreparedInlineMessage, error) {
	respBody, err := r.Request("savePreparedInlineMessage", params)
	if err != nil {
		return nil, err
	}

	var result models.PreparedInlineMessage
	if err := r.ParseResponse(respBody, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// https://core.telegram.org/bots/api#savepreparedkeyboardbutton
func (r *Requester) SavePreparedKeyboardButton(params models.SavePreparedKeyboardButtonParams) (*models.PreparedKeyboardButton, error) {
	respBody, err := r.Request("savePreparedKeyboardButton", params)
	if err != nil {
		return nil, err
	}

	var result models.PreparedKeyboardButton
	if err := r.ParseResponse(respBody, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
