package logic

import (
	"context"
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/trace"
	"go-zero-dandan/app/message/rpc/message"
	"go-zero-dandan/app/user/api/internal/svc"
	"go-zero-dandan/app/user/api/internal/types"
	"go-zero-dandan/common/constd"
	"go-zero-dandan/common/respd"
	"go-zero-dandan/common/utild"
	"go.opentelemetry.io/otel"
	oteltrace "go.opentelemetry.io/otel/trace"
	"strconv"
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
	fmt.Println("雪花id:", utild.MakeId())

	if check := utild.CheckIsPhone(phone); check == false {
		return nil, respd.FailCode(localizer, respd.ReqPhoneError, []string{})
	}
	code := strconv.Itoa(utild.Rand(1000, 9999))
	err = l.svcCtx.Redis.Set("verifyCode", phone, code, 300)
	if err != nil {
		fmt.Println("redis error,", err)
	}

	resp = &types.SuccessResp{Msg: respd.Msg(localizer, respd.Ok)}
	if l.svcCtx.Mode == constd.ModeDev {
		fmt.Println("code：", code)
		return resp, nil
	} else {
		_, rpcErr := l.svcCtx.MessageRpc.SendPhone(context.Background(), &message.SendPhoneReq{
			Phone:    phone,
			TempId:   1,
			TempData: []string{code, "5"},
		})
		if rpcErr != nil {
			return nil, respd.RpcFail(localizer, rpcErr)
		}
		return resp, nil
	}

	return resp, nil
}
func (l *GetPhoneVerifyCodeLogic) tracer(spanName string) {
	//创建一个链路
	tracer := otel.GetTracerProvider().Tracer(trace.TraceName)
	//开始，如果要下级，第一参数就是ctx，可以给下级start
	_, span := tracer.Start(l.ctx, spanName, oteltrace.WithSpanKind(oteltrace.SpanKindInternal))
	//结束方法
	defer span.End()
	//业务处理
}
