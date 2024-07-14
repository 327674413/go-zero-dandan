package logic

import (
	"context"
	"fmt"
	"go-zero-dandan/app/user/model"
	"go-zero-dandan/app/user/rpc/internal/svc"
	"go-zero-dandan/app/user/rpc/types/userRpc"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/utild/copier"
	"strings"
)

type GetUserPageLogic struct {
	*GetUserPageLogicGen
}

func NewGetUserPageLogic(ctx context.Context, svc *svc.ServiceContext) *GetUserPageLogic {
	return &GetUserPageLogic{
		GetUserPageLogicGen: NewGetUserPageLogicGen(ctx, svc),
	}
}

func (l *GetUserPageLogic) GetUserPage(req *userRpc.GetUserPageReq) (*userRpc.GetUserPageResp, error) {
	if err := l.initReq(req); err != nil {
		return nil, l.resd.Error(err)
	}
	userModel := model.NewUserMainModel(l.ctx, l.svc.SqlConn, l.req.PlatId)
	if l.hasReq.Match {
		strMatchs := []string{"phone", "nickname"}
		for _, field := range strMatchs {
			if item := l.req.Match[field]; item != nil {
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
	userList, err := userModel.Page(l.req.Page, l.req.Size).Select()
	if err != nil {
		return nil, l.resd.Error(err, resd.ErrMysqlSelect)
	}
	resList := make([]*userRpc.UserMainInfo, 0)
	err = copier.Copy(&resList, &userList)
	if err != nil {
		return nil, l.resd.Error(err, resd.ErrCopier)
	}
	return &userRpc.GetUserPageResp{
		List: resList,
	}, nil
}
