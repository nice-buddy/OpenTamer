# 国际化(i18n)实现说明

## 概述

OpenTamer 现已支持中英文 UI 语言切换。语言偏好在首次启动时自动检测,并可通过偏好设置菜单手动更改。

## 已实现的功能

### 1. 语言检测

- **自动检测**:首次启动时,系统语言将自动检测,使用:
  - macOS `defaults read -g AppleLanguages` 命令(主要方法)
  - 环境变量 `LANG` 和 `LC_ALL`(备用方法)

- **支持的语言**:
  - 英语 (`en`)
  - 中文 (`zh`)

- **默认行为**:如果检测到的语言不是中文,应用将默认使用英语。

### 2. 语言偏好设置

在偏好设置菜单的 **通用** 部分新增了语言偏好选项,包含以下选项:

- **自动(跟随系统语言)** - 自动检测并使用系统语言
- **English** - 强制使用英文界面
- **中文** - 强制使用中文界面

### 3. 持久化存储

- 语言偏好存储在 `~/Library/Application Support/OpenTamer/config.json`
- 首次用户:自动检测并保存语言
- 老用户:保留已有的语言偏好
- 将语言设置为 "" (空字符串)表示"自动"模式

### 4. 实现细节

#### 核心组件

**配置 (`internal/config/config.go`)**:
- `DetectLanguage()` - 检测系统语言
- 语言是 `GlobalPreferences` 结构的一部分
- 迁移逻辑确保现有配置在首次加载时获得语言检测

**本地化 (`internal/l10n/`)**:
- 定义语言常量
- `NormalizeLanguage()` 函数用于验证

**UI 本地化 (`internal/ui/l10n.go`)**:
- `EnglishLocale()` - 所有英文 UI 字符串
- `ChineseLocale()` - 所有中文 UI 字符串
- `LocaleForLanguage()` - 获取适当的本地化
- `LocaleForPreferences()` - 从偏好设置解析本地化(带自动检测回退)

**菜单桥接 (`internal/macos/menu/`)**:
- 在通用偏好设置中添加语言下拉菜单
- 三个选项:自动、English、中文
- 通过标准字符串偏好机制保存
- 界面层通过 `OpenTamerLocalized()` 统一调用 `OpenTamerL10nTranslate`
- 状态与安全原因通过 `OpenTamerL10nStatus` / `OpenTamerL10nSafety` 本地化

#### 界面文案覆盖

Objective-C 菜单栏界面中的全部用户可见文案均已接入翻译,包括:

- 主面板:标题、刷新、系统 CPU 摘要、Top/已跟踪/受管进程分区、空状态
- 应用菜单:状态、备注、规则、跟踪、降低优先级、限制 CPU、后台暂停
- 偏好设置:通用、统计与图表、高 CPU 提醒、策略防护的全部标题与选项
- 提示框:CPU 限制、提醒阈值、登录项错误
- 菜单栏状态:关闭状态、图表提示

`internal/macos/menu/menu_bar_l10n_test.go` 会自动扫描 Objective-C 源码:

- 每一个 `OpenTamerLocalized(@"key")` 都必须有中文翻译
- 禁止新增未本地化的菜单标题、按钮、提示框文案

#### 修改/添加的文件

**已修改**:
- `internal/config/config.go` - 语言检测和迁移
- `internal/core/rules.go` - 语言常量和规范化
- `internal/macos/menu/menu_bar_bridge.m` - 语言偏好 UI
- `internal/macos/menu/menu_bar_preferences_test.go` - 测试期望

**新增**:
- `internal/l10n/l10n.go` - 语言类型和规范化
- `internal/l10n/l10n_test.go` - 语言测试
- `internal/macos/menu/l10n_bridge_darwin.go` - l10n 函数的 C 桥接
- `internal/macos/menu/l10n_bridge.h` - l10n C 导出的头文件
- `internal/ui/l10n.go` - 两种语言的本地化定义
- `internal/ui/l10n_test.go` - 本地化测试
- `internal/ui/bridge_l10n.go` - UI 字符串的桥接函数
- `internal/ui/lang_detect.go` - 语言检测工具
- `internal/config/language_integration_test.go` - 集成测试

## 测试

### 测试覆盖

1. **语言检测测试**:验证系统语言被正确检测
2. **语言规范化测试**:确保只接受有效语言
3. **持久化测试**:确认语言偏好被正确保存和加载
4. **本地化测试**:验证所有本地化字符串被正确定义
5. **菜单测试**:验证语言选项出现在偏好设置菜单中

### 运行测试

```bash
# 运行所有测试
go test ./...

# 运行语言特定测试
go test -v -run TestLanguage ./internal/config
go test -v -run TestLocale ./internal/ui
```

## 使用说明

### 用户使用

1. **首次启动**:应用自动检测您的系统语言并设置
2. **更改语言**:
   - 点击 OpenTamer 菜单栏图标
   - 选择 "..." 按钮打开偏好设置
   - 进入 **通用** 部分
   - 从 **界面语言** 下拉菜单中选择您的首选语言
   - 可选择:自动(跟随系统语言)、English 或 中文

### 开发者使用

#### 添加新的语言字符串

1. 在 `internal/ui/l10n.go` 的 `Locale` 结构中添加字符串键
2. 在 `EnglishLocale()` 中添加英文翻译
3.在 `ChineseLocale()` 中添加中文翻译
4. 通过 `ui.LocaleForPreferences(preferences).YourNewKey` 使用该字符串

#### 添加新语言

1. 在 `internal/core/rules.go` 中添加新语言常量:
   ```go
   const LanguageSpanish = "es"
   ```
2. 更新 `NormalizeLanguage()` 以包含新语言
3. 在 `internal/ui/l10n.go` 中创建新的本地化函数:
   ```go
   func SpanishLocale() Locale { ... }
   ```
4. 更新 `LocaleForLanguage()` 以处理新语言
5. 在 `menu_bar_bridge.m` 中将新选项添加到菜单

## 技术说明

- 语言检测仅在首次启动时发生一次
- 后续启动使用保存的偏好设置
- 将语言设置为空字符串会恢复自动检测行为
- 更改语言后 UI 不会自动更新(需要重启应用)
- 所有本地化字符串都是编译时常量(无外部翻译文件)
