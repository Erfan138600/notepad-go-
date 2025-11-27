@echo off
cd /d "%~dp0"
set "PATH=%PATH%;C:\TDM-GCC-64\bin"
set "CGO_ENABLED=1"

echo Building...
go build -buildvcs=false -ldflags="-s -w" -o notepad-app.exe > build-log.txt 2>&1
echo Exit code: %ERRORLEVEL% >> build-log.txt

if exist notepad-app.exe (
    echo SUCCESS > build-result.txt
) else (
    echo FAILED > build-result.txt
)

