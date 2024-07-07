package strd

import (
	"strconv"
	"strings"
	"unicode"
)

// Int64 字符串转int64
func Int64(str string) int64 {
	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return num
}

// FirstUpper 首字母转大写
func FirstUpper(s string) string {
	if len(s) == 0 {
		return s
	}

	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

// FirstLower 首字母转小写
func FirstLower(s string) string {
	if len(s) == 0 {
		return s
	}

	r := []rune(s)
	r[0] = unicode.ToLower(r[0])
	return string(r)
}

// SnakeToLowerCamel 蛇形转小写驼峰
func SnakeToLowerCamel(input string) string {
	var result string
	words := strings.Split(input, "_")
	for i, word := range words {
		if i == 0 {
			result += strings.ToLower(word)
		} else {
			result += strings.Title(strings.ToLower(word))
		}
	}
	return result
}
