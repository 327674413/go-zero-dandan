package wechat

type WxpubConf struct {
	Appid  string
	Secret string
	Token  string
	AESKey string
}
type JssdkBuildParams struct {
	Url         string
	JsApiList   []string
	OpenTagList []string
}
type JssdkBuildResult struct {
	AppId       string
	Beta        bool
	Debug       bool
	JsApiList   []string
	NonceStr    string
	OpenTagList []string
	Signature   string
	Timestamp   int64
	Url         string
}
