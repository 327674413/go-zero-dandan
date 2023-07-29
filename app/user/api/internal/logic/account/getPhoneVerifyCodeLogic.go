package account

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/trace"
	"go-zero-dandan/app/message/rpc/message"
	"go-zero-dandan/common/constd"
	"go-zero-dandan/common/resd"
	"go.opentelemetry.io/otel"
	oteltrace "go.opentelemetry.io/otel/trace"
	"strconv"

	"go-zero-dandan/app/user/api/internal/svc"
	"go-zero-dandan/app/user/api/internal/types"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/common/utild"
)

type GetPhoneVerifyCodeLogic struct {
	logx.Logger
	ctx        context.Context
	svcCtx     *svc.ServiceContext
	lang       *i18n.Localizer
	platId     int64
	platClasEm int64
}

func NewGetPhoneVerifyCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPhoneVerifyCodeLogic {
	localizer := ctx.Value("lang").(*i18n.Localizer)
	return &GetPhoneVerifyCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		lang:   localizer,
	}
}

func (l *GetPhoneVerifyCodeLogic) GetPhoneVerifyCode(req *types.GetPhoneVerifyCodeReq) (resp *types.SuccessResp, err error) {
	phone := *req.Phone

	//生成验证码
	code := strconv.Itoa(utild.Rand(1000, 9999))
	err = l.svcCtx.Redis.Set("verifyCode", phone, code, 300)
	if err != nil {
		return nil, resd.Error(err, resd.RedisSetErr)
	}
	currAt := fmt.Sprintf("%d", utild.GetStamp())
	err = l.svcCtx.Redis.Set("verifyCodeGetAt", phone, currAt, 60)
	if err != nil {
		return nil, resd.Error(err, resd.RedisSetErr)
	}
	resp = &types.SuccessResp{Msg: resd.Msg(l.lang, resd.Ok)}
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
			return nil, resd.RpcFail(l.lang, rpcErr)
		}
		return resp, nil
	}

	return resp, nil
}

func (l *GetPhoneVerifyCodeLogic) initPlat() (err error) {
	platClasEm := utild.AnyToInt64(l.ctx.Value("platClasEm"))
	if platClasEm == 0 {
		return resd.NewErrCtx(l.ctx, "token中未获取到platClasEm", resd.PlatClasErr)
	}
	platClasId := utild.AnyToInt64(l.ctx.Value("platId"))
	if platClasId == 0 {
		return resd.NewErrCtx(l.ctx, "token中未获取到platId", resd.PlatIdErr)
	}
	l.platId = platClasId
	l.platClasEm = platClasEm
	return nil
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
