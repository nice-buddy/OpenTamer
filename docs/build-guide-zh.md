# OpenTamer 构建指南

本指南将帮助您在 macOS 上构建 OpenTamer 应用。

## 系统要求

- macOS 14.0 (Sonoma) 或更高版本
- Go 1.21 或更高版本
- Xcode Command Line Tools
- just (命令运行器,可选但推荐)

## 安装构建工具

### 1. 安装 Xcode Command Line Tools

如果尚未安装:

```bash
xcode-select --install
```

### 2. 安装 Go

从 https://go.dev/dl/ 下载安装，或使用 Homebrew:

```bash
brew install go
```

验证安装:

```bash
go version  # 应该显示 go1.21 或更高版本
```

### 3. 安装 just (推荐)

just 是一个命令运行器，可以简化构建过程:

```bash
brew install just
```

或者下载预编译的二进制文件:

```bash
curl --proto '=https' --tlsv1.2 -sSf https://just.systems/install.sh | bash -s -- --to /usr/local/bin
```

## 构建方法

### 方法一: 使用 just (推荐)

这是最简单的方法:

```bash
cd OpenTamer

# 构建 app bundle
just build

# 构建并运行
just run
```

构建完成后，应用位于: `build/OpenTamer.app`

### 方法二: 手动构建脚本

如果没有安装 just，可以使用以下脚本:

```bash
cd OpenTamer

# 运行构建脚本
./scripts/build.sh
```

我们提供了 `scripts/build.sh` 脚本来复制 Justfile 的功能。

### 方法三: 直接使用 Go 命令

最基础的构建方法:

```bash
cd OpenTamer

# 创建 app bundle 目录结构
mkdir -p build/OpenTamer.app/Contents/{MacOS,Resources}

# 复制 Info.plist
cp packaging/Info.plist build/OpenTamer.app/Contents/Info.plist

# 复制图标
cp assets/opentamer-icon.icns build/OpenTamer.app/Contents/Resources/opentamer-icon.icns

# 设置环境变量
export MACOSX_DEPLOYMENT_TARGET=14.0
export CGO_CFLAGS="-O2 -g -mmacosx-version-min=14.0"
export CGO_LDFLAGS="-Wl,-no_warn_duplicate_libraries -mmacosx-version-min=14.0"

# 编译
go build -trimpath -o build/OpenTamer.app/Contents/MacOS/OpenTamer ./cmd/opentamer

# 签名 (开发签名)
codesign --force --sign - build/OpenTamer.app
```

## 代码签名

### 开发签名 (本地使用)

默认情况下，构建脚本会使用 ad-hoc 签名 (`-`)，这对本地开发和测试足够了:

```bash
codesign --force --sign - build/OpenTamer.app
```

### 正式签名 (分发使用)

如果要分发应用，需要使用开发者证书:

```bash
# 查看可用的签名身份
security find-identity -v -p codesigning

# 使用特定身份构建
export OPENTAMER_CODESIGN_IDENTITY="Developer ID Application: Your Name (TEAM_ID)"
just build
```

或手动签名:

```bash
codesign --force --timestamp --options runtime \
  --sign "Developer ID Application: Your Name (TEAM_ID)" \
  build/OpenTamer.app
```

## 运行应用

### 从构建目录运行

```bash
open build/OpenTamer.app
```

### 安装到 Applications 文件夹

```bash
cp -r build/OpenTamer.app /Applications/
open /Applications/OpenTamer.app
```

## 创建 DMG 分发包

### 使用 create-dmg

```bash
# 安装 create-dmg
brew install create-dmg

# 创建 DMG
create-dmg \
  --volname "OpenTamer" \
  --volicon "assets/opentamer-icon.icns" \
  --window-pos 200 120 \
  --window-size 600 400 \
  --icon-size 100 \
  --icon "OpenTamer.app" 175 190 \
  --hide-extension "OpenTamer.app" \
  --app-drop-link 425 190 \
  "build/OpenTamer-1.4.1.dmg" \
  "build/OpenTamer.app"
```

### 手动创建 DMG

```bash
# 创建临时 DMG
hdiutil create -size 50m -fs HFS+ -volname "OpenTamer" build/temp.dmg

# 挂载
hdiutil attach build/temp.dmg

# 复制应用
cp -r build/OpenTamer.app /Volumes/OpenTamer/

# 卸载
hdiutil detach /Volumes/OpenTamer

# 转换为压缩的只读 DMG
hdiutil convert build/temp.dmg -format UDZO -o build/OpenTamer-1.4.1.dmg

# 清理临时文件
rm build/temp.dmg
```

## 测试

### 运行单元测试

```bash
go test ./...
```

### 运行特定测试

```bash
# 测试语言功能
go test -v -run TestLanguage ./internal/config

# 测试本地化
go test -v -run TestLocale ./internal/ui
```

### 验证构建

```bash
# 运行验证脚本
./scripts/verify_i18n.sh
```

## 常见问题

### 问题: cgo 编译错误

确保安装了 Xcode Command Line Tools:

```bash
xcode-select --install
sudo xcodebuild -license accept
```

### 问题: 签名失败

如果遇到签名错误，可以使用 ad-hoc 签名:

```bash
codesign --force --sign - build/OpenTamer.app
```

### 问题: 应用无法运行

检查系统版本:

```bash
sw_vers  # 需要 macOS 14.0 或更高版本
```

如果使用旧版本 macOS，需要修改 `MACOSX_DEPLOYMENT_TARGET`。

### 问题: 权限错误

给予构建脚本执行权限:

```bash
chmod +x scripts/*.sh
```

## 清理构建

```bash
# 清理构建产物
rm -rf build/

# 清理 Go 缓存
go clean -cache
```

## 开发工作流

### 快速迭代

```bash
# 监视文件变化并自动重新构建
just run
```

### 完整的开发周期

```bash
# 1. 进行代码更改
# 2. 运行测试
go test ./...

# 3. 构建
just build

# 4. 运行
open build/OpenTamer.app

# 5. 如果需要，查看日志
log show --predicate 'subsystem == "org.opentamer.OpenTamer"' --last 5m
```

## 发布清单

准备发布新版本时:

1. 更新 `packaging/Info.plist` 中的版本号
2. 更新 `CHANGELOG.md`
3. 运行完整测试: `go test ./...`
4. 使用正式证书构建: `OPENTAMER_CODESIGN_IDENTITY="..." just build`
5. 创建 DMG: 参见上面的 DMG 创建说明
6. 公证应用 (如果要分发到 App Store 之外)
7. 创建 GitHub Release 并上传 DMG

## 更多信息

- 项目 README: `README.md`
- 贡献指南: 见 README.md 的 Contributing 部分
- i18n 实现文档: `docs/i18n-implementation-zh.md`
