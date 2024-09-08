package handlerfx

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"scaffold/pkg/restful"
)

var Module = fx.Options(
	fx.Provide(fx.Annotate(chi.NewRouter, fx.As(new(http.Handler)), fx.As(new(chi.Router)))),
	fx.Provide(restful.ZapToRequestLoggerAdaptor),
	fx.Invoke(restful.RegisterMiddlewares),
	fx.Invoke(restful.RegisterMetricsHandler),
	fx.Invoke(restful.RegisterHealthCheckHandler),
	fx.Invoke(PrintRoutesOnStart),
)

func PrintRoutesOnStart(lifecycle fx.Lifecycle, logger *zap.Logger, router chi.Router) {
	lifecycle.Append(fx.StartHook(func() {
		_ = chi.Walk(router, func(
			method string,
			route string,
			_ http.Handler,
			middlewares ...func(http.Handler) http.Handler,
		) error {
			logger.Debug("handler registered",
				zap.String("method", method),
				zap.String("route", route),
				zap.Int("middlewares", len(middlewares)),
			)
			return nil
		})
	}))
}
