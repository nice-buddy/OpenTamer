package config_test

import (
	"os"
	"path/filepath"
	"testing"
	
	"github.com/Digital-Shane/open-tamer/internal/config"
	"github.com/Digital-Shane/open-tamer/internal/core"
)

func TestLanguageDetectionAndPersistence(t *testing.T) {
	// Create a temporary config directory
	tmpDir := filepath.Join(os.TempDir(), "opentamer-lang-test")
	os.RemoveAll(tmpDir)
	defer os.RemoveAll(tmpDir)
	
	store := config.NewStore(tmpDir)
	
	// Test 1: First load should detect and persist language
	cfg, err := store.LoadConfig()
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}
	
	if cfg.Preferences.Language == "" {
		t.Error("Expected language to be detected and set, got empty string")
	}
	
	detectedLang := cfg.Preferences.Language
	t.Logf("Detected language: %s", detectedLang)
	
	// Test 2: Subsequent loads should preserve the language
	cfg2, err := store.LoadConfig()
	if err != nil {
		t.Fatalf("Failed to reload config: %v", err)
	}
	
	if cfg2.Preferences.Language != detectedLang {
		t.Errorf("Expected language %s to be preserved, got %s", detectedLang, cfg2.Preferences.Language)
	}
	
	// Test 3: Manual language change should persist
	cfg2.Preferences.Language = core.LanguageChinese
	if err := store.SaveConfig(cfg2); err != nil {
		t.Fatalf("Failed to save config: %v", err)
	}
	
	cfg3, err := store.LoadConfig()
	if err != nil {
		t.Fatalf("Failed to reload config after change: %v", err)
	}
	
	if cfg3.Preferences.Language != core.LanguageChinese {
		t.Errorf("Expected language to be %s, got %s", core.LanguageChinese, cfg3.Preferences.Language)
	}
	
	// Test 4: Setting to empty string (auto-detect) should be preserved
	cfg3.Preferences.Language = ""
	if err := store.SaveConfig(cfg3); err != nil {
		t.Fatalf("Failed to save config with empty language: %v", err)
	}
	
	cfg4, err := store.LoadConfig()
	if err != nil {
		t.Fatalf("Failed to reload config: %v", err)
	}
	
	// Empty string means "auto" - it's resolved at UI render time, not persisted
	if cfg4.Preferences.Language != "" {
		t.Errorf("Expected language to remain empty (auto), got %s", cfg4.Preferences.Language)
	}
}

func TestLanguageNormalization(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"en", core.LanguageEnglish},
		{"zh", core.LanguageChinese},
		{"EN", ""},  // Case-sensitive
		{"ZH", ""},  // Case-sensitive
		{"zh-CN", ""},
		{"zh-Hans", ""},
		{"fr", ""},
		{"", ""},
	}
	
	for _, tt := range tests {
		result := core.NormalizeLanguage(tt.input)
		if result != tt.expected {
			t.Errorf("NormalizeLanguage(%q) = %q, expected %q", tt.input, result, tt.expected)
		}
	}
}

func TestDetectLanguage(t *testing.T) {
	detected := config.DetectLanguage()
	
	// Should return either English or Chinese
	if detected != core.LanguageEnglish && detected != core.LanguageChinese {
		t.Errorf("DetectLanguage() = %q, expected either %q or %q", 
			detected, core.LanguageEnglish, core.LanguageChinese)
	}
	
	t.Logf("System language detected as: %s", detected)
}
