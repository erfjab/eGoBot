# Error Handling در eGoBot

این داکیومنت نحوه استفاده از سیستم error handling در eGoBot را توضیح می‌دهد.

## مفاهیم اصلی

### TelegramError

کلاس `TelegramError` برای نمایش خطاهای API تلگرام طراحی شده است:

```go
type TelegramError struct {
    ErrorCode   int
    Description string
    Parameters  *ResponseParameters
    Update      *models.Update
}
```

### Error Handlers

Error handlers توابعی هستند که وقتی خطا رخ می‌دهد، اجرا می‌شوند:

```go
type ErrorHandlerFunc func(*Bot, *models.Update, error) error
```

## کدهای خطای رایج

| کد | توضیح | مثال |
|-----|--------|--------|
| 400 | Bad Request | پارامترهای نامعتبر |
| 401 | Unauthorized | توکن اشتباه |
| 403 | Forbidden | ربات بلاک شده یا دسترسی ندارد |
| 404 | Not Found | چت یا پیام یافت نشد |
| 409 | Conflict | درخواست مخالف |
| 429 | Too Many Requests | محدودیت نرخ |
| 5xx | Server Error | خطای سرور تلگرام |

## استفاده پایه

### ثبت Error Handler برای همه خطاها

```go
bot.SetFallbackErrorHandler(func(b *core.Bot, u *models.Update, err error) error {
    log.Printf("Error occurred: %v", err)
    return nil
})
```

### ثبت Error Handler برای خطاهای تلگرام

```go
bot.OnTelegramError(func(b *core.Bot, u *models.Update, err error) error {
    if teleErr, ok := err.(*core.TelegramError); ok {
        log.Printf("Telegram Error [%d]: %s", teleErr.ErrorCode, teleErr.Description)
    }
    return nil
})
```

## Error Handlers اختصاصی

eGoBot از error handlers ماژولار برای انواع مختلف خطاهای تلگرام پشتیبانی می‌کند.

### خطاهای مربوط به پیام (Message Errors)

```go
// پیام یافت نشد
bot.OnMessageNotFound(func(b *core.Bot, u *models.Update, err error) error {
    log.Println("پیام یافت نشد")
    return nil
})

// پیام قابل ویرایش نیست
bot.OnMessageCantBeEdited(func(b *core.Bot, u *models.Update, err error) error {
    log.Println("پیام قابل ویرایش نیست")
    return nil
})

// پیام قابل حذف نیست
bot.OnMessageCantBeDeleted(func(b *core.Bot, u *models.Update, err error) error {
    log.Println("پیام قابل حذف نیست")
    return nil
})

// متن پیام خالی است
bot.OnMessageTextEmpty(func(b *core.Bot, u *models.Update, err error) error {
    b.SendMessage(&models.SendMessageParams{
        ChatID: u.Message.Chat.ID,
        Text:   "❌ متن پیام نمی‌تواند خالی باشد",
    })
    return nil
})

// پیام خیلی طولانی است
bot.OnMessageTooLong(func(b *core.Bot, u *models.Update, err error) error {
    b.SendMessage(&models.SendMessageParams{
        ChatID: u.Message.Chat.ID,
        Text:   "❌ پیام شما خیلی طولانی است",
    })
    return nil
})
```

### خطاهای مربوط به کاربر و چت (User/Chat Errors)

```go
// کاربر ربات را بلاک کرده
bot.OnBotBlocked(func(b *core.Bot, u *models.Update, err error) error {
    if u.Message != nil && u.Message.From != nil {
        log.Printf("کاربر %d ربات را بلاک کرد", u.Message.From.ID)
        // db.MarkUserAsBlocked(u.Message.From.ID)
    }
    return nil
})

// ربات از گروه اخراج شده
bot.OnBotKicked(func(b *core.Bot, u *models.Update, err error) error {
    log.Println("ربات از گروه اخراج شد")
    // db.RemoveChatFromDatabase()
    return nil
})

// چت یافت نشد
bot.OnChatNotFound(func(b *core.Bot, u *models.Update, err error) error {
    log.Println("چت یافت نشد")
    return nil
})
```

### خطاهای مربوط به رسانه (Media Errors)

```go
// شناسه فایل نامعتبر است
bot.OnInvalidFileID(func(b *core.Bot, u *models.Update, err error) error {
    b.SendMessage(&models.SendMessageParams{
        ChatID: u.Message.Chat.ID,
        Text:   "❌ فایل نامعتبر یا منقضی شده است",
    })
    return nil
})

// داده دکمه نامعتبر است
bot.OnButtonDataInvalid(func(b *core.Bot, u *models.Update, err error) error {
    if u.CallbackQuery != nil {
        b.AnswerCallbackQuery(u.CallbackQuery.ID, "❌ دکمه نامعتبر", true)
    }
    return nil
})
```

### مدیریت Rate Limiting (429)

```go
bot.OnRateLimitError(func(b *core.Bot, u *models.Update, err error) error {
    if teleErr, ok := err.(*core.TelegramError); ok {
        retryAfter := teleErr.Parameters.RetryAfter
        log.Printf("Rate limited! Retry after %d seconds", retryAfter)
        
        // صبر کردن و دوباره تلاش
        time.Sleep(time.Duration(retryAfter) * time.Second)
    }
    return nil
})
```

### خطاهای عمومی

```go
// Bad Request (400)
bot.OnBadRequest(func(b *core.Bot, u *models.Update, err error) error {
    log.Printf("Bad request: %v", err)
    return nil
})

// Forbidden (403)
bot.OnForbiddenError(func(b *core.Bot, u *models.Update, err error) error {
    log.Printf("Forbidden: %v", err)
    return nil
})

// Server Errors (5xx)
bot.OnError(core.ServerErrorFilter(), func(b *core.Bot, u *models.Update, err error) error {
    log.Printf("Server error: %v", err)
    return nil
})
```

## لیست کامل Error Handlers

| متد | توضیح | نمونه استفاده |
|-----|--------|---------------|
| `OnMessageNotFound` | پیام یافت نشد | ویرایش/حذف پیام حذف شده |
| `OnMessageCantBeEdited` | پیام قابل ویرایش نیست | ویرایش پیام قدیمی |
| `OnMessageCantBeDeleted` | پیام قابل حذف نیست | حذف پیام محافظت شده |
| `OnMessageTextEmpty` | متن پیام خالی است | ارسال پیام بدون متن |
| `OnMessageTooLong` | پیام خیلی طولانی است | پیام بیش از 4096 کاراکتر |
| `OnBotBlocked` | کاربر ربات را بلاک کرده | ارسال پیام به کاربر بلاک کننده |
| `OnBotKicked` | ربات از گروه اخراج شده | ارسال پیام به گروهی که ربات عضو نیست |
| `OnChatNotFound` | چت یافت نشد | ارسال به chat_id نامعتبر |
| `OnInvalidFileID` | فایل نامعتبر است | استفاده از file_id منقضی شده |
| `OnButtonDataInvalid` | داده دکمه نامعتبر است | callback_data بیش از 64 بایت |
| `OnRateLimitError` | محدودیت نرخ | بیش از حد درخواست |
| `OnBadRequest` | درخواست نامعتبر | پارامترهای اشتباه |
| `OnForbiddenError` | دسترسی ممنوع | عدم دسترسی |
| `OnTelegramError` | خطای تلگرام | همه خطاهای API |
| `SetFallbackErrorHandler` | خطاهای مدیریت نشده | Fallback |

## Error Filters سفارشی

می‌توانید فیلترهای سفارشی برای خطاهای خاص ایجاد کنید:

```go
// مدیریت خطاهای خاص بر اساس کد
bot.OnError(func(err error) bool {
    if teleErr, ok := err.(*core.TelegramError); ok {
        return teleErr.ErrorCode == 409 // Conflict
    }
    return false
}, func(b *core.Bot, u *models.Update, err error) error {
    log.Printf("Conflict error: %v", err)
    return nil
})

// مدیریت خطاهای سرور (5xx)
bot.OnError(core.ServerErrorFilter(), func(b *core.Bot, u *models.Update, err error) error {
    log.Printf("Server error occurred: %v", err)
    return nil
})
```

## Error Filters داخلی

eGoBot چندین فیلتر از پیش تعریف شده دارد:

- `TelegramErrorFilter()` - همه خطاهای تلگرام
- `RateLimitErrorFilter()` - خطاهای 429
- `BadRequestErrorFilter()` - خطاهای 400
- `ForbiddenErrorFilter()` - خطاهای 403
- `UnauthorizedErrorFilter()` - خطاهای 401
- `ConflictErrorFilter()` - خطاهای 409
- `ServerErrorFilter()` - خطاهای 5xx
- `AllErrorsFilter()` - همه خطاها

## متدهای کمکی TelegramError

```go
if teleErr, ok := err.(*core.TelegramError); ok {
    // بررسی نوع خطا
    if teleErr.IsRateLimitError() {
        // مدیریت rate limit
    }
    
    if teleErr.IsBadRequest() {
        // مدیریت bad request
    }
    
    if teleErr.IsForbidden() {
        // مدیریت forbidden
    }
    
    if teleErr.IsServerError() {
        // مدیریت server error
    }
}
```

## مثال کامل

```go
func setupErrorHandlers(bot *core.Bot) {
    // Rate limiting
    bot.OnRateLimitError(func(b *core.Bot, u *models.Update, err error) error {
        if teleErr, ok := err.(*core.TelegramError); ok {
            log.Printf("⚠️ Rate limit: retry after %d seconds", teleErr.Parameters.RetryAfter)
            time.Sleep(time.Duration(teleErr.Parameters.RetryAfter) * time.Second)
        }
        return nil
    })
    
    // Forbidden (user blocked bot)
    bot.OnForbiddenError(func(b *core.Bot, u *models.Update, err error) error {
        if u.Message != nil {
            log.Printf("User %d blocked the bot", u.Message.From.ID)
        }
        return nil
    })
    
    // Bad requests
    bot.OnBadRequest(func(b *core.Bot, u *models.Update, err error) error {
        if teleErr, ok := err.(*core.TelegramError); ok {
            log.Printf("❌ Bad request: %s", teleErr.Description)
        }
        return nil
    })
    
    // Generic Telegram errors
    bot.OnTelegramError(func(b *core.Bot, u *models.Update, err error) error {
        if teleErr, ok := err.(*core.TelegramError); ok {
            log.Printf("🔴 Telegram Error [%d]: %s", teleErr.ErrorCode, teleErr.Description)
        }
        return nil
    })
    
    // Fallback for all other errors
    bot.SetFallbackErrorHandler(func(b *core.Bot, u *models.Update, err error) error {
        log.Printf("💥 Unhandled error: %v", err)
        return nil
    })
}

func main() {
    bot := core.NewBot(os.Getenv("BOT_TOKEN"))
    
    setupErrorHandlers(bot)
    
    bot.OnCommand("start", func(b *core.Bot, u *models.Update) error {
        _, err := b.SendMessage(&models.SendMessageParams{
            ChatID: u.Message.Chat.ID,
            Text:   "Hello!",
        })
        return err // خطا به error handler منتقل می‌شود
    })
    
    bot.StartPolling(nil)
}
```

## نکات مهم

1. **ترتیب اجرا**: Error handlers به ترتیب ثبت اجرا می‌شوند. اولین handler که match شود، اجرا می‌شود.

2. **Fallback Handler**: اگر هیچ error handler خاصی match نشود، fallback handler اجرا می‌شود.

3. **Return nil**: اگر می‌خواهید خطا را مدیریت کنید، `nil` برگردانید. اگر `error` برگردانید، خطا log می‌شود.

4. **Error Propagation**: خطاهایی که از handlers برمی‌گردند، به error handlers بعدی pass نمی‌شوند.

5. **Thread Safety**: Error handlers در goroutine های مختلف اجرا می‌شوند (اگر `Async: true`)، پس باید thread-safe باشند.

## Best Practices

1. همیشه rate limiting را مدیریت کنید
2. خطاهای forbidden را log کنید تا بفهمید کاربران چه زمانی ربات را block می‌کنند
3. از fallback handler برای خطاهای غیرمنتظره استفاده کنید
4. خطاها را در production log کنید اما جزئیات را به کاربر نشان ندهید
5. برای خطاهای سرور، retry logic پیاده‌سازی کنید
