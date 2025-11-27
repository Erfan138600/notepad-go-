@echo off
echo در حال بررسی نصب GCC...
where gcc >nul 2>nul
if %ERRORLEVEL% NEQ 0 (
    echo.
    echo ⚠ خطا: GCC پیدا نشد!
    echo.
    echo برای ساخت برنامه، باید GCC را نصب کنید.
    echo.
    echo لطفا install-gcc.bat را اجرا کنید یا به صورت دستی:
    echo.
    echo روش 1: نصب MinGW-w64
    echo   1. از https://sourceforge.net/projects/mingw-w64/ دانلود کنید
    echo   2. فایل exe را اجرا و نصب کنید
    echo   3. مسیر bin (مثلا C:\mingw64\bin) را به PATH اضافه کنید
    echo   4. PowerShell را restart کنید
    echo.
    echo روش 2: نصب TDM-GCC
    echo   1. از https://jmeubank.github.io/tdm-gcc/ دانلود کنید
    echo   2. نصب کنید
    echo   3. PowerShell را restart کنید
    echo.
    echo بعد از نصب GCC، این اسکریپت را دوباره اجرا کنید.
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
    echo.
    if exist notepad-app.exe (
        for %%A in (notepad-app.exe) do echo ✓ اندازه فایل: %%~zA بایت
    )
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

