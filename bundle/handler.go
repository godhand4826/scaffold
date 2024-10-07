package bundlefx

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"

	"scaffold/pkg/restful"
)

func RegisterHandlers(
	routeHandlers []restful.RouteHandler,
	logger *zap.Logger,
	router chi.Router,
) {
	for _, rh := range routeHandlers {
		rh.AttachOn(router)
	}

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
}
