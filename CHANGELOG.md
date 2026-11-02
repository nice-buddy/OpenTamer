# Changelog

All notable changes to OpenTamer will be documented in this file.

## [1.5.0] - 2026-11-02

### Added
- Full internationalization (i18n) support with Chinese (Simplified) and English
- Auto-detect system language on first launch
- Language preference in settings (Auto/English/中文)
- All UI elements, menu items, and alerts now localized
- Added 50%, 75%, and 100% CPU limit presets

### Changed
- Menu bar icon width now dynamically adjusts to fit content (more compact)
- Improved menu bar width calculation for better visual balance

### Removed
- Removed 1% and 0.01% CPU limit presets (rarely used, confusing at low values)

### Technical
- Added comprehensive i18n test coverage
- Localization bridge for all Objective-C UI strings
- Build verification scripts for translation completeness

## [1.4.1] - 2024-12-15

### Fixed
- Release date correction

## [1.4.0] - 2024-12-15

### Added
- Background-only CPU limiting option
- Automatic release of CPU limit when app comes to foreground

## [1.3.0] - 2024-12-14

### Added
- Adaptive system theme support
- Live AppKit hot-reloading for theme changes
- Launch at login preference

## [1.2.0] - 2024-12-13

### Added
- Per-process vs system-normalized CPU display modes
- Preference to choose CPU calculation method

## [1.1.0] - 2024-12-12

### Changed
- Image size optimization in documentation

## [1.0.0] - 2024-12-11

### Added
- Initial public release
- Menu bar CPU monitoring
- App management rules (Track, Lower Priority, Limit CPU, Pause)
- High CPU alerts
- System safety guards
- Local JSON-based configuration
