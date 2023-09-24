setlocal
set version=%3
set GOARCH=%2
set GOOS=%1
set suffix=
if "%1"=="windows" (set suffix=.exe)
go build -o "%~dp0\bin\%GOOS%_%GOARCH%_%version%\mockserver%suffix%" "%~dp0\main.go"
endlocal