package logfx

import (
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"scaffold/pkg/log"
)

var Module = fx.Options(
	fx.WithLogger(NewFxLogger),
	fx.Provide(NewLogger),
)

func NewLogger(lifecycle fx.Lifecycle, config log.Config) (*zap.Logger, error) {
	logger, err := log.New(config)
	if err != nil {
		return nil, err
	}

	logger.Info("logger start", zap.String("level", logger.Level().String()))

	lifecycle.Append(fx.StopHook(func() {
		logger.Info("flushing logger buffer")
		_ = logger.Sync()
	}))

	return logger, err
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
