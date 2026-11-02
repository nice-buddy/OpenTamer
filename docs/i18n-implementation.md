# Internationalization (i18n) Implementation

## Overview

OpenTamer now supports UI language switching between Chinese and English. The language preference is automatically detected on first launch and can be manually changed through the preferences menu.

## Features Implemented

### 1. Language Detection

- **Automatic Detection**: On first launch, the system language is automatically detected using:
  - macOS `defaults read -g AppleLanguages` command (primary method)
  - Environment variables `LANG` and `LC_ALL` (fallback method)

- **Supported Languages**:
  - English (`en`)
  - Chinese (`zh`)

- **Default Behavior**: If the detected language is not Chinese, the app defaults to English.

### 2. Language Preference Setting

A new language preference has been added to the **General** section of the Preferences menu with the following options:

- **Auto (system language)** - Automatically detects and uses the system language
- **English** - Forces English UI
- **中文** - Forces Chinese UI

### 3. Persistent Storage

- The language preference is stored in `~/Library/Application Support/OpenTamer/config.json`
- First-time users: Language is auto-detected and saved
- Returning users: Language preference is preserved
- Setting language to "" (empty string) means "Auto" mode

### 4. Implementation Details

#### Core Components

**Configuration (`internal/config/config.go`)**:
- `DetectLanguage()` - Detects system language
- Language is part of `GlobalPreferences` struct
- Migration logic ensures existing configs get language detection on first load

**Localization (`internal/l10n/`)**:
- Language constants defined
- `NormalizeLanguage()` function for validation

**UI Locales (`internal/ui/l10n.go`)**:
- `EnglishLocale()` - All English UI strings
- `ChineseLocale()` - All Chinese UI strings
- `LocaleForLanguage()` - Gets appropriate locale
- `LocaleForPreferences()` - Resolves locale from preferences (with auto-detect fallback)

**Menu Bridge (`internal/macos/menu/`)**:
- Added language dropdown to General preferences
- Three options: Auto, English, 中文
- Saved via standard string preference mechanism

#### Files Modified/Added

**Modified**:
- `internal/config/config.go` - Language detection and migration
- `internal/core/rules.go` - Language constants and normalization
- `internal/macos/menu/menu_bar_bridge.m` - Language preference UI
- `internal/macos/menu/menu_bar_preferences_test.go` - Test expectations

**Added**:
- `internal/l10n/l10n.go` - Language type and normalization
- `internal/l10n/l10n_test.go` - Language tests
- `internal/macos/menu/l10n_bridge_darwin.go` - C bridge for l10n functions
- `internal/macos/menu/l10n_bridge.h` - Header for l10n C exports
- `internal/ui/l10n.go` - Locale definitions for both languages
- `internal/ui/l10n_test.go` - Locale tests
- `internal/ui/bridge_l10n.go` - Bridge functions for UI strings
- `internal/ui/lang_detect.go` - Language detection utilities
- `internal/config/language_integration_test.go` - Integration tests

## Testing

### Test Coverage

1. **Language Detection Tests**: Verify system language is correctly detected
2. **Language Normalization Tests**: Ensure only valid languages are accepted
3. **Persistence Tests**: Confirm language preference is saved and loaded correctly
4. **Locale Tests**: Validate all locale strings are properly defined
5. **Menu Tests**: Verify language option appears in preferences menu

### Running Tests

```bash
# Run all tests
go test ./...

# Run language-specific tests
go test -v -run TestLanguage ./internal/config
go test -v -run TestLocale ./internal/ui
```

## Usage

### For Users

1. **First Launch**: The app automatically detects your system language and sets it
2. **Changing Language**:
   - Click the OpenTamer menu bar icon
   - Select the "..." button to open Preferences
   - Go to **General** section
   - Select your preferred language from the **Language** dropdown
   - Choose from: Auto (system language), English, or 中文

### For Developers

#### Adding New Language Strings

1. Add the string key to the `Locale` struct in `internal/ui/l10n.go`
2. Add the English translation to `EnglishLocale()`
3. Add the Chinese translation to `ChineseLocale()`
4. Use the string via `ui.LocaleForPreferences(preferences).YourNewKey`

#### Adding a New Language

1. Add a new language constant in `internal/core/rules.go`:
   ```go
   const LanguageSpanish = "es"
   ```
2. Update `NormalizeLanguage()` to include the new language
3. Create a new locale function in `internal/ui/l10n.go`:
   ```go
   func SpanishLocale() Locale { ... }
   ```
4. Update `LocaleForLanguage()` to handle the new language
5. Add the new option to the menu in `menu_bar_bridge.m`

## Technical Notes

- Language detection happens only once on first launch
- Subsequent launches use the saved preference
- Setting language to empty string reverts to auto-detection behavior
- The UI does not automatically update when language is changed (requires app restart)
- All localized strings are compile-time constants (no external translation files)
