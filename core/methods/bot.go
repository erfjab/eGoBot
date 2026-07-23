package methods

import (
	"fmt"

	"github.com/erfjab/egobot/models"
)

// https://core.telegram.org/bots/api#getme
func (r *Requester) GetMe() (*models.User, error) {
	respBody, err := r.Request("getMe", nil)
	if err != nil {
		return nil, err
	}

	var user models.User
	if err := r.ParseResponse(respBody, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

// https://core.telegram.org/bots/api#logout
func (r *Requester) LogOut() (bool, error) {
	respBody, err := r.Request("logOut", nil)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#close
func (r *Requester) Close() (bool, error) {
	respBody, err := r.Request("close", nil)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#setmycommands
func (r *Requester) SetMyCommands(params models.SetMyCommandsParams) (bool, error) {
	if len(params.Commands) == 0 {
		return false, fmt.Errorf("commands cannot be empty")
	}

	respBody, err := r.Request("setMyCommands", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#deletemycommands
func (r *Requester) DeleteMyCommands(params models.DeleteMyCommandsParams) (bool, error) {
	respBody, err := r.Request("deleteMyCommands", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#getmycommands
func (r *Requester) GetMyCommands(params models.GetMyCommandsParams) ([]models.BotCommand, error) {
	respBody, err := r.Request("getMyCommands", params)
	if err != nil {
		return nil, err
	}

	var result []models.BotCommand
	if err := r.ParseResponse(respBody, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#setmyname
func (r *Requester) SetMyName(params models.SetMyNameParams) (bool, error) {
	respBody, err := r.Request("setMyName", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#getmyname
func (r *Requester) GetMyName(params models.GetMyNameParams) (*models.BotName, error) {
	respBody, err := r.Request("getMyName", params)
	if err != nil {
		return nil, err
	}

	var result models.BotName
	if err := r.ParseResponse(respBody, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// https://core.telegram.org/bots/api#setmydescription
func (r *Requester) SetMyDescription(params models.SetMyDescriptionParams) (bool, error) {
	respBody, err := r.Request("setMyDescription", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#getmydescription
func (r *Requester) GetMyDescription(params models.GetMyDescriptionParams) (*models.BotDescription, error) {
	respBody, err := r.Request("getMyDescription", params)
	if err != nil {
		return nil, err
	}

	var result models.BotDescription
	if err := r.ParseResponse(respBody, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// https://core.telegram.org/bots/api#setmyshortdescription
func (r *Requester) SetMyShortDescription(params models.SetMyShortDescriptionParams) (bool, error) {
	respBody, err := r.Request("setMyShortDescription", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#getmyshortdescription
func (r *Requester) GetMyShortDescription(params models.GetMyShortDescriptionParams) (*models.BotShortDescription, error) {
	respBody, err := r.Request("getMyShortDescription", params)
	if err != nil {
		return nil, err
	}

	var result models.BotShortDescription
	if err := r.ParseResponse(respBody, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// https://core.telegram.org/bots/api#setchatmenubutton
func (r *Requester) SetChatMenuButton(params models.SetChatMenuButtonParams) (bool, error) {
	respBody, err := r.Request("setChatMenuButton", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#getchatmenubutton
func (r *Requester) GetChatMenuButton(params models.GetChatMenuButtonParams) (*models.MenuButton, error) {
	respBody, err := r.Request("getChatMenuButton", params)
	if err != nil {
		return nil, err
	}

	var result models.MenuButton
	if err := r.ParseResponse(respBody, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// https://core.telegram.org/bots/api#setmydefaultadministratorrights
func (r *Requester) SetMyDefaultAdministratorRights(params models.SetMyDefaultAdministratorRightsParams) (bool, error) {
	respBody, err := r.Request("setMyDefaultAdministratorRights", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#getmydefaultadministratorrights
func (r *Requester) GetMyDefaultAdministratorRights(params models.GetMyDefaultAdministratorRightsParams) (*models.ChatAdministratorRights, error) {
	respBody, err := r.Request("getMyDefaultAdministratorRights", params)
	if err != nil {
		return nil, err
	}

	var result models.ChatAdministratorRights
	if err := r.ParseResponse(respBody, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// https://core.telegram.org/bots/api#setuseremojistatus
func (r *Requester) SetUserEmojiStatus(params models.SetUserEmojiStatusParams) (bool, error) {
	respBody, err := r.Request("setUserEmojiStatus", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#getbusinessconnection
func (r *Requester) GetBusinessConnection(businessConnectionID string) (*models.BusinessConnection, error) {
	params := map[string]interface{}{
		"business_connection_id": businessConnectionID,
	}

	respBody, err := r.Request("getBusinessConnection", params)
	if err != nil {
		return nil, err
	}

	var result models.BusinessConnection
	if err := r.ParseResponse(respBody, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// https://core.telegram.org/bots/api#getmanagedbottoken
func (r *Requester) GetManagedBotToken(userID int64) (string, error) {
	params := map[string]interface{}{
		"user_id": userID,
	}

	respBody, err := r.Request("getManagedBotToken", params)
	if err != nil {
		return "", err
	}

	var result string
	if err := r.ParseResponse(respBody, &result); err != nil {
		return "", err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#replacemanagedbottoken
func (r *Requester) ReplaceManagedBotToken(userID int64) (string, error) {
	params := map[string]interface{}{
		"user_id": userID,
	}

	respBody, err := r.Request("replaceManagedBotToken", params)
	if err != nil {
		return "", err
	}

	var result string
	if err := r.ParseResponse(respBody, &result); err != nil {
		return "", err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#getmanagedbotaccesssettings
func (r *Requester) GetManagedBotAccessSettings(userID int64) (*models.BotAccessSettings, error) {
	params := map[string]interface{}{
		"user_id": userID,
	}

	respBody, err := r.Request("getManagedBotAccessSettings", params)
	if err != nil {
		return nil, err
	}

	var result models.BotAccessSettings
	if err := r.ParseResponse(respBody, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// https://core.telegram.org/bots/api#setmanagedbotaccesssettings
func (r *Requester) SetManagedBotAccessSettings(isAccessRestricted bool, addedUserIDs []int64) (bool, error) {
	params := map[string]interface{}{
		"is_access_restricted": isAccessRestricted,
		"added_user_ids":       addedUserIDs,
	}

	respBody, err := r.Request("setManagedBotAccessSettings", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#getavailablegifts
func (r *Requester) GetAvailableGifts() (*models.Gifts, error) {
	respBody, err := r.Request("getAvailableGifts", nil)
	if err != nil {
		return nil, err
	}

	var result models.Gifts
	if err := r.ParseResponse(respBody, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// https://core.telegram.org/bots/api#sendgift
func (r *Requester) SendGift(params models.SendGiftParams) (bool, error) {
	if params.GiftID == "" {
		return false, fmt.Errorf("gift_id cannot be empty")
	}

	respBody, err := r.Request("sendGift", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#giftpremiumsubscription
func (r *Requester) GiftPremiumSubscription(params models.GiftPremiumSubscriptionParams) (bool, error) {
	respBody, err := r.Request("giftPremiumSubscription", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#verifyuser
func (r *Requester) VerifyUser(params models.VerifyUserParams) (bool, error) {
	respBody, err := r.Request("verifyUser", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#verifychat
func (r *Requester) VerifyChat(params models.VerifyChatParams) (bool, error) {
	respBody, err := r.Request("verifyChat", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#removeuserverification
func (r *Requester) RemoveUserVerification(userID int64) (bool, error) {
	params := map[string]interface{}{
		"user_id": userID,
	}

	respBody, err := r.Request("removeUserVerification", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#removechatverification
func (r *Requester) RemoveChatVerification(chatID interface{}) (bool, error) {
	params := map[string]interface{}{
		"chat_id": chatID,
	}

	respBody, err := r.Request("removeChatVerification", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#getmystarbalance
func (r *Requester) GetMyStarBalance() (*models.StarAmount, error) {
	respBody, err := r.Request("getMyStarBalance", nil)
	if err != nil {
		return nil, err
	}

	var result models.StarAmount
	if err := r.ParseResponse(respBody, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// https://core.telegram.org/bots/api#getstartransactions
func (r *Requester) GetStarTransactions(params models.GetStarTransactionsParams) (*models.StarTransactions, error) {
	respBody, err := r.Request("getStarTransactions", params)
	if err != nil {
		return nil, err
	}

	var result models.StarTransactions
	if err := r.ParseResponse(respBody, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// https://core.telegram.org/bots/api#refundstarpayment
func (r *Requester) RefundStarPayment(params models.RefundStarPaymentParams) (bool, error) {
	respBody, err := r.Request("refundStarPayment", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#edituserstarsubscription
func (r *Requester) EditUserStarSubscription(params models.EditUserStarSubscriptionParams) (bool, error) {
	respBody, err := r.Request("editUserStarSubscription", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#getuserpersonalchatmessages
func (r *Requester) GetUserPersonalChatMessages(params models.GetUserPersonalChatMessagesParams) ([]models.Message, error) {
	respBody, err := r.Request("getUserPersonalChatMessages", params)
	if err != nil {
		return nil, err
	}

	var result []models.Message
	if err := r.ParseResponse(respBody, &result); err != nil {
		return nil, err
	}
	return result, nil
}
