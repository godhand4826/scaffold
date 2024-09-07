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
	fx.Invoke(PrintRoutesOnStart),
	fx.Invoke(restful.RegisterMiddlewares),
	fx.Invoke(restful.RegisterMetricsHandler),
	fx.Invoke(restful.RegisterHealthCheckHandler),
	fx.Invoke(restful.RegisterHandlers),
)

func PrintRoutesOnStart(lifecycle fx.Lifecycle, logger *zap.Logger, router chi.Router) {
	lifecycle.Append(fx.StartHook(func() {
		chi.Walk(router, func(
			method string,
			route string,
			handler http.Handler,
			middlewares ...func(http.Handler) http.Handler,
		) error {
			logger.Info("handler registered",
				zap.String("method", method),
				zap.String("route", route),
				zap.Int("middlewares", len(middlewares)),
			)
			return nil
		})
	}))
}
