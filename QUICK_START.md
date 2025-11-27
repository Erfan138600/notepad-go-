# راهنمای سریع ساخت فایل exe

## مشکل: GCC پیدا نشد

برای ساخت برنامه‌های Fyne در ویندوز، نیاز به کامپایلر C (GCC) دارید.

## نصب سریع GCC

### روش 1: استفاده از TDM-GCC (پیشنهادی - ساده‌تر)

1. **دانلود**: از [اینجا](https://jmeubank.github.io/tdm-gcc/) دانلود کنید
2. **نصب**: فایل exe را اجرا و Next را بزنید (تنظیمات پیش‌فرض مناسب است)
3. **بررسی**: PowerShell را restart کنید و دستور زیر را اجرا کنید:
   ```powershell
   gcc --version
   ```
   باید نسخه GCC نمایش داده شود.

4. **ساخت**: حالا `build.bat` را اجرا کنید

### روش 2: استفاده از MinGW-w64

1. **دانلود**: از [اینجا](https://sourceforge.net/projects/mingw-w64/) دانلود کنید
2. **نصب**: 
   - فایل exe را اجرا کنید
   - Architecture را x86_64 انتخاب کنید
   - Threads را posix انتخاب کنید
   - بقیه تنظیمات را پیش‌فرض بگذارید
   - محل نصب را یادداشت کنید (معمولا `C:\mingw64`)

3. **اضافه کردن به PATH**:
   - روی This PC راست کلیک → Properties
   - Advanced system settings
   - Environment Variables
   - در بخش System variables، متغیر Path را انتخاب و Edit کنید
   - New را بزنید و مسیر `C:\mingw64\bin` را اضافه کنید
   - OK را بزنید

4. **بررسی**: PowerShell را restart کنید و دستور زیر را اجرا کنید:
   ```powershell
   gcc --version
   ```

5. **ساخت**: حالا `build.bat` را اجرا کنید

### روش 3: استفاده از Chocolatey (برای کاربران حرفه‌ای)

اگر Chocolatey نصب دارید:
```powershell
choco install mingw -y
```

سپس PowerShell را restart کنید و `build.bat` را اجرا کنید.

## بعد از نصب GCC

1. PowerShell را **بستن و دوباره باز کردن** (مهم!)
2. در پوشه پروژه، دستور زیر را اجرا کنید:
   ```powershell
   .\build.bat
   ```

یا دستی:
```powershell
$env:CGO_ENABLED=1
go build -buildvcs=false -ldflags="-s -w" -o notepad-app.exe
```

## بررسی نصب

برای اطمینان از نصب صحیح:
```powershell
gcc --version
go version
```

هر دو باید بدون خطا کار کنند.

## اگر هنوز مشکل دارید

1. مطمئن شوید PowerShell را restart کرده‌اید
2. بررسی کنید که GCC در PATH است:
   ```powershell
   where.exe gcc
   ```
3. اگر GCC پیدا نشد، مسیر bin را به PATH اضافه کنید

## تست برنامه

بعد از ساخت موفق، فایل `notepad-app.exe` ایجاد می‌شود. دوبار کلیک کنید تا اجرا شود!


