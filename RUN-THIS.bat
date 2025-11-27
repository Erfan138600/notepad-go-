@echo off
title ساخت فایل exe - لطفا صبر کنید...
chcp 65001 >nul
color 0A

echo.
echo ╔════════════════════════════════════════╗
echo ║     ساخت فایل exe - لطفا صبر کنید    ║
echo ╚════════════════════════════════════════╝
echo.

cd /d "%~dp0"

echo [1/4] تنظیم مسیرها...
set "PATH=%PATH%;C:\TDM-GCC-64\bin"
set "CGO_ENABLED=1"
timeout /t 1 /nobreak >nul

echo [2/4] بررسی GCC...
where gcc >nul 2>&1
if errorlevel 1 (
    color 0C
    echo.
    echo ❌ خطا: GCC پیدا نشد!
    echo.
    echo لطفا مطمئن شوید که TDM-GCC در مسیر زیر نصب شده:
    echo C:\TDM-GCC-64\bin
    echo.
    pause
    exit /b 1
)
echo    ✓ GCC پیدا شد
timeout /t 1 /nobreak >nul

echo [3/4] بررسی Go...
where go >nul 2>&1
if errorlevel 1 (
    color 0C
    echo.
    echo ❌ خطا: Go پیدا نشد!
    pause
    exit /b 1
)
echo    ✓ Go پیدا شد
timeout /t 1 /nobreak >nul

echo [4/4] کامپایل برنامه...
echo.
echo این مرحله ممکن است چند دقیقه طول بکشد...
echo لطفا صبر کنید...
echo.

go build -buildvcs=false -ldflags="-s -w" -o notepad-app.exe

if errorlevel 1 (
    color 0C
    echo.
    echo ❌ خطا در کامپایل!
    echo.
    echo لطفا خطاهای بالا را بررسی کنید.
    pause
    exit /b 1
)

if not exist notepad-app.exe (
    color 0C
    echo.
    echo ❌ فایل exe ساخته نشد!
    pause
    exit /b 1
)

color 0A
echo.
echo ╔════════════════════════════════════════╗
echo ║   ✓✓✓ موفق! فایل exe ساخته شد ✓✓✓   ║
echo ╚════════════════════════════════════════╝
echo.

for %%A in (notepad-app.exe) do (
    set size=%%~zA
    set /a sizeMB=%%~zA/1048576
)
echo فایل: notepad-app.exe
echo اندازه: %sizeMB% MB (~%size% بایت)
echo.
echo ✓ آماده برای استفاده!
echo.
echo برای اجرا: فایل notepad-app.exe را دوبار کلیک کنید
echo.

pause

