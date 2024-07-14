package logic

import (
	"context"
	"fmt"
	"go-zero-dandan/app/user/model"
	"go-zero-dandan/app/user/rpc/types/userRpc"
	"go-zero-dandan/app/user/rpc/user"
	"go-zero-dandan/common/constd"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/utild"
	"go-zero-dandan/common/utild/copier"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/user/rpc/internal/svc"
)

type RegByAccountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegByAccountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegByAccountLogic {
	return &RegByAccountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegByAccountLogic) RegByAccount(in *userRpc.RegByAccountReq) (*userRpc.LoginResp, error) {
	// 校验表单
	account := strings.TrimSpace(in.Account)
	if account == "" {
		return nil, resd.RpcErrEncode(resd.NewErrWithTempCtx(l.ctx, "缺少账号", resd.ErrReqFieldRequired1, resd.VarAccount))
	}
	password := strings.TrimSpace(in.Password)
	if utild.Strlen(password) < 6 {
		return nil, resd.RpcErrEncode(resd.NewErrWithTempCtx(l.ctx, "密码不符合安全要求", resd.ReqParamFormatErr1, resd.VarPassword))
	}
	// 验证是否已注册
	userModel := model.NewUserMainModel(l.ctx, l.svcCtx.SqlConn, in.PlatId)
	existUser, err := userModel.Ctx(l.ctx).Where("account = ?", account).Find()
	if err != nil {
		return nil, resd.RpcErrEncode(err)
	}
	if existUser != nil {
		return nil, resd.RpcErrEncode(resd.NewErrWithTempCtx(l.ctx, "用户已存在", resd.DataExist1, "Account"))
	}
	addData, err := l.prepareUserMain(in)
	if err != nil {
		return nil, resd.RpcErrEncode(resd.ErrorCtx(l.ctx, err))
	}
	addData.Account = account
	addData.Password = utild.Sha1(password)
	addData.Id = utild.MakeId()
	_, err = userModel.Insert(addData)
	if err != nil {
		return nil, resd.RpcErrEncode(resd.NewErrCtx(l.ctx, "新增用户失败", resd.MysqlInsertErr))
	}
	userInfo := &user.UserMainInfo{}
	if err := copier.Copy(&userInfo, addData); err != nil {
		return nil, resd.RpcErrEncode(resd.NewErrCtx(l.ctx, "copier失败", resd.CopierErr))
	}
	if in.IsLogin == 1 {
		//生成token并存入缓存
		s := fmt.Sprintf("%d-%d-%d", userInfo.Id, utild.GetStamp(), utild.Rand(11111, 99999))
		token := utild.Sha256(s)
		// 因userInfo的id转json是字符串，rpc取出来转化int64会报错，这里用rpc的结构体来缓存
		expireNum := l.svcCtx.Config.Conf.LoginTokenExSec
		err = l.svcCtx.Redis.SetDataExCtx(l.ctx, "userToken", token, userInfo, expireNum)
		nowStamp := utild.GetStamp()
		if err != nil {
			return nil, resd.RpcErrEncode(resd.NewErrCtx(l.ctx, "缓存失败", resd.RedisSetErr))
		}
		userInfo.PlatId = in.PlatId
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
