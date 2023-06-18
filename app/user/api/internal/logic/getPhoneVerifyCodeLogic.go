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
	"go.opentelemetry.io/otel/propagation"
	oteltrace "go.opentelemetry.io/otel/trace"
	"strconv"
	"time"
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
	logx.Error("测试日志")
	l.local()
	tracer := otel.GetTracerProvider().Tracer(trace.TraceName)
	spanCtx, span := tracer.Start(l.ctx, "testParent1", oteltrace.WithSpanKind(oteltrace.SpanKindProducer))
	carrier := &propagation.HeaderCarrier{}
	otel.GetTextMapPropagator().Inject(spanCtx, carrier)

	localizer := l.ctx.Value("lang").(*i18n.Localizer)
	fmt.Println("雪花id:", utild.MakeId())

	span.End()

	if check := utild.CheckIsPhone(phone); check == false {
		return nil, respd.FailCode(localizer, respd.ReqPhoneError, []string{})
	}
	code := strconv.Itoa(utild.Rand(1000, 9999))
	err = l.svcCtx.Redis.Set("verifyCode", phone, code, 300)
	if err != nil {
		fmt.Println("redis error,", err)
	}
	time.Sleep(1 * time.Second)
	defer span.End()

	wireContext := otel.GetTextMapPropagator().Extract(l.ctx, carrier)
	tracer2 := otel.GetTracerProvider().Tracer(trace.TraceName)
	_, span2 := tracer2.Start(wireContext, "testChild1", oteltrace.WithSpanKind(oteltrace.SpanKindConsumer))
	time.Sleep(200 * time.Microsecond)
	span2.End()

	tracer3 := otel.GetTracerProvider().Tracer(trace.TraceName)
	span3Ctx, span3 := tracer3.Start(spanCtx, "testChild2", oteltrace.WithSpanKind(oteltrace.SpanKindConsumer))
	time.Sleep(500 * time.Microsecond)
	span3.End()

	tracer4 := otel.GetTracerProvider().Tracer(trace.TraceName)
	_, span4 := tracer4.Start(span3Ctx, "testChild222", oteltrace.WithSpanKind(oteltrace.SpanKindInternal))
	time.Sleep(500 * time.Microsecond)
	span4.End()
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
func (l *GetPhoneVerifyCodeLogic) local() {
	tracer := otel.GetTracerProvider().Tracer(trace.TraceName)
	_, span := tracer.Start(l.ctx, "testParent2", oteltrace.WithSpanKind(oteltrace.SpanKindInternal))
	defer span.End()
	time.Sleep(1 * time.Second)
}
