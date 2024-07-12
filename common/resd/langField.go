package resd

type langField struct {
	Value string
	Label string
}

var LangFields = []*langField{
	{"PhoneNumber", "手机号码"},
	{"SmsTemp", "短信模版"},
	{"MessageSysConfig", "消息中心系统参数"},
	{"Data", "数据"},
	{"Image", "图片"},
	{"Id", "id"},
	{"Account", "账号"},
	{"VarPassword", "密码"},
	{"UpoladTask", "上传任务"},
}
