package lang

type langErr struct {
	Code  int
	Value string
	Label string
}

var Errs = []*langErr{
	{200, "Ok", "完成"},
	{201, "OkAsync", "提交成功"},
	// 4开头权限类
	{400, "ErrAuth", "权限异常"},
	{40101, "ErrAuthPlat", "无效应用"},
	{40102, "ErrAuthPlatExpired", "应用token已失效"},
	{40201, "ErrAuthUserNotLogin", "您还未登录"},
	{40202, "ErrLoginAcctOrPassInvalid", "账号或密码错误"},
	{40301, "ErrAuthOperateUser", "用户无操作权限"},
	{40302, "ErrAuthOperateState", "数据状态不允许操作"},
	{40501, "ErrReqRateLimit", "请求太频繁，请休息一会"},
	{40502, "ErrLimiterErr", "遇到了限流问题"},
	// 5开头系统类
	{500, "ErrSys", "遇到了点小问题"},
	{50001, "ErrUserMainInfo", "用户信息获取失败"},
	{50002, "ErrRpcService", "服务暂不可用"},
	// 501 数据库操作
	{501, "ErrMysql", "数据库操作失败"},
	{50101, "ErrMysqlInsert", "数据库插入失败"},
	{50102, "ErrMysqlDelete", "数据库删除失败"},
	{50103, "ErrMysqlUpdate", "数据库修改失败"},
	{50104, "ErrMysqlSave", "数据库修改失败"},
	{50105, "ErrMysqlSelect", "数据库查询失败"},
	{50106, "ErrMysqlStartTrans", "数据库开启事务失败"},
	{50107, "ErrMysqlCommit", "数据库提交事务失败"},
	{50108, "ErrMysqlRollback", "数据库回滚失败"},
	{50109, "ErrMysqlPrepareUpdate", "数据库更新失败"},
	// 502 redis操作
	{502, "ErrRedis", "redis异常"},
	{50201, "ErrRedisSet", "缓存写入失败"},
	{50202, "ErrRedisGet", "获取缓存失败"},
	{50203, "ErrRedisInc", "redis 数字自增异常"},
	{50204, "ErrRedisDec", "redis 数字自减异常"},
	{50205, "ErrRedisKeyNil", "redis key不存在"},
	{502010001, "ErrRedisSetVerifyCode", "redis set异常"},
	{502010002, "ErrRedisSetVerifyCodeInterval", "redis set异常"},
	{502010003, "ErrRedisSetUserLoginState", "redis set异常"},
	{502020001, "ErrRedisGetUserToken", "获取缓存失败"},

	// 503 mongo
	{50301, "ErrMongoInsert", "数据新增失败"},
	{50302, "ErrMongoUpdate", "数据更新失败"},
	{50302, "ErrMongoDelete", "数据删除失败"},
	{50304, "ErrMongoSelect", "数据查询失败"},
	{50305, "ErrMongoStrToId", "数据更新标识转化失败"},
	{50306, "ErrMongoIdHex", "不合法的ID标识"},
	// 504 kafka
	{504, "ErrMqPush", "MQ推送异常"},
	{52001, "ErrJsonEncode", "转换json失败"},
	{52002, "ErrJsonDecode", "解析json失败"},
	{521, "ErrCopier", "内部数据格式转换失败"},
	{510, "ErrRpcResDecode", "RpcResDecode"},
	{51001, "ErrRpcMissMeta", "缺失meta信息"},
	// 505 CURL
	{50501, "ErrCurlCreate", "创建http请求失败"},
	{50502, "ErrCurlSend", "发送http请求失败"},
	{50503, "ErrCurlSteamNotSupported", "不支持的流式请求"},
	{50504, "ErrCurlSteamScan", "获取流式数据出错"},
	// 506 io
	{50601, "ErrIoRead", "获取数据失败"},
	//业务类
	{600, "ErrReqKeyRequired", "未提供主键"},
	{601, "ErrReqFieldRequired1", "缺少参数{{.Field1}}"},
	{60101, "ErrReqFieldEmpty1", "{{.Field1}}不能为空"},
	{602, "ErrReqParam", "请求参数不正确"},
	{60201, "ErrReqParamFormat1", "{{.Field1}}不符合要求"},
	{603, "ErrConfigNotInit1", "未配置{{.Field1}}"},
	{604, "ErrPlatInvalid", "无效应用"},
	{60400001, "ErrPlatId", "应用id异常"},
	{60400002, "ErrPlatClas", "应用类型异常"},
	{605, "ErrReqWait", "请求太频繁"},
	{606, "ErrNotFound1", "{{.Field1}}不存在"},
	{607, "ErrDataExist1", "该{{.Field1}}已存在"},
	{60800001, "ErrNotSupportPhoneArea", "暂不支持的手机号码"},
	{60800002, "ErrNotSupportImageType", "不支持的图片格式"},
	{60800003, "ErrNotSupportFileType", "不支持的文件类型"},
	{60800004, "ErrUploadImageTypeLimited1", "图片格式仅支持：{{.Field1}}"},
	{60800005, "ErrUploadFileTypeLimited1", "文件格式仅支持：{{.Field1}}"},
	{60901, "ErrMergeFileChunkNotFound", "未找到合并的文件分片"},
	{62000001, "ErrAccountOrPassWrong", "账号或密码错误"},
	{62000002, "ErrVerifyCodeWrong", "验证码错误"},
	{62000003, "ErrVerifyCodeExpired", "请先获取验证码"},
	{62000004, "ErrUploadImageSizeLimited1", "图片大小不可超过{{.Field1}}"},
	{62000005, "ErrUploadFileSizeLimited1", "文件大小不可超过{{.Field1}}"},
	{62000010, "ErrUploadFileFail", "上传文件失败"},
	{62000030, "ErrMultipartUploadFileHashRequired", "分片上传必须提供文件sha1哈希"},
	{62000031, "ErrMultipartUploadNotComplete", "分片上传未全部完成"},
	{60600001, "ErrNotFoundUser", "未找到该用户"},
	{60200001, "ErrReqPhone", "请输入正确的手机号码"},
	{60500001, "ErrReqGetPhoneVerifyCodeWait", "获取验证码太频繁，请稍后再试"},
	{60500002, "ErrReqGetPhoneVerifyCodeDayLimit", "今日获取短信超出上限，请明日再试"},
	{60500003, "ErrReqGetPhoneVerifyCodeHourLimit", "一小时内获取短信超出上限，请稍后再试"},
	//630 im应用类
	{63000001, "ErrSocialAlreadyFriend", "已经是好友了"},
	{63000002, "ErrSocialAlreadyBlackMe", "对方已把你拉黑"},
	{63000003, "ErrSocialNotAddSelf", "不能添加自己为好友"},
	{63000004, "ErrCreateConversation", "创建会话失败"},
	{660, "ErrDataBiz", "业务数据异常"},
	//700 三方平台
	{70001, "ErrTrdNotFound", "三方服务未找到"},
	{70101, "ErrTrdSmsSend", "短信发送失败"},
	{72002, "ErrTrdDifyChatStream", "发送ai流式信息失败"},
}
