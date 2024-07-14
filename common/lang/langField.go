package lang

type langField struct {
	Value string
	Label string
}

var Fields = []*langField{
	{"VarPhoneNumber", "手机号码"},
	{"VarSmsTemp", "短信模版"},
	{"VarMessageSysConfig", "消息中心系统参数"},
	{"VarData", "数据"},
	{"VarImage", "图片"},
	{"VarId", "id"},
	{"VarAccount", "账号"},
	{"VarPassword", "密码"},
	{"VarUpoladTask", "上传任务"},
}
