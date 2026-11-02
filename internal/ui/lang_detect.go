package ui

import (
	"os"
	"strings"
	"sync"

	"github.com/Digital-Shane/open-tamer/internal/config"
	"github.com/Digital-Shane/open-tamer/internal/core"
)

func DetectLanguage() string {
	language := strings.ToLower(os.Getenv("LANG"))
	if language == "" {
		language = strings.ToLower(os.Getenv("LC_ALL"))
	}
	if strings.HasPrefix(language, "zh") {
		return core.LanguageChinese
	}
	return core.LanguageEnglish
}

var systemLanguage struct {
	once  sync.Once
	value string
}

func DetectSystemLanguage() string {
	systemLanguage.once.Do(func() {
		systemLanguage.value = config.DetectLanguage()
	})
	return systemLanguage.value
}
