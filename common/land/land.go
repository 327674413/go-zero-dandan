package land

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var LangCurr string
var LangAccept map[string]bool
var bundle *i18n.Bundle

func init() {
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.MustLoadMessageFile("../../../common/land/en_us.toml")
	bundle.MustLoadMessageFile("../../../common/land/zh_cn.toml")
	LangAccept = map[string]bool{
		"en_us": true,
		"zh_cn": true,
	}
	LangCurr = "zh_cn"
	fmt.Println("-----------------------i18n Init --------------------")

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
