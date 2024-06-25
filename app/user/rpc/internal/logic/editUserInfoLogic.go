package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/user/model"
	"go-zero-dandan/app/user/rpc/internal/svc"
	"go-zero-dandan/app/user/rpc/types/pb"
	"go-zero-dandan/common/dao"
	"go-zero-dandan/common/resd"
)

type EditUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	platId string
}

func NewEditUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditUserInfoLogic {
	return &EditUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *EditUserInfoLogic) EditUserInfo(in *pb.EditUserInfoReq) (*pb.SuccResp, error) {
	userInfoModel := model.NewUserInfoModel(l.svcCtx.SqlConn, l.platId)
	_, err := userInfoModel.WhereId(in.Id).Update(map[dao.TableField]any{
		model.UserInfo_GraduateFrom: in.GraduateFrom,
		model.UserInfo_BirthDate:    in.BirthDate,
	})
	if err != nil {
		return nil, resd.ErrorCtx(l.ctx, err, resd.MysqlUpdateErr)
	}
	/*userModel := model.NewUserMainModel(l.svcCtx.SqlConn, l.platId)
	data := utild.StructToStrMapExcept(*in, "sizeCache", "unknownFields", "state")
	err := userModel.Update(l.ctx, data)
	*/
	_, err = userInfoModel.FindById(in.Id)
	if err != nil {
		return nil, resd.ErrorCtx(l.ctx, err, resd.MysqlSelectErr)
	}

	return &pb.SuccResp{Code: 200}, nil
}

func (l *EditUserInfoLogic) Plat(platId string) *EditUserInfoLogic {
	l.platId = platId
	return l
}

func (l *EditUserInfoLogic) rpcFail(err error) (*pb.SuccResp, error) {
	return nil, resd.RpcErrEncode(err)
}
