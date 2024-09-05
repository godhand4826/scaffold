package loggerfx

import (
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"scaffold/pkg/logger"
)

var Module = fx.Options(
	fx.WithLogger(NewFxLogger),
	fx.Provide(logger.New),
	fx.Invoke(RegisterHooks),
)

func RegisterHooks(lifecycle fx.Lifecycle, logger *zap.Logger) {
	lifecycle.Append(fx.StartHook(func() {
		logger.Info("logger start", zap.String("level", logger.Level().String()))
	}))
	lifecycle.Append(fx.StopHook(func() {
		logger.Info("flushing logger buffer")
		_ = logger.Sync()
	}))
}

func NewFxLogger(fxVerbose bool, logger *zap.Logger) fxevent.Logger {
	if !fxVerbose {
		logger = logger.WithOptions(zap.IncreaseLevel(zapcore.ErrorLevel))
	}

	fxLogger := &fxevent.ZapLogger{
		Logger: logger,
	}
	fxLogger.UseLogLevel(zapcore.DebugLevel)

	return fxLogger
}
