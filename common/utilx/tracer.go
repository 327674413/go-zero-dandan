package utilx

import (
	"context"
	"github.com/zeromicro/go-zero/core/trace"
	"go.opentelemetry.io/otel"
	oteltrace "go.opentelemetry.io/otel/trace"
)

func Tracer(ctx context.Context, spanName string, actionFn func()) context.Context {
	//创建一个链路
	tracer := otel.GetTracerProvider().Tracer(trace.TraceName)
	//开始，如果要下级，第一参数就是ctx，可以给下级start
	currCtx, span := tracer.Start(ctx, spanName, oteltrace.WithSpanKind(oteltrace.SpanKindInternal))
	//结束方法
	defer span.End()
	//执行具体的业务
	actionFn()
	//创建子追踪时使用
	return currCtx
}
