@echo off
chcp 65001 >nul
echo ========================================
echo ساخت فایل exe
echo ========================================
echo.

cd /d "%~dp0"
set "PATH=%PATH%;C:\TDM-GCC-64\bin"

echo بررسی GCC...
where gcc >nul 2>&1
if errorlevel 1 (
    echo ❌ GCC پیدا نشد!
    echo لطفا مطمئن شوید TDM-GCC نصب شده است
    pause
    exit /b 1
)
echo ✓ GCC پیدا شد
echo.

echo تنظیم CGO...
set CGO_ENABLED=1
echo ✓ CGO فعال شد
echo.

echo شروع کامپایل...
echo.
go build -buildvcs=false -ldflags="-s -w" -o notepad-app.exe
if errorlevel 1 (
    echo.
    echo ❌ خطا در کامپایل!
    echo لطفا خطاهای بالا را بررسی کنید.
    echo.
    pause
    exit /b 1
)
echo ✓ کامپایل با موفقیت انجام شد!
echo.
echo کد خروجی: %ERRORLEVEL%

if exist notepad-app.exe (
    echo.
    echo ========================================
    echo ✓✓✓ موفق! فایل exe ساخته شد ✓✓✓
    echo ========================================
    echo.
    echo فایل: notepad-app.exe
    for %%A in (notepad-app.exe) do echo اندازه: %%~zA بایت
    echo.
    echo برای اجرا، فایل را دوبار کلیک کنید!
    echo.
) else (
    echo.
    echo ❌ خطا در ساخت فایل exe
    echo.
    echo لطفا مطمئن شوید:
    echo   - Go نصب شده است
    echo   - GCC نصب شده است  
    echo   - همه وابستگی‌ها دانلود شده است
    echo.
)

pause

