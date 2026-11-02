package ui

import (
	"fmt"

	"github.com/Digital-Shane/open-tamer/internal/core"
)

type Locale struct {
	Language                    string `json:"language"`
	StatusStarting              string `json:"statusStarting"`
	StatusAppsObservedFormat    string `json:"statusAppsObserved"`
	SystemCPU                   string `json:"systemCPU"`
	ProcessCPU                  string `json:"processCPU"`
	SystemCPULabel              string `json:"systemCPULabel"`
	ProcessCPULabel             string `json:"processCPULabel"`
	SystemCPUShort              string `json:"systemCPUShort"`
	ProcessCPUShort             string `json:"processCPUShort"`
	CurrentSuffix               string `json:"currentSuffix"`
	AverageSuffix               string `json:"averageSuffix"`
	AllProcesses                string `json:"allProcesses"`
	AllProcessesShort           string `json:"allProcessesShort"`
	TrackedProcesses            string `json:"trackedProcesses"`
	ManagedProcesses            string `json:"managedProcesses"`
	ManagedApps                 string `json:"managedApps"`
	NoProcessSamples            string `json:"noProcessSamples"`
	NoTrackedProcesses          string `json:"noTrackedProcesses"`
	NoProcessesObserved         string `json:"noProcessesObserved"`
	NoManagedApps               string `json:"noManagedApps"`
	SystemOther                 string `json:"systemOther"`
	SystemOtherNote             string `json:"systemOtherNote"`
	Unknown                     string `json:"unknown"`
	UnknownProcess              string `json:"unknownProcess"`
	Note                        string `json:"note"`
	HighCPUUsage                string `json:"highCPUUsage"`
	HighCPUBodyFormat           string `json:"highCPUBody"`
	Preferences                 string `json:"preferences"`
	General                     string `json:"general"`
	LanguageMenuItem            string `json:"languageMenuItem"`
	LanguageAuto                string `json:"languageAuto"`
	LanguageEnglish             string `json:"languageEnglish"`
	LanguageChinese             string `json:"languageChinese"`
	Theme                       string `json:"themeMenuItem"`
	ThemeSystem                 string `json:"themeSystem"`
	ThemeLight                  string `json:"themeLight"`
	ThemeDark                   string `json:"themeDark"`
	RuleTrackOnly               string `json:"ruleTrackOnly"`
	RulePauseInBackground       string `json:"rulePauseInBackground"`
	RuleLowerPriorityBackground string `json:"ruleLowerPriorityBackground"`
	RuleLowerPriorityAlways     string `json:"ruleLowerPriorityAlways"`
	RuleLimitCPU                string `json:"ruleLimitCPU"`
	LimitTargetPrefix           string `json:"limitTargetPrefix"`
	RuleBackgroundSuffix        string `json:"ruleBackgroundSuffix"`
	RuleAlwaysSuffix            string `json:"ruleAlwaysSuffix"`
	RuleHideAfterIdle           string `json:"ruleHideAfterIdle"`
	RuleQuitAfterIdle           string `json:"ruleQuitAfterIdle"`
}

func (locale Locale) StatusAppsObserved(count int) string {
	return fmt.Sprintf(locale.StatusAppsObservedFormat, count)
}

func (locale Locale) HighCPUBody(appName string, cpuPercent float64) string {
	return fmt.Sprintf(locale.HighCPUBodyFormat, appName, cpuPercent)
}

func EnglishLocale() Locale {
	return Locale{
		Language:                    core.LanguageEnglish,
		StatusStarting:              "Starting",
		StatusAppsObservedFormat:    "%d apps observed",
		SystemCPU:                   "System CPU",
		ProcessCPU:                  "Process CPU",
		SystemCPULabel:              "System CPU",
		ProcessCPULabel:             "Process CPU",
		SystemCPUShort:              "system CPU",
		ProcessCPUShort:             "process CPU",
		CurrentSuffix:               "current",
		AverageSuffix:               "60s avg",
		AllProcesses:                "All Processes",
		AllProcessesShort:           "All",
		TrackedProcesses:            "Tracked Processes",
		ManagedProcesses:            "Managed Processes",
		ManagedApps:                 "Managed Apps",
		NoProcessSamples:            "No process samples yet",
		NoTrackedProcesses:          "No tracked processes configured",
		NoProcessesObserved:         "No processes observed yet",
		NoManagedApps:               "No app rules configured",
		SystemOther:                 "System / Other",
		SystemOtherNote:             "Unlisted processes and system CPU not attributed to sampled apps",
		Unknown:                     "unknown",
		UnknownProcess:              "Unknown",
		Note:                        "Note:",
		HighCPUUsage:                "High CPU usage",
		HighCPUBodyFormat:           "%s is using %.0f%% CPU.",
		Preferences:                 "Preferences",
		General:                     "General",
		LanguageMenuItem:            "Language",
		LanguageAuto:                "Auto (system language)",
		LanguageEnglish:             "English",
		LanguageChinese:             "中文",
		Theme:                       "Theme",
		ThemeSystem:                 "System theme",
		ThemeLight:                  "Light",
		ThemeDark:                   "Dark",
		RuleTrackOnly:               "Track Only",
		RulePauseInBackground:       "Pause in Background",
		RuleLowerPriorityBackground: "Lower Priority in Background",
		RuleLowerPriorityAlways:     "Lower Priority Always",
		RuleLimitCPU:                "Limit CPU",
		LimitTargetPrefix:           "to",
		RuleBackgroundSuffix:        " in Background",
		RuleAlwaysSuffix:            " Always",
		RuleHideAfterIdle:           "Hide After Idle",
		RuleQuitAfterIdle:           "Quit After Idle",
	}
}

func ChineseLocale() Locale {
	return Locale{
		Language:                    core.LanguageChinese,
		StatusStarting:              "正在启动",
		StatusAppsObservedFormat:    "已观察 %d 个应用",
		SystemCPU:                   "系统 CPU",
		ProcessCPU:                  "进程 CPU",
		SystemCPULabel:              "系统 CPU",
		ProcessCPULabel:             "进程 CPU",
		SystemCPUShort:              "系统 CPU",
		ProcessCPUShort:             "进程 CPU",
		CurrentSuffix:               "当前",
		AverageSuffix:               "60秒平均",
		AllProcesses:                "所有进程",
		AllProcessesShort:           "全部",
		TrackedProcesses:            "已跟踪进程",
		ManagedProcesses:            "受管进程",
		ManagedApps:                 "受管应用",
		NoProcessSamples:            "暂无进程样本",
		NoTrackedProcesses:          "暂无已跟踪进程",
		NoProcessesObserved:         "暂未观察到进程",
		NoManagedApps:               "暂无应用规则",
		SystemOther:                 "系统/其他",
		SystemOtherNote:             "未列出进程及未归属到采样应用的系统 CPU",
		Unknown:                     "未知",
		UnknownProcess:              "未知",
		Note:                        "备注：",
		HighCPUUsage:                "CPU 占用过高",
		HighCPUBodyFormat:           "%s 的 CPU 占用为 %.0f%%。",
		Preferences:                 "偏好设置",
		General:                     "通用",
		LanguageMenuItem:            "界面语言",
		LanguageAuto:                "自动（跟随系统语言）",
		LanguageEnglish:             "English",
		LanguageChinese:             "中文",
		Theme:                       "外观",
		ThemeSystem:                 "跟随系统",
		ThemeLight:                  "浅色",
		ThemeDark:                   "深色",
		RuleTrackOnly:               "仅跟踪",
		RulePauseInBackground:       "后台暂停",
		RuleLowerPriorityBackground: "后台降低优先级",
		RuleLowerPriorityAlways:     "始终降低优先级",
		RuleLimitCPU:                "限制 CPU",
		LimitTargetPrefix:           "至",
		RuleBackgroundSuffix:        "（后台）",
		RuleAlwaysSuffix:            "（始终）",
		RuleHideAfterIdle:           "空闲后隐藏",
		RuleQuitAfterIdle:           "空闲后退出",
	}
}

func LocaleForLanguage(language string) Locale {
	if language == "" {
		language = DetectSystemLanguage()
	}
	if core.NormalizeLanguage(language) == core.LanguageChinese {
		return ChineseLocale()
	}
	return EnglishLocale()
}

func LocaleForPreferences(preferences core.GlobalPreferences) Locale {
	language := preferences.Language
	if language == "" {
		language = DetectSystemLanguage()
	}
	return LocaleForLanguage(language)
}
