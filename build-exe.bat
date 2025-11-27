@echo off
cd /d "%~dp0"
set PATH=%PATH%;C:\TDM-GCC-64\bin
set CGO_ENABLED=1

echo در حال کامپایل...
go build -buildvcs=false -ldflags="-s -w" -o notepad-app.exe

if exist notepad-app.exe (
    echo.
    echo ✓ موفق! فایل notepad-app.exe ساخته شد
    dir notepad-app.exe
) else (
    echo.
    echo ✗ خطا در ساخت فایل
    go build -buildvcs=false -ldflags="-s -w" -o notepad-app.exe 2>&1
)

pause
