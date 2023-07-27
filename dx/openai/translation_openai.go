package dx_openai

import (
	"fmt"

	dx_user "dongxian.com/http_link/dx/user"
)

const promptStr = ""

var _targetLanguages = []string{"en", "es", "fr", "ru", "de", "zh-tw", "ko", "ja"}

// 翻译文本为所有目标信息，然后返回 语言使用 ^^^^ 间隔开
func TranslateText(msg string, targetLanguages []string) string {
	translateStr := ""
	if targetLanguages == nil || len(targetLanguages) == 0 {
		targetLanguages = _targetLanguages
	}
	len := len(targetLanguages)
	for i := 0; i < len; i++ {
		var language = targetLanguages[i]
		translateStr += translateOne(msg, language)
		if i < len-1 {
			translateStr += "^^^^"
		}
	}
	return translateStr
}

// 翻译为目标语言
func translateOne(text string, targetLanguage string) string {
	// 'auto'
	prompt := fmt.Sprintf("Translate the following text from Chinese to '%s': '%s'", text, targetLanguage)
	return Req(prompt)
}

// 设置语言类型
func SetLanguages(user string, lgs []string) {
	dx_user.SetLanguage(user, lgs)
}
