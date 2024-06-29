package gogen

import (
	_ "embed"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
	"unicode"

	"github.com/zeromicro/go-zero/tools/goctl/api/parser/g4/gen/api"
	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	"github.com/zeromicro/go-zero/tools/goctl/config"
	"github.com/zeromicro/go-zero/tools/goctl/util/format"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
	"github.com/zeromicro/go-zero/tools/goctl/vars"
)

//go:embed logic.tpl
var logicTemplate string

// ------danEditStart------
var logicGenTemplate string

// ------danEditEnd------
func genLogic(dir, rootPkg string, cfg *config.Config, api *spec.ApiSpec) error {
	for _, g := range api.Service.Groups {
		for _, r := range g.Routes {
			// ------danEditStart------
			// 下面第一个api的参数是我加的
			// ------danEditEnd------
			err := genLogicByRoute(api, dir, rootPkg, cfg, g, r)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// ------danEditStart------
// 下面第一个api *spec.ApiSpec 参数我加的
// ------danEditEnd------
func genLogicByRoute(api *spec.ApiSpec, dir, rootPkg string, cfg *config.Config, group spec.Group, route spec.Route) error {
	logic := getLogicName(route)
	goFile, err := format.FileNamingFormat(cfg.NamingFormat, logic)
	if err != nil {
		return err
	}
	imports := genLogicImports(route, rootPkg)
	var responseString string
	var returnString string
	var requestString string
	if len(route.ResponseTypeName()) > 0 {
		resp := responseGoTypeName(route, typesPacket)
		responseString = "(resp " + resp + ", err error)"
		returnString = "return"
	} else {
		responseString = "error"
		returnString = "return nil"
	}
	if len(route.RequestTypeName()) > 0 {
		requestString = "req *" + requestGoTypeName(route, typesPacket)
	}

	subDir := getLogicFolderPath(group, route)

	// ------danEditStart------
	os.Remove(path.Join(dir, subDir, goFile+"_gen.go"))
	hasUserInfo, mustUserInfo, userLoginInitVar := getDanUserLogin(api)
	defineVars, initVars := getDanGenVars(&getDanGenVarsReq{
		api:          api,
		reqKey:       strings.Replace(requestString, "req *types.", "", -1),
		hasUserInfo:  hasUserInfo,
		mustUserInfo: mustUserInfo,
	})
	importsForGen := imports
	if !shallImportTypesPackageForGen(route) {
		importsForGen = strings.Replace(imports, fmt.Sprintf("\"%s\"\n", pathx.JoinPackages(rootPkg, typesDir)), "", -1)
	}

	genFile(fileGenConfig{
		dir:             dir,
		subdir:          subDir,
		filename:        goFile + "_gen.go",
		templateName:    "logicTemplate",
		category:        category,
		templateFile:    logicGenTemplateFile,
		builtinTemplate: logicGenTemplate,
		data: map[string]string{
			"pkgName":          subDir[strings.LastIndex(subDir, "/")+1:],
			"imports":          importsForGen,
			"logic":            strings.Title(logic) + "Gen",
			"function":         strings.Title(strings.TrimSuffix(logic, "Logic")),
			"responseType":     responseString,
			"returnString":     returnString,
			"request":          requestString,
			"danInitReqFields": initVars,
			"danLogicVars":     defineVars,
			"danUserLoginVars": userLoginInitVar,
		},
	})
	//现在的空文件用不到logx，移除
	imports = strings.Replace(imports, fmt.Sprintf("\"%s/core/logx\"", vars.ProjectOpenSourceURL), "", -1)
	// ------danEditEnd------

	return genFile(fileGenConfig{
		dir:             dir,
		subdir:          subDir,
		filename:        goFile + ".go",
		templateName:    "logicTemplate",
		category:        category,
		templateFile:    logicTemplateFile,
		builtinTemplate: logicTemplate,
		data: map[string]string{
			"pkgName":      subDir[strings.LastIndex(subDir, "/")+1:],
			"imports":      imports,
			"logic":        strings.Title(logic),
			"function":     strings.Title(strings.TrimSuffix(logic, "Logic")),
			"responseType": responseString,
			"returnString": returnString,
			"request":      requestString,
		},
	})
}

// ------danEditStart------
func toFirstUpper(s string) string {
	if len(s) == 0 {
		return s
	}

	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}
func getDanUserLogin(api *spec.ApiSpec) (hasUserInfo, mustUserInfo bool, userLoginFlag string) {
	middlewares := getMiddleware(api)
	midds := make([]string, 0)
	for _, item := range middlewares {
		switch item {
		case "UserTokenMiddleware":
			mustUserInfo = true
			midds = append(midds, "mustUserInfo true")
		case "UserInfoMiddleware":
			hasUserInfo = true
			midds = append(midds, "hasUserInfo true")
		}
	}
	return hasUserInfo, mustUserInfo, strings.Join(midds, "\n")
}

type getDanGenVarsReq struct {
	api          *spec.ApiSpec
	reqKey       string
	hasUserInfo  bool
	mustUserInfo bool
}

func getDanGenVars(params *getDanGenVarsReq) (defineVars, initVars string) {
	hasStr := ""
	for _, tp := range params.api.Types {
		if tp.Name() == params.reqKey {
			obj, ok := tp.(spec.DefineStruct)
			if !ok {
				return "unspport struct type: " + tp.Name(), "unspport struct type: " + tp.Name()
			}
			for _, field := range obj.Members {
				fieldName := toFirstUpper(field.Name)
				fieldType := transReqType(field.Type.Name())
				defineVars += fmt.Sprintf("Req%s %s %s\n", fieldName, fieldType, field.Tag)
				hasStr += fmt.Sprintf("%s bool\n", fieldName)
				initVars += fmt.Sprintf(`
					if req.` + fieldName + `!= nil {
						l.Req` + fieldName + ` = *req.` + fieldName + `
						l.HasReq.` + fieldName + ` = true
					} else {
						l.HasReq.` + fieldName + ` = false
					}
				`)

				for _, tag := range field.Tags() {
					//fmt.Printf("key:%s,name:%s,options:%s\n", tag.Key, tag.Name, strings.Join(tag.Options, ","))
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
								initVars += fmt.Sprintf(`
									if l.HasReq.` + fieldName + `== false {
										return resd.NewErrWithTempCtx(l.ctx, "缺少参数` + fieldName + `", resd.ReqFieldRequired1, "*` + fieldName + `")
									}
								`)
								if fieldType == "string" {
									initVars += fmt.Sprintf(`
										if l.Req` + fieldName + `== "" {
											return resd.NewErrWithTempCtx(l.ctx, "` + fieldName + `不得为空", resd.ReqFieldEmpty1, "*` + fieldName + `")
										}
									`)
								}
							}
						}

					}
				}
			}
		}
		continue
	}
	defineVars += "HasReq struct{\n"
	defineVars += hasStr
	defineVars += "}\n"

	if params.hasUserInfo {
		initVars += "l.hasUserInfo = true\n"
	}
	if params.mustUserInfo {
		initVars += "l.mustUserInfo = true\n"
	}
	return defineVars, initVars
}

// ------danEditEnd------
func transReqType(typeName string) string {
	switch typeName {
	case "*string":
		return "string"
	case "*int64":
		return "int64"
	case "*bool":
		return "bool"
	}
	return "string"
}
func getLogicFolderPath(group spec.Group, route spec.Route) string {
	folder := route.GetAnnotation(groupProperty)
	if len(folder) == 0 {
		folder = group.GetAnnotation(groupProperty)
		if len(folder) == 0 {
			return logicDir
		}
	}
	folder = strings.TrimPrefix(folder, "/")
	folder = strings.TrimSuffix(folder, "/")
	return path.Join(logicDir, folder)
}

func genLogicImports(route spec.Route, parentPkg string) string {
	var imports []string
	imports = append(imports, `"context"`+"\n")
	imports = append(imports, fmt.Sprintf("\"%s\"", pathx.JoinPackages(parentPkg, contextDir)))
	if shallImportTypesPackage(route) {
		imports = append(imports, fmt.Sprintf("\"%s\"\n", pathx.JoinPackages(parentPkg, typesDir)))
	}
	imports = append(imports, fmt.Sprintf("\"%s/core/logx\"", vars.ProjectOpenSourceURL))
	return strings.Join(imports, "\n\t")
}

func onlyPrimitiveTypes(val string) bool {
	fields := strings.FieldsFunc(val, func(r rune) bool {
		return r == '[' || r == ']' || r == ' '
	})

	for _, field := range fields {
		if field == "map" {
			continue
		}
		// ignore array dimension number, like [5]int
		if _, err := strconv.Atoi(field); err == nil {
			continue
		}
		if !api.IsBasicType(field) {
			return false
		}
	}

	return true
}

// ------danEditStart------
func shallImportTypesPackageForGen(route spec.Route) bool {
	if len(route.RequestTypeName()) > 0 {
		return true
	}
	return false
}

// ------danEditEnd------
func shallImportTypesPackage(route spec.Route) bool {
	if len(route.RequestTypeName()) > 0 {
		return true
	}

	respTypeName := route.ResponseTypeName()
	if len(respTypeName) == 0 {
		return false
	}

	if onlyPrimitiveTypes(respTypeName) {
		return false
	}

	return true
}
