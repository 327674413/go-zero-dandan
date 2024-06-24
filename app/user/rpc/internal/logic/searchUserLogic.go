package logic

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/user/model"
	"go-zero-dandan/app/user/rpc/internal/svc"
	"go-zero-dandan/app/user/rpc/types/pb"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/utild/copier"
	"strings"
)

type SearchUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchUserLogic {
	return &SearchUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchUserLogic) SearchUser(in *pb.SearchUserReq) (*pb.SearchUserResp, error) {
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
	return &pb.SearchUserResp{
		List: resList,
	}, nil
}
func (l *SearchUserLogic) checkReqParams(in *pb.SearchUserReq) error {
	if in.PlatId == "" {
		return resd.NewRpcErrWithTempCtx(l.ctx, "参数缺少platId", resd.ReqFieldRequired1, "platId")
	}
	return nil
}
