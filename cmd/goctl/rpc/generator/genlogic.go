package generator

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/collection"
	apiParser "github.com/zeromicro/go-zero/tools/goctl/api/parser"
	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	conf "github.com/zeromicro/go-zero/tools/goctl/config"
	"github.com/zeromicro/go-zero/tools/goctl/rpc/parser"
	"github.com/zeromicro/go-zero/tools/goctl/util"
	"github.com/zeromicro/go-zero/tools/goctl/util/format"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
	"github.com/zeromicro/go-zero/tools/goctl/util/stringx"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

const logicFunctionTemplate = `{{if .hasComment}}{{.comment}}{{end}}
func (l *{{.logicName}}) {{.method}} ({{if .hasReq}}in {{.request}}{{if .stream}},stream {{.streamBody}}{{end}}{{else}}stream {{.streamBody}}{{end}}) ({{if .hasReply}}{{.response}},{{end}} error) {
	// todo: add your logic here and delete this line
	
	return {{if .hasReply}}&{{.responseType}}{},{{end}} nil
}
`

//go:embed logic.tpl
var logicTemplate string

// GenLogic generates the logic file of the rpc service, which corresponds to the RPC definition items in proto.
func (g *Generator) GenLogic(ctx DirContext, proto parser.Proto, cfg *conf.Config,
	c *ZRpcContext) error {
	if !c.Multiple {
		return g.genLogicInCompatibility(ctx, proto, cfg)
	}

	return g.genLogicGroup(ctx, proto, cfg)
}

type rpcFieldInfo struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Comment string `json:"comment"`
}

// ------danEditStart------
func JsonFile(jsonData any, savePathName string) error {

	if filepath.Ext(savePathName) == "" {
		savePathName = savePathName + ".json"
	}
	jsonStr, err := json.Marshal(jsonData)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return os.WriteFile(savePathName, jsonStr, 0644)
}

//------danEditEnd------

func (g *Generator) genLogicInCompatibility(ctx DirContext, proto parser.Proto,
	cfg *conf.Config) error {
	dir := ctx.GetLogic()
	service := proto.Service[0].Service.Name
	for _, rpc := range proto.Service[0].RPC {
		logicName := fmt.Sprintf("%sLogic", stringx.From(rpc.Name).ToCamel())
		logicFilename, err := format.FileNamingFormat(cfg.NamingFormat, rpc.Name+"_logic")
		if err != nil {
			return err
		}
		// ------danEditStart------
		filename := filepath.Join(dir.Filename, logicFilename+"_gen.go")
		imports := collection.NewSet()
		imports.AddStr(fmt.Sprintf(`"%v"`, ctx.GetSvc().Package))
		imports.AddStr(fmt.Sprintf(`"%v"`, ctx.GetPb().Package))
		api, err := apiParser.Parse(filepath.Join(filepath.Dir(proto.Src), service+".rpc"))
		if err != nil {
			return err
		}
		defineVars, initVars := getDanGenVars(&getDanGenVarsReq{
			api:    api,
			reqKey: fmt.Sprintf("%v", rpc.RequestType),
		})
		err = g.genLogicFunction(service, proto.PbPackage, logicName, rpc, filepath.Join(dir.Filename, logicFilename+".go"), strings.Join(imports.KeysStr(), pathx.NL))
		if err != nil {
			return err
		}

		text, err := pathx.LoadTemplate(category, logicTemplateFileFile, logicTemplate)
		if err != nil {
			return err
		}
		err = util.With("logic").GoFmt(true).Parse(text).SaveTo(map[string]any{
			"logicName":      fmt.Sprintf("%sLogic", stringx.From(rpc.Name).ToCamel()),
			"functions":      "", //functions,
			"packageName":    "logic",
			"hasReq":         !rpc.StreamsRequest,
			"request":        fmt.Sprintf("*%s.%s", proto.PbPackage, parser.CamelCase(rpc.RequestType)),
			"imports":        strings.Join(imports.KeysStr(), pathx.NL),
			"dandDefineVars": defineVars,
			"danInitVars":    initVars,
		}, filename, true)
		//------danEditEnd------
		if err != nil {
			return err
		}
	}
	return nil
}

// ------danEditStart------
type getDanGenVarsReq struct {
	api    *spec.ApiSpec
	reqKey string
}

// ------danEditEnd------
func toFirstUpper(s string) string {
	if len(s) == 0 {
		return s
	}

	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}
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

// ------danEditStart------
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
	return defineVars, initVars
}

// ------danEditEnd------
func (g *Generator) genLogicGroup(ctx DirContext, proto parser.Proto, cfg *conf.Config) error {
	dir := ctx.GetLogic()
	for _, item := range proto.Service {
		serviceName := item.Name
		for _, rpc := range item.RPC {
			var (
				err           error
				filename      string
				logicName     string
				logicFilename string
				packageName   string
			)

			logicName = fmt.Sprintf("%sLogic", stringx.From(rpc.Name).ToCamel())
			childPkg, err := dir.GetChildPackage(serviceName)
			if err != nil {
				return err
			}

			serviceDir := filepath.Base(childPkg)
			nameJoin := fmt.Sprintf("%s_logic", serviceName)
			packageName = strings.ToLower(stringx.From(nameJoin).ToCamel())
			logicFilename, err = format.FileNamingFormat(cfg.NamingFormat, rpc.Name+"_logic")
			if err != nil {
				return err
			}

			filename = filepath.Join(dir.Filename, serviceDir, logicFilename+".go")
			// ------danEditStart------
			//functions单独用来做模版了
			//functions, err := g.genLogicFunction(serviceName, proto.PbPackage, logicName, rpc)
			//if err != nil {
			//	return err
			//}

			imports := collection.NewSet()
			imports.AddStr(fmt.Sprintf(`"%v"`, ctx.GetSvc().Package))
			imports.AddStr(fmt.Sprintf(`"%v"`, ctx.GetPb().Package))
			text, err := pathx.LoadTemplate(category, logicTemplateFileFile, logicTemplate)
			if err != nil {
				return err
			}

			if err = util.With("logic").GoFmt(true).Parse(text).SaveTo(map[string]any{
				"logicName":   logicName,
				"functions":   "", //functions,
				"packageName": packageName,
				"imports":     strings.Join(imports.KeysStr(), pathx.NL),
			}, filename, false); err != nil {
				return err
			}
			// ------danEditEnd------
		}
	}
	return nil
}

// ------danEditStart------
func (g *Generator) genLogicFunction(serviceName, goPackage, logicName string,
	rpc *parser.RPC, fileName string, imports string) error {
	text, err := pathx.LoadTemplate(category, logicFuncTemplateFileFile, logicFunctionTemplate)
	if err != nil {
		return err
	}
	comment := parser.GetComment(rpc.Doc())
	streamServer := fmt.Sprintf("%s.%s_%s%s", goPackage, parser.CamelCase(serviceName),
		parser.CamelCase(rpc.Name), "Server")
	return util.With("func").GoFmt(true).Parse(text).SaveTo(map[string]any{
		"packageName":  "logic",
		"logicName":    logicName,
		"method":       parser.CamelCase(rpc.Name),
		"hasReq":       !rpc.StreamsRequest,
		"request":      fmt.Sprintf("*%s.%s", goPackage, parser.CamelCase(rpc.RequestType)),
		"hasReply":     !rpc.StreamsRequest && !rpc.StreamsReturns,
		"response":     fmt.Sprintf("*%s.%s", goPackage, parser.CamelCase(rpc.ReturnsType)),
		"responseType": fmt.Sprintf("%s.%s", goPackage, parser.CamelCase(rpc.ReturnsType)),
		"stream":       rpc.StreamsRequest || rpc.StreamsReturns,
		"streamBody":   streamServer,
		"hasComment":   len(comment) > 0,
		"comment":      comment,
		"imports":      imports,
	}, fileName, false)
}

//下面这些是老的方法，用来当自己的模版，不用下面了
//func (g *Generator) genLogicFunction(serviceName, goPackage, logicName string,
//	rpc *parser.RPC) (string,
//	error) {
//	functions := make([]string, 0)
//	text, err := pathx.LoadTemplate(category, logicFuncTemplateFileFile, logicFunctionTemplate)
//	if err != nil {
//		return "", err
//	}
//
//	comment := parser.GetComment(rpc.Doc())
//	streamServer := fmt.Sprintf("%s.%s_%s%s", goPackage, parser.CamelCase(serviceName),
//		parser.CamelCase(rpc.Name), "Server")
//	buffer, err := util.With("fun").Parse(text).Execute(map[string]any{
//		"logicName":    logicName,
//		"method":       parser.CamelCase(rpc.Name),
//		"hasReq":       !rpc.StreamsRequest,
//		"request":      fmt.Sprintf("*%s.%s", goPackage, parser.CamelCase(rpc.RequestType)),
//		"hasReply":     !rpc.StreamsRequest && !rpc.StreamsReturns,
//		"response":     fmt.Sprintf("*%s.%s", goPackage, parser.CamelCase(rpc.ReturnsType)),
//		"responseType": fmt.Sprintf("%s.%s", goPackage, parser.CamelCase(rpc.ReturnsType)),
//		"stream":       rpc.StreamsRequest || rpc.StreamsReturns,
//		"streamBody":   streamServer,
//		"hasComment":   len(comment) > 0,
//		"comment":      comment,
//	})
//	if err != nil {
//		return "", err
//	}
//
//	functions = append(functions, buffer.String())
//	return strings.Join(functions, pathx.NL), nil
//}

// ------danEditEnd------
