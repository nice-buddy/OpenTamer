# OpenTamer 构建和打包指南

本文档提供 OpenTamer 的完整构建和打包流程说明。

## 快速开始

```bash
# 1. 克隆或进入项目目录
cd OpenTamer

# 2. 构建应用
./scripts/build.sh

# 3. 运行应用
open build/OpenTamer.app

# 4. (可选) 创建 DMG 安装包
./scripts/create-dmg.sh
```

## 系统要求

- **操作系统**: macOS 14.0 (Sonoma) 或更高版本
- **Go**: 1.21 或更高版本
- **Xcode Command Line Tools**: 必需

## 准备工作

### 1. 安装 Xcode Command Line Tools

```bash
xcode-select --install
```

验证安装：

```bash
xcode-select -p
# 应该输出类似: /Library/Developer/CommandLineTools
```

### 2. 验证 Go 环境

```bash
go version
# 应该显示 go1.21 或更高版本
```

如果未安装 Go:

```bash
# 使用 Homebrew
brew install go

# 或从官网下载
# https://go.dev/dl/
```

### 3. 克隆项目（如果尚未克隆）

```bash
git clone https://github.com/Digital-Shane/OpenTamer.git
cd OpenTamer
```

## 构建方式

### 方式一：使用构建脚本（推荐）

我们提供了便捷的构建脚本：

```bash
# 构建应用
./scripts/build.sh

# 构建完成后，应用位于: build/OpenTamer.app
```

脚本会自动完成：
- 检查依赖
- 创建 app bundle 结构
- 复制资源文件（Info.plist、图标）
- 编译 Go 代码
- 代码签名（ad-hoc 签名用于本地开发）
- 验证构建

### 方式二：使用 just（如果已安装）

```bash
# 安装 just（可选）
brew install just

# 构建
just build

# 构建并运行
just run
```

### 方式三：手动构建

如果你想了解构建细节：

```bash
# 1. 创建目录结构
mkdir -p build/OpenTamer.app/Contents/{MacOS,Resources}

# 2. 复制资源
cp packaging/Info.plist build/OpenTamer.app/Contents/Info.plist
cp assets/opentamer-icon.icns build/OpenTamer.app/Contents/Resources/opentamer-icon.icns

# 3. 设置环境变量
export MACOSX_DEPLOYMENT_TARGET=14.0
export CGO_CFLAGS="-O2 -g -mmacosx-version-min=14.0"
export CGO_LDFLAGS="-Wl,-no_warn_duplicate_libraries -mmacosx-version-min=14.0"

# 4. 编译
go build -trimpath -o build/OpenTamer.app/Contents/MacOS/OpenTamer ./cmd/opentamer

# 5. 签名
codesign --force --sign - build/OpenTamer.app
```

## 运行和测试

### 运行应用

```bash
# 从构建目录运行
open build/OpenTamer.app

# 或者安装到 Applications 后运行
cp -r build/OpenTamer.app /Applications/
open /Applications/OpenTamer.app
```

### 运行测试

```bash
# 运行所有测试
go test ./...

# 运行语言功能测试
go test -v -run TestLanguage ./internal/config

# 运行本地化测试
go test -v -run TestLocale ./internal/ui

# 验证 i18n 实现
./scripts/verify_i18n.sh
```

### 查看应用日志

```bash
# 实时查看最近 5 分钟的日志
log show --predicate 'subsystem == "org.opentamer.OpenTamer"' --last 5m

# 持续监控日志
log stream --predicate 'subsystem == "org.opentamer.OpenTamer"'
```

## 创建分发包

### 创建 DMG 安装包

```bash
# 确保已构建应用
./scripts/build.sh

# 创建 DMG
./scripts/create-dmg.sh

# DMG 文件位于: build/OpenTamer-1.4.1.dmg
```

### DMG 内容

创建的 DMG 包含：
- OpenTamer.app
- Applications 文件夹快捷方式（方便拖放安装）

### 测试 DMG

```bash
# 挂载 DMG
open build/OpenTamer-1.4.1.dmg

# 从 DMG 中拖动应用到 Applications 文件夹测试安装流程
```

## 代码签名

### 开发签名（本地使用）

默认情况下，构建脚本使用 ad-hoc 签名，适合本地开发：

```bash
./scripts/build.sh
# 使用 ad-hoc 签名 (-)
```

验证签名：

```bash
codesign --verify --verbose build/OpenTamer.app
```

### 正式签名（公开分发）

如果要分发给其他用户，需要使用开发者证书：

```bash
# 1. 查看可用的签名身份
security find-identity -v -p codesigning

# 2. 设置签名身份并构建
export OPENTAMER_CODESIGN_IDENTITY="Developer ID Application: Your Name (TEAM_ID)"
./scripts/build.sh

# 3. 验证签名
codesign --verify --verbose build/OpenTamer.app
spctl --assess --verbose build/OpenTamer.app
```

### 公证（Notarization）

如果要在 Gatekeeper 保护下分发（推荐）：

```bash
# 1. 使用开发者证书构建
export OPENTAMER_CODESIGN_IDENTITY="Developer ID Application: Your Name (TEAM_ID)"
./scripts/build.sh

# 2. 创建 DMG
./scripts/create-dmg.sh

# 3. 提交公证
xcrun notarytool submit build/OpenTamer-1.4.1.dmg \
  --apple-id "your@email.com" \
  --team-id "TEAM_ID" \
  --password "app-specific-password" \
  --wait

# 4. 附加公证票据
xcrun stapler staple build/OpenTamer-1.4.1.dmg

# 5. 验证
spctl --assess --type open --context context:primary-signature --verbose build/OpenTamer-1.4.1.dmg
```

## 版本管理

### 更新版本号

编辑 `packaging/Info.plist`:

```xml
<key>CFBundleShortVersionString</key>
<string>1.5.0</string>  <!-- 用户可见版本 -->
<key>CFBundleVersion</key>
<string>1</string>        <!-- 构建号 -->
```

### 更新变更日志

在发布前更新 `CHANGELOG.md`。

## 发布流程

完整的发布检查清单：

```bash
# 1. 更新版本号
# 编辑 packaging/Info.plist

# 2. 更新变更日志
# 编辑 CHANGELOG.md

# 3. 运行所有测试
go test ./...
./scripts/verify_i18n.sh

# 4. 使用正式证书构建
export OPENTAMER_CODESIGN_IDENTITY="Developer ID Application: Your Name (TEAM_ID)"
./scripts/build.sh

# 5. 测试应用
open build/OpenTamer.app
# 手动测试核心功能和语言切换

# 6. 创建 DMG
./scripts/create-dmg.sh

# 7. 公证 DMG (可选但推荐)
# 见上面的公证说明

# 8. 创建 Git 标签
git tag -a v1.5.0 -m "Release version 1.5.0"
git push origin v1.5.0

# 9. 创建 GitHub Release
gh release create v1.5.0 build/OpenTamer-1.5.0.dmg \
  --title "OpenTamer 1.5.0" \
  --notes "查看 CHANGELOG.md 了解详细更新内容"
```

## 清理构建

```bash
# 清理构建产物
rm -rf build/

# 清理 Go 缓存和模块缓存
go clean -cache -modcache

# 深度清理（包括测试缓存）
go clean -cache -testcache -modcache
```

## 故障排除

### 问题：编译错误

```bash
# 检查 Go 版本
go version  # 需要 1.21+

# 检查 CGO
go env CGO_ENABLED  # 应该是 1

# 重新安装 Command Line Tools
sudo rm -rf /Library/Developer/CommandLineTools
xcode-select --install
```

### 问题：签名失败

```bash
# 使用 ad-hoc 签名
unset OPENTAMER_CODESIGN_IDENTITY
./scripts/build.sh

# 或手动签名
codesign --force --sign - build/OpenTamer.app
```

### 问题：应用无法启动

```bash
# 检查系统版本
sw_vers
# ProductVersion 应该 >= 14.0

# 查看崩溃报告
open ~/Library/Logs/DiagnosticReports/

# 查看控制台日志
Console.app  # 然后搜索 "OpenTamer"
```

### 问题：DMG 创建失败

```bash
# 确保应用已构建
ls -la build/OpenTamer.app

# 手动卸载可能挂载的卷
hdiutil detach /Volumes/OpenTamer 2>/dev/null || true

# 重试
./scripts/create-dmg.sh
```

### 问题：Gatekeeper 阻止运行

对于本地开发构建：

```bash
# 移除隔离属性
xattr -dr com.apple.quarantine build/OpenTamer.app

# 或在系统设置中允许
# 系统设置 > 隐私与安全性 > 安全性 > "仍要打开"
```

## 开发技巧

### 快速迭代开发

```bash
# 监视文件变化，使用 watchexec
brew install watchexec
watchexec -w internal -w cmd -e go './scripts/build.sh && open build/OpenTamer.app'
```

### 性能分析

```bash
# CPU 性能分析
go test -cpuprofile=cpu.prof -bench=. ./internal/...
go tool pprof cpu.prof

# 内存分析
go test -memprofile=mem.prof -bench=. ./internal/...
go tool pprof mem.prof
```

### 代码覆盖率

```bash
# 生成覆盖率报告
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## 相关文档

- [README.md](README.md) - 项目概述
- [docs/build-guide-zh.md](docs/build-guide-zh.md) - 详细构建指南
- [docs/i18n-implementation-zh.md](docs/i18n-implementation-zh.md) - 国际化实现说明
- [CHANGELOG.md](CHANGELOG.md) - 版本变更历史

## 获取帮助

- **Issue**: https://github.com/Digital-Shane/OpenTamer/issues
- **Discussions**: https://github.com/Digital-Shane/OpenTamer/discussions

---

**提示**: 如果您是首次构建，强烈建议使用 `./scripts/build.sh` 脚本，它会自动处理所有必要步骤。
