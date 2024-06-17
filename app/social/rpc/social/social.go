// Code generated by goctl. DO NOT EDIT.
// Source: social.proto

package social

import (
	"context"

	"go-zero-dandan/app/social/rpc/types/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	FriendListReq         = pb.FriendListReq
	FriendListResp        = pb.FriendListResp
	FriendOnlineResp      = pb.FriendOnlineResp
	FriendPutInHandleReq  = pb.FriendPutInHandleReq
	FriendPutInHandleResp = pb.FriendPutInHandleResp
	FriendPutInListReq    = pb.FriendPutInListReq
	FriendPutInListResp   = pb.FriendPutInListResp
	FriendPutInReq        = pb.FriendPutInReq
	FriendPutInResp       = pb.FriendPutInResp
	FriendRequests        = pb.FriendRequests
	Friends               = pb.Friends
	GroupCreateReq        = pb.GroupCreateReq
	GroupCreateResp       = pb.GroupCreateResp
	GroupListReq          = pb.GroupListReq
	GroupListResp         = pb.GroupListResp
	GroupMembers          = pb.GroupMembers
	GroupOnlineResp       = pb.GroupOnlineResp
	GroupPutInHandleReq   = pb.GroupPutInHandleReq
	GroupPutInHandleResp  = pb.GroupPutInHandleResp
	GroupPutinListReq     = pb.GroupPutinListReq
	GroupPutinListResp    = pb.GroupPutinListResp
	GroupPutinReq         = pb.GroupPutinReq
	GroupPutinResp        = pb.GroupPutinResp
	GroupRequests         = pb.GroupRequests
	GroupUsersReq         = pb.GroupUsersReq
	GroupUsersResp        = pb.GroupUsersResp
	Groups                = pb.Groups

	Social interface {
		FriendPutIn(ctx context.Context, in *FriendPutInReq, opts ...grpc.CallOption) (*FriendPutInResp, error)
		FriendPutInHandle(ctx context.Context, in *FriendPutInHandleReq, opts ...grpc.CallOption) (*FriendPutInHandleResp, error)
		FriendPutInList(ctx context.Context, in *FriendPutInListReq, opts ...grpc.CallOption) (*FriendPutInListResp, error)
		FriendList(ctx context.Context, in *FriendListReq, opts ...grpc.CallOption) (*FriendListResp, error)
		FriendOnlineList(ctx context.Context, in *FriendListReq, opts ...grpc.CallOption) (*FriendOnlineResp, error)
		GroupCreate(ctx context.Context, in *GroupCreateReq, opts ...grpc.CallOption) (*GroupCreateResp, error)
		GroupPutin(ctx context.Context, in *GroupPutinReq, opts ...grpc.CallOption) (*GroupPutinResp, error)
		GroupPutinList(ctx context.Context, in *GroupPutinListReq, opts ...grpc.CallOption) (*GroupPutinListResp, error)
		GroupPutInHandle(ctx context.Context, in *GroupPutInHandleReq, opts ...grpc.CallOption) (*GroupPutInHandleResp, error)
		GroupList(ctx context.Context, in *GroupListReq, opts ...grpc.CallOption) (*GroupListResp, error)
		GroupUsers(ctx context.Context, in *GroupUsersReq, opts ...grpc.CallOption) (*GroupUsersResp, error)
		GroupOnlineUserList(ctx context.Context, in *GroupUsersReq, opts ...grpc.CallOption) (*GroupOnlineResp, error)
	}

	defaultSocial struct {
		cli zrpc.Client
	}
)

func NewSocial(cli zrpc.Client) Social {
	return &defaultSocial{
		cli: cli,
	}
}

func (m *defaultSocial) FriendPutIn(ctx context.Context, in *FriendPutInReq, opts ...grpc.CallOption) (*FriendPutInResp, error) {
	client := pb.NewSocialClient(m.cli.Conn())
	return client.FriendPutIn(ctx, in, opts...)
}

func (m *defaultSocial) FriendPutInHandle(ctx context.Context, in *FriendPutInHandleReq, opts ...grpc.CallOption) (*FriendPutInHandleResp, error) {
	client := pb.NewSocialClient(m.cli.Conn())
	return client.FriendPutInHandle(ctx, in, opts...)
}

func (m *defaultSocial) FriendPutInList(ctx context.Context, in *FriendPutInListReq, opts ...grpc.CallOption) (*FriendPutInListResp, error) {
	client := pb.NewSocialClient(m.cli.Conn())
	return client.FriendPutInList(ctx, in, opts...)
}

func (m *defaultSocial) FriendList(ctx context.Context, in *FriendListReq, opts ...grpc.CallOption) (*FriendListResp, error) {
	client := pb.NewSocialClient(m.cli.Conn())
	return client.FriendList(ctx, in, opts...)
}

func (m *defaultSocial) FriendOnlineList(ctx context.Context, in *FriendListReq, opts ...grpc.CallOption) (*FriendOnlineResp, error) {
	client := pb.NewSocialClient(m.cli.Conn())
	return client.FriendOnlineList(ctx, in, opts...)
}

func (m *defaultSocial) GroupCreate(ctx context.Context, in *GroupCreateReq, opts ...grpc.CallOption) (*GroupCreateResp, error) {
	client := pb.NewSocialClient(m.cli.Conn())
	return client.GroupCreate(ctx, in, opts...)
}

func (m *defaultSocial) GroupPutin(ctx context.Context, in *GroupPutinReq, opts ...grpc.CallOption) (*GroupPutinResp, error) {
	client := pb.NewSocialClient(m.cli.Conn())
	return client.GroupPutin(ctx, in, opts...)
}

func (m *defaultSocial) GroupPutinList(ctx context.Context, in *GroupPutinListReq, opts ...grpc.CallOption) (*GroupPutinListResp, error) {
	client := pb.NewSocialClient(m.cli.Conn())
	return client.GroupPutinList(ctx, in, opts...)
}

func (m *defaultSocial) GroupPutInHandle(ctx context.Context, in *GroupPutInHandleReq, opts ...grpc.CallOption) (*GroupPutInHandleResp, error) {
	client := pb.NewSocialClient(m.cli.Conn())
	return client.GroupPutInHandle(ctx, in, opts...)
}

func (m *defaultSocial) GroupList(ctx context.Context, in *GroupListReq, opts ...grpc.CallOption) (*GroupListResp, error) {
	client := pb.NewSocialClient(m.cli.Conn())
	return client.GroupList(ctx, in, opts...)
}

func (m *defaultSocial) GroupUsers(ctx context.Context, in *GroupUsersReq, opts ...grpc.CallOption) (*GroupUsersResp, error) {
	client := pb.NewSocialClient(m.cli.Conn())
	return client.GroupUsers(ctx, in, opts...)
}

func (m *defaultSocial) GroupOnlineUserList(ctx context.Context, in *GroupUsersReq, opts ...grpc.CallOption) (*GroupOnlineResp, error) {
	client := pb.NewSocialClient(m.cli.Conn())
	return client.GroupOnlineUserList(ctx, in, opts...)
}
