package utild

import (
	"regexp"
)

func CheckIsPhone(phone string) bool {
	// 定义一个正则表达式，用于判断手机号格式是否正确
	phoneRegexp := regexp.MustCompile(`^1[3-9]\d{9}$`)

	// 判断一个字符串是否为手机号
	if phoneRegexp.MatchString(phone) {
		return true
	} else {
		return false
	}
}
