#!/bin/bash
# OpenTamer 构建脚本
# 这个脚本复制了 Justfile 中的构建逻辑

set -euo pipefail

# 配置
APP="build/OpenTamer.app"
CONTENTS="$APP/Contents"
CODESIGN_IDENTITY="${OPENTAMER_CODESIGN_IDENTITY:--}"

echo "=== OpenTamer 构建脚本 ==="
echo

# 检查依赖
echo "检查依赖..."
if ! command -v go &> /dev/null; then
    echo "错误: 未找到 Go。请安装 Go 1.21 或更高版本。"
    exit 1
fi

GO_VERSION=$(go version | awk '{print $3}' | sed 's/go//')
echo "✓ Go 版本: $GO_VERSION"

if ! command -v codesign &> /dev/null; then
    echo "错误: 未找到 codesign。请安装 Xcode Command Line Tools。"
    exit 1
fi
echo "✓ codesign 已安装"

# 清理旧的构建
echo
echo "清理旧构建..."
rm -rf "$APP"

# 创建目录结构
echo "创建 app bundle 目录结构..."
mkdir -p "$CONTENTS/MacOS" "$CONTENTS/Resources"

# 复制资源文件
echo "复制资源文件..."
if [ ! -f "packaging/Info.plist" ]; then
    echo "错误: 未找到 packaging/Info.plist"
    exit 1
fi
cp packaging/Info.plist "$CONTENTS/Info.plist"

if [ ! -f "assets/opentamer-icon.icns" ]; then
    echo "错误: 未找到 assets/opentamer-icon.icns"
    exit 1
fi
cp assets/opentamer-icon.icns "$CONTENTS/Resources/opentamer-icon.icns"

# 设置编译环境
echo "设置编译环境..."
export MACOSX_DEPLOYMENT_TARGET=14.0
export CGO_CFLAGS="${CGO_CFLAGS:--O2 -g} -mmacosx-version-min=$MACOSX_DEPLOYMENT_TARGET"
export CGO_LDFLAGS="${CGO_LDFLAGS:-} -Wl,-no_warn_duplicate_libraries -mmacosx-version-min=$MACOSX_DEPLOYMENT_TARGET"

echo "  MACOSX_DEPLOYMENT_TARGET: $MACOSX_DEPLOYMENT_TARGET"
echo "  CGO_CFLAGS: $CGO_CFLAGS"

# 编译
echo
echo "编译 OpenTamer..."
go build -trimpath -o "$CONTENTS/MacOS/OpenTamer" ./cmd/opentamer

if [ $? -ne 0 ]; then
    echo "错误: 编译失败"
    exit 1
fi
echo "✓ 编译成功"

# 代码签名
echo
echo "代码签名..."
if [ "$CODESIGN_IDENTITY" = "-" ]; then
    echo "  使用 ad-hoc 签名 (仅用于本地开发)"
    codesign --force --sign "$CODESIGN_IDENTITY" "$APP"
else
    echo "  使用开发者证书: $CODESIGN_IDENTITY"
    codesign --force --timestamp --options runtime --sign "$CODESIGN_IDENTITY" "$APP"
fi

if [ $? -ne 0 ]; then
    echo "错误: 签名失败"
    exit 1
fi
echo "✓ 签名成功"

# 验证签名
echo
echo "验证签名..."
codesign --verify --verbose "$APP"
if [ $? -eq 0 ]; then
    echo "✓ 签名验证通过"
else
    echo "⚠ 签名验证失败，但可能仍可运行"
fi

# 显示信息
echo
echo "=== 构建完成 ==="
echo
echo "应用位置: $(pwd)/$APP"
echo "应用大小: $(du -sh "$APP" | cut -f1)"
echo
echo "运行应用:"
echo "  open $APP"
echo
echo "安装到 Applications:"
echo "  cp -r $APP /Applications/"
echo

# 可选: 自动运行
if [ "${OPENTAMER_AUTO_RUN:-}" = "1" ]; then
    echo "自动运行应用..."
    pkill -f 'OpenTamer.app' || true
    sleep 0.5
    open "$APP"
fi
