package logic

import (
	"context"
	"fmt"
	"go-zero-dandan/app/user/model"
	"go-zero-dandan/app/user/rpc/internal/svc"
	"go-zero-dandan/app/user/rpc/types/pb"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/utild/copier"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserPageLogic {
	return &GetUserPageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserPageLogic) GetUserPage(in *pb.GetUserPageReq) (*pb.GetUserPageResp, error) {
	if err := l.checkReqParams(in); err != nil {
		return nil, err
	}
	userModel := model.NewUserMainModel(l.svcCtx.SqlConn, in.PlatId)
	if in.Match != nil {
		strMatchs := []string{"phone", "nickname"}
		for _, field := range strMatchs {
			if item := in.Match[field]; item != nil {
				v := strings.TrimSpace(*item.Str)
				if v == "" {
					if item.IsFuzzy == nil || !*item.IsFuzzy {
						userModel.Where(fmt.Sprintf("%s = ''", field))
					}
				} else {
					if item.IsFuzzy != nil && *item.IsFuzzy {
						userModel.Where(fmt.Sprintf("%s like ?", field), "%"+v+"%")
					} else {
						userModel.Where(fmt.Sprintf("%s = ?", field), v)
					}
				}
			}
		}

	}
	userList, err := userModel.Page(in.Page, in.Size).Select()
	if err != nil {
		return nil, resd.NewRpcErrCtx(l.ctx, err.Error(), resd.MysqlSelectErr)
	}
	resList := make([]*pb.UserMainInfo, 0)
	err = copier.Copy(&resList, &userList)
	if err != nil {
		return nil, resd.NewRpcErrCtx(l.ctx, err.Error(), resd.CopierErr)
	}
	return &pb.GetUserPageResp{
		List: resList,
	}, nil
}
func (l *GetUserPageLogic) checkReqParams(in *pb.GetUserPageReq) error {
	if in.PlatId == "" {
		return resd.NewRpcErrWithTempCtx(l.ctx, "参数缺少platId", resd.ReqFieldRequired1, "platId")
	}
	return nil
}
