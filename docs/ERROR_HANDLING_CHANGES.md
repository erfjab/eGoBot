# تغییرات Error Handling در eGoBot

## فایل‌های اضافه شده

### 1. `egobot/core/errors.go` ⭐
فایل اصلی error handling که شامل:
- **TelegramError struct**: برای نمایش خطاهای Telegram API
- **ResponseParameters struct**: شامل اطلاعات `retry_after` و `migrate_to_chat_id`
- **ErrorHandlerFunc**: type برای error handler functions
- **ErrorHandlers**: مدیریت و اجرای error handlers
- **Error Filters**: فیلترهای از پیش تعریف شده برای انواع خطاها
- **Helper Methods**: متدهای کمکی مثل `IsRateLimitError()`, `IsBadRequest()`, etc.

### 2. `egobot/core/errors_test.go` ✅
تست‌های کامل برای error handling شامل:
- تست TelegramError و متدهای آن
- تست ErrorHandlers و filters
- تست ترتیب اجرای handlers
- همه تست‌ها pass شده‌اند!

### 3. `examples/error_handling_example.go` 📝
مثال پایه برای استفاده از error handling:
- نحوه ثبت error handlers مختلف
- مدیریت rate limiting (429)
- مدیریت forbidden errors (403)
- مدیریت bad requests (400)
- استفاده از fallback handler

### 4. `examples/advanced_error_handling.go` 🚀
مثال پیشرفته برای یک ربات واقعی:
- مدیریت کامل تمام انواع خطاها
- شمارنده برای tracking errors
- پیام‌های فارسی برای کاربران
- دستورات تست و نمایش وضعیت

### 5. `docs/ERROR_HANDLING.md` 📚
داکیومنت کامل فارسی شامل:
- توضیح مفاهیم اصلی
- لیست کدهای خطای رایج
- مثال‌های استفاده
- Best practices
- نکات مهم

## تغییرات در فایل‌های موجود

### `egobot/core/bot.go` 🔧
- اضافه شدن فیلد `errorHandlers` به Bot struct
- متدهای جدید:
  - `OnError()`: ثبت error handler با filter سفارشی
  - `OnTelegramError()`: ثبت handler برای خطاهای تلگرام
  - `OnRateLimitError()`: ثبت handler برای rate limiting
  - `OnBadRequest()`: ثبت handler برای bad requests
  - `OnForbiddenError()`: ثبت handler برای forbidden errors
  - `SetFallbackErrorHandler()`: ثبت fallback handler

### `egobot/core/handler.go` 🔄
- تغییر متد `Process()` برای pass کردن خطاها به error handlers
- خطاها حالا به جای فقط log شدن، به error handlers منتقل می‌شوند

### `egobot/core/registrar.go` 📋
- اضافه شدن interface `ErrorHandlerRegistrar`
- توسعه امکانات registrar برای error handling

## ویژگی‌های اصلی

### ✨ Error Types
- **TelegramError**: خطاهای API تلگرام با error code و description
- **ResponseParameters**: شامل اطلاعات اضافی مثل retry_after

### 🎯 Error Filters
فیلترهای از پیش تعریف شده:
- `TelegramErrorFilter()`: همه خطاهای تلگرام
- `RateLimitErrorFilter()`: خطاهای 429
- `BadRequestErrorFilter()`: خطاهای 400
- `ForbiddenErrorFilter()`: خطاهای 403
- `UnauthorizedErrorFilter()`: خطاهای 401
- `ConflictErrorFilter()`: خطاهای 409
- `ServerErrorFilter()`: خطاهای 5xx
- `AllErrorsFilter()`: همه خطاها
- `ErrorCodeFilter(code)`: فیلتر سفارشی برای کد خاص

### 🔨 Helper Methods
- `IsRateLimitError()`: بررسی rate limit
- `IsBadRequest()`: بررسی bad request
- `IsForbidden()`: بررسی forbidden
- `IsUnauthorized()`: بررسی unauthorized
- `IsNotFound()`: بررسی not found
- `IsConflict()`: بررسی conflict
- `IsServerError()`: بررسی server error

## نحوه استفاده

```go
bot := core.NewBot(token)

// مدیریت rate limiting
bot.OnRateLimitError(func(b *core.Bot, u *models.Update, err error) error {
    if teleErr, ok := err.(*core.TelegramError); ok {
        time.Sleep(time.Duration(teleErr.Parameters.RetryAfter) * time.Second)
    }
    return nil
})

// مدیریت bad requests
bot.OnBadRequest(func(b *core.Bot, u *models.Update, err error) error {
    log.Printf("Bad request: %v", err)
    return nil
})

// Fallback برای همه خطاها
bot.SetFallbackErrorHandler(func(b *core.Bot, u *models.Update, err error) error {
    log.Printf("Error: %v", err)
    return nil
})
```

## کدهای خطای رایج

| کد  | نام | توضیح |
|-----|-----|-------|
| 400 | Bad Request | پارامترهای نامعتبر |
| 401 | Unauthorized | توکن اشتباه |
| 403 | Forbidden | دسترسی محدود یا کاربر ربات را بلاک کرده |
| 404 | Not Found | چت یا پیام یافت نشد |
| 409 | Conflict | instance دیگری از ربات اجراست |
| 429 | Too Many Requests | بیش از حد درخواست |
| 5xx | Server Error | خطای سرور تلگرام |

## مزایا

✅ مدیریت حرفه‌ای خطاها  
✅ جداسازی error handling از business logic  
✅ قابلیت customize کردن response به کاربر  
✅ Tracking و monitoring خطاها  
✅ Retry logic برای rate limiting  
✅ Type-safe با استفاده از Go  
✅ تست‌های کامل  
✅ داکیومنت فارسی کامل  

## تست‌ها

تمام تست‌ها با موفقیت pass شده‌اند:
```
✅ TestTelegramError
✅ TestTelegramErrorWithParams
✅ TestTelegramErrorMigrate
✅ TestErrorCodeChecks (9 sub-tests)
✅ TestIsTelegramError
✅ TestErrorHandlers
✅ TestErrorFilters (9 sub-tests)
✅ TestErrorHandlerOrder
```

## مستندات

برای اطلاعات بیشتر، [داکیومنت Error Handling](docs/ERROR_HANDLING.md) را مطالعه کنید.
