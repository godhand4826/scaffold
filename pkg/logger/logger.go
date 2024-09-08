package logger

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	Level                string
	EnableConsoleEncoder bool
	EnableCaller         bool
	EnableUTC            bool
}

func New(config Config) (*zap.Logger, error) {
	level, err := zapcore.ParseLevel(config.Level)
	if err != nil {
		return nil, err
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeCaller = zapcore.FullCallerEncoder
	encoderConfig.CallerKey = "cs"
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.ConsoleSeparator = " "

	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	if config.EnableUTC {
		encoderConfig.EncodeTime = ISO8601UTCTimeEncoder
	}

	encoder := zapcore.NewJSONEncoder(encoderConfig)
	if config.EnableConsoleEncoder {
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	return zap.New(
		zapcore.NewCore(encoder, zapcore.Lock(os.Stdout), level),
		zap.WithCaller(config.EnableCaller),
	), nil
}

func ISO8601UTCTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	zapcore.ISO8601TimeEncoder(t.UTC(), enc)
}
