package constd

const (
	SocialFriendStateEmNoRelat = 0  //无关系，friendApply无数据
	SocialFriendStateEmApply   = 1  //申请中
	SocialFriendStateEmPass    = 2  //已通过
	SocialFriendStateEmSelf    = 99 //本人自己
	SocialFriendStateEmCancel  = -1 //已撤销
	SocialFriendStateEmReject  = -2 //已撤销
	SocialFriendStateEmBlack   = -3 //已拉黑
	SocialFriendStateEmDelete  = -4 //已删除
)
const (
	SocialFriendSourceEmSearchPhoneBySelf   = 1   //自己搜索手机号加对方
	SocialFriendSourceEmSearchPhoneByOther  = -1  //对方搜索手机号加自己
	SocialFriendSourceEmSearchPhoneBySystem = -99 //系统设置加的
)

const (
	SocialFriendApplyRecordTypeEmApply  = 1 //申请添加
	SocialFriendApplyRecordTypeEmPass   = 2 //通过申请
	SocialFriendApplyRecordTypeEmReject = 3 //拒绝申请
	SocialFriendApplyRecordTypeEmChat   = 4 //询问回复
	SocialFriendApplyRecordTypeEmCancel = 5 //撤销,预留，其实没这种业务
)
