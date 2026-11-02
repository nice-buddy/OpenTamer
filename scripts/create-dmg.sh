#!/bin/bash
# 创建 OpenTamer DMG 安装包

set -euo pipefail

# 配置
APP="build/OpenTamer.app"
VERSION=$(defaults read "$(pwd)/$APP/Contents/Info.plist" CFBundleShortVersionString 2>/dev/null || echo "1.4.1")
DMG_NAME="OpenTamer-${VERSION}.dmg"
DMG_PATH="build/$DMG_NAME"
TEMP_DMG="build/temp.dmg"
VOLUME_NAME="OpenTamer"

echo "=== 创建 OpenTamer DMG ==="
echo

# 检查应用是否存在
if [ ! -d "$APP" ]; then
    echo "错误: 未找到 $APP"
    echo "请先运行: ./scripts/build.sh"
    exit 1
fi

echo "应用: $APP"
echo "版本: $VERSION"
echo "DMG: $DMG_PATH"
echo

# 清理旧文件
echo "清理旧 DMG..."
rm -f "$TEMP_DMG" "$DMG_PATH"

# 方法 1: 使用 create-dmg (如果已安装)
if command -v create-dmg &> /dev/null; then
    echo "使用 create-dmg 工具..."
    create-dmg \
        --volname "$VOLUME_NAME" \
        --volicon "assets/opentamer-icon.icns" \
        --window-pos 200 120 \
        --window-size 600 400 \
        --icon-size 100 \
        --icon "OpenTamer.app" 175 190 \
        --hide-extension "OpenTamer.app" \
        --app-drop-link 425 190 \
        "$DMG_PATH" \
        "$APP" 2>/dev/null || true
    
    if [ -f "$DMG_PATH" ]; then
        echo "✓ DMG 创建成功"
    else
        echo "create-dmg 失败，使用备用方法..."
    fi
fi

# 方法 2: 使用 hdiutil (备用方法)
if [ ! -f "$DMG_PATH" ]; then
    echo "使用 hdiutil 创建 DMG..."
    
    # 创建临时挂载点
    TEMP_DIR=$(mktemp -d)
    
    # 创建临时 DMG
    echo "  创建临时 DMG..."
    hdiutil create -size 100m -fs HFS+ -volname "$VOLUME_NAME" "$TEMP_DMG" >/dev/null
    
    # 挂载
    echo "  挂载 DMG..."
    MOUNT_DIR=$(hdiutil attach "$TEMP_DMG" -readwrite -noverify -noautoopen | grep "/Volumes/$VOLUME_NAME" | awk '{print $3}')
    
    if [ -z "$MOUNT_DIR" ]; then
        echo "错误: 无法挂载 DMG"
        exit 1
    fi
    
    echo "  挂载点: $MOUNT_DIR"
    
    # 复制应用
    echo "  复制应用..."
    cp -r "$APP" "$MOUNT_DIR/"
    
    # 创建 Applications 快捷方式
    echo "  创建 Applications 链接..."
    ln -s /Applications "$MOUNT_DIR/Applications"
    
    # 设置背景（可选）
    # if [ -f "assets/dmg-background.png" ]; then
    #     mkdir "$MOUNT_DIR/.background"
    #     cp "assets/dmg-background.png" "$MOUNT_DIR/.background/"
    # fi
    
    # 等待文件系统同步
    sync
    sleep 2
    
    # 卸载
    echo "  卸载 DMG..."
    hdiutil detach "$MOUNT_DIR" >/dev/null
    
    # 转换为压缩的只读 DMG
    echo "  压缩 DMG..."
    hdiutil convert "$TEMP_DMG" -format UDZO -o "$DMG_PATH" >/dev/null
    
    # 清理
    rm -f "$TEMP_DMG"
    rm -rf "$TEMP_DIR"
    
    echo "✓ DMG 创建成功"
fi

# 验证 DMG
echo
echo "验证 DMG..."
if hdiutil verify "$DMG_PATH" >/dev/null 2>&1; then
    echo "✓ DMG 验证通过"
else
    echo "⚠ DMG 验证失败，但可能仍可使用"
fi

# 显示信息
echo
echo "=== DMG 创建完成 ==="
echo
echo "DMG 位置: $(pwd)/$DMG_PATH"
echo "DMG 大小: $(du -sh "$DMG_PATH" | cut -f1)"
echo
echo "测试 DMG:"
echo "  open $DMG_PATH"
echo
echo "上传到 GitHub Release:"
echo "  gh release create v$VERSION $DMG_PATH"
echo

# 显示 SHA256
if command -v shasum &> /dev/null; then
    echo "SHA256:"
    shasum -a 256 "$DMG_PATH"
    echo
fi
