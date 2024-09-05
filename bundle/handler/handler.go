package handlerfx

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"

	"scaffold/pkg/restful"
)

var Module = fx.Options(
	fx.Provide(fx.Annotate(chi.NewRouter, fx.As(new(http.Handler)), fx.As(fx.Self()))),
	fx.Provide(restful.ZapToRequestLoggerAdaptor),
	fx.Invoke(restful.RegisterMiddlewares),
	fx.Invoke(restful.RegisterMetricsHandler),
	fx.Invoke(restful.RegisterHealthCheckHandler),
	fx.Invoke(restful.RegisterHandlers),
)
