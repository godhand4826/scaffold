package pgfx

import (
	"context"
	"time"

	"go.uber.org/fx"
	"go.uber.org/zap"

	"scaffold/ent"
	"scaffold/pkg/pg"
)

var Module = fx.Options(
	fx.Provide(NewPostgres),
)

func NewPostgres(
	lifecycle fx.Lifecycle,
	shutdowner fx.Shutdowner,
	logger *zap.Logger,
	config pg.Config,
) (*ent.Client, error) {
	client, err := pg.New(config, logger)
	if err != nil {
		return nil, err
	}

	lifecycle.Append(fx.StartHook(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		logger.With(
			zap.String("host", config.Host),
			zap.String("port", config.Port),
			zap.String("database", config.Database),
		).Info("ping database")

		if err := client.Ping(ctx); err != nil {
			logger.With(zap.Error(err)).Error("failed to ping postgres server")

			_ = shutdowner.Shutdown(fx.ExitCode(1))
			return
		}
	}))

	lifecycle.Append(fx.StopHook(func() {
		_ = client.Close()

		logger.Info("postgres client stopped")
	}))

	return client, nil
}
