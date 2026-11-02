package ui

import "github.com/Digital-Shane/open-tamer/internal/core"

type bridgeString struct {
	english string
	chinese string
}

var bridgeStrings = map[string]bridgeString{
	"status_off":            {"Off", "关闭"},
	"refresh":               {"Refresh", "刷新"},
	"top_processes":         {"Top Processes", "Top 进程"},
	"tracked_processes":     {"Tracked Processes", "已跟踪进程"},
	"managed_processes":     {"Managed Processes", "受管进程"},
	"all_processes":         {"All Processes", "所有进程"},
	"all_processes_count":   {"All Processes (%lu) >", "所有进程（%lu）>"},
	"all_processes_menu":    {"All Processes (%lu, current %@)", "所有进程（%lu，当前 %@）"},
	"top_current":           {"Top Processes (current %@)", "Top 进程（当前 %@）"},
	"top_average":           {"Top Processes (60s avg %@)", "Top 进程（60 秒平均 %@）"},
	"tracked_current":       {"Tracked Processes (current %@)", "已跟踪进程（当前 %@）"},
	"managed_apps_current":  {"Managed Apps (current %@)", "受管应用（当前 %@）"},
	"empty_process_samples": {"No process samples yet", "暂无进程样本"},
	"empty_tracked":         {"No tracked processes configured", "暂无已跟踪进程"},
	"empty_tracked_panel":   {"No tracked processes", "暂无已跟踪进程"},
	"empty_observed":        {"No processes observed yet", "暂未观察到进程"},
	"empty_rules":           {"No app rules configured", "暂无应用规则"},
	"empty_managed":         {"No managed processes", "暂无受管进程"},
	"graph_hidden":          {"All graph lines hidden", "已隐藏全部图表曲线"},
	"graph_collecting":      {"Collecting CPU history", "正在收集 CPU 历史"},
	"system_cpu_summary":    {"System CPU: %@", "系统 CPU：%@"},
	"system_cpu_headline":   {"System CPU %@   %@", "系统 CPU %@   %@"},
	"system_cpu_label":      {"System CPU", "系统 CPU"},
	"process_cpu_label":     {"Process CPU", "进程 CPU"},
	"system_cpu_short":      {"system CPU", "系统 CPU"},
	"process_cpu_short":     {"process CPU", "进程 CPU"},
	"cpu_short":             {"CPU", "CPU"},
	"avg":                   {"Avg", "均值"},
	"app_fallback":          {"App", "应用"},
	"unknown":               {"unknown", "未知"},
	"unknown_process":       {"Unknown", "未知"},
	"rule_fallback":         {"Rule", "规则"},
	"status_prefix":         {"Status: %@", "状态：%@"},
	"note_prefix":           {"Note: %@", "备注：%@"},
	"rule_prefix":           {"Rule: %@", "规则：%@"},
	"label_value":           {"%@: %@", "%@：%@"},
	"untrack":               {"Untrack", "取消跟踪"},
	"pause_background":      {"Pause in Background", "后台暂停"},
	"limit_cpu":             {"Limit CPU", "限制 CPU"},
	"limit_cpu_menu":        {"Limit CPU...", "限制 CPU…"},
	"in_background":         {"In Background", "后台"},
	"always":                {"Always", "始终"},
	"custom":                {"Custom...", "自定义…"},
	"lower_priority":        {"Lower Priority", "降低优先级"},
	"lower_priority_menu":   {"Lower Priority...", "降低优先级…"},
	"track_in":              {"Track in", "跟踪到"},
	"track_in_menu":         {"Track in...", "跟踪到…"},
	"track_menu_bar":        {"The Menu Bar", "菜单栏"},
	"remove_rule":           {"Remove Rule", "移除规则"},
	"preferences":           {"Preferences", "偏好设置"},
	"general":               {"General", "通用"},
	"stats_graph":           {"Stats & Graph", "统计与图表"},
	"high_cpu_alerts":       {"High CPU Alerts", "高 CPU 提醒"},
	"policy_guards":         {"Policy Guards", "策略防护"},
	"show_menu_icon":        {"Show Menu Icon", "显示菜单栏图标"},
	"aggregate_by_name":     {"Aggregate By Name", "按名称聚合"},
	"cpu_display":           {"CPU Display", "CPU 显示"},
	"per_core_cpu":          {"Per-Core Process CPU", "按核心进程 CPU"},
	"system_normalized_cpu": {"System Normalized CPU", "系统归一化 CPU"},
	"theme":                 {"Theme", "外观"},
	"theme_system":          {"System theme", "跟随系统"},
	"theme_light":           {"Light", "浅色"},
	"theme_dark":            {"Dark", "深色"},
	"wake_grace":            {"Wake Grace", "唤醒宽限"},
	"stats_interval":        {"Stats Interval", "统计间隔"},
	"averaging_window":      {"Averaging Window", "平均窗口"},
	"cpu_graph_window":      {"CPU Graph Window", "CPU 图表窗口"},
	"top_sort":              {"Top Processes Sort", "Top 进程排序"},
	"current_cpu":           {"Current CPU", "当前 CPU"},
	"average_cpu":           {"60s Average", "60 秒平均"},
	"detection_enabled":     {"Detection Enabled", "启用检测"},
	"threshold":             {"Threshold", "阈值"},
	"alert_after":           {"Alert After", "提醒延迟"},
	"cooldown":              {"Cooldown", "冷却时间"},
	"disable_ac_battery":    {"Disable On AC Battery Above", "接通电源且电量高于时停用"},
	"disable_idle":          {"Disable When User Idle Longer Than", "用户空闲超过时停用"},
	"reset_defaults":        {"Reset Defaults", "恢复默认设置"},
	"quit_app":              {"Quit OpenTamer", "退出 OpenTamer"},
	"language":              {"Language", "界面语言"},
	"lang_auto":             {"Auto (system language)", "自动（跟随系统语言）"},
	"lang_english":          {"English", "English"},
	"lang_chinese":          {"中文", "中文"},
	"launch_at_login":       {"Launch at Login", "登录时启动"},
	"approve_login":         {"Approve OpenTamer in System Settings to launch it at login.", "请在系统设置中允许 OpenTamer 在登录时启动。"},
	"limit_bg_title":        {"Limit CPU in Background", "后台限制 CPU"},
	"limit_always_title":    {"Limit CPU Always", "始终限制 CPU"},
	"limit_prompt":          {"Enter a maximum CPU percentage. Minimum is 0.01%.", "输入最大 CPU 百分比，最低 0.01%。"},
	"apply":                 {"Apply", "应用"},
	"cancel":                {"Cancel", "取消"},
	"limit_not_applied":     {"CPU limit not applied", "未应用 CPU 限制"},
	"limit_min":             {"Enter a value of 0.01% or higher.", "请输入不低于 0.01% 的数值。"},
	"ok":                    {"OK", "好"},
	"threshold_title":       {"High CPU Alert Threshold", "高 CPU 提醒阈值"},
	"threshold_prompt":      {"Enter the system CPU percentage that should trigger alerts.", "输入触发提醒的系统 CPU 百分比。"},
	"threshold_not_applied": {"Alert threshold not applied", "未应用提醒阈值"},
	"threshold_min":         {"Enter a value greater than 0%.", "请输入大于 0% 的数值。"},
	"login_not_updated":     {"Login item not updated", "登录项未更新"},
	"login_error":           {"OpenTamer could not update its Login Items setting.", "OpenTamer 无法更新登录项设置。"},
	"dur_off":               {"Off", "关闭"},
	"dur_3s":                {"3 seconds", "3 秒"},
	"dur_5s":                {"5 seconds", "5 秒"},
	"dur_10s":               {"10 seconds", "10 秒"},
	"dur_30s":               {"30 seconds", "30 秒"},
	"dur_1m":                {"1 minute", "1 分钟"},
	"dur_2m":                {"2 minutes", "2 分钟"},
	"dur_5m":                {"5 minutes", "5 分钟"},
	"dur_10m":               {"10 minutes", "10 分钟"},
	"dur_15m":               {"15 minutes", "15 分钟"},
	"dur_30m":               {"30 minutes", "30 分钟"},
	"dur_1h":                {"1 hour", "1 小时"},
}

func BridgeString(language, key string) string {
	value, ok := bridgeStrings[key]
	if !ok {
		return key
	}
	if core.NormalizeLanguage(language) == core.LanguageChinese {
		return value.chinese
	}
	return value.english
}

type statusLabel struct {
	english string
	chinese string
}

var statusLabels = map[string]statusLabel{
	string(core.AppStatusUnmanaged):        {"Unmanaged", "未管理"},
	string(core.AppStatusObserved):         {"Observed", "已观察"},
	string(core.AppStatusEligible):         {"Eligible", "可接管"},
	string(core.AppStatusWaiting):          {"Waiting", "等待中"},
	string(core.AppStatusPaused):           {"Paused", "已暂停"},
	string(core.AppStatusPriorityLowered):  {"Priority lowered", "已降低优先级"},
	string(core.AppStatusCPULimited):       {"CPU limited", "已限制 CPU"},
	string(core.AppStatusTemporarilyAwake): {"Temporarily awake", "临时唤醒"},
	string(core.AppStatusBlockedBySafety):  {"Blocked by safety", "因安全策略阻止"},
	string(core.AppStatusBlockedByPolicy):  {"Blocked by policy", "因策略阻止"},
	string(core.AppStatusRestoring):        {"Restoring", "恢复中"},
	string(core.AppStatusUnsupported):      {"Unsupported", "不支持"},
}

func StatusLabel(language, status string) string {
	label, ok := statusLabels[status]
	if !ok {
		return status
	}
	if core.NormalizeLanguage(language) == core.LanguageChinese {
		return label.chinese
	}
	return label.english
}

type safetyLabel struct {
	english string
	chinese string
}

var safetyLabels = map[string]safetyLabel{
	string(core.SafetyReasonEssentialSystem):  {"Essential system process", "关键系统进程"},
	string(core.SafetyReasonAccessibility):    {"Accessibility process", "辅助功能进程"},
	string(core.SafetyReasonOpenTamerSelf):    {"OpenTamer process", "OpenTamer 进程"},
	string(core.SafetyReasonPrivilegedHelper): {"OpenTamer privileged helper", "OpenTamer 特权辅助进程"},
	string(core.SafetyReasonProtectedProcess): {"Protected process", "受保护进程"},
	string(core.SafetyReasonRootOwned):        {"Owned by root", "root 所有"},
	string(core.SafetyReasonUnknown):          {"Unknown process", "未知进程"},
	string(core.SafetyReasonSlowOnly):         {"Slow only", "仅可降速"},
}

func SafetyLabel(language, reason string) string {
	label, ok := safetyLabels[reason]
	if !ok {
		return reason
	}
	if core.NormalizeLanguage(language) == core.LanguageChinese {
		return label.chinese
	}
	return label.english
}
