# راهنمای ساخت برنامه

## نیازمندی‌های کامپایل

برای کامپایل برنامه با Fyne در ویندوز، به موارد زیر نیاز دارید:

### 1. نصب Go
از [اینجا](https://go.dev/dl/) دانلود و نصب کنید.

### 2. نصب کامپایلر C (GCC)

Fyne نیاز به CGO دارد که برای کار نیاز به کامپایلر C دارد.

#### روش 1: استفاده از MinGW-w64
1. از [اینجا](https://sourceforge.net/projects/mingw-w64/) دانلود کنید
2. یا از Chocolatey:
   ```powershell
   choco install mingw
   ```

#### روش 2: استفاده از TDM-GCC
از [اینجا](https://jmeubank.github.io/tdm-gcc/) دانلود و نصب کنید.

### 3. افزودن GCC به PATH
بعد از نصب، مسیر `bin` کامپایلر را به متغیر محیطی PATH اضافه کنید.

مثلا اگر MinGW در `C:\mingw64\bin` نصب شده:
```powershell
setx PATH "%PATH%;C:\mingw64\bin"
```

## ساخت برنامه

### روش 1: استفاده از اسکریپت
```bash
build.bat
```

### روش 2: دستی
```bash
set CGO_ENABLED=1
go build -ldflags="-s -w" -o notepad-app.exe
```

### روش 3: استفاده از Fyne CLI (پیشنهادی)
```bash
# نصب Fyne CLI
go install fyne.io/fyne/v2/cmd/fyne@latest

# ساخت برنامه
fyne package -os windows
```

## تست نصب GCC

برای اطمینان از نصب صحیح GCC:
```bash
gcc --version
```

اگر دستور بالا کار کرد، GCC به درستی نصب شده است.


