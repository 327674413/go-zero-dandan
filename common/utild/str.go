package utild

import "unicode/utf8"

func Strlen(str string) int {
	return utf8.RuneCountInString(str)
}
