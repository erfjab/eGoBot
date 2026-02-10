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

### مدیریت Forbidden Errors (403)

```go
bot.OnForbiddenError(func(b *core.Bot, u *models.Update, err error) error {
    if teleErr, ok := err.(*core.TelegramError); ok {
        log.Printf("Bot was blocked or doesn't have permission: %s", teleErr.Description)
        
        // ذخیره در دیتابیس که کاربر ربات را بلاک کرده
        if u.Message != nil {
            // db.MarkUserAsBlocked(u.Message.From.ID)
        }
    }
    return nil
})
```

### مدیریت Bad Request Errors (400)

```go
bot.OnBadRequest(func(b *core.Bot, u *models.Update, err error) error {
    if teleErr, ok := err.(*core.TelegramError); ok {
        log.Printf("Bad request: %s", teleErr.Description)
        
        // اطلاع به کاربر
        if u.Message != nil {
            b.SendMessage(&models.SendMessageParams{
                ChatID: u.Message.Chat.ID,
                Text:   "درخواست نامعتبر است. لطفاً دوباره تلاش کنید.",
            })
        }
    }
    return nil
})
```

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
