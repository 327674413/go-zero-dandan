package logic

import (
	"context"
	"go-zero-dandan/common/api"
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
	if check := util.CheckIsPhone(phone); check == false {
		return nil, api.Fail("请输入正确的手机号")
	}
	/*sendPhoneRes, err := l.svcCtx.MessageRpc.SendPhone(context.Background(), &message.SendPhoneReq{
		Phone: phone,
	})
	fmt.Println("RPC返回：", sendPhoneRes, err)
	if err != nil {
		return nil, err
	}*/
	return
}
