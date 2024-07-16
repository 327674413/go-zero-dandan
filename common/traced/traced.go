package traced

import (
	"context"
	"github.com/zeromicro/go-zero/core/contextx"
	"github.com/zeromicro/go-zero/core/trace"
	"go.opentelemetry.io/otel"
	oteltrace "go.opentelemetry.io/otel/trace"
)

// 有点忘了trace怎么加了， 到时要复习下

// NewTracerStart 创建一个子tracer，需要主动调用span.End()方法,如果继续要创建下级就用upperCtx来创建
func NewTracerStart(spanName string, zeroLogicCtx ...context.Context) (upperCtx context.Context, span oteltrace.Span) {
	//创建一个链路
	tracer := otel.GetTracerProvider().Tracer(trace.TraceName)
	//开始，如果要下级，第一参数就是ctx，可以给下级start
	var ctx context.Context
	if len(zeroLogicCtx) > 0 && zeroLogicCtx[0] != nil {
		ctx = contextx.ValueOnlyFrom(zeroLogicCtx[0])
	} else {
		ctx = context.Background()
	}
	return tracer.Start(ctx, spanName, oteltrace.WithSpanKind(oteltrace.SpanKindInternal))
}
func NewTarcerChild(spanName string, upperCtx context.Context) (childCtx context.Context, span oteltrace.Span) {
	tracer := otel.GetTracerProvider().Tracer(trace.TraceName)
	return tracer.Start(upperCtx, spanName, oteltrace.WithSpanKind(oteltrace.SpanKindInternal))
}
