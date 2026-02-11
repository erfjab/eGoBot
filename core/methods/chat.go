package methods

import (
	"egobot/models"
	"fmt"
)

// https://core.telegram.org/bots/api#getchat
func (r *Requester) GetChat(chatID interface{}) (*models.Chat, error) {
	params := map[string]interface{}{
		"chat_id": chatID,
	}

	respBody, err := r.Request("getChat", params)
	if err != nil {
		return nil, err
	}

	var chat models.Chat
	if err := r.ParseResponse(respBody, &chat); err != nil {
		return nil, err
	}

	return &chat, nil
}

// https://core.telegram.org/bots/api#leavechat
func (r *Requester) LeaveChat(chatID interface{}) (bool, error) {
	params := map[string]interface{}{
		"chat_id": chatID,
	}

	respBody, err := r.Request("leaveChat", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}

	return result, nil
}

// https://core.telegram.org/bots/api#getchatadministrators
func (r *Requester) GetChatAdministrators(chatID interface{}) ([]models.ChatMember, error) {
	params := map[string]interface{}{
		"chat_id": chatID,
	}

	respBody, err := r.Request("getChatAdministrators", params)
	if err != nil {
		return nil, err
	}

	var admins []models.ChatMember
	if err := r.ParseResponse(respBody, &admins); err != nil {
		return nil, err
	}

	return admins, nil
}

// https://core.telegram.org/bots/api#getchatmembercount
func (r *Requester) GetChatMemberCount(chatID interface{}) (int, error) {
	params := map[string]interface{}{
		"chat_id": chatID,
	}

	respBody, err := r.Request("getChatMemberCount", params)
	if err != nil {
		return 0, err
	}

	var count int
	if err := r.ParseResponse(respBody, &count); err != nil {
		return 0, err
	}

	return count, nil
}

// https://core.telegram.org/bots/api#getchatmember
func (r *Requester) GetChatMember(params *models.GetChatMemberParams) (*models.ChatMember, error) {
	if params == nil {
		return nil, fmt.Errorf("params cannot be nil")
	}

	respBody, err := r.Request("getChatMember", params)
	if err != nil {
		return nil, err
	}

	var member models.ChatMember
	if err := r.ParseResponse(respBody, &member); err != nil {
		return nil, err
	}

	return &member, nil
}

// https://core.telegram.org/bots/api#setchatphoto
func (r *Requester) SetChatPhoto(params models.SetChatPhotoParams) (bool, error) {
	if params.Photo == nil {
		return false, fmt.Errorf("photo cannot be nil")
	}

	respBody, err := r.Request("setChatPhoto", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#deletechatphoto
func (r *Requester) DeleteChatPhoto(params models.DeleteChatPhotoParams) (bool, error) {
	respBody, err := r.Request("deleteChatPhoto", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#setchattitle
func (r *Requester) SetChatTitle(params models.SetChatTitleParams) (bool, error) {
	if params.Title == "" {
		return false, fmt.Errorf("title cannot be empty")
	}

	respBody, err := r.Request("setChatTitle", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#setchatdescription
func (r *Requester) SetChatDescription(params models.SetChatDescriptionParams) (bool, error) {
	respBody, err := r.Request("setChatDescription", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#pinchatmessage
func (r *Requester) PinChatMessage(params *models.PinChatMessageParams) (bool, error) {
	if params == nil {
		return false, fmt.Errorf("params cannot be nil")
	}

	respBody, err := r.Request("pinChatMessage", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}

	return result, nil
}

// https://core.telegram.org/bots/api#unpinchatmessage
func (r *Requester) UnpinChatMessage(params *models.UnpinChatMessageParams) (bool, error) {
	if params == nil {
		return false, fmt.Errorf("params cannot be nil")
	}

	respBody, err := r.Request("unpinChatMessage", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}

	return result, nil
}

// https://core.telegram.org/bots/api#unpinallchatmessages
func (r *Requester) UnpinAllChatMessages(chatID interface{}) (bool, error) {
	params := map[string]interface{}{
		"chat_id": chatID,
	}

	respBody, err := r.Request("unpinAllChatMessages", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}

	return result, nil
}

// https://core.telegram.org/bots/api#banchatmember
func (r *Requester) BanChatMember(params *models.BanChatMemberParams) (bool, error) {
	if params == nil {
		return false, fmt.Errorf("params cannot be nil")
	}

	respBody, err := r.Request("banChatMember", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}

	return result, nil
}

// https://core.telegram.org/bots/api#unbanchatmember
func (r *Requester) UnbanChatMember(params *models.UnbanChatMemberParams) (bool, error) {
	if params == nil {
		return false, fmt.Errorf("params cannot be nil")
	}

	respBody, err := r.Request("unbanChatMember", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}

	return result, nil
}

// https://core.telegram.org/bots/api#restrictchatmember
func (r *Requester) RestrictChatMember(params *models.RestrictChatMemberParams) (bool, error) {
	if params == nil {
		return false, fmt.Errorf("params cannot be nil")
	}

	respBody, err := r.Request("restrictChatMember", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}

	return result, nil
}

// https://core.telegram.org/bots/api#promotechatmember
func (r *Requester) PromoteChatMember(params *models.PromoteChatMemberParams) (bool, error) {
	if params == nil {
		return false, fmt.Errorf("params cannot be nil")
	}

	respBody, err := r.Request("promoteChatMember", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}

	return result, nil
}

// https://core.telegram.org/bots/api#setchatadministratorcustomtitle
func (r *Requester) SetChatAdministratorCustomTitle(params *models.SetChatAdministratorCustomTitleParams) (bool, error) {
	if params == nil {
		return false, fmt.Errorf("params cannot be nil")
	}

	respBody, err := r.Request("setChatAdministratorCustomTitle", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}

	return result, nil
}

// https://core.telegram.org/bots/api#setchatpermissions
func (r *Requester) SetChatPermissions(params models.SetChatPermissionsParams) (bool, error) {
	respBody, err := r.Request("setChatPermissions", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#exportchatinvitelink
func (r *Requester) ExportChatInviteLink(chatID interface{}) (string, error) {
	params := map[string]interface{}{
		"chat_id": chatID,
	}

	respBody, err := r.Request("exportChatInviteLink", params)
	if err != nil {
		return "", err
	}

	var result string
	if err := r.ParseResponse(respBody, &result); err != nil {
		return "", err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#createchatinvitelink
func (r *Requester) CreateChatInviteLink(params models.CreateChatInviteLinkParams) (*models.ChatInviteLink, error) {
	respBody, err := r.Request("createChatInviteLink", params)
	if err != nil {
		return nil, err
	}

	var result models.ChatInviteLink
	if err := r.ParseResponse(respBody, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// https://core.telegram.org/bots/api#editchatinvitelink
func (r *Requester) EditChatInviteLink(params models.EditChatInviteLinkParams) (*models.ChatInviteLink, error) {
	if params.InviteLink == "" {
		return nil, fmt.Errorf("invite_link cannot be empty")
	}

	respBody, err := r.Request("editChatInviteLink", params)
	if err != nil {
		return nil, err
	}

	var result models.ChatInviteLink
	if err := r.ParseResponse(respBody, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// https://core.telegram.org/bots/api#revokechatinvitelink
func (r *Requester) RevokeChatInviteLink(params models.RevokeChatInviteLinkParams) (*models.ChatInviteLink, error) {
	if params.InviteLink == "" {
		return nil, fmt.Errorf("invite_link cannot be empty")
	}

	respBody, err := r.Request("revokeChatInviteLink", params)
	if err != nil {
		return nil, err
	}

	var result models.ChatInviteLink
	if err := r.ParseResponse(respBody, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// https://core.telegram.org/bots/api#approvechatjoinrequest
func (r *Requester) ApproveChatJoinRequest(params models.ApproveChatJoinRequestParams) (bool, error) {
	respBody, err := r.Request("approveChatJoinRequest", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#declinechatjoinrequest
func (r *Requester) DeclineChatJoinRequest(params models.DeclineChatJoinRequestParams) (bool, error) {
	respBody, err := r.Request("declineChatJoinRequest", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#banchatsenderchat
func (r *Requester) BanChatSenderChat(params models.BanChatSenderChatParams) (bool, error) {
	respBody, err := r.Request("banChatSenderChat", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#unbanchatsenderchat
func (r *Requester) UnbanChatSenderChat(params models.UnbanChatSenderChatParams) (bool, error) {
	respBody, err := r.Request("unbanChatSenderChat", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#setchatstickerset
func (r *Requester) SetChatStickerSet(params models.SetChatStickerSetParams) (bool, error) {
	if params.StickerSetName == "" {
		return false, fmt.Errorf("sticker_set_name cannot be empty")
	}

	respBody, err := r.Request("setChatStickerSet", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#deletechatstickerset
func (r *Requester) DeleteChatStickerSet(params models.DeleteChatStickerSetParams) (bool, error) {
	respBody, err := r.Request("deleteChatStickerSet", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#createforumtopic
func (r *Requester) CreateForumTopic(params models.CreateForumTopicParams) (*models.ForumTopic, error) {
	if params.Name == "" {
		return nil, fmt.Errorf("name cannot be empty")
	}

	respBody, err := r.Request("createForumTopic", params)
	if err != nil {
		return nil, err
	}

	var result models.ForumTopic
	if err := r.ParseResponse(respBody, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// https://core.telegram.org/bots/api#editforumtopic
func (r *Requester) EditForumTopic(params models.EditForumTopicParams) (bool, error) {
	respBody, err := r.Request("editForumTopic", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#closeforumtopic
func (r *Requester) CloseForumTopic(params models.CloseForumTopicParams) (bool, error) {
	respBody, err := r.Request("closeForumTopic", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#reopenforumtopic
func (r *Requester) ReopenForumTopic(params models.ReopenForumTopicParams) (bool, error) {
	respBody, err := r.Request("reopenForumTopic", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#deleteforumtopic
func (r *Requester) DeleteForumTopic(params models.DeleteForumTopicParams) (bool, error) {
	respBody, err := r.Request("deleteForumTopic", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#unpinallforumtopicmessages
func (r *Requester) UnpinAllForumTopicMessages(params models.UnpinAllForumTopicMessagesParams) (bool, error) {
	respBody, err := r.Request("unpinAllForumTopicMessages", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#editgeneralforumtopic
func (r *Requester) EditGeneralForumTopic(params models.EditGeneralForumTopicParams) (bool, error) {
	if params.Name == "" {
		return false, fmt.Errorf("name cannot be empty")
	}

	respBody, err := r.Request("editGeneralForumTopic", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#closegeneralforumtopic
func (r *Requester) CloseGeneralForumTopic(params models.CloseGeneralForumTopicParams) (bool, error) {
	respBody, err := r.Request("closeGeneralForumTopic", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#reopengeneralforumtopic
func (r *Requester) ReopenGeneralForumTopic(params models.ReopenGeneralForumTopicParams) (bool, error) {
	respBody, err := r.Request("reopenGeneralForumTopic", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#hidegeneralforumtopic
func (r *Requester) HideGeneralForumTopic(params models.HideGeneralForumTopicParams) (bool, error) {
	respBody, err := r.Request("hideGeneralForumTopic", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#unhidegeneralforumtopic
func (r *Requester) UnhideGeneralForumTopic(params models.UnhideGeneralForumTopicParams) (bool, error) {
	respBody, err := r.Request("unhideGeneralForumTopic", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}
	return result, nil
}

// https://core.telegram.org/bots/api#setmyprofilephoto
func (r *Requester) SetMyProfilePhoto(photo *models.InputProfilePhoto) (bool, error) {
	if photo == nil {
		return false, fmt.Errorf("photo cannot be nil")
	}

	params := map[string]interface{}{
		"photo": photo,
	}

	respBody, err := r.Request("setMyProfilePhoto", params)
	if err != nil {
		return false, err
	}

	var result bool
	err = r.ParseResponse(respBody, &result)
	return result, err
}

// https://core.telegram.org/bots/api#removemyprofilephoto
func (r *Requester) RemoveMyProfilePhoto() (bool, error) {
	respBody, err := r.Request("removeMyProfilePhoto", nil)
	if err != nil {
		return false, err
	}

	var result bool
	err = r.ParseResponse(respBody, &result)
	return result, err
}

// https://core.telegram.org/bots/api#getuserprofileaudios
func (r *Requester) GetUserProfileAudios(params *models.GetUserProfileAudiosParams) (*models.UserProfileAudios, error) {
	if params == nil || params.UserID == 0 {
		return nil, fmt.Errorf("user_id is required")
	}

	respBody, err := r.Request("getUserProfileAudios", params)
	if err != nil {
		return nil, err
	}

	var result models.UserProfileAudios
	err = r.ParseResponse(respBody, &result)
	return &result, err
}
