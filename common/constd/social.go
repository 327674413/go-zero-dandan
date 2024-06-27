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
