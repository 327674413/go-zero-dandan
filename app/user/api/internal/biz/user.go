package biz

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/message/rpc/message"
	"go-zero-dandan/app/user/api/internal/svc"
	"go-zero-dandan/app/user/api/internal/types"
	"go-zero-dandan/app/user/model"
	"go-zero-dandan/app/user/rpc/types/userRpc"
	"go-zero-dandan/common/constd"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/typed"
	"go-zero-dandan/common/utild"
	"go-zero-dandan/common/utild/copier"
	"strconv"
)

type UserBiz struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	resd   *resd.Resp
	meta   *typed.ReqMeta
}

func NewUserBiz(ctx context.Context, svcCtx *svc.ServiceContext, resp *resd.Resp, meta *typed.ReqMeta) *UserBiz {
	biz := &UserBiz{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		resd:   resp,
		meta:   meta,
	}
	return biz
}
func (t *UserBiz) defaultRegByPhone(regInfo *UserRegInfo) (res *types.UserInfoResp, err error) {
	unionModel := model.NewUserUnionModel(t.ctx, t.svcCtx.SqlConn)
	tx, err := unionModel.StartTrans()
	if err != nil {
		return nil, t.resd.Error(err)
	}
	unionInfo := &model.UserUnion{}
	unionInfo.Id = utild.MakeId()
	_, err = unionModel.TxInsert(tx, unionInfo)
	if err != nil {
		return nil, t.resd.Error(err)
	}
	userMain := &model.UserMain{
		Id:        unionInfo.Id,
		UnionId:   unionInfo.Id,
		StateEm:   constd.UserStateEmNormal,
		Phone:     regInfo.Phone,
		PhoneArea: regInfo.PhoneArea,
	}
	userMainModel := model.NewUserMainModel(t.ctx, t.svcCtx.SqlConn, t.meta.PlatId)
	if err != nil {
		return nil, t.resd.Error(err)
	}
	_, err = userMainModel.TxInsert(tx, userMain)
	if err != nil {
		return nil, t.resd.Error(err)
	}
	err = unionModel.Commit(tx)
	if err != nil {
		return nil, t.resd.Error(err)
	}
	res = &types.UserInfoResp{}
	copier.Copy(&res, userMain)
	return res, nil
}
func (t *UserBiz) RegByPhone(regInfo *UserRegInfo) (res *types.UserInfoResp, err error) {
	regByPhoneStrage := map[int64]func(*UserRegInfo) (*types.UserInfoResp, error){}
	if strateFunc, ok := regByPhoneStrage[t.meta.PlatClasEm]; ok {
		return strateFunc(regInfo)
	} else {
		return t.defaultRegByPhone(regInfo)
	}
	return nil, nil
}
func (t *UserBiz) SendPhoneVerifyCode(phone string, phoneArea string) (string, error) {
	//生成验证码
	code := strconv.Itoa(utild.Rand(1000, 9999))
	err := t.svcCtx.Redis.SetExCtx(t.ctx, "verifyCode", phone, code, 300)
	if err != nil {
		return "", t.resd.Error(err)
	}
	currAt := fmt.Sprintf("%d", utild.GetStamp())
	err = t.svcCtx.Redis.SetExCtx(t.ctx, "verifyCodeGetAt", phone, currAt, 60)
	if err != nil {
		return "", t.resd.Error(err)
	}
	if t.svcCtx.Mode == constd.ModeDev {
		return code, nil
	} else {
		tempId := "1"
		_, rpcErr := t.svcCtx.MessageRpc.SendPhone(context.Background(), &message.SendPhoneReq{
			Phone:    &phone,
			TempId:   &tempId,
			TempData: []string{code, "5"},
		})
		if rpcErr != nil {
			return "", resd.Error(rpcErr)
		}
		return code, nil
	}
}
func (t *UserBiz) CheckPhoneVerifyCode(phone string, phoneArea string, code string) error {
	targetCode, err := t.svcCtx.Redis.Get("verifyCode", phone)
	if err != nil {
		return resd.Error(err)
	}
	if targetCode == "" {
		return t.resd.NewErr(resd.ErrVerifyCodeExpired)
	}
	if targetCode != code {
		return t.resd.NewErr(resd.ErrVerifyCodeWrong)
	}
	_, err = t.svcCtx.Redis.Del("verifyCode", phone)

	if err != nil {
		return t.resd.NewErr(resd.ErrVerifyCodeExpired)
	}
	return nil
}
func (t *UserBiz) CreateLoginState(userInfo *types.UserInfoResp) (string, error) {
	s := fmt.Sprintf("%d-%d-%d", userInfo.Id, utild.GetStamp(), utild.Rand(11111, 99999))
	token := utild.Sha256(s)
	// 因userInfo的id转json是字符串，rpc取出来转化int64会报错，这里用rpc的结构体来缓存
	cacheData := &userRpc.UserMainInfo{
		Id:        userInfo.Id,
		UnionId:   userInfo.UnionId,
		Account:   userInfo.Account,
		Nickname:  userInfo.Nickname,
		Phone:     userInfo.Phone,
		PhoneArea: userInfo.PhoneArea,
		SexEm:     userInfo.SexEm,
		Email:     userInfo.Email,
		AvatarImg: userInfo.Avatar,
		PlatId:    userInfo.PlatId,
	}
	err := t.svcCtx.Redis.SetDataExCtx(t.ctx, "userToken", token, cacheData, t.svcCtx.Config.Conf.LoginTokenExSec)
	if err != nil {
		return "", resd.Error(err)
	}
	return token, nil
}
func (t *UserBiz) EditUserInfo(editUserInfoReq *userRpc.EditUserInfoReq) error {
	userRpc := t.svcCtx.UserRpc
	_, err := userRpc.EditUserInfo(t.ctx, editUserInfoReq)
	if err != nil {
		return resd.ErrorCtx(t.ctx, err)
	}
	return nil
}
