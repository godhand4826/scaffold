package examplefx

import (
	"go.uber.org/fx"

	"scaffold/internal/example"
	"scaffold/pkg/restful"
)

var Module = fx.Options(
	fx.Provide(fx.Annotate(
		example.NewRouteHandler,
		fx.As(new(restful.RouteHandler)),
		fx.ResultTags(`group:"route_handler"`),
	)),
)
