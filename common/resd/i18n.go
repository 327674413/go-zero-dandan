package resd

import (
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"go-zero-dandan/common/fmtd"
	"golang.org/x/text/language"
	"path/filepath"
	"strings"
)

var I18n *DanI18n
var msg map[int]string

type I18nConfig struct {
	LangPathList []string
	DefaultLang  string
}
type DanI18n struct {
	bundle        *i18n.Bundle
	acceptLangMap map[string]bool
	defaultLang   string
}
type Lang struct {
	localizer *i18n.Localizer
}

// NewI18n 使用I18nConfig，其中文件名必须是xxx.toml，xxx会作为NewLang的key
func NewI18n(conf *I18nConfig) (*DanI18n, error) {
	if len(conf.LangPathList) == 0 {
		return nil, errors.New("param LangPathList required")
	}
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	acceptLangMap := make(map[string]bool)
	for _, path := range conf.LangPathList {
		_, err := bundle.LoadMessageFile(path)
		if err != nil {
			return nil, err
		}
		// 获取文件名
		fileName := filepath.Base(path)
		// 去掉后缀
		acceptLangMap[strings.TrimSuffix(fileName, filepath.Ext(fileName))] = true
	}

	return &DanI18n{
		bundle:        bundle,
		acceptLangMap: acceptLangMap,
		defaultLang:   conf.DefaultLang,
	}, nil

}

func (t *DanI18n) NewLang(lang string) *Lang {
	if t == nil {
		fmtd.Error("未配置i18n：请检查入口文件、etc配置、config配置解析")
		return nil
	}
	//测试模版路径错误时的场景
	if _, ok := t.acceptLangMap[lang]; !ok {
		lang = t.defaultLang
	}
	return &Lang{localizer: i18n.NewLocalizer(t.bundle, lang)}
}

// Trans 将模版变量注入模版
func (t *Lang) Trans(temp string, tempData ...map[string]string) string {
	if t == nil {
		return ""
	}
	var data map[string]string
	if len(tempData) > 0 {
		data = tempData[0]
	} else {
		data = map[string]string{}
	}

	str, _, err := t.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
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

// Msg 只对Var开头的内容进行二次转义，*Var来开头不转并去掉*
func (t *Lang) Msg(msgCode int, tempDatas ...string) string {
	tempData := make([]string, 0)
	for _, v := range tempDatas {
		tempData = append(tempData, v)
	}
	m := make(map[string]string)
	for i, v := range tempData {
		key := "Field" + fmt.Sprint(i+1)
		if v == "" {
			m[key] = ""
		} else if len(v) > 4 && v[:4] == "*Var" {
			m[key] = v[1:]
		} else if len(v) > 3 && v[:3] == "Var" {
			m[key] = t.Trans(v)
		} else {
			m[key] = v
		}
	}
	if code, ok := msg[msgCode]; ok {
		return t.Trans(code, m)
	} else {
		return t.Trans(msg[ErrSys], m)
	}

}

//func (t *I18n) SuccCode(localize *i18n.Localizer, succCode int, tempData ...[]string) *resd.SuccInfo {
//	var tempD []string
//	if len(tempData) > 0 {
//		tempD = tempData[0]
//	}
//	text := t.Msg(succCode, tempD)
//	return &resd.SuccInfo{Result: true, Code: 200, Data: map[string]string{"msg": text}}
//}
