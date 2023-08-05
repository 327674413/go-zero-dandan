package land

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"go-zero-dandan/common/resd"
	"golang.org/x/text/language"
)

var LangCurr string
var LangAccept map[string]bool
var bundle *i18n.Bundle

func init() {
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	var err error
	//这里如何结偶未想明白
	//_, err = bundle.LoadMessageFile("../../../common/land/en_us.toml")
	//_, err = bundle.LoadMessageFile("../../../common/land/zh_cn.toml")
	bundle.MustLoadMessageFile("../../../common/land/en_us.toml")
	bundle.MustLoadMessageFile("../../../common/land/zh_cn.toml")
	LangAccept = map[string]bool{
		"en_us": true,
		"zh_cn": true,
	}
	LangCurr = "zh_cn"
	if err == nil {
		fmt.Println("-----------------------i18n Init --------------------")
	} else {
		fmt.Println("-----------------------i18n Load lang fail:", err, " --------------------")
	}

}
func Set(lang string) *i18n.Localizer {
	currLang := "zh_cn"
	_, ok := LangAccept[lang]
	if ok {
		currLang = lang
	}
	return i18n.NewLocalizer(bundle, currLang)
}

func Trans(localize *i18n.Localizer, temp string, tempData ...map[string]string) string {
	var data map[string]string
	if len(tempData) > 0 {
		data = tempData[0]
	} else {
		data = map[string]string{}
	}
	str, _, err := localize.LocalizeWithTag(&i18n.LocalizeConfig{
		MessageID: temp,
		DefaultMessage: &i18n.Message{
			ID: temp,
		},
		TemplateData: data,
	})
	if err != nil {
		return temp
	} else {
		return str
	}
	return str
	/*return localize.MustLocalize(&i18n.LocalizeConfig{
		MessageID: temp,
		DefaultMessage: &i18n.Message{
			ID: temp,
		},
		TemplateData: data,
	})*/
}
func Msg(localize *i18n.Localizer, msgCode int, tempDataArr ...[]string) string {
	tempData := make([]string, 0)
	if len(tempDataArr) > 0 {
		tempData = tempDataArr[0]
	}
	m := make(map[string]string)
	for i, v := range tempData {
		key := "Field" + fmt.Sprint(i+1)
		m[key] = getMsg(localize, v)
	}
	if code, ok := resd.Msg[msgCode]; ok {
		return Trans(localize, code, m)
	} else {
		return Trans(localize, resd.Msg[resd.SysErr], m)
	}

}

func getMsg(localize *i18n.Localizer, tempCode string, tempData ...map[string]string) string {
	return Trans(localize, tempCode, tempData...)
}
func SuccCode(localize *i18n.Localizer, succCode int, tempData ...[]string) *resd.SuccInfo {
	var tempD []string
	if len(tempData) > 0 {
		tempD = tempData[0]
	}
	text := Msg(localize, succCode, tempD)
	return &resd.SuccInfo{Result: true, Code: 200, Data: map[string]string{"msg": text}}
}
