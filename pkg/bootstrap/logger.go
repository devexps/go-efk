package bootstrap

import (
	"os"

	"github.com/devexps/go-efk/api/gen/go/common/conf"

	"github.com/devexps/go-micro/v2/log"
	"github.com/devexps/go-micro/v2/middleware/tracing"

	fluentLogger "github.com/devexps/go-micro/log/fluent/v2"
)

type LoggerType string

const (
	LoggerTypeStd    LoggerType = "std"
	LoggerTypeFluent LoggerType = "fluent"
)

// NewLoggerProvider creates a new logger provider
func NewLoggerProvider(cfg *conf.Logger, serviceInfo *ServiceInfo) log.Logger {
	l := NewLogger(cfg)

	return log.With(
		l,
		"service.id", serviceInfo.Id,
		"service.name", serviceInfo.Name,
		"service.version", serviceInfo.Version,
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"trace_id", tracing.TraceID(),
		"span_id", tracing.SpanID(),
	)
}

// NewLogger creates a new logger
func NewLogger(cfg *conf.Logger) log.Logger {
	if cfg == nil {
		return NewStdLogger()
	}

	switch LoggerType(cfg.Type) {
	default:
		fallthrough
	case LoggerTypeStd:
		return NewStdLogger()
	case LoggerTypeFluent:
		return NewFluentLogger(cfg)
	}
}

// NewStdLogger creates a new logger - Micro built-in, console output
func NewStdLogger() log.Logger {
	return log.NewStdLogger(os.Stdout)
}

// NewFluentLogger creates a new logger - Fluent
func NewFluentLogger(cfg *conf.Logger) log.Logger {
	wrapped, err := fluentLogger.NewLogger(cfg.Fluent.Endpoint)
	if err != nil {
		panic("create fluent logger failed")
		return nil
	}
	return wrapped
}
