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
	"go-zero-dandan/common/constd"
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
		logx.Error(err)
		return nil, resd.FailCode(t.lang, resd.MysqlStartTransErr)
	}
	unionInfo := &model.UserUnion{}
	unionInfo.Id = utild.MakeId()
	_, err = unionModel.TxInsert(tx, t.ctx, unionInfo)
	if err != nil {
		logx.Error(err)
		return nil, resd.FailCode(t.lang, resd.MysqlInsertErr)
	}
	userMain := &model.UserMain{
		Id:        unionInfo.Id,
		UnionId:   unionInfo.Id,
		StateEm:   constd.UserStateEmNormal,
		Phone:     regInfo.Phone,
		PhoneArea: regInfo.PhoneArea,
	}
	userMainModel := model.NewUserMainModel(t.svcCtx.SqlConn, t.platId)
	_, err = userMainModel.TxInsert(tx, t.ctx, userMain)
	if err != nil {
		logx.Error(err)
		return nil, resd.FailCode(t.lang, resd.MysqlInsertErr)
	}
	err = tx.Commit()
	if err != nil {
		logx.Error(err)
		return nil, resd.FailCode(t.lang, resd.MysqlCommitErr)
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
		return "", resd.FailCode(t.lang, resd.RedisSetErr)
	}
	currAt := fmt.Sprintf("%d", utild.GetStamp())
	err = t.svcCtx.Redis.Set("verifyCodeGetAt", phone, currAt, 60)
	if err != nil {
		return "", resd.FailCode(t.lang, resd.RedisSetErr)
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
		return resd.FailCode(t.lang, resd.RedisGetErr)
	}
	if targetCode == "" {
		return resd.FailCode(t.lang, resd.VerifyCodeExpired)
	}
	if targetCode != code {
		return resd.FailCode(t.lang, resd.VerifyCodeWrong)
	}
	_, err = t.svcCtx.Redis.Del("verifyCode", phone)

	if err != nil {
		return resd.FailCode(t.lang, resd.VerifyCodeExpired)
	}
	return nil
}
func (t *UserBiz) initPlat() (err error) {
	platClasEm := utild.AnyToInt64(t.ctx.Value("platClasEm"))
	if platClasEm == 0 {
		return resd.FailCode(t.lang, resd.PlatClasErr)
	}
	platClasId := utild.AnyToInt64(t.ctx.Value("platId"))
	if platClasId == 0 {
		return resd.FailCode(t.lang, resd.PlatIdErr)
	}
	t.platId = platClasId
	t.platClasEm = platClasEm
	return nil
}
