
{{if .hasComment}}{{.comment}}{{end}}
func (s *{{.server}}Server) {{.method}} ({{if .notStream}}ctx context.Context,{{if .hasReq}} in {{.request}}{{end}}{{else}}{{if .hasReq}} in {{.request}},{{end}}stream {{.streamBody}}{{end}}) ({{if .notStream}}{{.response}},{{end}}error) {
	l := {{.logicPkg}}.New{{.logicName}}({{if .notStream}}ctx,{{else}}stream.Context(),{{end}}s.svcCtx)
	resp,err := l.{{.method}}({{if .hasReq}}in{{if .stream}} ,stream{{end}}{{else}}{{if .stream}}stream{{end}}{{end}})
    if err != nil{
        danErr,ok := resd.AssertErr(err)
        if ok {
            byt,err := json.Marshal(danErr)
            if err == nil{
                return nil,errors.New(string(byt))
            }
        }
        return nil,err
    }
    return resp,err
}
