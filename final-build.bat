@echo off
cd /d "%~dp0"
set "PATH=%PATH%;C:\TDM-GCC-64\bin"
set "CGO_ENABLED=1"
go build -buildvcs=false -ldflags="-s -w" -o notepad-app.exe > build-status.txt 2>&1
echo ExitCode: %ERRORLEVEL% >> build-status.txt
if exist notepad-app.exe (
    echo SUCCESS >> build-status.txt
    dir notepad-app.exe >> build-status.txt
) else (
    echo FAILED >> build-status.txt
)

