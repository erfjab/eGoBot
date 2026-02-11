package methods

import (
	"egobot/models"
	"errors"
	"fmt"
)

// https://core.telegram.org/bots/api#sendmessage
func (r *Requester) SendMessage(params *models.SendMessageParams) (*models.Message, error) {
	if params == nil {
		return nil, fmt.Errorf("params cannot be nil")
	}

	if params.Text == "" {
		return nil, fmt.Errorf("text cannot be empty")
	}

	respBody, err := r.Request("sendMessage", params)
	if err != nil {
		return nil, err
	}

	var message models.Message
	if err := r.ParseResponse(respBody, &message); err != nil {
		return nil, err
	}

	return &message, nil
}

// https://core.telegram.org/bots/api#sendphoto
func (r *Requester) SendPhoto(params *models.SendPhotoParams) (*models.Message, error) {
	if params == nil {
		return nil, fmt.Errorf("params cannot be nil")
	}

	if params.Photo == nil {
		return nil, fmt.Errorf("photo cannot be nil")
	}

	respBody, err := r.Request("sendPhoto", params)
	if err != nil {
		return nil, err
	}

	var message models.Message
	if err := r.ParseResponse(respBody, &message); err != nil {
		return nil, err
	}

	return &message, nil
}

// https://core.telegram.org/bots/api#senddocument
func (r *Requester) SendDocument(params *models.SendDocumentParams) (*models.Message, error) {
	if params == nil {
		return nil, fmt.Errorf("params cannot be nil")
	}

	if params.Document == nil {
		return nil, fmt.Errorf("document cannot be nil")
	}

	respBody, err := r.Request("sendDocument", params)
	if err != nil {
		return nil, err
	}

	var message models.Message
	if err := r.ParseResponse(respBody, &message); err != nil {
		return nil, err
	}

	return &message, nil
}

// https://core.telegram.org/bots/api#sendvideo
func (r *Requester) SendVideo(params *models.SendVideoParams) (*models.Message, error) {
	if params == nil {
		return nil, fmt.Errorf("params cannot be nil")
	}

	if params.Video == nil {
		return nil, fmt.Errorf("video cannot be nil")
	}

	respBody, err := r.Request("sendVideo", params)
	if err != nil {
		return nil, err
	}

	var message models.Message
	if err := r.ParseResponse(respBody, &message); err != nil {
		return nil, err
	}

	return &message, nil
}

// https://core.telegram.org/bots/api#sendaudio
func (r *Requester) SendAudio(params *models.SendAudioParams) (*models.Message, error) {
	if params == nil {
		return nil, fmt.Errorf("params cannot be nil")
	}

	if params.Audio == nil {
		return nil, fmt.Errorf("audio cannot be nil")
	}

	respBody, err := r.Request("sendAudio", params)
	if err != nil {
		return nil, err
	}

	var message models.Message
	if err := r.ParseResponse(respBody, &message); err != nil {
		return nil, err
	}

	return &message, nil
}

// https://core.telegram.org/bots/api#sendanimation
func (r *Requester) SendAnimation(params models.SendAnimationParams) (*models.Message, error) {
	if params.Animation == nil {
		return nil, errors.New("animation is required")
	}

	respBody, err := r.Request("sendAnimation", params)
	if err != nil {
		return nil, err
	}

	var result models.Message
	if err := r.ParseResponse(respBody, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// https://core.telegram.org/bots/api#sendvoice
func (r *Requester) SendVoice(params models.SendVoiceParams) (*models.Message, error) {
	if params.Voice == nil {
		return nil, errors.New("voice is required")
	}

	respBody, err := r.Request("sendVoice", params)
	if err != nil {
		return nil, err
	}

	var result models.Message
	if err := r.ParseResponse(respBody, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// https://core.telegram.org/bots/api#sendvideonote
func (r *Requester) SendVideoNote(params models.SendVideoNoteParams) (*models.Message, error) {
	if params.VideoNote == nil {
		return nil, errors.New("video_note is required")
	}

	respBody, err := r.Request("sendVideoNote", params)
	if err != nil {
		return nil, err
	}

	var result models.Message
	if err := r.ParseResponse(respBody, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// https://core.telegram.org/bots/api#sendmediagroup
func (r *Requester) SendMediaGroup(params models.SendMediaGroupParams) ([]models.Message, error) {
	if len(params.Media) == 0 {
		return nil, errors.New("media is required")
	}

	respBody, err := r.Request("sendMediaGroup", params)
	if err != nil {
		return nil, err
	}

	var result []models.Message
	if err := r.ParseResponse(respBody, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// https://core.telegram.org/bots/api#sendlocation
func (r *Requester) SendLocation(params *models.SendLocationParams) (*models.Message, error) {
	if params == nil {
		return nil, fmt.Errorf("params cannot be nil")
	}

	respBody, err := r.Request("sendLocation", params)
	if err != nil {
		return nil, err
	}

	var message models.Message
	if err := r.ParseResponse(respBody, &message); err != nil {
		return nil, err
	}

	return &message, nil
}

// https://core.telegram.org/bots/api#sendvenue
func (r *Requester) SendVenue(params models.SendVenueParams) (*models.Message, error) {
	if params.Title == "" || params.Address == "" {
		return nil, errors.New("title and address are required")
	}

	respBody, err := r.Request("sendVenue", params)
	if err != nil {
		return nil, err
	}

	var result models.Message
	if err := r.ParseResponse(respBody, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// https://core.telegram.org/bots/api#sendcontact
func (r *Requester) SendContact(params *models.SendContactParams) (*models.Message, error) {
	if params == nil {
		return nil, fmt.Errorf("params cannot be nil")
	}

	respBody, err := r.Request("sendContact", params)
	if err != nil {
		return nil, err
	}

	var message models.Message
	if err := r.ParseResponse(respBody, &message); err != nil {
		return nil, err
	}

	return &message, nil
}

// https://core.telegram.org/bots/api#sendpoll
func (r *Requester) SendPoll(params *models.SendPollParams) (*models.Message, error) {
	if params == nil {
		return nil, fmt.Errorf("params cannot be nil")
	}

	respBody, err := r.Request("sendPoll", params)
	if err != nil {
		return nil, err
	}

	var message models.Message
	if err := r.ParseResponse(respBody, &message); err != nil {
		return nil, err
	}

	return &message, nil
}

// https://core.telegram.org/bots/api#senddice
func (r *Requester) SendDice(params models.SendDiceParams) (*models.Message, error) {
	respBody, err := r.Request("sendDice", params)
	if err != nil {
		return nil, err
	}

	var result models.Message
	if err := r.ParseResponse(respBody, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// https://core.telegram.org/bots/api#sendchecklist
func (r *Requester) SendChecklist(params models.SendChecklistParams) (*models.Message, error) {
	if params.BusinessConnectionID == "" {
		return nil, errors.New("business_connection_id is required")
	}

	respBody, err := r.Request("sendChecklist", params)
	if err != nil {
		return nil, err
	}

	var result models.Message
	if err := r.ParseResponse(respBody, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// https://core.telegram.org/bots/api#sendpaidmedia
func (r *Requester) SendPaidMedia(params models.SendPaidMediaParams) (*models.Message, error) {
	if params.StarCount <= 0 || len(params.Media) == 0 {
		return nil, errors.New("star_count and media are required")
	}

	respBody, err := r.Request("sendPaidMedia", params)
	if err != nil {
		return nil, err
	}

	var result models.Message
	if err := r.ParseResponse(respBody, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// https://core.telegram.org/bots/api#sendsticker
func (r *Requester) SendSticker(params models.SendStickerParams) (*models.Message, error) {
	if params.Sticker == nil {
		return nil, errors.New("sticker is required")
	}

	respBody, err := r.Request("sendSticker", params)
	if err != nil {
		return nil, err
	}

	var result models.Message
	if err := r.ParseResponse(respBody, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// https://core.telegram.org/bots/api#sendmessagedraft
func (r *Requester) SendMessageDraft(params models.SendMessageDraftParams) (bool, error) {
	if params.DraftID == 0 || params.Text == "" {
		return false, errors.New("draft_id and text are required")
	}

	respBody, err := r.Request("sendMessageDraft", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}

	return result, nil
}

// https://core.telegram.org/bots/api#editmessagetext
func (r *Requester) EditMessageText(params *models.EditMessageTextParams) (*models.Message, error) {
	if params == nil {
		return nil, fmt.Errorf("params cannot be nil")
	}

	if params.Text == "" {
		return nil, fmt.Errorf("text cannot be empty")
	}

	respBody, err := r.Request("editMessageText", params)
	if err != nil {
		return nil, err
	}

	var message models.Message
	if err := r.ParseResponse(respBody, &message); err != nil {
		return nil, err
	}

	return &message, nil
}

// https://core.telegram.org/bots/api#editmessagecaption
func (r *Requester) EditMessageCaption(params *models.EditMessageCaptionParams) (*models.Message, error) {
	if params == nil {
		return nil, fmt.Errorf("params cannot be nil")
	}

	respBody, err := r.Request("editMessageCaption", params)
	if err != nil {
		return nil, err
	}

	var message models.Message
	if err := r.ParseResponse(respBody, &message); err != nil {
		return nil, err
	}

	return &message, nil
}

// https://core.telegram.org/bots/api#editmessagemedia
func (r *Requester) EditMessageMedia(params *models.EditMessageMediaParams) (*models.Message, error) {
	if params == nil {
		return nil, fmt.Errorf("params cannot be nil")
	}

	if params.Media == nil {
		return nil, fmt.Errorf("media is required")
	}

	respBody, err := r.Request("editMessageMedia", params)
	if err != nil {
		return nil, err
	}

	var message models.Message
	if err := r.ParseResponse(respBody, &message); err != nil {
		return nil, err
	}

	return &message, nil
}

// https://core.telegram.org/bots/api#editmessagereplymarkup
func (r *Requester) EditMessageReplyMarkup(params *models.EditMessageReplyMarkupParams) (*models.Message, error) {
	if params == nil {
		return nil, fmt.Errorf("params cannot be nil")
	}

	respBody, err := r.Request("editMessageReplyMarkup", params)
	if err != nil {
		return nil, err
	}

	var message models.Message
	if err := r.ParseResponse(respBody, &message); err != nil {
		return nil, err
	}

	return &message, nil
}

// https://core.telegram.org/bots/api#editmessagelivelocation
func (r *Requester) EditMessageLiveLocation(params *models.EditMessageLiveLocationParams) (*models.Message, error) {
	if params == nil {
		return nil, fmt.Errorf("params cannot be nil")
	}

	respBody, err := r.Request("editMessageLiveLocation", params)
	if err != nil {
		return nil, err
	}

	var message models.Message
	if err := r.ParseResponse(respBody, &message); err != nil {
		return nil, err
	}

	return &message, nil
}

// https://core.telegram.org/bots/api#stopmessagelivelocation
func (r *Requester) StopMessageLiveLocation(params *models.StopMessageLiveLocationParams) (*models.Message, error) {
	if params == nil {
		return nil, fmt.Errorf("params cannot be nil")
	}

	respBody, err := r.Request("stopMessageLiveLocation", params)
	if err != nil {
		return nil, err
	}

	var message models.Message
	if err := r.ParseResponse(respBody, &message); err != nil {
		return nil, err
	}

	return &message, nil
}

// https://core.telegram.org/bots/api#editmessagechecklist
func (r *Requester) EditMessageChecklist(params *models.EditMessageChecklistParams) (*models.Message, error) {
	if params == nil {
		return nil, fmt.Errorf("params cannot be nil")
	}

	if params.BusinessConnectionID == "" {
		return nil, fmt.Errorf("business_connection_id is required")
	}

	respBody, err := r.Request("editMessageChecklist", params)
	if err != nil {
		return nil, err
	}

	var message models.Message
	if err := r.ParseResponse(respBody, &message); err != nil {
		return nil, err
	}

	return &message, nil
}

// https://core.telegram.org/bots/api#stoppoll
func (r *Requester) StopPoll(params *models.StopPollParams) (*models.Poll, error) {
	if params == nil {
		return nil, fmt.Errorf("params cannot be nil")
	}

	respBody, err := r.Request("stopPoll", params)
	if err != nil {
		return nil, err
	}

	var poll models.Poll
	if err := r.ParseResponse(respBody, &poll); err != nil {
		return nil, err
	}

	return &poll, nil
}

// https://core.telegram.org/bots/api#deletemessage
func (r *Requester) DeleteMessage(params *models.DeleteMessageParams) (bool, error) {
	if params == nil {
		return false, fmt.Errorf("params cannot be nil")
	}

	respBody, err := r.Request("deleteMessage", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}

	return result, nil
}

// https://core.telegram.org/bots/api#forwardmessage
func (r *Requester) ForwardMessage(params *models.ForwardMessageParams) (*models.Message, error) {
	if params == nil {
		return nil, fmt.Errorf("params cannot be nil")
	}

	respBody, err := r.Request("forwardMessage", params)
	if err != nil {
		return nil, err
	}

	var message models.Message
	if err := r.ParseResponse(respBody, &message); err != nil {
		return nil, err
	}

	return &message, nil
}

// https://core.telegram.org/bots/api#copymessage
func (r *Requester) CopyMessage(params *models.CopyMessageParams) (*models.MessageID, error) {
	if params == nil {
		return nil, fmt.Errorf("params cannot be nil")
	}

	respBody, err := r.Request("copyMessage", params)
	if err != nil {
		return nil, err
	}

	var messageID models.MessageID
	if err := r.ParseResponse(respBody, &messageID); err != nil {
		return nil, err
	}

	return &messageID, nil
}

// https://core.telegram.org/bots/api#copymessages
func (r *Requester) CopyMessages(params models.CopyMessagesParams) ([]models.MessageID, error) {
	if len(params.MessageIDs) == 0 {
		return nil, errors.New("message_ids is required")
	}

	respBody, err := r.Request("copyMessages", params)
	if err != nil {
		return nil, err
	}

	var result []models.MessageID
	if err := r.ParseResponse(respBody, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// https://core.telegram.org/bots/api#forwardmessages
func (r *Requester) ForwardMessages(params models.ForwardMessagesParams) ([]models.MessageID, error) {
	if len(params.MessageIDs) == 0 {
		return nil, errors.New("message_ids is required")
	}

	respBody, err := r.Request("forwardMessages", params)
	if err != nil {
		return nil, err
	}

	var result []models.MessageID
	if err := r.ParseResponse(respBody, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// https://core.telegram.org/bots/api#deletemessages
func (r *Requester) DeleteMessages(params models.DeleteMessagesParams) (bool, error) {
	if len(params.MessageIDs) == 0 {
		return false, errors.New("message_ids is required")
	}

	respBody, err := r.Request("deleteMessages", params)
	if err != nil {
		return false, err
	}

	var result bool
	if err := r.ParseResponse(respBody, &result); err != nil {
		return false, err
	}

	return result, nil
}
