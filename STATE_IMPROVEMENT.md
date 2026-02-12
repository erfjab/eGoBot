# بهبود State Management - بارگذاری خودکار State و Data

## تغییرات اعمال شده

### قبل از این تغییرات:
```go
func handleMessage(bot *core.Bot, update *models.Update, ctx *core.Context) error {
    userID := update.Message.From.ID
    userManager := bot.StateManager.ForUser(userID)
    
    // نیاز به لود جداگانه state
    currentState, err := userManager.GetState(context.Background())
    if err != nil {
        return err
    }
    
    // نیاز به لود جداگانه data
    stateData, err := userManager.GetData(context.Background())
    if err != nil {
        return err
    }
    
    // استفاده از state و data
    name := stateData["name"]
}
```

### بعد از این تغییرات:
```go
func handleMessage(bot *core.Bot, update *models.Update, ctx *core.Context) error {
    // state و data به صورت خودکار در ctx موجود است!
    
    // دریافت state name
    stateName := ctx.GetStateName()
    
    // دریافت data
    stateData := ctx.GetStateData()
    if stateData != nil {
        name := stateData["name"]
    }
    
    // کار با data...
}
```

## چگونه کار میکند؟

وقتی یک handler با state filter تعریف میکنید:

```go
bot.AddHandler(
    core.TextFilter(),
    handleNameInput,
    state.InState(WaitingForName), // state filter
)
```

سیستم به صورت خودکار:

1. ✅ userID را شناسایی میکند
2. ✅ با یک بار فراخوانی `GetContext()`، هم state و هم data را لود میکند
3. ✅ state و data را در Context قرار میدهد
4. ✅ چک میکند که state با filter مطابقت دارد یا نه
5. ✅ اگر مطابقت داشت، handler را با Context پر شده اجرا میکند

## API های جدید Context

### `ctx.GetStateName()`
state name فعلی کاربر را برمیگرداند:

```go
stateName := ctx.GetStateName()
// مثال: "user.waiting_for_name"
```

### `ctx.GetStateData()`
data ذخیره شده کاربر را برمیگرداند:

```go
stateData := ctx.GetStateData()
if stateData != nil {
    name := stateData["name"].(string)
    age := stateData["age"].(string)
}
```

### دسترسی به State Object
اگر نیاز به object کامل State داشتید:

```go
if stateObj := ctx.Get("state_obj"); stateObj != nil {
    currentState := stateObj.(*state.State)
    // استفاده از currentState
}
```

## مزایا

### 1. Performance بهتر ⚡
- قبلاً: **2 query** به storage (یکی برای state، یکی برای data)
- الان: **1 query** به storage (با GetContext هر دو با هم)

### 2. کد تمیزتر 🧹
- نیازی به کد تکراری برای لود state و data نیست
- handler ها ساده‌تر و خواناتر میشن

### 3. کمتر احتمال خطا 🛡️
- مدیریت خطا توسط سیستم انجام میشه
- احتمال فراموشی لود data یا state صفر میشه

### 4. سازگاری با Middleware 🔄
- Middleware ها هم به state و data دسترسی دارن
- میشه در middleware ها از این اطلاعات استفاده کرد

## مثال کامل

برای مثال کامل، فایل [`examples/state_data_example.go`](examples/state_data_example.go) را مشاهده کنید.

## یادداشت‌های مهم

⚠️ **این تغییر فقط برای handlerهایی که state filter دارند اعمال میشه.**

اگر handler بدون state filter تعریف کنید:
```go
bot.AddHandler(core.TextFilter(), handleMessage)
```

در این حالت، state و data به صورت خودکار لود **نمیشه** و Context خالی خواهد بود.

---

## تغییرات فنی

### فایل‌های تغییر یافته:

1. **[core/handler.go](core/handler.go)**
   - استفاده از `GetContext` به جای `GetState`
   - افزودن state و data به Context
   - افزودن import برای `storage`

2. **[core/middleware.go](core/middleware.go)**
   - افزودن تابع `NewMiddlewareChainWithContext`

3. **[core/context.go](core/context.go)**
   - افزودن متد `GetStateData()`
   - افزودن متد `GetStateName()`

4. **[examples/state_data_example.go](examples/state_data_example.go)** (جدید)
   - مثال کامل استفاده از state management بهبود یافته
