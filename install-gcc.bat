@echo off
echo در حال بررسی Chocolatey...
where choco >nul 2>nul
if %ERRORLEVEL% EQU 0 (
    echo Chocolatey پیدا شد. در حال نصب MinGW...
    choco install mingw -y
    echo.
    echo در حال رفرش PATH...
    refreshenv
    echo.
    echo GCC نصب شد! حالا می‌توانید build.bat را اجرا کنید.
) else (
    echo Chocolatey پیدا نشد.
    echo.
    echo لطفا یکی از روش‌های زیر را انتخاب کنید:
    echo.
    echo روش 1: نصب Chocolatey و MinGW
    echo   1. از https://chocolatey.org نصب کنید
    echo   2. سپس این اسکریپت را دوباره اجرا کنید
    echo.
    echo روش 2: دانلود دستی MinGW
    echo   1. از https://sourceforge.net/projects/mingw-w64/ دانلود کنید
    echo   2. نصب کنید و مسیر bin را به PATH اضافه کنید
    echo.
    echo روش 3: استفاده از TDM-GCC
    echo   1. از https://jmeubank.github.io/tdm-gcc/ دانلود کنید
    echo   2. نصب کنید و مسیر bin را به PATH اضافه کنید
    echo.
)

pause


