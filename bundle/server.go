package bundlefx

import (
	"context"
	"net/http"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func StartServer(
	lifecycle fx.Lifecycle,
	shutdowner fx.Shutdowner,
	server *http.Server,
	logger *zap.Logger,
) {
	lifecycle.Append(fx.StartHook(func() {
		go func() {
			logger.Info("HTTP server start", zap.String("addr", server.Addr))
			if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				logger.Error("HTTP server unexpected stopped", zap.Error(err))
			} else {
				logger.Info("HTTP server stopped")
			}
			_ = shutdowner.Shutdown(fx.ExitCode(1))
		}()
	}))
	lifecycle.Append(fx.StopHook(func(ctx context.Context) error {
		logger.Info("stopping HTTP server")
		return server.Shutdown(ctx)
	}))
}
