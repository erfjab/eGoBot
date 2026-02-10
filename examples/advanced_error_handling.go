package main

import (
	"egobot/egobot/core"
	"egobot/egobot/models"
	"fmt"
	"log"
	"os"
	"time"
)

// این مثال نشان می‌دهد چگونه error handling پیشرفته را در یک ربات واقعی پیاده‌سازی کنیم

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
	// این خطا زمانی رخ می‌دهد که تعداد درخواست‌ها بیش از حد مجاز باشد
	bot.OnRateLimitError(func(b *core.Bot, u *models.Update, err error) error {
		if teleErr, ok := err.(*core.TelegramError); ok {
			rateLimitCount++
			retryAfter := 3 // default
			
			if teleErr.Parameters != nil && teleErr.Parameters.RetryAfter > 0 {
				retryAfter = teleErr.Parameters.RetryAfter
			}
			
			log.Printf("⏰ [Rate Limit #%d] Waiting %d seconds before retry", 
				rateLimitCount, retryAfter)
			
			// صبر کردن قبل از retry
			time.Sleep(time.Duration(retryAfter) * time.Second)
			
			// اطلاع به کاربر
			if u != nil && u.Message != nil {
				b.SendMessage(&models.SendMessageParams{
					ChatID: u.Message.Chat.ID,
					Text: fmt.Sprintf(
						"⏳ ربات در حال حاضر بسیار شلوغ است.\n" +
						"لطفاً %d ثانیه صبر کنید و دوباره تلاش کنید.",
						retryAfter,
					),
				})
			}
		}
		return nil
	})

	// 2. مدیریت Forbidden Errors (403)
	// این خطا معمولاً زمانی رخ می‌دهد که کاربر ربات را بلاک کرده باشد
	bot.OnForbiddenError(func(b *core.Bot, u *models.Update, err error) error {
		if teleErr, ok := err.(*core.TelegramError); ok {
			blockedUsersCount++
			
			if u != nil && u.Message != nil && u.Message.From != nil {
				log.Printf("🚫 [User Blocked #%d] User %d (@%s) has blocked the bot or restricted access",
					blockedUsersCount,
					u.Message.From.ID,
					u.Message.From.Username,
				)
				
				// در اینجا می‌توانید کاربر را در دیتابیس به عنوان blocked علامت‌گذاری کنید
				// db.MarkUserAsBlocked(u.Message.From.ID)
			} else {
				log.Printf("🚫 [Forbidden] %s", teleErr.Description)
			}
		}
		return nil
	})

	// 3. مدیریت Bad Request Errors (400)
	// این خطا زمانی رخ می‌دهد که پارامترهای درخواست نامعتبر باشند
	bot.OnBadRequest(func(b *core.Bot, u *models.Update, err error) error {
		if teleErr, ok := err.(*core.TelegramError); ok {
			log.Printf("❌ [Bad Request] %s", teleErr.Description)
			
			// بررسی خطاهای خاص
			if u != nil && u.Message != nil {
				var userMsg string
				
				// پیام‌های مختلف بر اساس نوع خطا
				switch {
				case containsString(teleErr.Description, "message text is empty"):
					userMsg = "❌ متن پیام نمی‌تواند خالی باشد."
					
				case containsString(teleErr.Description, "message is too long"):
					userMsg = "❌ پیام شما خیلی طولانی است. لطفاً کوتاه‌تر کنید."
					
				case containsString(teleErr.Description, "chat not found"):
					userMsg = "❌ چت مورد نظر یافت نشد."
					
				case containsString(teleErr.Description, "message to delete not found"):
					userMsg = "❌ پیام مورد نظر برای حذف یافت نشد."
					
				case containsString(teleErr.Description, "message can't be edited"):
					userMsg = "❌ این پیام قابل ویرایش نیست."
					
				case containsString(teleErr.Description, "message to edit not found"):
					userMsg = "❌ پیام مورد نظر برای ویرایش یافت نشد."
					
				default:
					userMsg = "❌ درخواست نامعتبر است. لطفاً دوباره تلاش کنید."
				}
				
				b.SendMessage(&models.SendMessageParams{
					ChatID: u.Message.Chat.ID,
					Text:   userMsg,
				})
			}
		}
		return nil
	})

	// 4. مدیریت Not Found Errors (404)
	bot.OnError(core.ErrorCodeFilter(core.ErrorCodeNotFound), func(b *core.Bot, u *models.Update, err error) error {
		if teleErr, ok := err.(*core.TelegramError); ok {
			log.Printf("🔍 [Not Found] %s", teleErr.Description)
			
			if u != nil && u.Message != nil {
				b.SendMessage(&models.SendMessageParams{
					ChatID: u.Message.Chat.ID,
					Text:   "❌ مورد درخواستی یافت نشد.",
				})
			}
		}
		return nil
	})

	// 5. مدیریت Conflict Errors (409)
	// این خطا معمولاً زمانی رخ می‌دهد که دو instance از ربات با یک token اجرا شوند
	bot.OnError(core.ConflictErrorFilter(), func(b *core.Bot, u *models.Update, err error) error {
		if teleErr, ok := err.(*core.TelegramError); ok {
			log.Printf("⚡ [Conflict] %s", teleErr.Description)
			log.Println("⚠️ WARNING: Another instance of the bot might be running!")
		}
		return nil
	})

	// 6. مدیریت Server Errors (5xx)
	// این خطاها مربوط به سرور تلگرام هستند
	bot.OnError(core.ServerErrorFilter(), func(b *core.Bot, u *models.Update, err error) error {
		if teleErr, ok := err.(*core.TelegramError); ok {
			log.Printf("🔴 [Server Error %d] %s", teleErr.ErrorCode, teleErr.Description)
			
			if u != nil && u.Message != nil {
				b.SendMessage(&models.SendMessageParams{
					ChatID: u.Message.Chat.ID,
					Text: "⚠️ سرور تلگرام در حال حاضر مشکل دارد.\n" +
						  "لطفاً چند لحظه دیگر دوباره تلاش کنید.",
				})
			}
		}
		return nil
	})

	// 7. Fallback handler برای همه خطاهای دیگر
	bot.SetFallbackErrorHandler(func(b *core.Bot, u *models.Update, err error) error {
		// بررسی نوع خطا
		if teleErr, ok := err.(*core.TelegramError); ok {
			log.Printf("💥 [Unhandled Telegram Error %d] %s", teleErr.ErrorCode, teleErr.Description)
		} else {
			log.Printf("💥 [Unhandled Error] %v", err)
		}
		
		// اطلاع به کاربر
		if u != nil && u.Message != nil {
			b.SendMessage(&models.SendMessageParams{
				ChatID: u.Message.Chat.ID,
				Text: "⚠️ خطایی رخ داده است.\n" +
					  "لطفاً بعداً دوباره تلاش کنید.",
			})
		}
		
		return nil
	})
	
	log.Println("✅ Error handlers configured successfully")
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
				"/test_error - تست error handling",
		})
		return err
	})

	// دستور /status
	bot.OnCommand("status", func(b *core.Bot, u *models.Update) error {
		statusText := fmt.Sprintf(
			"📊 *وضعیت ربات*\n\n"+
				"▫️ Rate Limit Errors: `%d`\n"+
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

	// دستور /test_error - برای تست error handling
	bot.OnCommand("test_error", func(b *core.Bot, u *models.Update) error {
		// ارسال پیام خالی برای تست bad request error
		_, err := b.SendMessage(&models.SendMessageParams{
			ChatID: u.Message.Chat.ID,
			Text:   "", // متن خالی باعث bad request error می‌شود
		})
		
		// خطا به error handler منتقل می‌شود
		return err
	})
	
	log.Println("✅ Command handlers configured successfully")
}

// تابع کمکی برای بررسی وجود substring
func containsString(str, substr string) bool {
	for i := 0; i <= len(str)-len(substr); i++ {
		if str[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
