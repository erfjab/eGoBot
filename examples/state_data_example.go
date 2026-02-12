package examples

import (
	"context"
	"fmt"
	"log"

	"github.com/erfjab/egobot/core"
	"github.com/erfjab/egobot/models"
	"github.com/erfjab/egobot/state"
)

// مثال استفاده از state و data که به صورت خودکار لود میشه

// تعریف state group
var UserStates = state.NewStateGroup("user")

var (
	WaitingForName  = UserStates.Add("waiting_for_name")
	WaitingForAge   = UserStates.Add("waiting_for_age")
	WaitingForEmail = UserStates.Add("waiting_for_email")
)

// این handler وقتی که user در حالت WaitingForName هست اجرا میشه
// state و data به صورت خودکار در ctx موجود هستن
func handleNameInput(bot *core.Bot, update *models.Update, ctx *core.Context) error {
	// دیگر نیازی به لود جداگانه state و data نیست!
	// به صورت خودکار توی ctx موجوده
	
	// دریافت state name
	stateName := ctx.GetStateName()
	fmt.Printf("Current state: %s\n", stateName)
	
	// دریافت data که قبلا ذخیره شده
	stateData := ctx.GetStateData()
	if stateData != nil {
		if previousData, ok := stateData["some_key"]; ok {
			fmt.Printf("Previous data: %v\n", previousData)
		}
	}
	
	// پردازش ورودی user
	userName := update.Message.Text
	
	// ذخیره اسم در data
	userID := update.Message.From.ID
	userManager := bot.StateManager.ForUser(userID)
	
	// به‌روزرسانی data
	err := userManager.SetData(context.Background(), map[string]interface{}{
		"name": userName,
	})
	if err != nil {
		return err
	}
	
	// تغییر state به مرحله بعد
	err = userManager.SetState(context.Background(), WaitingForAge)
	if err != nil {
		return err
	}
	
	bot.SendMessage(&models.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "اسم شما ثبت شد! حالا سن خود را وارد کنید:",
	})
	return nil
}

// این handler برای دریافت سن
func handleAgeInput(bot *core.Bot, update *models.Update, ctx *core.Context) error {
	// دریافت data که قبلا ذخیره شده (شامل name)
	stateData := ctx.GetStateData()
	userName := ""
	if stateData != nil {
		if name, ok := stateData["name"].(string); ok {
			userName = name
		}
	}
	
	// پردازش ورودی user
	userAge := update.Message.Text
	
	// ذخیره سن در data
	userID := update.Message.From.ID
	userManager := bot.StateManager.ForUser(userID)
	
	err := userManager.SetData(context.Background(), map[string]interface{}{
		"age": userAge,
	})
	if err != nil {
		return err
	}
	
	// تغییر state
	err = userManager.SetState(context.Background(), WaitingForEmail)
	if err != nil {
		return err
	}
	
	message := fmt.Sprintf("سلام %s! سن شما ثبت شد. حالا ایمیل خود را وارد کنید:", userName)
	bot.SendMessage(&models.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   message,
	})
	return nil
}

// این handler برای دریافت ایمیل و نهایی کردن ثبت‌نام
func handleEmailInput(bot *core.Bot, update *models.Update, ctx *core.Context) error {
	// دریافت تمام data ذخیره شده
	stateData := ctx.GetStateData()
	
	userName := ""
	userAge := ""
	if stateData != nil {
		if name, ok := stateData["name"].(string); ok {
			userName = name
		}
		if age, ok := stateData["age"].(string); ok {
			userAge = age
		}
	}
	
	userEmail := update.Message.Text
	
	// ذخیره ایمیل
	userID := update.Message.From.ID
	userManager := bot.StateManager.ForUser(userID)
	
	err := userManager.SetData(context.Background(), map[string]interface{}{
		"email": userEmail,
	})
	if err != nil {
		return err
	}
	
	// پاک کردن state چون کار تمام شد
	err = userManager.ClearState(context.Background())
	if err != nil {
		log.Printf("Error clearing state: %v", err)
	}
	
	// نمایش اطلاعات نهایی
	finalMessage := fmt.Sprintf(
		"ثبت‌نام کامل شد!\n\nنام: %s\nسن: %s\nایمیل: %s",
		userName, userAge, userEmail,
	)
	
	bot.SendMessage(&models.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   finalMessage,
	})
	
	// حالا میتونی data رو هم پاک کنی یا نگه داری
	// err = userManager.ClearData(context.Background())
	
	return nil
}

// ثبت handlerها
func RegisterStateHandlers(bot *core.Bot) {
	// Handler برای دریافت اسم
	bot.AddHandler(
		core.TextFilter(),
		handleNameInput,
		state.InState(WaitingForName),
	)
	
	// Handler برای دریافت سن
	bot.AddHandler(
		core.TextFilter(),
		handleAgeInput,
		state.InState(WaitingForAge),
	)
	
	// Handler برای دریافت ایمیل
	bot.AddHandler(
		core.TextFilter(),
		handleEmailInput,
		state.InState(WaitingForEmail),
	)
}

// مزایای این رویکرد:
// 1. دیگر نیازی به لود جداگانه state و data در هر handler نیست
// 2. state و data به صورت خودکار و یکجا لود میشن
// 3. کد تمیزتر و خواناتر میشه
// 4. performance بهتر میشه چون یک بار query میزنیم به جای دو بار
// 5. در middleware ها هم به state و data دسترسی داری

