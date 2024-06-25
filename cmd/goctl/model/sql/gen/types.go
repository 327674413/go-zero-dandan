package gen

import (
	"fmt"
	"github.com/zeromicro/go-zero/tools/goctl/model/sql/template"
	"github.com/zeromicro/go-zero/tools/goctl/util"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
	"github.com/zeromicro/go-zero/tools/goctl/util/stringx"
)

func genTypes(table Table, methods string, withCache bool) (string, error) {
	fields := table.Fields
	fieldsString, err := genFields(table, fields)
	if err != nil {
		return "", err
	}

	text, err := pathx.LoadTemplate(category, typesTemplateFile, template.Types)
	if err != nil {
		return "", err
	}
	// ------danEditStart------
	constDatabaseFields := ""
	for _, field := range fields {
		constDatabaseFields += fmt.Sprintf("%s_%s dao.TableField = \"%s\"\n", table.Name.ToCamel(), field.Name.ToCamel(), field.NameOriginal)
	}
	// ------danEditEnd------
	output, err := util.With("types").
		Parse(text).
		Execute(map[string]any{
			"withCache":             withCache,
			"method":                methods,
			"upperStartCamelObject": table.Name.ToCamel(),
			"lowerStartCamelObject": stringx.From(table.Name.ToCamel()).Untitle(),
			"fields":                fieldsString,
			"data":                  table,
			// ------danEditStart------
			"constDatabaseFields": constDatabaseFields,
			// ------danEditEnd------
		})
	if err != nil {
		return "", err
	}

	return output.String(), nil
}
