#!/usr/bin/env bash
# -------- 遇到错误时立即退出 --------
set -e

# 进入项目根目录，保证相对路径正确
# "$(dirname "$0")是脚本所在目录
cd "$(dirname "$0")"/..

# -------- 配置 --------
APP_NAME=ExcelExporter
APP_VERSION=v2.3
APP_SRC="./src"
APP_DIST="./build/release/${APP_VERSION}"

# -------- 创建输出目录 --------
mkdir -p "$APP_DIST"

# -------- 执行全部单元测试 --------
echo "Running all tests..."
if ! go test "$APP_SRC"/...; then
    echo "Tests failed! Build aborted."
    exit 1
fi
echo "All tests passed."
printf '\n'

# -------- 开始跨平台构建 --------
CGO_ENABLED=0
# -------- macOS --------
echo "Building macOS amd64..."
GOOS=darwin GOARCH=amd64 \
go build -buildvcs=true -o "${APP_DIST}/${APP_NAME}-darwin-amd64_${APP_VERSION}" "$APP_SRC"

# -------- freebsd --------
echo "Building freebsd amd64..."
GOOS=freebsd GOARCH=amd64 \
go build -buildvcs=true -o "${APP_DIST}/${APP_NAME}-freebsd-amd64_${APP_VERSION}" "$APP_SRC"

echo "Building freebsd arm64..."
GOOS=freebsd GOARCH=arm64 \
go build -buildvcs=true -o "${APP_DIST}/${APP_NAME}-freebsd-arm64_${APP_VERSION}" "$APP_SRC"

# -------- Linux --------
echo "Building Linux amd64..."
GOOS=linux GOARCH=amd64 \
go build -buildvcs=true -o "${APP_DIST}/${APP_NAME}-linux-amd64_${APP_VERSION}" "$APP_SRC"

echo "Building Linux arm64..."
GOOS=linux GOARCH=arm64 \
go build -buildvcs=true -o "${APP_DIST}/${APP_NAME}-linux-arm64_${APP_VERSION}" "$APP_SRC"

# -------- Windows --------
echo "Building Windows amd64..."
GOOS=windows GOARCH=amd64 \
go build -buildvcs=true -o "${APP_DIST}/${APP_NAME}-windows-amd64_${APP_VERSION}.exe" "$APP_SRC"

# -------- openbsd --------
echo "Building openbsd amd64..."
GOOS=openbsd GOARCH=amd64 \
go build -buildvcs=true -o "${APP_DIST}/${APP_NAME}-openbsd-amd64_${APP_VERSION}" "$APP_SRC"

printf '\n'

# -------- 开始打包源代码 --------
echo "Packaging source code..."
tar -czf "$APP_DIST/source_$APP_VERSION.tar.gz" -C "$APP_SRC" .
echo "Package source code finish[source_$APP_VERSION.tar.gz]."

echo "Done."
