// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"
	"time"

	friend "go-zero-dandan/app/im/api/internal/handler/friend"
	group "go-zero-dandan/app/im/api/internal/handler/group"
	"go-zero-dandan/app/im/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.LangMiddleware, serverCtx.UserInfoMiddleware, serverCtx.UserTokenMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/friend/createFriendApply",
					Handler: friend.CreateFriendApplyHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/friend/operateMyRecvFriendApply",
					Handler: friend.OperateMyRecvFriendApplyHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/friend/getMyFriendApplyRecvPage",
					Handler: friend.GetMyFriendApplyRecvPageHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/friend/getMyFriendList",
					Handler: friend.GetMyFriendListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/friend/searchNewFriendPage",
					Handler: friend.SearchNewFriendPageHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/social/v1"),
		rest.WithTimeout(30000*time.Millisecond),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.LangMiddleware, serverCtx.UserInfoMiddleware, serverCtx.UserTokenMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/group/createGroup",
					Handler: group.CreateGroupHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/group/createGroupMemberApply",
					Handler: group.CreateGroupMemberApplyHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/group/operateGroupMemberApply",
					Handler: group.OperateGroupMemberApplyHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/group/getMyGroupApplyRecvList",
					Handler: group.GetMyGroupApplyRecvListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/groups/getMyGroupList",
					Handler: group.GetMyGroupListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/group/getGroupMemberList",
					Handler: group.GetGroupMemberListHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/social/v1"),
		rest.WithTimeout(30000*time.Millisecond),
	)
}
