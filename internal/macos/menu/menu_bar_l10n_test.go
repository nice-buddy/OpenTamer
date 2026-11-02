package menu

import (
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/Digital-Shane/open-tamer/internal/ui"
)

func TestMenuBridgeLocalizesEveryUserVisibleString(t *testing.T) {
	source := readMenuBridgeSource(t)

	localizedKey := regexp.MustCompile(`OpenTamerLocalized\(@"([^"]+)"\)`)
	matches := localizedKey.FindAllStringSubmatch(source, -1)
	if len(matches) < 50 {
		t.Fatalf("expected the menu bridge to localize many strings, found %d", len(matches))
	}
	seen := make(map[string]bool, len(matches))
	for _, match := range matches {
		key := match[1]
		if seen[key] {
			continue
		}
		seen[key] = true
		if got := ui.BridgeString("zh", key); got == key {
			t.Errorf("missing Chinese translation for key %q", key)
		}
		if got := ui.BridgeString("en", key); got == "" {
			t.Errorf("missing English translation for key %q", key)
		}
	}
}

func TestMenuBridgeHasNoHardCodedUserVisibleTitles(t *testing.T) {
	source := readMenuBridgeSource(t)

	patterns := []struct {
		name string
		re   *regexp.Regexp
	}{
		{"menu item title", regexp.MustCompile(`initWithTitle:@"([A-Za-z][^"]*)"`)},
		{"alert button", regexp.MustCompile(`addButtonWithTitle:@"([A-Za-z][^"]*)"`)},
		{"alert message", regexp.MustCompile(`messageText = @"([A-Za-z][^"]*)"`)},
		{"alert detail", regexp.MustCompile(`informativeText = @"([A-Za-z][^"]*)"`)},
		{"disabled item", regexp.MustCompile(`disabledItemWithTitle:@"([A-Za-z][^"]*)"`)},
		{"command item", regexp.MustCompile(`commandItemWithTitle:@"([A-Za-z][^"]*)"`)},
		{"label title", regexp.MustCompile(`title:@"([A-Za-z][^"]*)"`)},
		{"formatted title", regexp.MustCompile(`stringWithFormat:@"([A-Za-z][^"]*)"`)},
		{"tooltip", regexp.MustCompile(`toolTip = @"([A-Za-z][^"]*)"`)},
	}

	for _, pattern := range patterns {
		for _, match := range pattern.re.FindAllStringSubmatch(source, -1) {
			value := match[1]
			if value == "OpenTamer" {
				continue
			}
			if strings.HasPrefix(value, "pref-") ||
				strings.HasPrefix(value, "rule|") ||
				strings.HasPrefix(value, "disable-rule|") ||
				strings.HasPrefix(value, "graph-window|") {
				continue
			}
			t.Errorf("%s is not localized: %q", pattern.name, value)
		}
	}
}

func TestLanguagePreferenceIncludesAutoOption(t *testing.T) {
	source := readMenuBridgeSource(t)

	if !strings.Contains(source, `labels:@[OpenTamerLocalized(@"lang_auto"), OpenTamerLocalized(@"lang_english"), OpenTamerLocalized(@"lang_chinese")]`) {
		t.Fatal("language preference should expose auto, English, and Chinese choices")
	}
	if !strings.Contains(source, `values:@[@"", @"en", @"zh"]`) {
		t.Fatal("auto language choice should persist an empty value")
	}
	start := strings.Index(source, "- (void)addStringPreferenceWithTitle:")
	end := strings.Index(source, "- (void)addNullableFloatPreferenceWithTitle:")
	if start < 0 || end <= start {
		t.Fatal("could not locate string preference helper")
	}
	stringPrefBlock := source[start:end]
	if strings.Contains(stringPrefBlock, "value.length == 0") {
		t.Fatal("string preference helper must allow empty auto values")
	}
}

func readMenuBridgeSource(t *testing.T) string {
	t.Helper()
	payload, err := os.ReadFile("menu_bar_bridge.m")
	if err != nil {
		t.Fatalf("read menu bridge: %v", err)
	}
	return string(payload)
}
