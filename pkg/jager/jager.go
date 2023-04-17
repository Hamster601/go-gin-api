package jager

import (
	"io"
	"log"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jagerlog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"
)

func NewGlobalTestTracer() (opentracing.Tracer, io.Closer, error) {
	cfg := jaegercfg.Configuration{
		ServiceName: "test",
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans: true,
		},
	}

	// 示例 Logger 和 Metric 分别使用 github.com/uber/jaeger-client-go/log 和 github.com/uber/jaeger-lib/metrics
	jLogger := jagerlog.StdLogger
	jMetricsFactory := metrics.NullFactory

	// 初始化 Tracer 实例
	tracer, closer, err := cfg.NewTracer(
		jaegercfg.Logger(jLogger),
		jaegercfg.Metrics(jMetricsFactory),
		// 设置最大 Tag 长度，根据情况设置
		jaegercfg.MaxTagValueLength(65535),
	)
	if err != nil {
		log.Printf("Could not initialize jaeger tracer: %s", err.Error())
		panic(err)
	}
	return tracer, closer, err
}
