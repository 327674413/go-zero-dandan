package logic

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/user/model"
	"go-zero-dandan/app/user/rpc/internal/svc"
	"go-zero-dandan/app/user/rpc/types/pb"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/utild"
)

type EditUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	platId int64
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
	data, err := utild.MakeModelData(*in, "Id,GraduateFrom,BirthDate", true)
	if err != nil {
		return nil, resd.RpcEncodeTempErr(resd.Err)
	}
	_, err = userInfoModel.Update(data)
	/*userModel := model.NewUserMainModel(l.svcCtx.SqlConn, l.platId)
	data := utild.StructToStrMapExcept(*in, "sizeCache", "unknownFields", "state")
	err := userModel.Update(l.ctx, data)
	*/
	find, err := userInfoModel.Find(in.Id)
	fmt.Println(find)
	if err != nil {
		logx.Error(err)
		return nil, resd.RpcEncodeTempErr(resd.MysqlUpdateErr)
	}

	return &pb.SuccResp{Code: 200}, nil
}

func (l *EditUserInfoLogic) Plat(platId int64) *EditUserInfoLogic {
	l.platId = platId
	return l
}
