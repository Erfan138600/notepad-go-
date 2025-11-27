@echo off
echo در حال بررسی نصب GCC...
where gcc >nul 2>nul
if %ERRORLEVEL% NEQ 0 (
    echo.
    echo ⚠ خطا: GCC پیدا نشد!
    echo.
    echo لطفا ابتدا GCC را نصب کنید:
    echo   1. install-gcc.bat را اجرا کنید
    echo   2. یا به صورت دستی MinGW/TDM-GCC را نصب کنید
    echo   3. سپس دوباره این اسکریپت را اجرا کنید
    echo.
    pause
    exit /b 1
)

echo ✓ GCC پیدا شد
echo.

echo در حال دانلود وابستگی‌ها...
go mod download
if %ERRORLEVEL% NEQ 0 (
    echo ✗ خطا در دانلود وابستگی‌ها
    pause
    exit /b 1
)

echo.
echo در حال کامپایل برنامه...
set CGO_ENABLED=1
go build -buildvcs=false -ldflags="-s -w" -o notepad-app.exe

if %ERRORLEVEL% EQU 0 (
    echo.
    echo ✓✓✓ برنامه با موفقیت ساخته شد! ✓✓✓
    echo.
    echo ✓ فایل خروجی: notepad-app.exe
    echo ✓ اندازه فایل: 
    for %%A in (notepad-app.exe) do echo    %%~zA بایت
    echo.
    echo برای اجرا، فایل notepad-app.exe را دوبار کلیک کنید.
    echo.
) else (
    echo.
    echo ✗ خطا در ساخت برنامه!
    echo.
    echo لطفا مطمئن شوید که:
    echo   - Go به درستی نصب شده است
    echo   - GCC به درستی نصب شده و در PATH است
    echo.
)

pause


