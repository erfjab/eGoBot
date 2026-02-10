package main

import (
	"egobot/egobot/core"
	"egobot/egobot/models"
	"fmt"
	"log"
	"os"
	"time"
)

// این مثال نشان می‌دهد چگونه error handling پیشرفته و ماژولار را در یک ربات واقعی پیاده‌سازی کنیم

var (
	// شمارنده برای تعداد rate limit errors
	rateLimitCount = 0
	
	// شمارنده برای تعداد کاربرانی که ربات را بلاک کرده‌اند
	blockedUsersCount = 0
)

func main() {
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatal("BOT_TOKEN environment variable is not set")
	}

	bot := core.NewBot(token)
	
	// تنظیم error handlers
	setupErrorHandlers(bot)
	
	// تنظیم command handlers
	setupCommandHandlers(bot)
	
	// شروع polling
	log.Println("🚀 Bot started with advanced error handling!")
	bot.StartPolling(&core.PollingOptions{
		Timeout: 30,
		Async:   true,
		OnStart: func() {
			log.Println("✅ Polling started successfully")
		},
		OnError: func(err error) {
			log.Printf("⚠️ Polling error: %v", err)
		},
	})
}

func setupErrorHandlers(bot *core.Bot) {
	// 1. مدیریت Rate Limit Errors (429)
	bot.OnRateLimitError(func(b *core.Bot, u *models.Update, err error) error {
		if teleErr, ok := err.(*core.TelegramError); ok {
			rateLimitCount++
			retryAfter := 3 // default
			
			if teleErr.Parameters != nil && teleErr.Parameters.RetryAfter > 0 {
				retryAfter = teleErr.Parameters.RetryAfter
			}
			
			log.Printf("⏰ [Rate Limit #%d] Waiting %d seconds", rateLimitCount, retryAfter)
			time.Sleep(time.Duration(retryAfter) * time.Second)
			
			if u != nil && u.Message != nil {
				b.SendMessage(&models.SendMessageParams{
					ChatID: u.Message.Chat.ID,
					Text: fmt.Sprintf("⏳ ربات مشغول است. %d ثانیه صبر کنید.", retryAfter),
				})
			}
		}
		return nil
	})

	// 2. مدیریت Bot Blocked - کاربر ربات را بلاک کرده
	bot.OnBotBlocked(func(b *core.Bot, u *models.Update, err error) error {
		blockedUsersCount++
		if u != nil && u.Message != nil && u.Message.From != nil {
			log.Printf("🚫 [Blocked #%d] User %d (@%s)", 
				blockedUsersCount, u.Message.From.ID, u.Message.From.Username)
			// db.MarkUserAsBlocked(u.Message.From.ID)
		}
		return nil
	})

	// 3. مدیریت Bot Kicked - ربات از گروه اخراج شده
	bot.OnBotKicked(func(b *core.Bot, u *models.Update, err error) error {
		log.Printf("👋 Bot was kicked from a chat")
		// db.RemoveChatFromDatabase()
		return nil
	})

	// 4. مدیریت Message Not Found - پیام یافت نشد
	bot.OnMessageNotFound(func(b *core.Bot, u *models.Update, err error) error {
		log.Printf("🔍 Message not found")
		if u != nil && u.Message != nil {
			b.SendMessage(&models.SendMessageParams{
				ChatID: u.Message.Chat.ID,
				Text:   "❌ پیام مورد نظر یافت نشد.",
			})
		}
		return nil
	})

	// 5. مدیریت Message Can't Be Edited - پیام قابل ویرایش نیست
	bot.OnMessageCantBeEdited(func(b *core.Bot, u *models.Update, err error) error {
		log.Printf("✏️ Message can't be edited")
		if u != nil && u.Message != nil {
			b.SendMessage(&models.SendMessageParams{
				ChatID: u.Message.Chat.ID,
				Text:   "❌ این پیام قابل ویرایش نیست.",
			})
		}
		return nil
	})

	// 6. مدیریت Message Can't Be Deleted - پیام قابل حذف نیست
	bot.OnMessageCantBeDeleted(func(b *core.Bot, u *models.Update, err error) error {
		log.Printf("🗑️ Message can't be deleted")
		if u != nil && u.Message != nil {
			b.SendMessage(&models.SendMessageParams{
				ChatID: u.Message.Chat.ID,
				Text:   "❌ این پیام قابل حذف نیست.",
			})
		}
		return nil
	})

	// 7. مدیریت Message Text Empty - متن پیام خالی است
	bot.OnMessageTextEmpty(func(b *core.Bot, u *models.Update, err error) error {
		log.Printf("📝 Empty message text")
		if u != nil && u.Message != nil {
			b.SendMessage(&models.SendMessageParams{
				ChatID: u.Message.Chat.ID,
				Text:   "❌ متن پیام نمی‌تواند خالی باشد.",
			})
		}
		return nil
	})

	// 8. مدیریت Message Too Long - پیام خیلی طولانی است
	bot.OnMessageTooLong(func(b *core.Bot, u *models.Update, err error) error {
		log.Printf("📏 Message too long")
		if u != nil && u.Message != nil {
			b.SendMessage(&models.SendMessageParams{
				ChatID: u.Message.Chat.ID,
				Text:   "❌ پیام شما خیلی طولانی است. لطفاً کوتاه‌تر کنید.",
			})
		}
		return nil
	})

	// 9. مدیریت Chat Not Found - چت یافت نشد
	bot.OnChatNotFound(func(b *core.Bot, u *models.Update, err error) error {
		log.Printf("💬 Chat not found")
		if u != nil && u.Message != nil {
			b.SendMessage(&models.SendMessageParams{
				ChatID: u.Message.Chat.ID,
				Text:   "❌ چت مورد نظر یافت نشد.",
			})
		}
		return nil
	})

	// 10. مدیریت Invalid File ID - شناسه فایل نامعتبر است
	bot.OnInvalidFileID(func(b *core.Bot, u *models.Update, err error) error {
		log.Printf("📎 Invalid file_id")
		if u != nil && u.Message != nil {
			b.SendMessage(&models.SendMessageParams{
				ChatID: u.Message.Chat.ID,
				Text:   "❌ فایل مورد نظر نامعتبر یا منقضی شده است.",
			})
		}
		return nil
	})

	// 11. مدیریت Button Data Invalid - داده دکمه نامعتبر است
	bot.OnButtonDataInvalid(func(b *core.Bot, u *models.Update, err error) error {
		log.Printf("🔘 Invalid button data")
		if u != nil && u.CallbackQuery != nil {
			b.AnswerCallbackQuery(u.CallbackQuery.ID, "❌ دکمه نامعتبر است", true)
		}
		return nil
	})

	// 12. مدیریت Bad Request (عمومی)
	bot.OnBadRequest(func(b *core.Bot, u *models.Update, err error) error {
		if teleErr, ok := err.(*core.TelegramError); ok {
			log.Printf("❌ [Bad Request] %s", teleErr.Description)
			if u != nil && u.Message != nil {
				b.SendMessage(&models.SendMessageParams{
					ChatID: u.Message.Chat.ID,
					Text:   "❌ درخواست نامعتبر است.",
				})
			}
		}
		return nil
	})

	// 13. مدیریت Forbidden Errors (عمومی)
	bot.OnForbiddenError(func(b *core.Bot, u *models.Update, err error) error {
		if teleErr, ok := err.(*core.TelegramError); ok {
			log.Printf("🚫 [Forbidden] %s", teleErr.Description)
		}
		return nil
	})

	// 14. مدیریت Conflict Errors
	bot.OnError(core.ConflictErrorFilter(), func(b *core.Bot, u *models.Update, err error) error {
		log.Printf("⚡ [Conflict] Another bot instance might be running!")
		return nil
	})

	// 15. مدیریت Server Errors
	bot.OnError(core.ServerErrorFilter(), func(b *core.Bot, u *models.Update, err error) error {
		if teleErr, ok := err.(*core.TelegramError); ok {
			log.Printf("🔴 [Server Error %d] %s", teleErr.ErrorCode, teleErr.Description)
			if u != nil && u.Message != nil {
				b.SendMessage(&models.SendMessageParams{
					ChatID: u.Message.Chat.ID,
					Text:   "⚠️ سرور تلگرام مشکل دارد. بعداً تلاش کنید.",
				})
			}
		}
		return nil
	})

	// 16. Fallback handler برای خطاهای مدیریت نشده
	bot.SetFallbackErrorHandler(func(b *core.Bot, u *models.Update, err error) error {
		if teleErr, ok := err.(*core.TelegramError); ok {
			log.Printf("💥 [Unhandled] [%d] %s", teleErr.ErrorCode, teleErr.Description)
		} else {
			log.Printf("💥 [Unhandled] %v", err)
		}
		
		if u != nil && u.Message != nil {
			b.SendMessage(&models.SendMessageParams{
				ChatID: u.Message.Chat.ID,
				Text:   "⚠️ خطایی رخ داده است. بعداً تلاش کنید.",
			})
		}
		return nil
	})
	
	log.Println("✅ Error handlers configured (16 handlers)")
}

func setupCommandHandlers(bot *core.Bot) {
	// دستور /start
	bot.OnCommand("start", func(b *core.Bot, u *models.Update) error {
		_, err := b.SendMessage(&models.SendMessageParams{
			ChatID: u.Message.Chat.ID,
			Text: "👋 سلام!\n\n" +
				"به ربات نمونه error handling خوش آمدید.\n\n" +
				"دستورات موجود:\n" +
				"/start - شروع\n" +
				"/status - نمایش وضعیت\n" +
				"/test_empty - تست empty text\n" +
				"/test_edit - تست edit error",
		})
		return err
	})

	// دستور /status
	bot.OnCommand("status", func(b *core.Bot, u *models.Update) error {
		statusText := fmt.Sprintf(
			"📊 *وضعیت ربات*\n\n"+
				"▫️ Rate Limits: `%d`\n"+
				"▫️ Blocked Users: `%d`\n",
			rateLimitCount,
			blockedUsersCount,
		)
		
		_, err := b.SendMessage(&models.SendMessageParams{
			ChatID:    u.Message.Chat.ID,
			Text:      statusText,
			ParseMode: "Markdown",
		})
		return err
	})

	// دستور /test_empty - تست empty text error
	bot.OnCommand("test_empty", func(b *core.Bot, u *models.Update) error {
		_, err := b.SendMessage(&models.SendMessageParams{
			ChatID: u.Message.Chat.ID,
			Text:   "", // خالی - باعث error می‌شود
		})
		return err // خطا به error handler منتقل می‌شود
	})

	// دستور /test_edit - تست edit error
	bot.OnCommand("test_edit", func(b *core.Bot, u *models.Update) error {
		// تلاش برای ویرایش پیامی که وجود ندارد
		_, err := b.EditMessageText(&models.EditMessageTextParams{
			ChatID:    u.Message.Chat.ID,
			MessageID: 99999, // پیام وجود ندارد
			Text:      "Test",
		})
		return err // خطا به error handler منتقل می‌شود
	})
	
	log.Println("✅ Command handlers configured")
}
