package main

import (
	"egobot/egobot/core"
	"egobot/egobot/models"
	"log"
	"os"
)

// مثال ساده و مینیمال برای error handling

func main() {
	bot := core.NewBot(os.Getenv("BOT_TOKEN"))

	// ثبت error handlers به صورت ماژولار

	// خطاهای مربوط به پیام‌ها
	bot.OnMessageNotFound(handleMessageNotFound)
	bot.OnMessageCantBeEdited(handleMessageCantBeEdited)
	bot.OnMessageTextEmpty(handleMessageTextEmpty)
	bot.OnMessageTooLong(handleMessageTooLong)

	// خطاهای مربوط به کاربر/چت
	bot.OnBotBlocked(handleBotBlocked)
	bot.OnBotKicked(handleBotKicked)
	bot.OnChatNotFound(handleChatNotFound)

	// خطاهای مربوط به رسانه
	bot.OnInvalidFileID(handleInvalidFileID)
	bot.OnButtonDataInvalid(handleButtonDataInvalid)

	// خطاهای تلگرام عمومی
	bot.OnRateLimitError(handleRateLimit)
	bot.OnBadRequest(handleBadRequest)
	bot.OnForbiddenError(handleForbidden)

	// Fallback برای خطاهای مدیریت نشده
	bot.SetFallbackErrorHandler(handleUnknownError)

	// ثبت command handlers
	bot.OnCommand("start", func(b *core.Bot, u *models.Update) error {
		_, err := b.SendMessage(&models.SendMessageParams{
			ChatID: u.Message.Chat.ID,
			Text:   "سلام! ربات با error handling ماژولار 👋",
		})
		return err
	})

	log.Println("🤖 Bot started!")
	bot.StartPolling(nil)
}

// Error Handlers - هر کدام یک مسئولیت مشخص دارند

func handleMessageNotFound(b *core.Bot, u *models.Update, err error) error {
	log.Printf("🔍 Message not found")
	if u != nil && u.Message != nil {
		b.SendMessage(&models.SendMessageParams{
			ChatID: u.Message.Chat.ID,
			Text:   "❌ پیام یافت نشد",
		})
	}
	return nil
}

func handleMessageCantBeEdited(b *core.Bot, u *models.Update, err error) error {
	log.Printf("✏️ Can't edit message")
	if u != nil && u.Message != nil {
		b.SendMessage(&models.SendMessageParams{
			ChatID: u.Message.Chat.ID,
			Text:   "❌ پیام قابل ویرایش نیست",
		})
	}
	return nil
}

func handleMessageTextEmpty(b *core.Bot, u *models.Update, err error) error {
	log.Printf("📝 Empty text")
	if u != nil && u.Message != nil {
		b.SendMessage(&models.SendMessageParams{
			ChatID: u.Message.Chat.ID,
			Text:   "❌ متن پیام خالی است",
		})
	}
	return nil
}

func handleMessageTooLong(b *core.Bot, u *models.Update, err error) error {
	log.Printf("📏 Message too long")
	if u != nil && u.Message != nil {
		b.SendMessage(&models.SendMessageParams{
			ChatID: u.Message.Chat.ID,
			Text:   "❌ پیام خیلی طولانی است",
		})
	}
	return nil
}

func handleBotBlocked(b *core.Bot, u *models.Update, err error) error {
	if u != nil && u.Message != nil && u.Message.From != nil {
		log.Printf("🚫 User %d blocked bot", u.Message.From.ID)
		// می‌توانید در دیتابیس ذخیره کنید
	}
	return nil
}

func handleBotKicked(b *core.Bot, u *models.Update, err error) error {
	log.Printf("👋 Bot kicked from chat")
	// می‌توانید اطلاعات گروه را حذف کنید
	return nil
}

func handleChatNotFound(b *core.Bot, u *models.Update, err error) error {
	log.Printf("💬 Chat not found")
	return nil
}

func handleInvalidFileID(b *core.Bot, u *models.Update, err error) error {
	log.Printf("📎 Invalid file_id")
	if u != nil && u.Message != nil {
		b.SendMessage(&models.SendMessageParams{
			ChatID: u.Message.Chat.ID,
			Text:   "❌ فایل نامعتبر است",
		})
	}
	return nil
}

func handleButtonDataInvalid(b *core.Bot, u *models.Update, err error) error {
	log.Printf("🔘 Invalid button data")
	if u != nil && u.CallbackQuery != nil {
		b.AnswerCallbackQuery(u.CallbackQuery.ID, "❌ دکمه نامعتبر", true)
	}
	return nil
}

func handleRateLimit(b *core.Bot, u *models.Update, err error) error {
	log.Printf("⏰ Rate limited")
	// می‌توانید منطق retry اضافه کنید
	return nil
}

func handleBadRequest(b *core.Bot, u *models.Update, err error) error {
	if teleErr, ok := err.(*core.TelegramError); ok {
		log.Printf("❌ Bad request: %s", teleErr.Description)
	}
	return nil
}

func handleForbidden(b *core.Bot, u *models.Update, err error) error {
	if teleErr, ok := err.(*core.TelegramError); ok {
		log.Printf("🚫 Forbidden: %s", teleErr.Description)
	}
	return nil
}

func handleUnknownError(b *core.Bot, u *models.Update, err error) error {
	log.Printf("💥 Unknown error: %v", err)
	if u != nil && u.Message != nil {
		b.SendMessage(&models.SendMessageParams{
			ChatID: u.Message.Chat.ID,
			Text:   "⚠️ خطایی رخ داد",
		})
	}
	return nil
}
