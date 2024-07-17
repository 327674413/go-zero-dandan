package gogen

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"text/template"

	"github.com/zeromicro/go-zero/core/collection"
	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	"github.com/zeromicro/go-zero/tools/goctl/api/util"
	"github.com/zeromicro/go-zero/tools/goctl/pkg/golang"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
)

type fileGenConfig struct {
	dir             string
	subdir          string
	filename        string
	templateName    string
	category        string
	templateFile    string
	builtinTemplate string
	data            any
}

func genFile(c fileGenConfig) error {
	fp, created, err := util.MaybeCreateFile(c.dir, c.subdir, c.filename)
	if err != nil {
		return err
	}
	if !created {
		return nil
	}
	defer fp.Close()

	var text string
	if len(c.category) == 0 || len(c.templateFile) == 0 {
		text = c.builtinTemplate
	} else {
		text, err = pathx.LoadTemplate(c.category, c.templateFile, c.builtinTemplate)
		if err != nil {
			return err
		}
	}

	t := template.Must(template.New(c.templateName).Parse(text))
	buffer := new(bytes.Buffer)
	err = t.Execute(buffer, c.data)
	if err != nil {
		return err
	}

	code := golang.FormatCode(buffer.String())
	_, err = fp.WriteString(code)
	return err
}

func writeProperty(writer io.Writer, name, tag, comment string, tp spec.Type, indent int) error {
	util.WriteIndent(writer, indent)
	var err error
	if len(comment) > 0 {
		comment = strings.TrimPrefix(comment, "//")
		comment = "//" + comment
		_, err = fmt.Fprintf(writer, "%s %s %s %s\n", strings.Title(name), tp.Name(), tag, comment)
	} else {
		_, err = fmt.Fprintf(writer, "%s %s %s\n", strings.Title(name), tp.Name(), tag)
	}

	return err
}

func getAuths(api *spec.ApiSpec) []string {
	authNames := collection.NewSet()
	for _, g := range api.Service.Groups {
		jwt := g.GetAnnotation("jwt")
		if len(jwt) > 0 {
			authNames.Add(jwt)
		}
	}
	return authNames.KeysStr()
}

func getJwtTrans(api *spec.ApiSpec) []string {
	jwtTransList := collection.NewSet()
	for _, g := range api.Service.Groups {
		jt := g.GetAnnotation(jwtTransKey)
		if len(jt) > 0 {
			jwtTransList.Add(jt)
		}
	}
	return jwtTransList.KeysStr()
}

func getMiddleware(api *spec.ApiSpec) []string {
	result := collection.NewSet()
	for _, g := range api.Service.Groups {
		middleware := g.GetAnnotation("middleware")
		if len(middleware) > 0 {
			for _, item := range strings.Split(middleware, ",") {
				result.Add(strings.TrimSpace(item))
			}
		}
	}

	return result.KeysStr()
}

func responseGoTypeName(r spec.Route, pkg ...string) string {
	if r.ResponseType == nil {
		return ""
	}

	resp := golangExpr(r.ResponseType, pkg...)
	switch r.ResponseType.(type) {
	case spec.DefineStruct:
		if !strings.HasPrefix(resp, "*") {
			return "*" + resp
		}
	}

	return resp
}

func requestGoTypeName(r spec.Route, pkg ...string) string {
	if r.RequestType == nil {
		return ""
	}

	return golangExpr(r.RequestType, pkg...)
}

func golangExpr(ty spec.Type, pkg ...string) string {
	switch v := ty.(type) {
	case spec.PrimitiveType:
		return v.RawName
	case spec.DefineStruct:
		if len(pkg) > 1 {
			panic("package cannot be more than 1")
		}

		if len(pkg) == 0 {
			return v.RawName
		}

		return fmt.Sprintf("%s.%s", pkg[0], strings.Title(v.RawName))
	case spec.ArrayType:
		if len(pkg) > 1 {
			panic("package cannot be more than 1")
		}

		if len(pkg) == 0 {
			return v.RawName
		}

		return fmt.Sprintf("[]%s", golangExpr(v.Value, pkg...))
	case spec.MapType:
		if len(pkg) > 1 {
			panic("package cannot be more than 1")
		}

		if len(pkg) == 0 {
			return v.RawName
		}

		return fmt.Sprintf("map[%s]%s", v.Key, golangExpr(v.Value, pkg...))
	case spec.PointerType:
		if len(pkg) > 1 {
			panic("package cannot be more than 1")
		}

		if len(pkg) == 0 {
			return v.RawName
		}

		return fmt.Sprintf("*%s", golangExpr(v.Type, pkg...))
	case spec.InterfaceType:
		return v.RawName
	}

	return ""
}

// ------danEditStart------
func DanValidateInitCode(fieldName, fieldType string, tags []*spec.Tag) string {
	code := ""
	for _, tag := range tags {
		if tag.Key == "check" {
			//zero用的structtag包解析，会把第一个参数放到Name属性里，所以重新组装属性集合
			if tag.Name == "" {
				continue
			}
			opts := make([]string, 0)
			opts = append(opts, tag.Name)
			opts = append(opts, tag.Options...)
			for _, check := range opts {
				switch check {
				case "required":
					code += fmt.Sprintf(`
									if l.hasReq.` + fieldName + `== false {
										return l.resd.NewErrWithTemp(resd.ErrReqFieldRequired1, "` + fieldName + `")
									}
								`)
					if fieldType == "string" {
						code += fmt.Sprintf(`
										if l.req.` + fieldName + `== "" {
											return l.resd.NewErrWithTemp(resd.ErrReqFieldEmpty1, "` + fieldName + `")
										}
									`)
					} else if fieldType == "int64" {
						code += fmt.Sprintf(`
										if l.req.` + fieldName + `== "" {
											return l.resd.NewErrWithTemp(resd.ErrReqFieldEmpty1, "` + fieldName + `")
										}
									`)
					} else if len(fieldType) > 2 && fieldType[:2] == "[]" {
						code += fmt.Sprintf(`
										if len(l.req.` + fieldName + `) == 0 {
											return l.resd.NewErrWithTemp(resd.ErrReqFieldEmpty1, "` + fieldName + `")
										}
									`)
					} else if len(fieldType) > 4 && fieldType[:4] == "map[" {
						code += fmt.Sprintf(`
										if len(l.req.` + fieldName + `) == 0 {
											return l.resd.NewErrWithTemp(resd.ErrReqFieldEmpty1, "` + fieldName + `")
										}
									`)
					}
				}
			}

		}
	}
	return code
}

// ------danEditEnd------
