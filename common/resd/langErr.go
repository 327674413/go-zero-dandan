package resd

type langErr struct {
	Code  int
	Value string
	Label string
}

var LangErrs = []*langErr{
	{200, "Ok", "完成"},
	{201, "OkAsync", "提交成功"},
	{400, "Auth", "权限异常"},
	{40101, "AuthPlatErr", "无效应用"},
	{40201, "AuthUserNotLoginErr", "您还未登录"},
	{40301, "AuthOperateUserErr", "用户无操作权限"},
	{40302, "AuthOperateStateErr", "数据状态不允许操作"},
	{500, "SysErr", "遇到了点小问题"},
	{50001, "UserMainInfoErr", "系统异常：用户信息获取失败"},
	{501, "MysqlErr", "数据库操作失败"},
	{50101, "MysqlInsertErr", "数据库插入失败"},
	{50102, "MysqlDeleteErr", "数据库删除失败"},
	{50103, "MysqlUpdateErr", "数据库更新失败"},
	{50104, "MysqlSelectErr", "数据库查询失败"},
	{50106, "MysqlStartTransErr", "数据库开启事务失败"},
	{50107, "MysqlCommitErr", "数据库提交事务失败"},
	{50108, "MysqlRollbackErr", "数据库回滚失败"},
	{502, "RedisErr", "redis异常"},
	{50201, "RedisSetErr", "缓存写入失败"},
	{50202, "RedisGetErr", "获取缓存失败"},
	{50203, "RedisIncErr", "redis 数字自增异常"},
	{50204, "RedisDecErr", "redis 数字自减异常"},
	{502010001, "RedisSetVerifyCodeErr", "redis set异常"},
	{502010002, "RedisSetVerifyCodeIntervalErr", "redis set异常"},
	{502010003, "RedisSetUserLoginStateErr", "redis set异常"},
	{502020001, "RedisGetUserTokenErr", "获取缓存失败"},
	{503, "DataBizErr", "业务数据异常"},
	{504, "MqPushErr", "MQ推送异常"},
	{521, "CopierErr", "内部数据格式转换失败"},
	{510, "RpcResDecodeErr", "RpcResDecodeErr"},
	{600, "ReqKeyRequired", "未提供主键"},
	{601, "ReqFieldRequired1", "缺少参数{{.Field1}}"},
	{60101, "ReqFieldEmpty1", "{{.Field1}}不能为空"},
	{602, "ReqParamErr", "请求参数不正确"},
	{60201, "ReqParamFormatErr1", "{{.Field1}}不符合要求"},
	{603, "ConfigNotInit1", "未配置{{.Field1}}"},
	{604, "PlatInvalid", "无效应用"},
	{60400001, "PlatIdErr", "应用id异常"},
	{60400002, "PlatClasErr", "应用类型异常"},
	{605, "ReqWait", "请求太频繁"},
	{606, "NotFound1", "{{.Field1}}不存在"},
	{607, "DataExist1", "该{{.Field1}}已存在"},
	{60800001, "NotSupportPhoneArea", "暂不支持的手机号码"},
	{60800002, "NotSupportImageType", "不支持的图片格式"},
	{60800003, "NotSupportFileType", "不支持的文件类型"},
	{60800004, "UploadImageTypeLimited1", "图片格式仅支持：{{.Field1}}"},
	{60800005, "UploadFileTypeLimited1", "文件格式仅支持：{{.Field1}}"},
	{60901, "MergeFileChunkNotFound", "未找到合并的文件分片"},
	{62000001, "AccountOrPassWrong", "账号或密码错误"},
	{62000002, "VerifyCodeWrong", "验证码错误"},
	{62000003, "VerifyCodeExpired", "请先获取验证码"},
	{62000004, "UploadImageSizeLimited1", "图片大小不可超过{{.Field1}}"},
	{62000005, "UploadFileSizeLimited1", "文件大小不可超过{{.Field1}}"},
	{62000010, "UploadFileFail", "上传文件失败"},
	{62000030, "MultipartUploadFileHashRequired", "分片上传必须提供文件sha1哈希"},
	{62000031, "MultipartUploadNotComplete", "分片上传未全部完成"},
	{60600001, "NotFoundUser", "未找到该用户"},
	{60200001, "ReqPhoneErr", "请输入正确的手机号码"},
	{60500001, "ReqGetPhoneVerifyCodeWait", "获取验证码太频繁，请稍后再试"},
	{60500002, "ReqGetPhoneVerifyCodeDayLimit", "今日获取短信超出上限，请明日再试"},
	{60500003, "ReqGetPhoneVerifyCodeHourLimit", "一小时内获取短信超出上限，请稍后再试"},
	{63000001, "SocialAlreadyFriend", "已经是好友了"},
	{63000002, "SocialAlreadyBlackMe", "对方已把你拉黑"},
	{63000003, "SocialNotAddSelf", "不能添加自己为好友"},
	{700, "TrdSmsSendErr", "TrdSmsSendErr"},
}
