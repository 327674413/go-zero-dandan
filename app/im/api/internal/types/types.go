// Code generated by goctl. DO NOT EDIT.
package types

type FriendInfo struct {
	Id           string `json:"id"`
	FriendUid    string `json:"friendUid"`
	FriendName   string `json:"friendName"`
	FriendAlias  string `json:"friendAlias"`
	FriendIcon   string `json:"friendIcon"`
	FriendSexEm  int64  `json:"friendSexEm"`
	SourceEm     int64  `json:"sourcEm"`
	FriendRemark string `json:"friendRemark"`
}

type FriendApply struct {
	Id              string `json:"id"`
	UserId          string `json:"userId"`
	FriendUid       string `json:"friendUid"`
	ApplyLastMsg    string `json:"applyLastMsg"`
	ApplyLastAt     int64  `json:"applyLastAt"`
	OperateMsg      string `json:"operateMsg"`
	OperateAt       int64  `json:"operateAt"`
	StateEm         int64  `json:"stateEm"`
	UserName        string `json:"userName"`
	UserSex         int64  `json:"userSex"`
	UserAvatarImg   string `json:"userAvatarImg"`
	FriendName      string `json:"friendName"`
	FriendSex       int64  `json:"friendSex"`
	FriendAvatarImg string `json:"friendAvatarImg"`
}

type GroupInfo struct {
	Id          string         `json:"id"`
	Name        string         `json:"name"`
	StateEm     int64          `json:"stateEm"`
	TypeEm      int64          `json:"typeEm"`
	IsVerify    int64          `json:"isVerify"`
	NotiContent string         `json:"notiContent"`
	NotiUid     string         `json:"notiUid"`
	MemberList  []*GroupMember `json:"memberList"`
	CreateAt    int64          `json:"create_at"`
}

type GroupMember struct {
	Id           string `json:"id"`
	GroupId      string `json:"groupId"`
	UserId       string `json:"userId"`
	RoleLevel    int64  `json:"roleLevel"`
	JoinAt       int64  `json:"joinAt"`
	JoinSourceEm int64  `json:"joinSourceEm"`
	InviteUid    string `json:"inviteUid"`
	OperateUid   string `json:"operateUid"`
}

type GroupMemberApply struct {
	Id             string `json:"id"`
	GroupId        string `json:"groupId"`
	UserId         string `json:"userId"`
	ApplyMsg       string `json:"applyMsg"`
	ApplyAt        int64  `json:"applyAt"`
	OperateUid     string `json:"operateUid"`
	OperateMsg     string `json:"operateMsg"`
	OperateAt      int64  `json:"operateAt"`
	OperateStateEm int64  `json:"operateStateEm"`
	JoinSourceEm   int64  `json:"joinSourceEm"`
	InviteUid      string `json:"inviteUid"`
}

type ResultResp struct {
	Code    bool   `json:"result"`
	Content string `json:"content"`
}

type CreateFriendApplyReq struct {
	ApplyMsg  *string `json:"applyMsg,optional"`
	FriendUid *string `json:"friendUid,optional" check:"required"`
	SourceEm  *int64  `json:"sourceEm,optional" check:"required"`
}

type CreateGroupReq struct {
	Name *string `json:"name,optional"`
}

type CreateGroupResp struct {
	GroupId *string `json:"groupId,optional"`
}

type OperateMyRecvFriendApplyReq struct {
	ApplyId        *string `json:"applyId,optional" check:"required"`
	OperateStateEm *int64  `json:"operateStateEm,optional" check:"required"`
	OperateMsg     *string `json:"operaeteMsg,optional"`
}

type CreateGroupMemberApplyReq struct {
	GroupId   *string `json:"groupId,optional"`
	GroupCode *string `json:"groupCode,optional"`
	ApplyMsg  *string `json:"applyMsg,optional"`
	SourceEm  *int64  `json:"sourceEm,optional"`
}

type CreateGroupMemberApplyResp struct {
	ApplyId *string `json:"applyId,optional"`
}

type OperateGroupMemberApplyReq struct {
	ApplyId        *string `json:"applyId,optional"`
	OpreateStateEm *int64  `json:"operateStateEm,optional"` // 处理结果
	OpreateMsg     *string `json:"operateMsg,optional"`
}

type GetGroupMemberListReq struct {
	GroupId *string `json:"groupId,optional"`
}

type GetMyFriendApplyRecvPageReq struct {
	Page *int64 `json:"page,optional"`
	Size *int64 `json:"size,optional"`
}

type FriendApplyListResp struct {
	List []*FriendApply `json:"list"`
}

type FriendListResp struct {
	List []*FriendInfo `json:"list"`
}

type GroupListResp struct {
	List []*GroupInfo `json:"list"`
}

type GroupMemberListResp struct {
	List []*GroupMember `json:"List"`
}

type GroupApplyListResp struct {
	List []*GroupMemberApply `json:"List"`
}

type IdReq struct {
	Id *string `json:"id,optional" check:"required"`
}

type NewFriendInfo struct {
	Id        string `json:"id"`
	Nickname  string `json:"nickname"`
	AvatarImg string `json:"avatarImg"`
	Signature string `json:"signature"`
	StateEm   int64  `json:"stateEm"`
}

type SearchNewFriendReq struct {
	Keyword *string `json:"keyword,optional"`
}

type SearchNewFriendResp struct {
	List []*NewFriendInfo `json:"list"`
}

type GetMySysMsgUnreadNumReq struct {
	MsgClasEm *int64 `json:"msgClasEm,optional"`
}

type GetMySysMsgUnreadNumResp struct {
	Unread map[int64]int64 `json:"unread"`
}

type SetMySysMsgReadByClasReq struct {
	MsgClasEms []int64 `json:"msgClasEms,optional" check:"required"`
}

type SetMySysMsgReadByIdReq struct {
	MsgClasEm *int64  `json:"msgClasEm,optional" check:"required"`
	Ids       []int64 `json:"ids,optional" check:"required"`
}

type ChatLog struct {
	Id             string `json:"id"`
	ConversationId string `json:"conversationId"`
	SendId         string `json:"sendId"`
	RecvId         string `json:"recvId"`
	MsgType        int64  `json:"msgType"`
	MsgContent     string `json:"msgContent"`
	ChatType       int64  `json:"chatType"`
	SendTime       string `json:"sendTime"`
	SendAtMs       int64  `json:"sendAtMs"`
	MsgState       int64  `json:"msgState"`
}

type Conversation struct {
	ConversationId string   `json:"conversationId"`
	ChatType       int64    `json:"chatType"`
	TargetId       string   `json:"targetId"`
	IsShow         bool     `json:"isShow"`
	ReadSeq        int64    `json:"readSeq"`
	DeleteSeq      int64    `json:"deleteSeq"`
	Total          int64    `json:"total"`
	Unread         int64    `json:"unread"`
	LastMsg        *ChatLog `json:"lastMsg"`
	LastAtMs       int64    `json:"lastAtMs"`
}

type GetChatLogReq struct {
	ConversationId *string `json:"conversationId,optional" check:"required"`
	StartSendAt    *int64  `json:"startSendAt,optional"`
	EndSendAt      *int64  `json:"endSendAt,optional"`
	Count          *int64  `json:"count,optional"`
}

type GetChatLogResp struct {
	List []*ChatLog `json:"list"`
}

type GetConversationListResp struct {
	Conversations map[string]*Conversation `json:"conversations"`
}

type UpdateConversationListReq struct {
	ConversationList map[string]*Conversation `json:"conversationList,optional" check:"required"`
}

type SetUpMyConversationReq struct {
	RecvId   *string `json:"recvId,optional" check:"required"`
	ChatType *int64  `json:"chatType,optional" check:"required"`
}
