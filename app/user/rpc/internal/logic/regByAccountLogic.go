package logic

import (
	"context"
	"fmt"
	"go-zero-dandan/app/user/model"
	"go-zero-dandan/app/user/rpc/internal/svc"
	"go-zero-dandan/app/user/rpc/types/userRpc"
	"go-zero-dandan/app/user/rpc/user"
	"go-zero-dandan/common/constd"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/utild"
	"go-zero-dandan/common/utild/copier"
)

type RegByAccountLogic struct {
	*RegByAccountLogicGen
}

func NewRegByAccountLogic(ctx context.Context, svc *svc.ServiceContext) *RegByAccountLogic {
	return &RegByAccountLogic{
		RegByAccountLogicGen: NewRegByAccountLogicGen(ctx, svc),
	}
}

func (l *RegByAccountLogic) RegByAccount(req *userRpc.RegByAccountReq) (*userRpc.LoginResp, error) {
	if err := l.initReq(req); err != nil {
		return nil, l.resd.Error(err)
	}
	// 校验表单
	if l.req.Account == "" {
		return nil, l.resd.NewErrWithTemp(resd.ErrReqFieldRequired1, resd.VarAccount)
	}
	if utild.Strlen(l.req.Password) < 6 {
		return nil, l.resd.NewErrWithTemp(resd.ErrReqParamFormat1, resd.VarPassword)
	}
	// 验证是否已注册
	userModel := model.NewUserMainModel(l.ctx, l.svc.SqlConn, l.meta.PlatId)
	existUser, err := userModel.Ctx(l.ctx).Where("account = ?", l.req.Account).Find()
	if err != nil {
		return nil, l.resd.Error(err)
	}
	if existUser != nil {
		return nil, l.resd.NewErrWithTemp(resd.ErrDataExist1, resd.VarAccount)
	}
	addData, err := l.prepareUserMain(req)
	if err != nil {
		return nil, l.resd.Error(err)
	}
	addData.Account = l.req.Account
	addData.Password = utild.Sha1(l.req.Password)
	addData.Id = utild.MakeId()
	_, err = userModel.Insert(addData)
	if err != nil {
		return nil, l.resd.NewErr(resd.ErrMysqlInsert)
	}
	userInfo := &user.UserMainInfo{}
	if err := copier.Copy(&userInfo, addData); err != nil {
		return nil, l.resd.NewErr(resd.ErrCopier)
	}
	if l.req.IsLogin == 1 {
		//生成token并存入缓存
		s := fmt.Sprintf("%d-%d-%d", userInfo.Id, utild.GetStamp(), utild.Rand(11111, 99999))
		token := utild.Sha256(s)
		// 因userInfo的id转json是字符串，rpc取出来转化int64会报错，这里用rpc的结构体来缓存
		expireNum := l.svc.Config.Conf.LoginTokenExSec
		err = l.svc.Redis.SetDataExCtx(l.ctx, "userToken", token, userInfo, expireNum)
		nowStamp := utild.GetStamp()
		if err != nil {
			return nil, l.resd.Error(err, resd.ErrRedisSet)
		}
		userInfo.PlatId = l.meta.PlatId
		return &userRpc.LoginResp{
			Id:       addData.Id,
			Token:    token,
			ExpireAt: nowStamp + int64(expireNum),
			UserInfo: userInfo,
		}, nil
	} else {
		return &userRpc.LoginResp{
			Id: addData.Id,
		}, nil
	}

}

func (l *RegByAccountLogic) prepareUserMain(in *userRpc.RegByAccountReq) (*model.UserMain, error) {
	addData := &model.UserMain{
		StateEm: constd.UserStateEmNormal,
	}
	if in.AvatarImg != nil {
		addData.AvatarImg = *in.AvatarImg
	}
	if in.Phone != nil {
		addData.Phone = *in.Phone
	}
	if in.PhoneArea != nil {
		addData.PhoneArea = *in.PhoneArea
	}
	if in.Email != nil {
		addData.Email = *in.Email
	}
	if in.SexEm != nil {
		addData.SexEm = *in.SexEm
	}
	return addData, nil
}
