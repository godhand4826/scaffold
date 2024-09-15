package webfx

import (
	"go.uber.org/fx"

	"scaffold/internal/web"
	"scaffold/pkg/restful"
)

var Module = fx.Options(
	fx.Provide(fx.Annotate(
		web.NewRouteHandler,
		fx.As(new(restful.RouteHandler)),
		fx.ResultTags(`group:"route_handler"`),
	)),
)
