package account

import (
	"context"
	"fmt"
	"go-zero-dandan/app/message/rpc/message"
	"go-zero-dandan/common/constd"
	"go-zero-dandan/common/resd"
	"strconv"

	"go-zero-dandan/app/user/api/internal/svc"
	"go-zero-dandan/app/user/api/internal/types"

	"go-zero-dandan/common/utild"
)

type GetPhoneVerifyCodeLogic struct {
	*GetPhoneVerifyCodeLogicGen
}

func NewGetPhoneVerifyCodeLogic(ctx context.Context, svc *svc.ServiceContext) *GetPhoneVerifyCodeLogic {
	return &GetPhoneVerifyCodeLogic{
		GetPhoneVerifyCodeLogicGen: NewGetPhoneVerifyCodeLogicGen(ctx, svc),
	}
}

func (l *GetPhoneVerifyCodeLogic) GetPhoneVerifyCode(req *types.GetPhoneVerifyCodeReq) (resp *types.SuccessResp, err error) {
	if err := l.initReq(req); err != nil {
		return nil, l.resd.Error(err)
	}
	//生成验证码
	code := strconv.Itoa(utild.Rand(1000, 9999))
	err = l.svc.Redis.SetExCtx(l.ctx, "verifyCode", l.req.Phone, code, 300)
	if err != nil {
		return nil, l.resd.Error(err, resd.ErrRedisSet)
	}
	currAt := fmt.Sprintf("%d", utild.GetStamp())
	err = l.svc.Redis.SetExCtx(l.ctx, "verifyCodeGetAt", l.req.Phone, currAt, 60)
	if err != nil {
		return nil, l.resd.Error(err, resd.ErrRedisSet)
	}
	resp = &types.SuccessResp{Msg: ""}
	if l.svc.Config.Mode == constd.ModeDev {
		resp.Msg = code
		return resp, nil
	} else {
		tempId := "1"
		_, err = l.svc.MessageRpc.SendPhone(context.Background(), &message.SendPhoneReq{
			Phone:    &l.req.Phone,
			TempId:   &tempId,
			TempData: []string{code, "5"},
		})
		if err != nil {
			return nil, l.resd.Error(err)
		}
		return resp, nil
	}

}
