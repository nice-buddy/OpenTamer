# OpenTamer 快速开始指南

> 从零开始构建、运行和分发 OpenTamer

## 🚀 5 分钟快速构建

```bash
# 1. 进入项目目录
cd OpenTamer

# 2. 构建应用（约 10 秒）
./scripts/build.sh

# 3. 运行应用
open build/OpenTamer.app
```

就这么简单！✨

## 📋 前置要求

- macOS 14.0+
- Go 1.21+
- Xcode Command Line Tools

### 安装 Command Line Tools

```bash
xcode-select --install
```

### 验证环境

```bash
# 检查 Go
go version  # 应该显示 go1.21 或更高

# 检查 Xcode Tools
xcode-select -p  # 应该显示 /Library/Developer/CommandLineTools
```

## 🛠️ 构建选项

### 选项 1: 使用脚本（推荐）

```bash
./scripts/build.sh
```

自动完成所有步骤，包括编译、签名和验证。

### 选项 2: 使用 just

如果已安装 just:

```bash
brew install just  # 首次安装
just build         # 构建
just run           # 构建并运行
```

### 选项 3: 手动构建

```bash
mkdir -p build/OpenTamer.app/Contents/{MacOS,Resources}
cp packaging/Info.plist build/OpenTamer.app/Contents/Info.plist
cp assets/opentamer-icon.icns build/OpenTamer.app/Contents/Resources/
export MACOSX_DEPLOYMENT_TARGET=14.0
go build -trimpath -o build/OpenTamer.app/Contents/MacOS/OpenTamer ./cmd/opentamer
codesign --force --sign - build/OpenTamer.app
```

## 📦 创建安装包

```bash
# 确保已构建应用
./scripts/build.sh

# 创建 DMG
./scripts/create-dmg.sh

# 结果: build/OpenTamer-1.4.1.dmg
```

## 🧪 测试

```bash
# 运行所有测试
go test ./...

# 验证 i18n 功能
./scripts/verify_i18n.sh

# 测试特定功能
go test -v -run TestLanguage ./internal/config
```

## 🌍 语言切换功能

新增的国际化功能：

1. **自动检测**: 首次启动自动检测系统语言
2. **手动切换**: 偏好设置 > 通用 > 界面语言
3. **支持语言**: 
   - 自动（跟随系统）
   - English
   - 中文

### 测试语言切换

```bash
# 运行应用
open build/OpenTamer.app

# 1. 点击菜单栏图标
# 2. 点击 "..." 按钮
# 3. 选择 "General" (通用)
# 4. 在 "Language" (界面语言) 中选择语言
# 5. 重启应用查看效果
```

## 📝 构建输出

成功构建后你会看到：

```
=== 构建完成 ===

应用位置: /path/to/OpenTamer/build/OpenTamer.app
应用大小: 5.7M

运行应用:
  open build/OpenTamer.app

安装到 Applications:
  cp -r build/OpenTamer.app /Applications/
```

## 🔧 常见问题

### Q: 编译错误 "cgo: C compiler not found"

**A**: 安装 Xcode Command Line Tools

```bash
xcode-select --install
```

### Q: 应用无法打开 "已损坏"

**A**: 移除隔离属性

```bash
xattr -dr com.apple.quarantine build/OpenTamer.app
```

### Q: 构建脚本无法执行

**A**: 添加执行权限

```bash
chmod +x scripts/*.sh
```

### Q: Go 版本过低

**A**: 升级 Go

```bash
brew upgrade go
# 或从 https://go.dev/dl/ 下载
```

## 📚 更多文档

- [BUILD-CN.md](../BUILD-CN.md) - 完整构建指南
- [i18n-implementation-zh.md](i18n-implementation-zh.md) - 国际化实现细节
- [build-guide-zh.md](build-guide-zh.md) - 详细构建说明

## 🎯 下一步

构建成功后：

1. **本地使用**: `open build/OpenTamer.app`
2. **安装**: `cp -r build/OpenTamer.app /Applications/`
3. **分发**: 创建 DMG 并上传到 GitHub Release

## 💡 提示

- 首次构建可能需要下载依赖，耗时较长
- 后续构建会使用缓存，速度很快
- 使用 ad-hoc 签名足够本地开发
- 公开分发需要开发者证书

## 🆘 获取帮助

遇到问题？

1. 查看[完整构建指南](../BUILD-CN.md)
2. 搜索 [Issues](https://github.com/Digital-Shane/OpenTamer/issues)
3. 提交新 Issue

---

**祝构建顺利！** 🎉
