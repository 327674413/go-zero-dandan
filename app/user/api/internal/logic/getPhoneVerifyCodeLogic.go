package logic

import (
	"context"
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"go-zero-dandan/app/message/rpc/message"
	"go-zero-dandan/common/errd"
	"go-zero-dandan/common/land"
	"go-zero-dandan/common/util"

	"go-zero-dandan/app/user/api/internal/svc"
	"go-zero-dandan/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPhoneVerifyCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPhoneVerifyCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPhoneVerifyCodeLogic {
	return &GetPhoneVerifyCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPhoneVerifyCodeLogic) GetPhoneVerifyCode(req *types.GetPhoneVerifyCodeReq) (resp *types.SuccessResp, err error) {
	phone := *req.Phone
	localizer := l.ctx.Value("lang").(*i18n.Localizer)
	phoneTextData := map[string]string{"Field": land.Trans(localizer, "PhoneNumbera")}
	fmt.Println("雪花id:", util.MakeId())
	if check := util.CheckIsPhone(phone); check == false {
		return nil, errd.FailCode(localizer, errd.ReqPhoneError, phoneTextData)
	}
	sendPhoneRes, err := l.svcCtx.MessageRpc.SendPhone(context.Background(), &message.SendPhoneReq{
		Phone: phone,
	})
	fmt.Println("RPC返回：", sendPhoneRes)
	if err != nil {
		return nil, errd.Fail(err.Error())
	}
	return
}
