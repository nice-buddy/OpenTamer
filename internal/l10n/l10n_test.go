package l10n

import "testing"

func TestNormalizeLanguage(t *testing.T) {
	if got := NormalizeLanguage("zh"); got != LanguageChinese {
		t.Fatalf("NormalizeLanguage(\"zh\") = %q, want %q", got, LanguageChinese)
	}
	if got := NormalizeLanguage("en"); got != LanguageEnglish {
		t.Fatalf("NormalizeLanguage(\"en\") = %q, want %q", got, LanguageEnglish)
	}
	if got := NormalizeLanguage("fr"); got != LanguageEnglish {
		t.Fatalf("NormalizeLanguage(\"fr\") = %q, want %q", got, LanguageEnglish)
	}
	if got := NormalizeLanguage(""); got != LanguageEnglish {
		t.Fatalf("NormalizeLanguage(\"\") = %q, want %q", got, LanguageEnglish)
	}
}
