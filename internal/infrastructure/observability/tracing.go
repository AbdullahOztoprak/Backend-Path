package observability

import (
    "context"
    "github.com/opentracing/opentracing-go"
    "github.com/opentracing/opentracing-go/ext"
    "github.com/uber/jaeger-client-go"
    "github.com/uber/jaeger-client-go/config"
    "log"
)

type Tracer struct {
    tracer opentracing.Tracer
}

func NewTracer(serviceName string) *Tracer {
    cfg := config.Configuration{
        ServiceName: serviceName,
        Sampler: &config.SamplerConfig{
            Type:  "const",
            Param: 1,
        },
        Reporter: &config.ReporterConfig{
            LogSpans:            true,
            BufferFlushInterval: 1,
        },
    }

    tracer, closer, err := cfg.NewTracer()
    if err != nil {
        log.Fatalf("could not initialize jaeger tracer: %v", err)
    }
    opentracing.SetGlobalTracer(tracer)

    return &Tracer{tracer: tracer}
}

func (t *Tracer) StartSpan(operationName string, ctx context.Context) opentracing.Span {
    span, _ := opentracing.StartSpanFromContext(ctx, operationName)
    return span
}

func (t *Tracer) FinishSpan(span opentracing.Span, err error) {
    if err != nil {
        ext.Error.Set(span, true)
        span.SetTag("error.message", err.Error())
    }
    span.Finish()
}