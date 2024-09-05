package restfulfx

import (
	"context"
	"net/http"

	"go.uber.org/fx"
	"go.uber.org/zap"

	server "scaffold/pkg/restful"
)

var Module = fx.Options(
	fx.Provide(fx.Annotate(server.New, fx.ParamTags("", `name:"server_addr"`))),
	fx.Invoke(RegisterHooks),
)

func RegisterHooks(
	lifecycle fx.Lifecycle,
	shutdowner fx.Shutdowner,
	server *http.Server,
	logger *zap.Logger,
) {
	lifecycle.Append(fx.StartHook(func() {
		logger.Info("HTTP server start", zap.String("addr", server.Addr))

		go func() {
			err := server.ListenAndServe()
			if err != nil && err != http.ErrServerClosed {
				logger.Error("HTTP server unexpected stopped", zap.Error(err))
			} else {
				logger.Info("HTTP server stopped", zap.Error(err))
			}
			_ = shutdowner.Shutdown(fx.ExitCode(1))
		}()
	}))
	lifecycle.Append(fx.StopHook(func(ctx context.Context) error {
		logger.Info("stopping HTTP server")
		return server.Shutdown(ctx)
	}))
}
