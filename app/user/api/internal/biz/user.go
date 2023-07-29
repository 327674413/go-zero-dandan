package biz

import (
	"context"
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/message/rpc/message"
	"go-zero-dandan/app/user/api/internal/svc"
	"go-zero-dandan/app/user/api/internal/types"
	"go-zero-dandan/app/user/model"
	"go-zero-dandan/app/user/rpc/types/pb"
	"go-zero-dandan/common/constd"
	"go-zero-dandan/common/dao"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/utild"
	"strconv"
)

type UserBiz struct {
	logx.Logger
	ctx        context.Context
	svcCtx     *svc.ServiceContext
	lang       *i18n.Localizer
	platId     int64
	platClasEm int64
}

func NewUserBiz(ctx context.Context, svcCtx *svc.ServiceContext) *UserBiz {
	localizer := ctx.Value("lang").(*i18n.Localizer)
	biz := &UserBiz{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		lang:   localizer,
	}
	biz.initPlat()
	return biz
}
func (t *UserBiz) defaultRegByPhone(regInfo *UserRegInfo) (res *types.UserInfoResp, err error) {
	unionModel := model.NewUserUnionModel(t.svcCtx.SqlConn, t.platId)
	db, _ := t.svcCtx.SqlConn.RawDB()
	tx, err := db.BeginTx(t.ctx, nil)
	if err != nil {
		return nil, resd.Error(err, resd.MysqlStartTransErr)
	}
	unionInfo := &model.UserUnion{}
	unionInfo.Id = utild.MakeId()
	data, err := dao.PrepareData(unionInfo)
	if err != nil {
		return nil, resd.Error(err)
	}
	_, err = unionModel.Ctx(t.ctx).TxInsert(tx, data)
	if err != nil {
		return nil, resd.Error(err, resd.MysqlInsertErr)
	}
	userMain := &model.UserMain{
		Id:        unionInfo.Id,
		UnionId:   unionInfo.Id,
		StateEm:   constd.UserStateEmNormal,
		Phone:     regInfo.Phone,
		PhoneArea: regInfo.PhoneArea,
	}
	userMainModel := model.NewUserMainModel(t.svcCtx.SqlConn, t.platId)
	data, err = dao.PrepareData(userMain)
	if err != nil {
		return nil, resd.Error(err)
	}
	_, err = userMainModel.Ctx(t.ctx).TxInsert(tx, data)
	if err != nil {
		return nil, resd.Error(err)
	}
	err = tx.Commit()
	if err != nil {
		logx.Error(err)
		return nil, resd.Error(err)
	}
	res = &types.UserInfoResp{}
	utild.Copy(&res, userMain)
	return res, nil
}
func (t *UserBiz) RegByPhone(regInfo *UserRegInfo) (res *types.UserInfoResp, err error) {
	regByPhoneStrage := map[int64]func(*UserRegInfo) (*types.UserInfoResp, error){}
	if strateFunc, ok := regByPhoneStrage[t.platClasEm]; ok {
		return strateFunc(regInfo)
	} else {
		return t.defaultRegByPhone(regInfo)
	}
	return nil, nil
}
func (t *UserBiz) SendPhoneVerifyCode(phone string, phoneArea string) (string, error) {
	//生成验证码
	code := strconv.Itoa(utild.Rand(1000, 9999))
	err := t.svcCtx.Redis.Set("verifyCode", phone, code, 300)
	if err != nil {
		return "", resd.Error(err, resd.RedisSetErr)
	}
	currAt := fmt.Sprintf("%d", utild.GetStamp())
	err = t.svcCtx.Redis.Set("verifyCodeGetAt", phone, currAt, 60)
	if err != nil {
		return "", resd.Error(err, resd.RedisSetErr)
	}
	if t.svcCtx.Mode == constd.ModeDev {
		return code, nil
	} else {
		_, rpcErr := t.svcCtx.MessageRpc.SendPhone(context.Background(), &message.SendPhoneReq{
			Phone:    phone,
			TempId:   1,
			TempData: []string{code, "5"},
		})
		if rpcErr != nil {
			return "", resd.RpcFail(t.lang, rpcErr)
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
		return resd.NewErr("验证码失效", resd.VerifyCodeExpired)
	}
	if targetCode != code {
		return resd.NewErr("验证码失败", resd.VerifyCodeWrong)
	}
	_, err = t.svcCtx.Redis.Del("verifyCode", phone)

	if err != nil {
		return resd.NewErr("验证码过期", resd.VerifyCodeExpired)
	}
	return nil
}
func (t *UserBiz) CreateLoginState(userInfo *model.UserMain) (string, error) {
	s := fmt.Sprintf("%d-%d-%d", userInfo.Id, utild.GetStamp(), utild.Rand(11111, 99999))
	token := utild.Sha256(s)
	err := t.svcCtx.Redis.SetData("userToken", token, userInfo, t.svcCtx.Config.Conf.LoginTokenExSec)
	if err != nil {
		return "", resd.Error(err)
	}
	return token, nil
}
func (t *UserBiz) EditUserInfo(editUserInfoReq *pb.EditUserInfoReq) error {
	userRpc := t.svcCtx.UserRpc
	fmt.Println("nickname:", editUserInfoReq.Nickname, "avatar:", editUserInfoReq.Avatar, "email:", editUserInfoReq.Email, "sexEm:", editUserInfoReq.SexEm)
	_, err := userRpc.EditUserInfo(t.ctx, editUserInfoReq)
	if err != nil {
		return resd.RpcFail(t.lang, err)
	}
	return nil
}

func (t *UserBiz) initPlat() (err error) {
	platClasEm := utild.AnyToInt64(t.ctx.Value("platClasEm"))
	if platClasEm == 0 {
		return resd.NewErrCtx(t.ctx, "token中未获取到platClasEm", resd.PlatClasErr)
	}
	platClasId := utild.AnyToInt64(t.ctx.Value("platId"))
	if platClasId == 0 {
		return resd.NewErrCtx(t.ctx, "token中未获取到platId", resd.PlatIdErr)
	}
	t.platId = platClasId
	t.platClasEm = platClasEm
	return nil
}
