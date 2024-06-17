{{if .hasComment}}{{.comment}}{{end}}
func (l *{{.logicName}}) {{.method}} ({{if .hasReq}}in {{.request}}{{if .stream}},stream {{.streamBody}}{{end}}{{else}}stream {{.streamBody}}{{end}}) ({{if .hasReply}}{{.response}},{{end}} error) {
	if err := l.checkReqParams(in); err != nil {
        return nil, err
    }
	
	return {{if .hasReply}}&{{.responseType}}{},{{end}} nil
}
func (l *{{.logicName}}) checkReqParams(in {{.request}}) error {
	if in.PlatId == "" {
		return resd.NewRpcErrWithTempCtx(l.ctx, "参数缺少platId", resd.ReqFieldRequired1, "platId")
	}
	return nil
}