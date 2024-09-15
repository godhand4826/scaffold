package log

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var defaultLogger *zap.Logger

func init() {
	logger, err := New(Config{})
	if err != nil {
		panic(err)
	}
	defaultLogger = logger.With(zap.Bool("default", true))
}

type ctxKey int

const loggerKey ctxKey = 0

func WithLogger(ctx context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}

func Get(ctx context.Context) *zap.Logger {
	if logger := ctx.Value(loggerKey); logger != nil {
		return logger.(*zap.Logger)
	}

	// fallback to default logger
	return defaultLogger
}

func With(ctx context.Context, fields ...zapcore.Field) context.Context {
	return WithLogger(ctx, Get(ctx).With(fields...))
}
