@echo off
setlocal

REM -------- 回到项目根目录 --------
REM -------- ~dp0 指的是当前bat文件的路径 --------
cd %~dp0\..

REM -------- 配置 --------
set APP_NAME=ExcelExporter
set APP_VERSION=v2.3
set APP_SRC=.\src
set APP_DIST=.\build\release\%APP_VERSION%


REM -------- 创建输出目录 --------
if not exist %APP_DIST% (
    mkdir %APP_DIST%
)

REM -------- 执行全部单元测试 --------
echo Running all tests...
go test %APP_SRC%\...
@REM %ERRORLEVEL% 是 一个系统变量，表示 上一条命令的退出状态码（exit code）
IF %ERRORLEVEL% NEQ 0 (
    echo Tests failed! Build aborted.
    exit /b 1
    pause
)
echo All tests passed.

REM -------- 开始跨平台构建 --------

set CGO_ENABLED=0

REM -------- macOS --------
echo Building macOS amd64...
set GOOS=darwin
set GOARCH=amd64
set APP_PATH=%APP_DIST%\%APP_NAME%-darwin-amd64_%APP_VERSION%
go build -o %APP_PATH% %APP_SRC%

REM -------- freebsd --------
echo Building freebsd amd64...
set GOOS=freebsd
set GOARCH=amd64
set APP_PATH=%APP_DIST%\%APP_NAME%-freebsd-amd64_%APP_VERSION%
go build -buildvcs=true -o %APP_PATH% %APP_SRC%

echo Building freebsd arm64...
set GOOS=freebsd
set GOARCH=arm64
set APP_PATH=%APP_DIST%\%APP_NAME%-freebsd-arm64_%APP_VERSION%
go build -buildvcs=true -o %APP_PATH% %APP_SRC%

REM -------- Linux --------
echo Building Linux amd64...
set GOOS=linux
set GOARCH=amd64
set APP_PATH=%APP_DIST%\%APP_NAME%-linux-amd64_%APP_VERSION%
go build -buildvcs=true -o %APP_PATH% %APP_SRC%

echo Building Linux arm64...
set GOOS=linux
set GOARCH=arm64
set APP_PATH=%APP_DIST%\%APP_NAME%-linux-arm64_%APP_VERSION%
go build -buildvcs=true -o %APP_PATH% %APP_SRC%

REM -------- Windows --------
echo Building Windows amd64...
set GOOS=windows
set GOARCH=amd64
set APP_PATH=%APP_DIST%\%APP_NAME%-windows-amd64_%APP_VERSION%.exe
go build -buildvcs=true -o %APP_PATH% %APP_SRC%

REM -------- openbsd --------
echo Building openbsd amd64...
set GOOS=openbsd
set GOARCH=amd64
set APP_PATH=%APP_DIST%\%APP_NAME%-openbsd-amd64_%APP_VERSION%
go build -buildvcs=true -ldflags -o %APP_PATH% %APP_SRC%


echo Done.
endlocal

pause