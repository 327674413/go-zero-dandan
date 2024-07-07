package gen

import (
	"github.com/zeromicro/go-zero/tools/goctl/model/sql/template"
	"github.com/zeromicro/go-zero/tools/goctl/util"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
	"strings"
)

// ------danEditStart------
func toLowerCamelCase(input string) string {
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

// ------danEditEnd------
func genTag(table Table, in string) (string, error) {
	if in == "" {
		return in, nil
	}

	text, err := pathx.LoadTemplate(category, tagTemplateFile, template.Tag)
	if err != nil {
		return "", err
	}
	output, err := util.With("tag").Parse(text).Execute(map[string]any{
		"field": in,
		// ------danEditStart------
		"jsonField": toLowerCamelCase(in),
		// ------danEditEnd------
		"data": table,
	})
	if err != nil {
		return "", err
	}

	return output.String(), nil
}
