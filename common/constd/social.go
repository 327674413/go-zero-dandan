package constd

const (
	SocialFriendStateEmNoRelat = 0  //无关系，friendApply无数据
	SocialFriendStateEmApply   = 1  //申请中
	SocialFriendStateEmPass    = 2  //已通过
	SocialFriendStateEmReject  = -1 //已拒绝
	SocialFriendStateEmBlack   = -2 //已拉黑
	SocialFriendStateEmDelete  = -3 //已删除
)
