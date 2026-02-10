package methods

import (
	"egobot/egobot/models"
	"fmt"
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


