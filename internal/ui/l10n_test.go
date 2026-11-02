package ui

import (
	"strings"
	"testing"

	"github.com/Digital-Shane/open-tamer/internal/core"
)

func TestLocaleForPreferencesDefaultsToEnglish(t *testing.T) {
	if got := LocaleForPreferences(core.GlobalPreferences{Language: core.LanguageEnglish}); got != EnglishLocale() {
		t.Fatalf("locale = %#v, want English", got)
	}
	if got := LocaleForPreferences(core.GlobalPreferences{Language: core.LanguageChinese}); got != ChineseLocale() {
		t.Fatalf("locale = %#v, want Chinese", got)
	}
	if got := LocaleForPreferences(core.GlobalPreferences{Language: "fr"}); got != EnglishLocale() {
		t.Fatalf("unsupported language should fall back to English, got %#v", got)
	}
}

func TestChineseLocaleCoversUIStrings(t *testing.T) {
	locale := ChineseLocale()
	if !strings.Contains(locale.StatusAppsObserved(3), "3") {
		t.Fatal("Chinese status should include count")
	}
	if locale.AllProcesses == "" || locale.Preferences == "" || locale.Language == "" {
		t.Fatal("Chinese locale should cover UI strings")
	}
	if got := locale.HighCPUBody("Safari", 87.6); got != "Safari 的 CPU 占用为 88%。" {
		t.Fatalf("high CPU body = %q", got)
	}
}

func TestDetectLanguageUsesEnvironment(t *testing.T) {
	t.Setenv("LANG", "zh_CN.UTF-8")
	t.Setenv("LC_ALL", "")
	if got := DetectLanguage(); got != core.LanguageChinese {
		t.Fatalf("DetectLanguage() = %q, want zh", got)
	}
	t.Setenv("LANG", "en_US.UTF-8")
	if got := DetectLanguage(); got != core.LanguageEnglish {
		t.Fatalf("DetectLanguage() = %q, want en", got)
	}
	t.Setenv("LANG", "")
	t.Setenv("LC_ALL", "C")
	if got := DetectLanguage(); got != core.LanguageEnglish {
		t.Fatalf("DetectLanguage() = %q, want en", got)
	}
}
