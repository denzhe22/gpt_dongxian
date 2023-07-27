// 定义使用者
package dx_user

var languagesMap = make(map[string][]string)

func SetLanguage(user string, languages []string) {
	languagesMap[user] = languages
}

func GetLanguage(user string) []string {
	return languagesMap[user]
}

// 是否已经设置过了
func HasSeted(user string) bool {
	return GetLanguage(user) != nil && len(GetLanguage(user)) > 0
}
