package logic

import (
	"context"
	"go-zero-dandan/app/user/model"
	"go-zero-dandan/app/user/rpc/internal/svc"
	"go-zero-dandan/app/user/rpc/types/userRpc"
	"go-zero-dandan/common/dao"
	"go-zero-dandan/common/resd"
)

type EditUserInfoLogic struct {
	*EditUserInfoLogicGen
}

func NewEditUserInfoLogic(ctx context.Context, svc *svc.ServiceContext) *EditUserInfoLogic {
	return &EditUserInfoLogic{
		EditUserInfoLogicGen: NewEditUserInfoLogicGen(ctx, svc),
	}
}

func (l *EditUserInfoLogic) EditUserInfo(req *userRpc.EditUserInfoReq) (*userRpc.SuccResp, error) {
	if err := l.initReq(req); err != nil {
		return nil, l.resd.Error(err)
	}
	userInfoModel := model.NewUserInfoModel(l.ctx, l.svc.SqlConn, l.meta.PlatId)
	_, err := userInfoModel.WhereId(l.req.Id).Update(map[dao.TableField]any{
		model.UserInfo_GraduateFrom: l.req.GraduateFrom,
		model.UserInfo_BirthDate:    l.req.BirthDate,
	})
	if err != nil {
		return nil, l.resd.Error(err, resd.ErrMysqlUpdate)
	}
	/*userModel := model.NewUserMainModel(l.ctx,l.svcCtx.SqlConn, l.platId)
	data := utild.StructToStrMapExcept(*in, "sizeCache", "unknownFields", "state")
	err := userModel.Update(l.ctx, data)
	*/
	_, err = userInfoModel.FindById(l.req.Id)
	if err != nil {
		return nil, l.resd.Error(err, resd.ErrMysqlSelect)
	}

	return &userRpc.SuccResp{Code: 200}, nil
}
