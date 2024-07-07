package parsed

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"os"
	"strings"
)

type GoStructField struct {
	Name    string
	Type    string
	Comment string
}

func GetGoStructFieldsByCode(code string, structName string) (fields []*GoStructField) {
	fields = make([]*GoStructField, 0)
	// 解析代码字符串
	fs := token.NewFileSet()
	node, err := parser.ParseFile(fs, "", code, parser.ParseComments)
	if err != nil {
		return
	}
	return getGoStructFields(node, structName)
}
func GetGoStructFieldsByFile(filePath string, structName string) (fields []*GoStructField) {
	fields = make([]*GoStructField, 0)
	// 打开并读取Go文件
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()
	// 解析Go文件
	fs := token.NewFileSet()
	node, err := parser.ParseFile(fs, filePath, nil, parser.ParseComments)
	if err != nil {
		return
	}
	return getGoStructFields(node, structName)
}

func getGoStructFields(node *ast.File, structName string) (fields []*GoStructField) {
	fields = make([]*GoStructField, 0)
	// 遍历AST节点
	ast.Inspect(node, func(n ast.Node) bool {
		// 查找类型声明节点
		ts, ok := n.(*ast.TypeSpec)
		if !ok {
			return true
		}

		// 匹配结构体名称
		if ts.Name.Name != structName {
			return true
		}

		// 查找结构体类型
		structType, ok := ts.Type.(*ast.StructType)
		if !ok {
			return true
		}

		// 遍历结构体字段
		for _, field := range structType.Fields.List {
			// 跳过非导出字段（小写开头的字段）
			if !field.Names[0].IsExported() {
				continue
			}

			// 获取字段名
			fieldName := field.Names[0].Name
			// 获取字段类型
			// 获取字段类型名称
			var fieldType string
			if ident, ok := field.Type.(*ast.Ident); ok {
				fieldType = ident.Name
			} else if starExpr, ok := field.Type.(*ast.StarExpr); ok {
				if ident, ok := starExpr.X.(*ast.Ident); ok {
					fieldType = "*" + ident.Name
				}
			} else {
				fieldType = types.ExprString(field.Type)
			}

			// 获取字段注释
			fieldComment := ""
			if field.Comment != nil {
				fieldComment = strings.TrimSpace(field.Comment.Text())
			}

			// 添加字段信息到结果列表
			fields = append(fields, &GoStructField{
				Name:    fieldName,
				Type:    fieldType,
				Comment: fieldComment,
			})
		}
		return false
	})
	return
}
