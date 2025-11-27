$ErrorActionPreference = "Continue"
$OutputEncoding = [System.Text.Encoding]::UTF8

Write-Host "========================================" -ForegroundColor Cyan
Write-Host "ساخت فایل exe" -ForegroundColor Cyan
Write-Host "========================================" -ForegroundColor Cyan
Write-Host ""

# اضافه کردن GCC به PATH
$env:PATH += ";C:\TDM-GCC-64\bin"

# بررسی GCC
Write-Host "بررسی GCC..." -ForegroundColor Yellow
$gccPath = (Get-Command gcc -ErrorAction SilentlyContinue)
if (-not $gccPath) {
    Write-Host "❌ GCC پیدا نشد!" -ForegroundColor Red
    Write-Host "لطفا مطمئن شوید TDM-GCC نصب شده است" -ForegroundColor Red
    Read-Host "برای خروج Enter را بزنید"
    exit 1
}
Write-Host "✓ GCC پیدا شد: $($gccPath.Source)" -ForegroundColor Green
Write-Host ""

# تنظیم CGO
$env:CGO_ENABLED = "1"
Write-Host "✓ CGO فعال شد" -ForegroundColor Green
Write-Host ""

# تغییر به پوشه پروژه
Set-Location $PSScriptRoot

# کامپایل
Write-Host "شروع کامپایل..." -ForegroundColor Yellow
Write-Host ""

$buildOutput = go build -buildvcs=false -ldflags="-s -w" -o notepad-app.exe 2>&1

if ($LASTEXITCODE -eq 0 -and (Test-Path "notepad-app.exe")) {
    Write-Host ""
    Write-Host "========================================" -ForegroundColor Green
    Write-Host "✓✓✓ موفق! فایل exe ساخته شد ✓✓✓" -ForegroundColor Green
    Write-Host "========================================" -ForegroundColor Green
    Write-Host ""
    $file = Get-Item "notepad-app.exe"
    Write-Host "فایل: $($file.Name)" -ForegroundColor Cyan
    Write-Host "اندازه: $([math]::Round($file.Length / 1MB, 2)) MB ($($file.Length) بایت)" -ForegroundColor Cyan
    Write-Host "برای اجرا، فایل را دوبار کلیک کنید!" -ForegroundColor Green
    Write-Host ""
} else {
    Write-Host ""
    Write-Host "❌ خطا در ساخت فایل exe" -ForegroundColor Red
    Write-Host ""
    Write-Host "خروجی کامپایل:" -ForegroundColor Yellow
    Write-Host $buildOutput -ForegroundColor Red
    Write-Host ""
}

Read-Host "برای خروج Enter را بزنید"

