package l10n

type Language string

const (
	LanguageEnglish Language = "en"
	LanguageChinese Language = "zh"
)

func NormalizeLanguage(language string) Language {
	if language == string(LanguageChinese) {
		return LanguageChinese
	}
	return LanguageEnglish
}
