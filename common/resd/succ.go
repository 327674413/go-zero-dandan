package resd

import "github.com/nicksnyder/go-i18n/v2/i18n"

type SuccInfo struct {
	Result bool `json:"result"`
	Code   int  `json:"code"`
	Data   any  `json:"data"`
}

func Succ(data any) *SuccInfo {
	return &SuccInfo{Result: true, Code: Ok, Data: data}
}
func SuccAsync(data any) *SuccInfo {
	return &SuccInfo{Result: true, Code: OkAsync, Data: data}
}

func SuccCode(localize *i18n.Localizer, succCode int, tempData ...[]string) *SuccInfo {
	var tempD []string
	if len(tempData) > 0 {
		tempD = tempData[0]
	}
	text := Msg(localize, succCode, tempD)
	return &SuccInfo{Result: true, Code: 200, Data: map[string]string{"msg": text}}
}
