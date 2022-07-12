package zapcore

import "context"

type traceIdKey struct{}

var ctxTraceId = traceIdKey{}

// A TraceIdEncoder serializes a Context to a primitive type.
type TraceIdEncoder func(context.Context, PrimitiveArrayEncoder)

func ContextTraceIdEncoder(ctx context.Context, enc PrimitiveArrayEncoder) {
	if ctx == nil {
		return
	}
	traceId := ctx.Value(ctxTraceId)
	if traceId == nil {
		return
	}
	enc.AppendString(traceId.(string))
}
