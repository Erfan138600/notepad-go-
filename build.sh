#!/bin/bash

echo "در حال دانلود وابستگی‌ها..."
go mod download

echo ""
echo "در حال کامپایل برنامه..."
go build -ldflags="-s -w" -o notepad-app

if [ $? -eq 0 ]; then
    echo ""
    echo "✓ برنامه با موفقیت ساخته شد!"
    echo "✓ فایل خروجی: notepad-app"
else
    echo ""
    echo "✗ خطا در ساخت برنامه!"
    echo "لطفا مطمئن شوید که Go نصب شده است."
fi


