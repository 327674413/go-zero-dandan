// Code generated by goctl. DO NOT EDIT.
// Source: social.proto

package server

import (
	"context"
	"encoding/json"
	"errors"
	"go-zero-dandan/app/social/rpc/internal/logic"
	"go-zero-dandan/app/social/rpc/internal/svc"
	"go-zero-dandan/app/social/rpc/types/socialRpc"
	"go-zero-dandan/common/resd"
)

type SocialServer struct {
	svcCtx *svc.ServiceContext
	socialRpc.UnimplementedSocialServer
}

func NewSocialServer(svcCtx *svc.ServiceContext) *SocialServer {
	return &SocialServer{
		svcCtx: svcCtx,
	}
}

func (s *SocialServer) CreateFriendApply(ctx context.Context, in *socialRpc.CreateFriendApplyReq) (*socialRpc.CreateFriendApplyResp, error) {
	l := logic.NewCreateFriendApplyLogic(ctx, s.svcCtx)
	resp, err := l.CreateFriendApply(in)
	if err != nil {
		danErr, ok := resd.AssertErr(err)
		if ok {
			byt, err := json.Marshal(danErr)
			if err == nil {
				return nil, errors.New(string(byt))
			}
		}
		return nil, err
	}
	return resp, err
}

func (s *SocialServer) OperateFriendApply(ctx context.Context, in *socialRpc.OperateFriendApplyReq) (*socialRpc.ResultResp, error) {
	l := logic.NewOperateFriendApplyLogic(ctx, s.svcCtx)
	resp, err := l.OperateFriendApply(in)
	if err != nil {
		danErr, ok := resd.AssertErr(err)
		if ok {
			byt, err := json.Marshal(danErr)
			if err == nil {
				return nil, errors.New(string(byt))
			}
		}
		return nil, err
	}
	return resp, err
}

func (s *SocialServer) GetUserRecvFriendApplyPage(ctx context.Context, in *socialRpc.GetUserRecvFriendApplyPageReq) (*socialRpc.FriendApplyPageResp, error) {
	l := logic.NewGetUserRecvFriendApplyPageLogic(ctx, s.svcCtx)
	resp, err := l.GetUserRecvFriendApplyPage(in)
	if err != nil {
		danErr, ok := resd.AssertErr(err)
		if ok {
			byt, err := json.Marshal(danErr)
			if err == nil {
				return nil, errors.New(string(byt))
			}
		}
		return nil, err
	}
	return resp, err
}

func (s *SocialServer) GetUserFriendList(ctx context.Context, in *socialRpc.GetUserFriendListReq) (*socialRpc.FriendListResp, error) {
	l := logic.NewGetUserFriendListLogic(ctx, s.svcCtx)
	resp, err := l.GetUserFriendList(in)
	if err != nil {
		danErr, ok := resd.AssertErr(err)
		if ok {
			byt, err := json.Marshal(danErr)
			if err == nil {
				return nil, errors.New(string(byt))
			}
		}
		return nil, err
	}
	return resp, err
}

func (s *SocialServer) GetFriendOnline(ctx context.Context, in *socialRpc.GetFriendOnlineReq) (*socialRpc.FriendOnlineResp, error) {
	l := logic.NewGetFriendOnlineLogic(ctx, s.svcCtx)
	resp, err := l.GetFriendOnline(in)
	if err != nil {
		danErr, ok := resd.AssertErr(err)
		if ok {
			byt, err := json.Marshal(danErr)
			if err == nil {
				return nil, errors.New(string(byt))
			}
		}
		return nil, err
	}
	return resp, err
}

func (s *SocialServer) GetUserRelation(ctx context.Context, in *socialRpc.GetUserRelationReq) (*socialRpc.GetUserRelationResp, error) {
	l := logic.NewGetUserRelationLogic(ctx, s.svcCtx)
	resp, err := l.GetUserRelation(in)
	if err != nil {
		danErr, ok := resd.AssertErr(err)
		if ok {
			byt, err := json.Marshal(danErr)
			if err == nil {
				return nil, errors.New(string(byt))
			}
		}
		return nil, err
	}
	return resp, err
}

func (s *SocialServer) CreateGroup(ctx context.Context, in *socialRpc.CreateGroupReq) (*socialRpc.CreateGroupResp, error) {
	l := logic.NewCreateGroupLogic(ctx, s.svcCtx)
	resp, err := l.CreateGroup(in)
	if err != nil {
		danErr, ok := resd.AssertErr(err)
		if ok {
			byt, err := json.Marshal(danErr)
			if err == nil {
				return nil, errors.New(string(byt))
			}
		}
		return nil, err
	}
	return resp, err
}

func (s *SocialServer) CreateGroupMemberApply(ctx context.Context, in *socialRpc.CreateGroupMemberApplyReq) (*socialRpc.CreateGroupMemberApplyResp, error) {
	l := logic.NewCreateGroupMemberApplyLogic(ctx, s.svcCtx)
	resp, err := l.CreateGroupMemberApply(in)
	if err != nil {
		danErr, ok := resd.AssertErr(err)
		if ok {
			byt, err := json.Marshal(danErr)
			if err == nil {
				return nil, errors.New(string(byt))
			}
		}
		return nil, err
	}
	return resp, err
}

func (s *SocialServer) GetUserGroupMemberApplyList(ctx context.Context, in *socialRpc.GetUserGroupMemberApplyListReq) (*socialRpc.GroupMemberApplyListResp, error) {
	l := logic.NewGetUserGroupMemberApplyListLogic(ctx, s.svcCtx)
	resp, err := l.GetUserGroupMemberApplyList(in)
	if err != nil {
		danErr, ok := resd.AssertErr(err)
		if ok {
			byt, err := json.Marshal(danErr)
			if err == nil {
				return nil, errors.New(string(byt))
			}
		}
		return nil, err
	}
	return resp, err
}

func (s *SocialServer) OperateGroupMemberApply(ctx context.Context, in *socialRpc.OperateGroupMemberApplyReq) (*socialRpc.ResultResp, error) {
	l := logic.NewOperateGroupMemberApplyLogic(ctx, s.svcCtx)
	resp, err := l.OperateGroupMemberApply(in)
	if err != nil {
		danErr, ok := resd.AssertErr(err)
		if ok {
			byt, err := json.Marshal(danErr)
			if err == nil {
				return nil, errors.New(string(byt))
			}
		}
		return nil, err
	}
	return resp, err
}

func (s *SocialServer) GetUserGroupList(ctx context.Context, in *socialRpc.GetUserGroupListReq) (*socialRpc.GroupListResp, error) {
	l := logic.NewGetUserGroupListLogic(ctx, s.svcCtx)
	resp, err := l.GetUserGroupList(in)
	if err != nil {
		danErr, ok := resd.AssertErr(err)
		if ok {
			byt, err := json.Marshal(danErr)
			if err == nil {
				return nil, errors.New(string(byt))
			}
		}
		return nil, err
	}
	return resp, err
}

func (s *SocialServer) GetGroupMemberList(ctx context.Context, in *socialRpc.GetGroupMemberListReq) (*socialRpc.GroupMemberListResp, error) {
	l := logic.NewGetGroupMemberListLogic(ctx, s.svcCtx)
	resp, err := l.GetGroupMemberList(in)
	if err != nil {
		danErr, ok := resd.AssertErr(err)
		if ok {
			byt, err := json.Marshal(danErr)
			if err == nil {
				return nil, errors.New(string(byt))
			}
		}
		return nil, err
	}
	return resp, err
}

func (s *SocialServer) GetGroupUserOnline(ctx context.Context, in *socialRpc.GetGroupUserOnlineReq) (*socialRpc.GroupUserOnlineResp, error) {
	l := logic.NewGetGroupUserOnlineLogic(ctx, s.svcCtx)
	resp, err := l.GetGroupUserOnline(in)
	if err != nil {
		danErr, ok := resd.AssertErr(err)
		if ok {
			byt, err := json.Marshal(danErr)
			if err == nil {
				return nil, errors.New(string(byt))
			}
		}
		return nil, err
	}
	return resp, err
}
